package service

import "github.com/ursulgwopp/simbir-go/internal/custom_errors"

func (s *Service) Hesoyam(accountId int, userId int, isAdmin bool) error {
	if accountId != userId && !isAdmin {
		return custom_errors.ErrAccessDenied
	}

	return s.repo.Hesoyam(accountId)
}
