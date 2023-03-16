package handler

import (
	"encoding/json"
	"main/model"
	"main/service"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	Service *service.UserService
}

func (handler *UserHandler) CreateUser(writer http.ResponseWriter, req *http.Request) {
	var user model.User
	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.Service.Create(&user)

	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *UserHandler) GetOne(writer http.ResponseWriter, req *http.Request) {
	username := mux.Vars(req)["username"]
	user, err := handler.Service.GetOne(username)

	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	json.NewEncoder(writer).Encode(user)
}
