package service

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ursulgwopp/simbir-go/internal/custom_errors"
	"github.com/ursulgwopp/simbir-go/internal/models"
)

func (s *Service) SignUp(req models.AccountRequest) (int, error) {
	// VALIDATE ACCOUNT INFO
	if err := validateAccountRequest(req); err != nil {
		return -1, err
	}

	// CHECK IF USERNAME IS UNIQUE
	if err := validateUsernameUniqueness(s.repo.CheckUsernameExists(req.Username)); err != nil {
		return -1, err
	}

	// HASH PASSWORD
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
	// CHECK IF ACCOUNT ID EXISTS
	if err := validateAccountId(s.repo.CheckAccountIdExists(accountId)); err != nil {
		return models.AccountResponse{}, err
	}

	return s.repo.GetAccount(accountId)
}

func (s *Service) UpdateAccount(accountId int, req models.AccountRequest) error {
	// CHECK IF ACCOUNT ID EXISTS
	if err := validateAccountId(s.repo.CheckAccountIdExists(accountId)); err != nil {
		return err
	}

	// VALIDATE ACCOUNT INFO
	if err := validateAccountRequest(req); err != nil {
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

	return s.repo.UpdateAccount(accountId, req)
}
