package dtos

import (
	"time"
)

type SearchedFlightDto struct {
	Id               string    `json:"id"`
	FlighDateAndTime time.Time `json:"dateTime"`
	StartingPoint    string    `json:"startingPoint"`
	Destination      string    `json:"destination"`
	Price            int       `json:"price"`
	Seats            int       `json:"seats"`
	TotalPrice       int       `json:"totalPrice"`
}
