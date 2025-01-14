package service

import "github.com/ursulgwopp/simbir-go/internal/custom_errors"

func (s *Service) Hesoyam(accountId int, userId int, isAdmin bool) error {
	if err := validateAccountId(s, accountId); err != nil {
		return err
	}

	if err := validateAccountId(s, userId); err != nil {
		return err
	}

	if accountId != userId && !isAdmin {
		return custom_errors.ErrAccessDenied
	}

	return s.repo.Hesoyam(accountId)
}
