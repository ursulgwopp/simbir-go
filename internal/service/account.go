package service

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ursulgwopp/simbir-go/internal/custom_errors"
	"github.com/ursulgwopp/simbir-go/internal/models"
)

func (s *Service) SignUp(req models.AccountRequest) (int, error) {
	if err := validateUsername(req.Username); err != nil {
		return -1, err
	}

	if err := validatePassword(req.Password); err != nil {
		return -1, err
	}

	exists, err := s.repo.CheckUsernameExists(req.Username)
	if err != nil {
		return -1, err
	}

	if exists {
		return -1, custom_errors.ErrUsernameExists
	}

	req.Password = generatePasswordHash(req.Password)

	return s.repo.SignUp(req)
}

func (s *Service) SignIn(req models.AccountRequest) (string, error) {
	req.Password = generatePasswordHash(req.Password)

	tokenInfo, err := s.repo.SignIn(req)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return "", custom_errors.ErrInvalidUsernameOrPassword
		}

		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		TokenInfo: tokenInfo,
	})

	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func (s *Service) SignOut(token string) error {
	return s.repo.SignOut(token)
}

func (s *Service) GetAccount(accountId int) (models.AccountResponse, error) {
	exists, err := s.repo.CheckAccountIdExists(accountId)
	if err != nil {
		return models.AccountResponse{}, err
	}

	if !exists {
		return models.AccountResponse{}, custom_errors.ErrIdNotFound
	}

	return s.repo.GetAccount(accountId)
}

func (s *Service) UpdateAccount(accountId int, req models.AccountRequest) error {
	exists, err := s.repo.CheckAccountIdExists(accountId)
	if err != nil {
		return err
	}

	if !exists {
		return custom_errors.ErrIdNotFound
	}

	if err := validateUsername(req.Username); err != nil {
		return err
	}

	if err := validatePassword(req.Password); err != nil {
		return err
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

	return s.repo.UpdateAccount(accountId, req)
}
