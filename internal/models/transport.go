package models

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
