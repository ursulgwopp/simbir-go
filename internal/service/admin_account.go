package service

import (
	"github.com/ursulgwopp/simbir-go/internal/custom_errors"
	"github.com/ursulgwopp/simbir-go/internal/models"
)

func (s *Service) AdminCreateAccount(req models.AdminAccountRequest) (int, error) {
	if err := validateUsername(req.Username); err != nil {
		return -1, custom_errors.ErrInvalidParams
	}

	if err := validatePassword(req.Password); err != nil {
		return -1, custom_errors.ErrInvalidParams
	}

	if req.Balance < 0 {
		return -1, custom_errors.ErrInvalidParams
	}

	exists, err := s.repo.CheckUsernameExists(req.Username)
	if err != nil {
		return -1, err
	}

	if exists {
		return -1, custom_errors.ErrUsernameExists
	}

	req.Password = generatePasswordHash(req.Password)

	return s.repo.AdminCreateAccount(req)
}

func (s *Service) AdminListAccounts(from int, count int) ([]models.AdminAccountResponse, error) {
	if from < 0 || count < 0 {
		return []models.AdminAccountResponse{}, custom_errors.ErrInvalidParams
	}

	return s.repo.AdminListAccounts(from, count)
}

func (s *Service) AdminGetAccount(accountId int) (models.AdminAccountResponse, error) {
	exists, err := s.repo.CheckAccountIdExists(accountId)
	if err != nil {
		return models.AdminAccountResponse{}, err
	}

	if !exists {
		return models.AdminAccountResponse{}, custom_errors.ErrIdNotFound
	}

	return s.repo.AdminGetAccount(accountId)
}

func (s *Service) AdminUpdateAccount(accountId int, req models.AdminAccountRequest) error {
	exists, err := s.repo.CheckAccountIdExists(accountId)
	if err != nil {
		return err
	}

	if !exists {
		return custom_errors.ErrIdNotFound
	}

	if err := validateUsername(req.Username); err != nil {
		return custom_errors.ErrInvalidParams
	}

	if err := validatePassword(req.Password); err != nil {
		return custom_errors.ErrInvalidParams
	}

	if req.Balance < 0 {
		return custom_errors.ErrInvalidParams
	}

	equal, err := s.repo.CheckUsernameIsEqualToOld(accountId, req.Username)
	if err != nil {
		return err
	}

	if !equal {
		exists, err := s.repo.CheckUsernameExists(req.Username)
		if err != nil {
			return err
		}

		if exists {
			return custom_errors.ErrUsernameExists
		}
	}

	req.Password = generatePasswordHash(req.Password)

	return s.repo.AdminUpdateAccount(accountId, req)
}

func (s *Service) AdminDeleteAccount(accountId int) error {
	has, err := s.repo.CheckAccountIdHasActiveRents(accountId)
	if err != nil {
		return err
	}

	if has {
		return custom_errors.ErrCanNotDelete
	}

	exists, err := s.repo.CheckAccountIdExists(accountId)
	if err != nil {
		return err
	}

	if !exists {
		return custom_errors.ErrIdNotFound
	}

	return s.repo.AdminDeleteAccount(accountId)
}
