package service

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ursulgwopp/simbir-go/internal/custom_errors"
	"github.com/ursulgwopp/simbir-go/internal/models"
)

func (s *Service) SignUp(req models.AccountRequest) (int, error) {
	if err := validateAccountRequest(req); err != nil {
		return -1, err
	}

	if err := validateUsernameUniqueness(s.repo.CheckUsernameExists(req.Username)); err != nil {
		return -1, err
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
	if err := validateAccountId(s.repo.CheckAccountIdExists(accountId)); err != nil {
		return models.AccountResponse{}, err
	}

	return s.repo.GetAccount(accountId)
}

func (s *Service) UpdateAccount(accountId int, req models.AccountRequest) error {
	if err := validateAccountId(s.repo.CheckAccountIdExists(accountId)); err != nil {
		return err
	}

	if err := validateAccountRequest(req); err != nil {
		return err
	}

	equal, err1 := s.repo.CheckUsernameIsEqualToOld(accountId, req.Username)
	exists, err2 := s.repo.CheckUsernameExists(req.Username)
	if err := validateUpdatedUsernameUniqueness(equal, err1, exists, err2); err != nil {
		return err
	}

	req.Password = generatePasswordHash(req.Password)

	return s.repo.UpdateAccount(accountId, req)
}
