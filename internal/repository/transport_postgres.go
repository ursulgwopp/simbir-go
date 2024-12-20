package repository

import (
	"github.com/ursulgwopp/simbir-go/internal/models"
)

func (r *PostgresRepository) CreateTransport(ownerId int, req models.TransportRequest) (int, error) {
	var id int

	query := `INSERT INTO transports (owner_id, can_be_rented, transport_type, model, color, identifier, description, latitude, longitude, minute_price, day_price) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`
	if err := r.db.QueryRow(query, ownerId, req.CanBeRented, req.TransportType, req.Model, req.Color, req.Identifier, req.Description, req.Latitude, req.Longitude, req.MinutePrice, req.DayPrice).Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}

func (r *PostgresRepository) GetTransport(transportId int) (models.TransportResponse, error) {
	var transport models.TransportResponse
	transport.Id = transportId

	query := `SELECT can_be_rented, transport_type, model, color, identifier, description, latitude, longitude, minute_price, day_price FROM transports WHERE id = $1`
	if err := r.db.QueryRow(query, transportId).Scan(&transport.CanBeRented, &transport.TransportType, &transport.Model, &transport.Color, &transport.Identifier, &transport.Description, &transport.Latitude, &transport.Longitude, &transport.MinutePrice, &transport.DayPrice); err != nil {
		return models.TransportResponse{}, err
	}

	return transport, nil
}

func (r *PostgresRepository) UpdateTransport(transportId int, req models.TransportRequest) error {
	query := `UPDATE transports SET can_be_rented = $1, model = $2, color = $3, identifier = $4, description = $5, latitude = $6, longitude = $7, minute_price = $8, day_price = $9 WHERE id = $10`
	_, err := r.db.Exec(query, req.CanBeRented, req.Model, req.Color, req.Identifier, req.Description, req.Latitude, req.Longitude, req.MinutePrice, req.DayPrice, transportId)

	return err
}

func (r *PostgresRepository) DeleteTransport(transportId int) error {
	query := `DELETE FROM transports WHERE id = $1`
	_, err := r.db.Exec(query, transportId)

	return err
}
