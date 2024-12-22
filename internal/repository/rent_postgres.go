package repository

import (
	"database/sql"
	"math"

	"github.com/ursulgwopp/simbir-go/internal/models"
)

func (r *PostgresRepository) GetAvailableTransport(latitude float64, longitude float64, radius float64, transportType string) ([]models.TransportResponse, error) {
	var query string
	var rows *sql.Rows
	var err error

	if transportType == "All" {
		query = `SELECT id, can_be_rented, transport_type, model, color, identifier, description, latitude, longitude, minute_price, day_price FROM transports ORDER BY id`
		rows, err = r.db.Query(query)
	} else {
		query = `SELECT id, can_be_rented, transport_type, model, color, identifier, description, latitude, longitude, minute_price, day_price FROM transports WHERE transport_type = $1 ORDER BY id`
		rows, err = r.db.Query(query, transportType)
	}

	if err != nil {
		return []models.TransportResponse{}, err
	}
	defer rows.Close()

	var transports []models.TransportResponse
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
		return nil, err
	}

	return transports, nil
}

func (r *PostgresRepository) GetRent(rentId int) (models.RentResponse, error) {
	var rent models.RentResponse
	rent.Id = rentId

	query := `SELECT transport_id, user_id, time_start, time_end, price_of_unit, price_type, final_price FROM rents WHERE id = $1`
	if err := r.db.QueryRow(query, rentId).Scan(&rent.TransportId, &rent.UserId, &rent.TimeStart, &rent.TimeEnd, &rent.PriceOfUnit, &rent.PriceType, &rent.FinalPrice); err != nil {
		return models.RentResponse{}, err
	}

	return rent, nil
}

func (r *PostgresRepository) GetTransportHistory(transportId int) ([]models.RentResponse, error) {
	query := `SELECT id, user_id, time_start, time_end, price_of_unit, price_type, final_price FROM rents WHERE transport_id = $1`
	rows, err := r.db.Query(query, transportId)
	if err != nil {
		return []models.RentResponse{}, err
	}
	defer rows.Close()

	var rents []models.RentResponse
	for rows.Next() {
		var rent models.RentResponse
		rent.TransportId = transportId
		if err := rows.Scan(&rent.Id, &rent.UserId, &rent.TimeStart, &rent.TimeEnd, &rent.PriceOfUnit, &rent.PriceType, &rent.FinalPrice); err != nil {
			return []models.RentResponse{}, err
		}

		rents = append(rents, rent)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return rents, nil
}

func (r *PostgresRepository) GetUserHistory(accountId int) ([]models.RentResponse, error) {
	query := `SELECT id, transport_id, time_start, time_end, price_of_unit, price_type, final_price FROM rents WHERE user_id = $1`
	rows, err := r.db.Query(query, accountId)
	if err != nil {
		return []models.RentResponse{}, err
	}
	defer rows.Close()

	var rents []models.RentResponse
	for rows.Next() {
		var rent models.RentResponse
		rent.UserId = accountId
		if err := rows.Scan(&rent.Id, &rent.TransportId, &rent.TimeStart, &rent.TimeEnd, &rent.PriceOfUnit, &rent.PriceType, &rent.FinalPrice); err != nil {
			return []models.RentResponse{}, err
		}

		rents = append(rents, rent)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return rents, nil
}

// StartRent implements service.Repository.
func (*PostgresRepository) StartRent(transportId int, rentType string) (int, error) {
	panic("unimplemented")
}

// StopRent implements service.Repository.
func (*PostgresRepository) StopRent(rentId int, latitude float64, longitude float64) error {
	panic("unimplemented")
}

func IsAvailable(lat1, long1, lat2, long2, radius float64) bool {
	x := lat1 - lat2
	y := long1 - long2
	r := math.Sqrt(x*x + y*y)

	return r < radius
}
