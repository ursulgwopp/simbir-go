package service

import (
	"github.com/ursulgwopp/simbir-go/internal/custom_errors"
	"github.com/ursulgwopp/simbir-go/internal/models"
)

func (s *Service) AdminListAccounts(from int, count int) ([]models.AdminAccountResponse, error) {
	if err := validatePagination(from, count); err != nil {
		return []models.AdminAccountResponse{}, custom_errors.ErrInvalidParams
	}

	return s.repo.AdminListAccounts(from, count)
}

func (s *Service) AdminGetAccount(accountId int) (models.AdminAccountResponse, error) {
	if err := validateAccountId(s.repo.CheckAccountIdExists(accountId)); err != nil {
		return models.AdminAccountResponse{}, err
	}

	return s.repo.AdminGetAccount(accountId)
}

func (s *Service) AdminCreateAccount(req models.AdminAccountRequest) (int, error) {
	if err := validateAdminAccountRequest(req); err != nil {
		return -1, err
	}

	if err := validateUsernameUniqueness(s.repo.CheckUsernameExists(req.Username)); err != nil {
		return -1, err
	}

	req.Password = generatePasswordHash(req.Password)

	return s.repo.AdminCreateAccount(req)
}

func (s *Service) AdminUpdateAccount(accountId int, req models.AdminAccountRequest) error {
	if err := validateAccountId(s.repo.CheckAccountIdExists(accountId)); err != nil {
		return err
	}

	if err := validateAdminAccountRequest(req); err != nil {
		return err
	}

	equal, err1 := s.repo.CheckUsernameIsEqualToOld(accountId, req.Username)
	exists, err2 := s.repo.CheckUsernameExists(req.Username)
	if err := validateUpdatedUsernameUniqueness(equal, err1, exists, err2); err != nil {
		return err
	}

	req.Password = generatePasswordHash(req.Password)

	return s.repo.AdminUpdateAccount(accountId, req)
}

func (s *Service) AdminDeleteAccount(accountId int) error {
	if err := validateAccountId(s.repo.CheckAccountIdExists(accountId)); err != nil {
		return err
	}

	if err := validateAccountDeletion(s.repo.CheckAccountIdHasActiveRents(accountId)); err != nil {
		return err
	}

	return s.repo.AdminDeleteAccount(accountId)
}
