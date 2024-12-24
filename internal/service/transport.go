package service

import (
	"github.com/ursulgwopp/simbir-go/internal/models"
)

func (s *Service) GetTransport(transportId int) (models.TransportResponse, error) {
	if err := validateTransportId(s.repo.CheckTransportIdExists(transportId)); err != nil {
		return models.TransportResponse{}, err
	}

	return s.repo.GetTransport(transportId)
}

func (s *Service) CreateTransport(ownerId int, req models.TransportRequest) (int, error) {
	if err := validateAccountId(s.repo.CheckAccountIdExists(ownerId)); err != nil {
		return -1, err
	}

	if err := validateTransportType(req.TransportType); err != nil {
		return -1, err
	}

	if err := validateTransportRequest(req); err != nil {
		return -1, err
	}

	return s.repo.CreateTransport(ownerId, req)
}

func (s *Service) UpdateTransport(userId int, transportId int, req models.TransportRequest) error {
	if err := validateAccountId(s.repo.CheckAccountIdExists(userId)); err != nil {
		return err
	}

	if err := validateTransportId(s.repo.CheckTransportIdExists(transportId)); err != nil {
		return err
	}

	ownerId, err := s.repo.CheckOwnerId(transportId)
	if err := validateTransportOwner(userId, ownerId, err); err != nil {
		return err
	}

	if err := validateTransportRequest(req); err != nil {
		return err
	}

	return s.repo.UpdateTransport(transportId, req)
}

func (s *Service) DeleteTransport(userId int, transportId int) error {
	if err := validateAccountId(s.repo.CheckAccountIdExists(userId)); err != nil {
		return err
	}

	if err := validateTransportId(s.repo.CheckTransportIdExists(transportId)); err != nil {
		return err
	}

	ownerId, err := s.repo.CheckOwnerId(transportId)
	if err := validateTransportOwner(userId, ownerId, err); err != nil {
		return err
	}

	if err := validateTransportDeletion(s.repo.CheckTransportIdHasActiveRents(transportId)); err != nil {
		return err
	}

	return s.repo.DeleteTransport(transportId)
}
