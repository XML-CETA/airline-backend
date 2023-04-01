package dtos

import (
	"errors"
	"main/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (flightDto *FlightDto) Repackage() (model.Flight, error) {
	id, err := primitive.ObjectIDFromHex(flightDto.Id)

	if err != nil {
		err = errors.New("Provided HexStringID is not a valid ObjectID!")
		return model.Flight{}, err
	}

	flight := model.Flight{
		Id:               id,
		FlighDateAndTime: flightDto.FlighDateAndTime,
		StartingPoint:    flightDto.StartingPoint,
		Destination:      flightDto.Destination,
		Price:            flightDto.Price,
		Seats:            flightDto.Seats,
		RemainingSeats:   flightDto.RemainingSeats,
	}
	return flight, nil
}
