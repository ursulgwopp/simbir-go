package service

import "github.com/ursulgwopp/simbir-go/internal/models"

type Repository interface {
	SignUp(req models.AuthRequest) (int, error)
	SignIn(req models.AuthRequest) (models.TokenInfo, error)
	SignOut(token string) error

	GetAccount(accountId int) (models.AccountResponse, error)
	UpdateAccount(accountId int, req models.AccountResponse)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
