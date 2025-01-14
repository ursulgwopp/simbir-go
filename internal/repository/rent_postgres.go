package repository

import (
	"context"
	"database/sql"
	"errors"
	"math"
	"time"

	"github.com/ursulgwopp/simbir-go/internal/models"
)

func (r *PostgresRepository) GetAvailableTransport(latitude float64, longitude float64, radius float64, transportType string) ([]models.TransportResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var transports []models.TransportResponse
	var query string
	var rows *sql.Rows
	var err error

	if transportType == "All" {
		query = `SELECT id, can_be_rented, transport_type, model, color, identifier, description, latitude, longitude, minute_price, day_price FROM transports WHERE can_be_rented = true ORDER BY id`
		rows, err = r.db.QueryContext(ctx, query)
	} else {
		query = `SELECT id, can_be_rented, transport_type, model, color, identifier, description, latitude, longitude, minute_price, day_price FROM transports WHERE transport_type = $1 AND can_be_rented = true ORDER BY id`
		rows, err = r.db.QueryContext(ctx, query, transportType)
	}

	if err != nil {
		return []models.TransportResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var transport models.TransportResponse
		if err := rows.Scan(&transport.Id, &transport.CanBeRented, &transport.TransportType, &transport.Model, &transport.Color, &transport.Identifier, &transport.Description, &transport.Latitude, &transport.Longitude, &transport.MinutePrice, &transport.DayPrice); err != nil {
			return []models.TransportResponse{}, err
		}

		if IsAvailable(latitude, longitude, transport.Latitude, transport.Longitude, radius) {
			transports = append(transports, transport)
		}
	}

	if err := rows.Err(); err != nil {
		return []models.TransportResponse{}, err
	}

	return transports, nil
}

func (r *PostgresRepository) GetRent(rentId int) (models.RentResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var rent models.RentResponse = models.RentResponse{Id: rentId}
	query := `SELECT transport_id, user_id, time_start, time_end, price_of_unit, price_type, final_price, is_active FROM rents WHERE id = $1`
	if err := r.db.QueryRowContext(ctx, query, rentId).Scan(&rent.TransportId, &rent.UserId, &rent.TimeStart, &rent.TimeEnd, &rent.PriceOfUnit, &rent.PriceType, &rent.FinalPrice, &rent.IsActive); err != nil {
		return models.RentResponse{}, err
	}

	return rent, nil
}

func (r *PostgresRepository) GetTransportHistory(transportId int) ([]models.RentResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var rents []models.RentResponse
	query := `SELECT id, user_id, time_start, time_end, price_of_unit, price_type, final_price, is_active FROM rents WHERE transport_id = $1`
	rows, err := r.db.QueryContext(ctx, query, transportId)
	if err != nil {
		return []models.RentResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var rent models.RentResponse
		rent.TransportId = transportId
		if err := rows.Scan(&rent.Id, &rent.UserId, &rent.TimeStart, &rent.TimeEnd, &rent.PriceOfUnit, &rent.PriceType, &rent.FinalPrice, &rent.IsActive); err != nil {
			return []models.RentResponse{}, err
		}

		rents = append(rents, rent)
	}

	if err := rows.Err(); err != nil {
		return []models.RentResponse{}, err
	}

	return rents, nil
}

func (r *PostgresRepository) GetUserHistory(accountId int) ([]models.RentResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var rents []models.RentResponse
	query := `SELECT id, transport_id, time_start, time_end, price_of_unit, price_type, final_price, is_active FROM rents WHERE user_id = $1`
	rows, err := r.db.QueryContext(ctx, query, accountId)
	if err != nil {
		return []models.RentResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var rent models.RentResponse = models.RentResponse{UserId: accountId}
		if err := rows.Scan(&rent.Id, &rent.TransportId, &rent.TimeStart, &rent.TimeEnd, &rent.PriceOfUnit, &rent.PriceType, &rent.FinalPrice, &rent.IsActive); err != nil {
			return []models.RentResponse{}, err
		}

		rents = append(rents, rent)
	}

	if err := rows.Err(); err != nil {
		return []models.RentResponse{}, err
	}

	return rents, nil
}

