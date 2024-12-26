package service

import (
	"github.com/ursulgwopp/simbir-go/internal/models"
)

func (s *Service) AdminCreateTransport(req models.AdminTransportRequest) (int, error) {
	if err := validateAccountId(s, req.OwnerId); err != nil {
		return -1, err
	}

	if err := validateTransportType(req.TransportType); err != nil {
		return -1, err
	}

	if err := validateAdminTransportRequest(req); err != nil {
		return -1, err
	}

	return s.repo.AdminCreateTransport(req)
}

func (s *Service) AdminDeleteTransport(transportId int) error {
	if err := validateTransportId(s, transportId); err != nil {
		return err
	}

	if err := validateTransportDeletion(s, transportId); err != nil {
		return err
	}

	return s.repo.AdminDeleteTransport(transportId)
}

func (s *Service) AdminGetTransport(transportId int) (models.AdminTransportResponse, error) {
	if err := validateTransportId(s, transportId); err != nil {
		return models.AdminTransportResponse{}, err
	}

	return s.repo.AdminGetTransport(transportId)
}

func (s *Service) AdminListTransports(from int, count int, transportType string) ([]models.AdminTransportResponse, error) {
	if err := validatePagination(from, count); err != nil {
		return []models.AdminTransportResponse{}, err
	}

	if err := validateTransportType(transportType); err != nil {
		return []models.AdminTransportResponse{}, err
	}

	return s.repo.AdminListTransports(from, count, transportType)
}

func (s *Service) AdminUpdateTransport(transportId int, req models.AdminTransportRequest) error {
	if err := validateTransportId(s, transportId); err != nil {
		return err
	}

	if err := validateTransportType(req.TransportType); err != nil {
		return err
	}

	if err := validateAdminTransportRequest(req); err != nil {
		return err
	}

	return s.repo.AdminUpdateTransport(transportId, req)
}
