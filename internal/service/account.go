package service

import "github.com/ursulgwopp/simbir-go/internal/models"

// GetAccount implements transport.Service.
func (*Service) GetAccount(accountId int) (models.AccountResponse, error) {
	panic("unimplemented")
}

// SignIn implements transport.Service.
func (*Service) SignIn(req models.AuthRequest) (string, error) {
	panic("unimplemented")
}

// SignOut implements transport.Service.
func (*Service) SignOut(token string) error {
	panic("unimplemented")
}

// SignUp implements transport.Service.
func (*Service) SignUp(req models.AuthRequest) (int, error) {
	panic("unimplemented")
}

// UpdateAccount implements transport.Service.
func (*Service) UpdateAccount(accountId int, req models.AccountResponse) {
	panic("unimplemented")
}
