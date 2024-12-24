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
