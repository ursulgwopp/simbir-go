package service

import "github.com/ursulgwopp/simbir-go/internal/models"

// AdminDeleteRent implements transport.Service.
func (*Service) AdminDeleteRent(rentId int) error {
	panic("unimplemented")
}

// AdminGetRent implements transport.Service.
func (*Service) AdminGetRent(rentId int) (models.RentResponse, error) {
	panic("unimplemented")
}

// AdminGetTransportHistory implements transport.Service.
func (*Service) AdminGetTransportHistory(transportId int) ([]models.RentResponse, error) {
	panic("unimplemented")
}

// AdminGetUserHistory implements transport.Service.
func (*Service) AdminGetUserHistory(userId int) ([]models.RentResponse, error) {
	panic("unimplemented")
}

// AdminStartRent implements transport.Service.
func (*Service) AdminStartRent(req models.RentRequest) (int, error) {
	panic("unimplemented")
}

// AdminStopRent implements transport.Service.
func (*Service) AdminStopRent(rentId int, latitude float64, longitude float64) error {
	panic("unimplemented")
}

// AdminUpdateRent implements transport.Service.
func (*Service) AdminUpdateRent(rentId int, req models.RentRequest) error {
	panic("unimplemented")
}
