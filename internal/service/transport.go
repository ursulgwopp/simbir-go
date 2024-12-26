package service

import (
	"github.com/ursulgwopp/simbir-go/internal/models"
)

func (s *Service) GetTransport(transportId int) (models.TransportResponse, error) {
	if err := validateTransportId(s, transportId); err != nil {
		return models.TransportResponse{}, err
	}

	return s.repo.GetTransport(transportId)
}

func (s *Service) CreateTransport(ownerId int, req models.TransportRequest) (int, error) {
	if err := validateAccountId(s, ownerId); err != nil {
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
	if err := validateAccountId(s, userId); err != nil {
		return err
	}

	if err := validateTransportId(s, transportId); err != nil {
		return err
	}

	if err := validateTransportOwner(s, userId, transportId); err != nil {
		return err
	}

	if err := validateTransportRequest(req); err != nil {
		return err
	}

	return s.repo.UpdateTransport(transportId, req)
}

func (s *Service) DeleteTransport(userId int, transportId int) error {
	if err := validateAccountId(s, userId); err != nil {
		return err
	}

	if err := validateTransportId(s, transportId); err != nil {
		return err
	}

	if err := validateTransportOwner(s, userId, transportId); err != nil {
		return err
	}

	if err := validateTransportDeletion(s, transportId); err != nil {
		return err
	}

	return s.repo.DeleteTransport(transportId)
}
