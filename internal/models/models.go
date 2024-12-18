package models

import "github.com/dgrijalva/jwt-go"

type AccountRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AccountResponse struct {
	Id       int     `json:"id"`
	Username string  `json:"username"`
	Balance  float64 `json:"balance"`
}

type TokenInfo struct {
	AccountId int  `json:"account_id"`
	IsAdmin   bool `json:"is_admin"`
}

type TokenClaims struct {
	jwt.StandardClaims
	TokenInfo
}
