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

	AdminListAccounts(from int, count int) ([]models.AdminAccountResponse, error)
	AdminGetAccount(accountId int) (models.AdminAccountResponse, error)
	AdminCreateAccount(req models.AdminAccountRequest) (int, error)
	AdminUpdateAccount(accountId int, req models.AdminAccountRequest) error
	AdminDeleteAccount(accountId int) error

	Hesoyam(accountId int) error

	CreateTransport(ownerId int, req models.TransportRequest) (int, error)
	GetTransport(transportId int) (models.TransportResponse, error)
	UpdateTransport(transportId int, req models.TransportRequest) error
	DeleteTransport(transportId int) error

	GetAvailableTransport(latitude float64, longitude float64, radius float64, transportType string) ([]models.TransportResponse, error)
	GetRent(rentId int) (models.RentResponse, error)
	GetUserHistory(accountId int) ([]models.RentResponse, error)
	GetTransportHistory(transportId int) ([]models.RentResponse, error)
	StartRent(transportId int, rentType string) (int, error)
	StopRent(rentId int, latitude float64, longitude float64) error

	AdminListTransports(from int, count int, transportType string) ([]models.AdminTransportResponse, error)
	AdminGetTransport(transportId int) (models.AdminTransportResponse, error)
	AdminCreateTransport(req models.AdminTransportRequest) (int, error)
	AdminUpdateTransport(transportId int, req models.AdminTransportRequest) error
	AdminDeleteTransport(transportId int) error

	CheckUsernameExists(username string) (bool, error)
	CheckUsernameIsEqualToOld(accountId int, username string) (bool, error)
	CheckTokenIsValid(token string) (bool, error)
	CheckAccountIdExists(accountId int) (bool, error)
	CheckOwnerId(transportId int) (int, error)
	CheckTransportIdExists(transportId int) (bool, error)
	CheckRentIdExists(rentId int) (bool, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
