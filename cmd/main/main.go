package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/ursulgwopp/simbir-go/configs"
	"github.com/ursulgwopp/simbir-go/internal/repository"
	"github.com/ursulgwopp/simbir-go/internal/server"
	"github.com/ursulgwopp/simbir-go/internal/service"
	"github.com/ursulgwopp/simbir-go/internal/transport"
)

// @title SimbirGO
// @version 1.0

// @host localhost:2024
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	if err := configs.InitConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(configs.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repo := repository.NewPostgresRepository(db)
	service := service.NewService(repo)
	transport := transport.NewTransport(service)

	srv := &server.Server{}
	go func() {
		if err := srv.Run(viper.GetString("port"), transport.InitRoutes()); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("error running http server: %s", err.Error())
		}
	}()

	go startWithdrawingProcess(repo, db)

	logrus.Print("App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("App Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func startWithdrawingProcess(repo service.Repository, db *sqlx.DB) {
	// Calculate the duration until the next minute
	now := time.Now()
	nextMinute := now.Truncate(time.Minute).Add(time.Minute)
	durationUntilNextMinute := nextMinute.Sub(now)

	// Wait until the next minute
	time.Sleep(durationUntilNextMinute)

	// Now set up a ticker to withdraw every minute
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := withdrawFromBalance(repo, db)
			if err != nil {
				log.Panic(err)
			} else {
				log.Println("successfully")
			}
		}
	}
}

type WithdrawInfo struct {
	RentId      int
	UserId      int
	PriceOfUnit int
}

func withdrawFromBalance(repo service.Repository, db *sqlx.DB) error {
	var users []WithdrawInfo
	query := `SELECT id, user_id, price_of_unit FROM rents WHERE is_active = TRUE AND price_type = 'Minutes'`
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var user WithdrawInfo
		if err := rows.Scan(&user.RentId, &user.UserId, &user.PriceOfUnit); err != nil {
			return err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	for _, user := range users {
		var balance int
		query := `SELECT balance FROM accounts WHERE id = $1`
		if err := db.QueryRow(query, user.UserId).Scan(&balance); err != nil {
			return err
		}

		if balance < user.PriceOfUnit {
			repo.StopRent(user.RentId, 61, 31)
			continue
		}

		query = `UPDATE accounts SET balance = balance - $1 WHERE id = $2`
		_, err := db.Exec(query, user.PriceOfUnit, user.UserId)
		if err != nil {
			return err
		}
	}

	return nil
}
