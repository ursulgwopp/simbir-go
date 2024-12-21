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

type AdminAccountRequest struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
	IsAdmin  bool    `json:"isAdmin"`
}

type AdminAccountResponse struct {
	Id       int     `json:"id"`
	Username string  `json:"username"`
	Balance  float64 `json:"balance"`
	IsAdmin  bool    `json:"isAdmin"`
}

type TransportRequest struct {
	CanBeRented   bool    `json:"canBeRented"`
	TransportType string  `json:"transportType"`
	Model         string  `json:"model"`
	Color         string  `json:"color"`
	Identifier    string  `json:"identifier"`
	Description   string  `json:"description"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	MinutePrice   float64 `json:"minutePrice"`
	DayPrice      float64 `json:"dayPrice"`
}

type TransportResponse struct {
	Id            int     `json:"id"`
	CanBeRented   bool    `json:"canBeRented"`
	TransportType string  `json:"transportType"`
	Model         string  `json:"model"`
	Color         string  `json:"color"`
	Identifier    string  `json:"identifier"`
	Description   string  `json:"description,omitempty"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	MinutePrice   float64 `json:"minutePrice,omitempty"`
	DayPrice      float64 `json:"dayPrice,omitempty"`
}

type AdminTransportRequest struct {
	OwnerId       int     `json:"ownerId"`
	CanBeRented   bool    `json:"canBeRented"`
	TransportType string  `json:"transportType"`
	Model         string  `json:"model"`
	Color         string  `json:"color"`
	Identifier    string  `json:"identifier"`
	Description   string  `json:"description"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	MinutePrice   float64 `json:"minutePrice"`
	DayPrice      float64 `json:"dayPrice"`
}

type AdminTransportResponse struct {
	Id            int     `json:"id"`
	OwnerId       int     `json:"ownerId"`
	CanBeRented   bool    `json:"canBeRented"`
	TransportType string  `json:"transportType"`
	Model         string  `json:"model"`
	Color         string  `json:"color"`
	Identifier    string  `json:"identifier"`
	Description   string  `json:"description,omitempty"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	MinutePrice   float64 `json:"minutePrice,omitempty"`
	DayPrice      float64 `json:"dayPrice,omitempty"`
}
