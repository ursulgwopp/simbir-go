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
	if err := validateAccountId(s.repo.CheckAccountIdExists(userId)); err != nil {
		return models.RentResponse{}, err
	}

	if err := validateRentId(s.repo.CheckRentIdExists(rentId)); err != nil {
		return models.RentResponse{}, err
	}

	rent, err := s.repo.GetRent(rentId)
	if err != nil {
		return models.RentResponse{}, err
	}

	ownerId, err := s.repo.CheckOwnerId(rent.TransportId)
	if err := validateRentAccess(userId, ownerId, rent.UserId, err); err != nil {
		return models.RentResponse{}, err
	}

	return rent, nil
}

func (s *Service) GetTransportHistory(userId int, transportId int) ([]models.RentResponse, error) {
	if err := validateAccountId(s.repo.CheckAccountIdExists(userId)); err != nil {
		return []models.RentResponse{}, err
	}

	if err := validateTransportId(s.repo.CheckTransportIdExists(transportId)); err != nil {
		return []models.RentResponse{}, err
	}

	ownerId, err := s.repo.CheckOwnerId(transportId)
	if err := validateRentOwner(userId, ownerId, err); err != nil {
		return []models.RentResponse{}, err
	}

	return s.repo.GetTransportHistory(transportId)
}

func (s *Service) GetUserHistory(accountId int) ([]models.RentResponse, error) {
	if err := validateAccountId(s.repo.CheckAccountIdExists(accountId)); err != nil {
		return []models.RentResponse{}, err
	}

	return s.repo.GetUserHistory(accountId)
}

func (s *Service) StartRent(userId int, transportId int, rentType string) (int, error) {
	if err := validateAccountId(s.repo.CheckAccountIdExists(userId)); err != nil {
		return -1, err
	}

	ownerId, err := s.repo.CheckOwnerId(transportId)
	if err := validateTransportOwner1(userId, ownerId, err); err != nil {
		return -1, err
	}

	if err := validateTransportId(s.repo.CheckTransportIdExists(transportId)); err != nil {
		return -1, err
	}

	if err := validateTransportIsAvailable(s.repo.CheckTransportIsAvailable(transportId)); err != nil {
		return -1, err
	}

	if err := validateRentType(rentType); err != nil {
		return -1, err
	}

	return s.repo.StartRent(userId, transportId, rentType)
}

func (s *Service) StopRent(userId int, rentId int, latitude float64, longitude float64) error {
	if err := validateAccountId(s.repo.CheckAccountIdExists(userId)); err != nil {
		return err
	}

	ownerId, err := s.repo.CheckRentOwnerId(rentId)
	if err := validateRentOwner(userId, ownerId, err); err != nil {
		return err
	}

	if latitude < 0 || longitude < 0 {
		return custom_errors.ErrInvalidParams
	}

	if err := validateRentId(s.repo.CheckRentIdExists(rentId)); err != nil {
		return err
	}

	if err := validateRentIsActive(s.repo.CheckRentIsActive(rentId)); err != nil {
		return err
	}

	return s.repo.StopRent(rentId, latitude, longitude)
}
