package service

import (
	"github.com/ursulgwopp/simbir-go/internal/custom_errors"
	"github.com/ursulgwopp/simbir-go/internal/models"
)

func (s *Service) GetAvailableTransport(latitude float64, longitude float64, radius float64, transportType string) ([]models.TransportResponse, error) {
	if latitude < 0 || longitude < 0 || radius < 0 {
		return []models.TransportResponse{}, custom_errors.ErrInvalidParams
	}

	if err := validateTransportType(transportType); err != nil {
		return []models.TransportResponse{}, custom_errors.ErrInvalidParams
	}

	return s.repo.GetAvailableTransport(latitude, longitude, radius, transportType)
}

func (s *Service) GetRent(userId int, rentId int) (models.RentResponse, error) {
	if err := validateAccountId(s, userId); err != nil {
		return models.RentResponse{}, err
	}

	if err := validateRentId(s, rentId); err != nil {
		return models.RentResponse{}, err
	}

	rent, err := s.repo.GetRent(rentId)
	if err != nil {
		return models.RentResponse{}, err
	}

	if err := validateRentAccess(s, userId, rent.TransportId, rent.UserId); err != nil {
		return models.RentResponse{}, err
	}

	return rent, nil
}

func (s *Service) GetTransportHistory(userId int, transportId int) ([]models.RentResponse, error) {
	if err := validateAccountId(s, userId); err != nil {
		return []models.RentResponse{}, err
	}

	if err := validateTransportId(s, transportId); err != nil {
		return []models.RentResponse{}, err
	}

	if err := validateTransportOwner(s, userId, transportId); err != nil {
		return []models.RentResponse{}, err
	}

	return s.repo.GetTransportHistory(transportId)
}

func (s *Service) GetUserHistory(accountId int) ([]models.RentResponse, error) {
	if err := validateAccountId(s, accountId); err != nil {
		return []models.RentResponse{}, err
	}

	return s.repo.GetUserHistory(accountId)
}

func (s *Service) StartRent(userId int, transportId int, rentType string) (int, error) {
	if err := validateAccountId(s, userId); err != nil {
		return -1, err
	}

	if err := validateTransportOwner(s, userId, transportId); err == nil {
		return -1, custom_errors.ErrCanNotRent
	}

	if err := validateTransportId(s, transportId); err != nil {
		return -1, err
	}

	if err := validateTransportIsAvailable(s, transportId); err != nil {
		return -1, err
	}

	if err := validateRentType(rentType); err != nil {
		return -1, err
	}

	return s.repo.StartRent(userId, transportId, rentType)
}

func (s *Service) StopRent(userId int, rentId int, latitude float64, longitude float64) error {
	if err := validateAccountId(s, userId); err != nil {
		return err
	}

	if err := validateRentOwner(s, userId, rentId); err != nil {
		return err
	}

	if latitude < 0 || longitude < 0 {
		return custom_errors.ErrInvalidParams
	}

	if err := validateRentId(s, rentId); err != nil {
		return err
	}

	if err := validateRentIsActive(s, rentId); err != nil {
		return err
	}

	return s.repo.StopRent(rentId, latitude, longitude)
}
