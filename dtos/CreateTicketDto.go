package dtos

import (
	"main/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateTicketDto struct {
	FlightId primitive.ObjectID `json:"flightId"`
	Amount   int                `json:"amount"`
	ApiKey   string             `json:"apiKey,omitempty"`
}

func (ticketDto *CreateTicketDto) Repackage() model.Ticket {
	ticket := model.Ticket{
		User:     "",
		FlightId: ticketDto.FlightId,
		Amount:   ticketDto.Amount,
	}
	return ticket
}
