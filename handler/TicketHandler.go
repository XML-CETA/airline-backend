package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"main/dtos"
	"main/service"
	"net/http"
)

type TicketHandler struct {
	Service *service.TicketService
	Auth    *AuthHandler
}

func (handler *TicketHandler) CreateTicket(writer http.ResponseWriter, req *http.Request) {
	var ticketDto dtos.CreateTicketDto
	err := json.NewDecoder(req.Body).Decode(&ticketDto)
	user, _ := handler.Auth.GetUsername(writer, req)
	ticket := ticketDto.Repackage()
	ticket.User = user
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.Service.Create(&ticket)

	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *TicketHandler) GetOne(writer http.ResponseWriter, req *http.Request) {
	temp := mux.Vars(req)["id"]
	id, _ := primitive.ObjectIDFromHex(temp)
	ticket, err := handler.Service.GetOne(id)

	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	json.NewEncoder(writer).Encode(ticket)
}

func (handler *TicketHandler) GetAll(writer http.ResponseWriter, req *http.Request) {

	username, _ := handler.Auth.GetUsername(writer, req)

	tickets, err := handler.Service.GetAll(username)

	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	json.NewEncoder(writer).Encode(tickets)

}
