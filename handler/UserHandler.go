package handler

import (
	"net/http"
	"encoding/json"
)

type UserHandler struct {

}


func (handler *UserHandler) CreateUser(writer http.ResponseWriter, req *http.Request) {
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode("Hello")
}
