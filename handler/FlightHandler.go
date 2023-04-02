package handler

import (
	"encoding/json"
	"main/dtos"
	"main/model"
	"main/service"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FlightHandler struct {
	Service *service.FlightService
}

func (handler *FlightHandler) CreateFlight(writer http.ResponseWriter, req *http.Request) {
	var flight model.Flight
	err := json.NewDecoder(req.Body).Decode(&flight)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.Service.Create(&flight)

	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *FlightHandler) GetOne(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	idO, _ := primitive.ObjectIDFromHex(id)
	flight, err := handler.Service.GetOne(idO)

	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	json.NewEncoder(writer).Encode(flight)
}

func (handler *FlightHandler) UpdateFlight(writer http.ResponseWriter, req *http.Request) {
	var dto dtos.FlightDto
	err := json.NewDecoder(req.Body).Decode(&dto)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	flight, err := dto.Repackage()

	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	err = handler.Service.Update(&flight)

	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *FlightHandler) DeleteFlight(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	idO, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.Service.Delete(idO)

	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *FlightHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	var flights []dtos.FlightDto
	flights, err := handler.Service.GetAll()

	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	json.NewEncoder(writer).Encode(flights)
}

func (handler *FlightHandler) GetAllUpcoming(writer http.ResponseWriter, req *http.Request) {
	var flights []dtos.FlightDto
	flights, err := handler.Service.GetAllUpcoming()

	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	json.NewEncoder(writer).Encode(flights)
}
