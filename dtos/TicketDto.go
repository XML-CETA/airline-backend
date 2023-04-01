package dtos

import (
	"main/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TicketDto struct {
	Id               primitive.ObjectID `json:"id,omitempty"`
	Amount           int                `json:"amount"`
	FlighDateAndTime time.Time          `json:"dateTime"`
	StartingPoint    string             `json:"startingPoint"`
	Destination      string             `json:"destination"`
	Price            int                `json:"price"`
}

func Construct(ticket model.Ticket, flight model.Flight) TicketDto {
	dto := TicketDto{
		Id:               ticket.Id,
		Amount:           ticket.Amount,
		FlighDateAndTime: flight.FlighDateAndTime,
		StartingPoint:    flight.StartingPoint,
		Destination:      flight.Destination,
		Price:            flight.Price,
	}
	return dto
}
