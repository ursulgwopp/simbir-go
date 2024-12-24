package service

import "github.com/ursulgwopp/simbir-go/internal/custom_errors"

func (s *Service) Hesoyam(accountId int, userId int, isAdmin bool) error {
	exists, err := s.repo.CheckAccountIdExists(accountId)
	if err != nil {
		return err
	}
	if !exists {
		return custom_errors.ErrIdNotFound
	}

	exists, err = s.repo.CheckAccountIdExists(userId)
	if err != nil {
		return err
	}
	if !exists {
		return custom_errors.ErrIdNotFound
	}

	if accountId != userId && !isAdmin {
		return custom_errors.ErrAccessDenied
	}

	return s.repo.Hesoyam(accountId)
}