func (r *PostgresRepository) StartRent(userId int, transportId int, rentType string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var minutePrice float64
	var dayPrice float64
	query := `SELECT minute_price, day_price FROM transports WHERE id = $1`
	if err := r.db.QueryRowContext(ctx, query, transportId).Scan(&minutePrice, &dayPrice); err != nil {
		return -1, err
	}

	var balance float64
	query = `SELECT balance FROM accounts WHERE id = $1`
	if err := r.db.QueryRowContext(ctx, query, userId).Scan(&balance); err != nil {
		return -1, err
	}

	var id int
	query = `INSERT INTO rents (transport_id, user_id, time_start, price_of_unit, price_type) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	if rentType == "Days" {
		if balance < dayPrice {
			return -1, errors.New("no money")
		}

		if err := r.db.QueryRowContext(ctx, query, transportId, userId, time.Now(), dayPrice, rentType).Scan(&id); err != nil {
			return -1, err
		}

		query = `UPDATE accounts SET balance = balance - $1 WHERE id = 2`
		_, err := r.db.ExecContext(ctx, query, dayPrice, userId)
		if err != nil {
			return -1, err
		}
	} else {
		if balance < minutePrice {
			return -1, errors.New("no money")
		}

		if err := r.db.QueryRowContext(ctx, query, transportId, userId, time.Now(), minutePrice, rentType).Scan(&id); err != nil {
			return -1, err
		}

		query = `UPDATE accounts SET balance = balance - $1 WHERE id = $2`
		_, err := r.db.ExecContext(ctx, query, minutePrice, userId)
		if err != nil {
			return -1, err
		}
	}

	query = `UPDATE transports SET can_be_rented = false WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, transportId)

	return id, err
}

func (r *PostgresRepository) StopRent(rentId int, latitude float64, longitude float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var rent models.RentResponse = models.RentResponse{TimeEnd: time.Now()}
	query := `SELECT transport_id, user_id, time_start, price_of_unit, price_type FROM rents WHERE id = $1`
	if err := r.db.QueryRowContext(ctx, query, rentId).Scan(&rent.TransportId, &rent.UserId, &rent.TimeStart, &rent.PriceOfUnit, &rent.PriceType); err != nil {
		return err
	}

	// update timeend
	query = `UPDATE rents SET time_end = $1, is_active = false WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, rent.TimeEnd, rentId)
	if err != nil {
		return err
	}

	// update final price
	if rent.PriceType == "Days" {
		days := float64(int(rent.TimeEnd.Sub(rent.TimeStart).Hours())/24) + 1
		rent.FinalPrice = days * rent.PriceOfUnit
	} else {
		minutes := float64(int(rent.TimeEnd.Sub(rent.TimeStart).Minutes())) + 1
		rent.FinalPrice = minutes * rent.PriceOfUnit
	}

	query = `UPDATE rents SET final_price = $1 WHERE id = $2`
	_, err = r.db.ExecContext(ctx, query, rent.FinalPrice, rentId)
	if err != nil {
		return err
	}

	// withdraw money
	// query = `UPDATE accounts SET balance = balance - $1 WHERE id = $2`
	// _, err = r.db.ExecContext(ctx, query, rent.FinalPrice, rent.UserId)
	// if err != nil {
	// 	return err
	// }

	// update canberented for car
	query = `UPDATE transports SET can_be_rented = true WHERE id = $1`
	_, err = r.db.ExecContext(ctx, query, rent.TransportId)
	if err != nil {
		return err
	}

	// update coordinates
	query = `UPDATE transports SET latitude = $1, longitude = $2 WHERE id = $3`
	_, err = r.db.ExecContext(ctx, query, latitude, longitude, rent.TransportId)
	if err != nil {
		return err
	}

	return nil
}

func IsAvailable(lat1, long1, lat2, long2, radius float64) bool {
	x := lat1 - lat2
	y := long1 - long2
	r := math.Sqrt(x*x + y*y)

	return r < radius
}
