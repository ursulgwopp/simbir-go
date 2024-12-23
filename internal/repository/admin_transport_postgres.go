package repository

import (
	"context"
	"database/sql"

	"github.com/ursulgwopp/simbir-go/internal/models"
)

func (r *PostgresRepository) AdminCreateTransport(req models.AdminTransportRequest) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var id int
	query := `INSERT INTO transports (owner_id, can_be_rented, transport_type, model, color, identifier, description, latitude, longitude, minute_price, day_price) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`
	if err := r.db.QueryRowContext(ctx, query, req.OwnerId, req.CanBeRented, req.TransportType, req.Model, req.Color, req.Identifier, req.Description, req.Latitude, req.Longitude, req.MinutePrice, req.DayPrice).Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}

func (r *PostgresRepository) AdminGetTransport(transportId int) (models.AdminTransportResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var transport models.AdminTransportResponse = models.AdminTransportResponse{Id: transportId}
	query := `SELECT owner_id, can_be_rented, transport_type, model, color, identifier, description, latitude, longitude, minute_price, day_price FROM transports WHERE id = $1`
	if err := r.db.QueryRowContext(ctx, query, transportId).Scan(&transport.OwnerId, &transport.CanBeRented, &transport.TransportType, &transport.Model, &transport.Color, &transport.Identifier, &transport.Description, &transport.Latitude, &transport.Longitude, &transport.MinutePrice, &transport.DayPrice); err != nil {
		return models.AdminTransportResponse{}, err
	}

	return transport, nil
}

func (r *PostgresRepository) AdminListTransports(from int, count int, transportType string) ([]models.AdminTransportResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var transports []models.AdminTransportResponse
	var query string
	var rows *sql.Rows
	var err error

	if transportType == "All" {
		query = `SELECT id, owner_id, can_be_rented, transport_type, model, color, identifier, description, latitude, longitude, minute_price, day_price FROM transports ORDER BY id OFFSET $1 LIMIT $2`
		rows, err = r.db.QueryContext(ctx, query, from, count)
	} else {
		query = `SELECT id, owner_id, can_be_rented, transport_type, model, color, identifier, description, latitude, longitude, minute_price, day_price FROM transports WHERE transport_type = $1 ORDER BY id OFFSET $2 LIMIT $3`
		rows, err = r.db.QueryContext(ctx, query, transportType, from, count)
	}

	if err != nil {
		return []models.AdminTransportResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var transport models.AdminTransportResponse
		if err := rows.Scan(&transport.Id, &transport.OwnerId, &transport.CanBeRented, &transport.TransportType, &transport.Model, &transport.Color, &transport.Identifier, &transport.Description, &transport.Latitude, &transport.Longitude, &transport.MinutePrice, &transport.DayPrice); err != nil {
			return []models.AdminTransportResponse{}, err
		}

		transports = append(transports, transport)
	}

	if err := rows.Err(); err != nil {
		return []models.AdminTransportResponse{}, err
	}

	return transports, nil
}

func (r *PostgresRepository) AdminUpdateTransport(transportId int, req models.AdminTransportRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	query := `UPDATE transports SET owner_id = $1, can_be_rented = $2, transport_type = $3, model = $4, color = $5, identifier = $6, description = $7, latitude = $8, longitude = $9, minute_price = $10, day_price = $11 WHERE id = $12`
	_, err := r.db.ExecContext(ctx, query, req.OwnerId, req.CanBeRented, req.TransportType, req.Model, req.Color, req.Identifier, req.Description, req.Latitude, req.Longitude, req.MinutePrice, req.DayPrice, transportId)

	return err
}

func (r *PostgresRepository) AdminDeleteTransport(transportId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	query := `DELETE FROM transports WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, transportId)

	return err
}
