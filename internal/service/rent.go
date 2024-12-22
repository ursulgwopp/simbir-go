package service

import (
	"github.com/ursulgwopp/simbir-go/internal/custom_errors"
	"github.com/ursulgwopp/simbir-go/internal/models"
)

func (s *Service) GetAvailableTransport(latitude float64, longitude float64, radius float64, transportType string) ([]models.TransportResponse, error) {
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
