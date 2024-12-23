package service

import (
	"github.com/ursulgwopp/simbir-go/internal/custom_errors"
	"github.com/ursulgwopp/simbir-go/internal/models"
)

func (s *Service) GetAvailableTransport(latitude float64, longitude float64, radius float64, transportType string) ([]models.TransportResponse, error) {
	if latitude < 0 || longitude < 0 || radius < 0 {
		return []models.TransportResponse{}, custom_errors.ErrInvalidParams
	}

	if transportType != "All" && transportType != "Car" && transportType != "Bike" && transportType != "Scooter" {
		return []models.TransportResponse{}, custom_errors.ErrInvalidParams
	}

	return s.repo.GetAvailableTransport(latitude, longitude, radius, transportType)
}

func (s *Service) GetRent(userId int, rentId int) (models.RentResponse, error) {
	exists, err := s.repo.CheckRentIdExists(rentId)
	if err != nil {
		return models.RentResponse{}, err
	}

	if !exists {
		return models.RentResponse{}, custom_errors.ErrIdNotFound
	}

	rent, err := s.repo.GetRent(rentId)
	if err != nil {
		return models.RentResponse{}, err
	}

	ownerId, err := s.repo.CheckOwnerId(rent.TransportId)
	if err != nil {
		return models.RentResponse{}, err
	}

	if userId != ownerId && userId != rent.UserId {
		return models.RentResponse{}, custom_errors.ErrAccessDenied
	}

	return rent, nil
}

func (s *Service) GetTransportHistory(userId int, transportId int) ([]models.RentResponse, error) {
	exists, err := s.repo.CheckTransportIdExists(transportId)
	if err != nil {
		return []models.RentResponse{}, err
	}

	if !exists {
		return []models.RentResponse{}, custom_errors.ErrIdNotFound
	}

	ownerId, err := s.repo.CheckOwnerId(transportId)
	if err != nil {
		return []models.RentResponse{}, err
	}

	if userId != ownerId {
		return []models.RentResponse{}, custom_errors.ErrAccessDenied
	}

	return s.repo.GetTransportHistory(transportId)
}

func (s *Service) GetUserHistory(accountId int) ([]models.RentResponse, error) {
	return s.repo.GetUserHistory(accountId)
}

func (s *Service) StartRent(userId int, transportId int, rentType string) (int, error) {
	ownerId, err := s.repo.CheckOwnerId(transportId)
	if err != nil {
		return -1, err
	}

	if ownerId == userId {
		return -1, custom_errors.ErrCanNotRent
	}

	exists, err := s.repo.CheckTransportIdExists(transportId)
	if err != nil {
		return -1, err
	}

	if !exists {
		return -1, custom_errors.ErrIdNotFound
	}

	is_available, err := s.repo.CheckTransportIsAvailable(transportId)
	if err != nil {
		return -1, err
	}

	if !is_available {
		return -1, custom_errors.ErrTransportNotAvailable
	}

	if rentType != "Minutes" && rentType != "Days" {
		return -1, custom_errors.ErrInvalidParams
	}

	return s.repo.StartRent(userId, transportId, rentType)
}

func (s *Service) StopRent(userId int, rentId int, latitude float64, longitude float64) error {
	ownerId, err := s.repo.CheckRentOwnerId(rentId)
	if err != nil {
		return err
	}

	if ownerId != userId {
		return custom_errors.ErrAccessDenied
	}

	if latitude < 0 || longitude < 0 {
		return custom_errors.ErrInvalidParams
	}

	exists, err := s.repo.CheckRentIdExists(rentId)
	if err != nil {
		return err
	}

	if !exists {
		return custom_errors.ErrIdNotFound
	}

	is_active, err := s.repo.CheckRentIsActive(rentId)
	if err != nil {
		return err
	}

	if !is_active {
		return custom_errors.ErrAlreadyStopped
	}

	return s.repo.StopRent(rentId, latitude, longitude)
}
