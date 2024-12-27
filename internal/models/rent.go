package models

import (
	"time"
)

type RentResponse struct {
	Id          int       `json:"id"`
	TransportId int       `json:"transportId"`
	UserId      int       `json:"userId"`
	TimeStart   time.Time `json:"timeStart"`
	TimeEnd     time.Time `json:"timeEnd,omitempty"`
	PriceOfUnit float64   `json:"priceOfUnit"`
	PriceType   string    `json:"priceType"`
	FinalPrice  float64   `json:"finalPrice,omitempty"`
	IsActive    bool      `json:"is_active"`
}

type RentRequest struct {
	TransportId int       `json:"transport_id"`
	UserId      int       `json:"user_id"`
	TimeStart   time.Time `json:"time_start"`
	TimeEnd     time.Time `json:"time_end,omitempty"`
	PriceOfUnit float64   `json:"priceOfUnit"`
	PriceType   string    `json:"priceType"`
	FinalPrice  float64   `json:"finalPrice,omitempty"`
}
