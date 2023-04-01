package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"time"
)

type Flight struct {
	Id               primitive.ObjectID `json:"id,omitempty"  bson:"_id,omitempty"`
	FlighDateAndTime time.Time          `json:"dateTime"`
	StartingPoint    string             `json:"startingPoint"`
	Destination      string             `json:"destination"`
	Price            int                `json:"price"`
	Seats            int                `json:"allSeats"`
	RemainingSeats   int                `json:"remainingSeats,omitempty"`
}

// // OVO POSLE IZBRISATI KAD SE SERVIS POPRAVI
// func (flight *Flight) Repackage() dtos.FlightDto {
// 	dto := dtos.FlightDto{
// 		Id:               flight.Id.Hex(),
// 		FlighDateAndTime: flight.FlighDateAndTime,
// 		StartingPoint:    flight.StartingPoint,
// 		Destination:      flight.Destination,
// 		Price:            flight.Price,
// 		Seats:            flight.Seats,
// 		RemainingSeats:   flight.RemainingSeats,
// 	}
// 	return dto
// }
