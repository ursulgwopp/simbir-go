package service

import (
	"github.com/ursulgwopp/simbir-go/internal/custom_errors"
	"github.com/ursulgwopp/simbir-go/internal/models"
)

func (s *Service) CreateTransport(ownerId int, req models.TransportRequest) (int, error) {
	return s.repo.CreateTransport(ownerId, req)
}

func (s *Service) GetTransport(transportId int) (models.TransportResponse, error) {
	exists, err := s.repo.CheckTransportIdExists(transportId)
	if err != nil {
		return models.TransportResponse{}, err
	}

	if !exists {
		return models.TransportResponse{}, custom_errors.ErrIdNotFound
	}

	return s.repo.GetTransport(transportId)
}

func (s *Service) UpdateTransport(userId int, transportId int, req models.TransportRequest) error {
	exists, err := s.repo.CheckTransportIdExists(transportId)
	if err != nil {
		return err
	}

	if !exists {
		return custom_errors.ErrIdNotFound
	}

	ownerId, err := s.repo.CheckOwnerId(transportId)
	if err != nil {
		return err
	}

	if userId != ownerId {
		return custom_errors.ErrAccessDenied
	}

	return s.repo.UpdateTransport(transportId, req)
}

func (s *Service) DeleteTransport(userId int, transportId int) error {
	exists, err := s.repo.CheckTransportIdExists(transportId)
	if err != nil {
		return err
	}

	if !exists {
		return custom_errors.ErrIdNotFound
	}

	ownerId, err := s.repo.CheckOwnerId(transportId)
	if err != nil {
		return err
	}

	if userId != ownerId {
		return custom_errors.ErrAccessDenied
	}

	return s.repo.DeleteTransport(transportId)
}
