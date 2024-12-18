package models

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AccountResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type TokenInfo struct {
	AccountId int  `json:"account_id"`
	IsAdmin   bool `json:"is_admin"`
}
