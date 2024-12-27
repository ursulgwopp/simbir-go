package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	go startMinutelyWithdrawingProcess(repo)
	go startDailyWithdrawingProcess(repo)

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

func startMinutelyWithdrawingProcess(repo service.Repository) {
	for {
		now := time.Now()
		nextMinute := now.Truncate(time.Minute).Add(time.Minute) // Get the start of the next minute

		// Wait until the next minute
		fmt.Printf("Waiting until %s...\n", nextMinute.Format("15:04:05"))
		time.Sleep(time.Until(nextMinute)) // Sleep until the next minute

		// Execute your withdrawal logic here
		fmt.Println("Withdrawing at:", time.Now())
		if err := repo.MinutelyPayment(); err != nil {
			log.Panic(err)
		}
	}
}

func startDailyWithdrawingProcess(repo service.Repository) {
	for {
		now := time.Now()
		nextMidnight := now.Truncate(24 * time.Hour).Add(21 * time.Hour)

		fmt.Printf("Waiting until %s...\n", nextMidnight.Format("15:04:05"))
		time.Sleep(time.Until(nextMidnight)) // Sleep until the next minute

		// Execute your withdrawal logic here
		fmt.Println("Withdrawing at:", time.Now())
		if err := repo.DailyPayment(); err != nil {
			log.Panic(err)
		}
	}
}
