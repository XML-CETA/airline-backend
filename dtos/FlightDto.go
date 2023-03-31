package dtos

import (
	"time"
)

type FlightDto struct {
	Id               string    `json:"id"`
	FlighDateAndTime time.Time `json:"dateTime"`
	StartingPoint    string    `json:"startingPoint"`
	Destination      string    `json:"destination"`
	Price            int       `json:"price"`
	Seats            int       `json:"allSeats"`
	RemainingSeats   int       `json:"remainingSeats,omitempty"`
}
