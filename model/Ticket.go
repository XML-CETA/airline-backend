package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ticket struct {
	Id       primitive.ObjectID `json:"id,omitempty"  bson:"_id,omitempty"`
	User     string             `json:"username"`
	FlightId primitive.ObjectID `json:"flightId"`
	Amount   int                `json:"amount"`
}
