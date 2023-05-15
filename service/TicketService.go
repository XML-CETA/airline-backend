package service

import (
	"errors"
	"main/dtos"
	"main/model"
	"main/repo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TicketService struct {
	Repo       *repo.TicketRepository
	FlightRepo *repo.FlightRepository
	UserRepo   *repo.UserRepository
}

func (service *TicketService) Create(ticket *model.Ticket) error {
	flight, err := service.FlightRepo.GetOne(ticket.FlightId)

	if err != nil {
		return errors.New("Flight does not exist!")
	}

	_, err = service.UserRepo.GetOne(ticket.User)

	if err != nil {
		return errors.New("User does not exist!")
	}

	if flight.RemainingSeats-ticket.Amount < 0 {
		return errors.New("Not enough tickets remaining!")
	}

	flight.RemainingSeats = flight.RemainingSeats - ticket.Amount
	err = service.FlightRepo.Update(&flight)

	if err != nil {
		return errors.New("Failed to update flight!")
	}

	return service.Repo.Create(ticket)
}

func (service *TicketService) GetOne(id primitive.ObjectID) (dtos.TicketDto, error) {
	ticket, err := service.Repo.GetOne(id)

	if err != nil {
		err = errors.New("Ticket does not exist!")
		return dtos.TicketDto{}, err
	}

	flight, err := service.FlightRepo.GetOne(ticket.FlightId)

	if err != nil {
		err = errors.New("Flight does not exist!")
		return dtos.TicketDto{}, err
	}

	dto := dtos.Construct(ticket, flight)

	return dto, err
}

func (service *TicketService) GetAll(username string) ([]dtos.TicketDto, error) {
	tickets, err := service.Repo.GetAll(username)

	if err != nil {
		return []dtos.TicketDto{}, err
	}

	var ticketList []dtos.TicketDto

	for _, ticket := range tickets {
		flight, err := service.FlightRepo.GetOne(ticket.FlightId)

		if err != nil {
			err = errors.New("Flight for Ticket ID -> {" + ticket.Id.String() + "} does not exist!")
			return []dtos.TicketDto{}, err
		}

		ticketList = append(ticketList, dtos.Construct(ticket, flight))
	}

	return ticketList, err
}

func (service *TicketService) DeleteByFlight(flight primitive.ObjectID) error {
	return service.Repo.DeleteByFlight(flight);
}
