package service

import (
	//"errors"

	"main/dtos"
	"main/model"
	"main/repo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FlightService struct {
	Repo *repo.FlightRepository
}

func (service *FlightService) Create(flight *model.Flight) error {
	flight.RemainingSeats = flight.Seats
	return service.Repo.Create(flight)
}

func (service *FlightService) GetOne(id primitive.ObjectID) (model.Flight, error) {
	return service.Repo.GetOne(id)
}

func (service *FlightService) Delete(id primitive.ObjectID) error {
	return service.Repo.Delete(id)
}

func (service *FlightService) Update(flight *model.Flight) error {
	return service.Repo.Update(flight)
}

func (service *FlightService) GetAll() ([]dtos.FlightDto, error) {
	return service.Repo.GetAll()
}

func (service *FlightService) GetAllUpcoming() ([]dtos.FlightDto, error) {
	return service.Repo.GetAllUpcoming()
}
