package service

import (
	"github.com/ursulgwopp/simbir-go/internal/models"
)

type Repository interface {
	SignUp(req models.AccountRequest) (int, error)
	SignIn(req models.AccountRequest) (models.TokenInfo, error)
	SignOut(token string) error

	GetAccount(accountId int) (models.AccountResponse, error)
	UpdateAccount(accountId int, req models.AccountRequest) error

	CheckUsernameExists(username string) (bool, error)
	CheckUsernameIsEqualToOld(accountId int, username string) (bool, error)
	CheckTokenIsValid(token string) (bool, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
