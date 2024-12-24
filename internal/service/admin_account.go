package service

import (
	"github.com/ursulgwopp/simbir-go/internal/custom_errors"
	"github.com/ursulgwopp/simbir-go/internal/models"
)

func (s *Service) AdminCreateAccount(req models.AdminAccountRequest) (int, error) {
	// VALIDATE ACCOUNT INFO
	if err := validateAdminAccountRequest(req); err != nil {
		return -1, err
	}

	// CHECK IF USERNAME IS UNIQUE
	if err := validateUsernameUniqueness(s.repo.CheckUsernameExists(req.Username)); err != nil {
		return -1, err
	}

	// HASH PASSWORD
	req.Password = generatePasswordHash(req.Password)

	return s.repo.AdminCreateAccount(req)
}

func (s *Service) AdminListAccounts(from int, count int) ([]models.AdminAccountResponse, error) {
	// VALIDATE INPUT PARAMS
	if err := validatePagination(from, count); err != nil {
		return []models.AdminAccountResponse{}, custom_errors.ErrInvalidParams
	}

	return s.repo.AdminListAccounts(from, count)
}

func (s *Service) AdminGetAccount(accountId int) (models.AdminAccountResponse, error) {
	// CHECK IF ACCOUNT ID EXISTS
	if err := validateAccountId(s.repo.CheckAccountIdExists(accountId)); err != nil {
		return models.AdminAccountResponse{}, err
	}

	return s.repo.AdminGetAccount(accountId)
}

func (s *Service) AdminUpdateAccount(accountId int, req models.AdminAccountRequest) error {
	// CHECK IF ACCOUNT ID EXISTS
	if err := validateAccountId(s.repo.CheckAccountIdExists(accountId)); err != nil {
		return err
	}

	// VALIDATE ACCOUNT INFO
	if err := validateAdminAccountRequest(req); err != nil {
		return err
	}

	// CHECK IF USERNAME IS EQUAL TO OLD
	// IN CASE IT IS NOT - CHECK IF USERNAME IS UNIQUE
	equal, err1 := s.repo.CheckUsernameIsEqualToOld(accountId, req.Username)
	exists, err2 := s.repo.CheckUsernameExists(req.Username)
	if err := validateUpdatedUsername(equal, err1, exists, err2); err != nil {
		return err
	}

	// HASH PASSWORD
	req.Password = generatePasswordHash(req.Password)

	return s.repo.AdminUpdateAccount(accountId, req)
}

func (s *Service) AdminDeleteAccount(accountId int) error {
	// CHECK IF ACCOUNT ID EXISTS
	if err := validateAccountId(s.repo.CheckAccountIdExists(accountId)); err != nil {
		return err
	}

	// CHECK IF ACCOUNT ID HAS ANY ACTIVE RENTS
	// IF SO - CAN NOT DELETE ACCOUNT
	if err := validateAccountDeletion(s.repo.CheckAccountIdHasActiveRents(accountId)); err != nil {
		return err
	}

	return s.repo.AdminDeleteAccount(accountId)
}
