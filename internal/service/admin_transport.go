package service

import (
	"github.com/ursulgwopp/simbir-go/internal/custom_errors"
	"github.com/ursulgwopp/simbir-go/internal/models"
)

func (s *Service) AdminCreateTransport(req models.AdminTransportRequest) (int, error) {
	exists, err := s.repo.CheckAccountIdExists(req.OwnerId)
	if err != nil {
		return -1, err
	}

	if !exists {
		return -1, custom_errors.ErrIdNotFound
	}

	if err := validateTransportType(req.TransportType); err != nil {
		return -1, err
	}

	if err := validateTransportProperties(req.Model, req.Color, req.Description, req.Identifier); err != nil {
		return -1, err
	}

	if req.Latitude < 0 || req.Longitude < 0 || req.MinutePrice < 0 || req.DayPrice < 0 {
		return -1, custom_errors.ErrInvalidParams
	}

	return s.repo.AdminCreateTransport(req)
}

func (s *Service) AdminDeleteTransport(transportId int) error {
	exists, err := s.repo.CheckTransportIdExists(transportId)
	if err != nil {
		return err
	}

	if !exists {
		return custom_errors.ErrIdNotFound
	}

	has, err := s.repo.CheckTransportIdHasActiveRents(transportId)
	if err != nil {
		return err
	}

	if has {
		return custom_errors.ErrCanNotDelete
	}

	return s.repo.AdminDeleteTransport(transportId)
}

func (s *Service) AdminGetTransport(transportId int) (models.AdminTransportResponse, error) {
	exists, err := s.repo.CheckTransportIdExists(transportId)
	if err != nil {
		return models.AdminTransportResponse{}, err
	}

	if !exists {
		return models.AdminTransportResponse{}, custom_errors.ErrIdNotFound
	}

	return s.repo.AdminGetTransport(transportId)
}
func (s *Service) AdminListTransports(from int, count int, transportType string) ([]models.AdminTransportResponse, error) {
	if from < 0 || count < 0 {
		return []models.AdminTransportResponse{}, custom_errors.ErrInvalidParams
	}

	if err := validateTransportType(transportType); err != nil {
		return []models.AdminTransportResponse{}, err
	}

	return s.repo.AdminListTransports(from, count, transportType)
}

func (s *Service) AdminUpdateTransport(transportId int, req models.AdminTransportRequest) error {
	exists, err := s.repo.CheckTransportIdExists(transportId)
	if err != nil {
		return err
	}

	if !exists {
		return custom_errors.ErrIdNotFound
	}

	if err := validateTransportType(req.TransportType); err != nil {
		return err
	}

	if err := validateTransportProperties(req.Model, req.Color, req.Description, req.Identifier); err != nil {
		return err
	}

	if req.Latitude < 0 || req.Longitude < 0 || req.MinutePrice < 0 || req.DayPrice < 0 {
		return custom_errors.ErrInvalidParams
	}

	return s.repo.AdminUpdateTransport(transportId, req)
}
