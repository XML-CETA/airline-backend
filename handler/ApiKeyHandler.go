package handler

import (
	"encoding/json"
	"main/dtos"
	"main/service"
	"net/http"
)

type ApiKeyHandler struct {
	Service *service.ApiKeyService
	Auth    *AuthHandler
}

func (handler *ApiKeyHandler) GenerateApiKey(writer http.ResponseWriter, req *http.Request) {
	var apiKeyDto dtos.ApiKeyDto
	user, role := handler.Auth.GetUsername(writer, req)
	err := json.NewDecoder(req.Body).Decode(&apiKeyDto)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	apiKey, err := handler.Service.GenerateApiKey(user, role, &apiKeyDto)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(apiKey)
}
