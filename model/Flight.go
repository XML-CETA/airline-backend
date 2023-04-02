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
