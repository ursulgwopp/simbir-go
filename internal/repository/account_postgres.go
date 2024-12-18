package repository

import "github.com/ursulgwopp/simbir-go/internal/models"

// GetAccount implements service.Repository.
func (*PostgresRepository) GetAccount(accountId int) (models.AccountResponse, error) {
	panic("unimplemented")
}

// SignIn implements service.Repository.
func (*PostgresRepository) SignIn(req models.AuthRequest) (models.TokenInfo, error) {
	panic("unimplemented")
}

// SignOut implements service.Repository.
func (*PostgresRepository) SignOut(token string) error {
	panic("unimplemented")
}

// SignUp implements service.Repository.
func (*PostgresRepository) SignUp(req models.AuthRequest) (int, error) {
	panic("unimplemented")
}

// UpdateAccount implements service.Repository.
func (*PostgresRepository) UpdateAccount(accountId int, req models.AccountResponse) {
	panic("unimplemented")
}
