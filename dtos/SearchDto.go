package dtos

import "time"

type SearchDto struct {
	FlighDateAndTime time.Time `json:"dateTime"`
	StartingPoint    string    `json:"startingPoint"`
	Destination      string    `json:"destination"`
	NeededSeats      int       `json:"neededSeats"`
}
