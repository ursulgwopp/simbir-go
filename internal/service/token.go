package service

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/ursulgwopp/simbir-go/internal/models"
)

func (s *Service) CheckTokenIsValid(token string) (bool, error) {
	return s.repo.CheckTokenIsValid(token)
}

func (s *Service) ParseToken(token string) (models.TokenInfo, error) {
	token_, err := jwt.ParseWithClaims(token, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return models.TokenInfo{}, err
	}

	claims, ok := token_.Claims.(*models.TokenClaims)
	if !ok {
		return models.TokenInfo{}, errors.New("token claims are not of type tokenClaims")
	}

	return models.TokenInfo{AccountId: claims.AccountId, IsAdmin: claims.IsAdmin}, nil
}
