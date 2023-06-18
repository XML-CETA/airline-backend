package handler

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"main/dtos"
	"main/service"
	"main/utils"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TicketHandler struct {
	Service *service.TicketService
	Auth    *AuthHandler
}

func (handler *TicketHandler) CreateTicket(writer http.ResponseWriter, req *http.Request) {
	var ticketDto dtos.CreateTicketDto
	err := json.NewDecoder(req.Body).Decode(&ticketDto)
	var user string
	if ticketDto.ApiKey != "" {
		apiKey, claims, err := handler.ParseApiKey(ticketDto.ApiKey)
		if err != nil {
			writer.WriteHeader(http.StatusUnauthorized)
			writer.Write([]byte(err.Error()))
			return
		}
		if claims.CustomClaims["goal"] != "buyingTickets" || !apiKey.Valid {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		user = claims.CustomClaims["username"]
	} else {
		user, _ = handler.Auth.GetUsername(writer, req)
	}
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

func (handler *TicketHandler) ParseApiKey(apiKey string) (*jwt.Token, *utils.Claims, error) {
	parsed, err := jwt.ParseWithClaims(apiKey, &utils.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("SECRETAPIKEYFORBOOKINGAPP"), nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		return nil, nil, err
	}

	claims, ok := parsed.Claims.(*utils.Claims)
	if !ok {
		return nil, nil, fmt.Errorf("failed to extract claims from token")
	}

	return parsed, claims, nil
}
