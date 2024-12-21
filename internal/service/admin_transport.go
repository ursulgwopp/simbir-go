package service

import (
	"github.com/ursulgwopp/simbir-go/internal/custom_errors"
	"github.com/ursulgwopp/simbir-go/internal/models"
)

func (s *Service) AdminCreateTransport(req models.AdminTransportRequest) (int, error) {
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

	return s.repo.AdminUpdateTransport(transportId, req)
}
