package repository

import "github.com/ursulgwopp/simbir-go/internal/models"

// AdminDeleteRent implements service.Repository.
func (*PostgresRepository) AdminDeleteRent(rentId int) error {
	panic("unimplemented")
}

// AdminGetRent implements service.Repository.
func (*PostgresRepository) AdminGetRent(rentId int) (models.RentResponse, error) {
	panic("unimplemented")
}

// AdminGetTransportHistory implements service.Repository.
func (*PostgresRepository) AdminGetTransportHistory(transportId int) ([]models.RentResponse, error) {
	panic("unimplemented")
}

// AdminGetUserHistory implements service.Repository.
func (*PostgresRepository) AdminGetUserHistory(userId int) ([]models.RentResponse, error) {
	panic("unimplemented")
}

// AdminStartRent implements service.Repository.
func (*PostgresRepository) AdminStartRent(req models.RentRequest) (int, error) {
	panic("unimplemented")
}

// AdminStopRent implements service.Repository.
func (*PostgresRepository) AdminStopRent(rentId int, latitude float64, longitude float64) error {
	panic("unimplemented")
}

// AdminUpdateRent implements service.Repository.
func (*PostgresRepository) AdminUpdateRent(rentId int, req models.RentRequest) error {
	panic("unimplemented")
}
