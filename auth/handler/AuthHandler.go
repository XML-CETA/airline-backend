package handler

import (
	"encoding/json"
	"main/auth/generator"
	"main/auth/service"
	"main/dtos"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	AuthService *service.AuthService
}

func (handler *AuthHandler) Login(writer http.ResponseWriter, req *http.Request) {
	var loginDto dtos.LoginDto
	err := json.NewDecoder(req.Body).Decode(&loginDto)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	var token string
	token, err = handler.AuthService.Login(&loginDto)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(token)
}

func (handler *AuthHandler) Authorize(protectedEndpoint http.HandlerFunc, expectedRole string) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		authorizationHeader := request.Header.Get("Authorization")
		if authorizationHeader == "" {
			http.Error(writer, "You are unauthorized", http.StatusUnauthorized)
			return
		}
		token, claims := handler.ParseJwt(authorizationHeader)
		if !token.Valid {
			http.Error(writer, "Token is not valid anymore", http.StatusUnauthorized)
			return
		}

		if claims.CustomClaims["role"] != expectedRole {
			http.Error(writer, "You are not authorized for this endpoint", http.StatusForbidden)
			return
		}
		protectedEndpoint(writer, request)

	})
}

func (handler *AuthHandler) ParseJwt(authorizationHeader string) (*jwt.Token, *generator.Claims) {
	tokenString := strings.TrimSpace(strings.Split(authorizationHeader, "Bearer")[1])
	token, _ := jwt.ParseWithClaims(tokenString, &generator.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	}, jwt.WithLeeway(5*time.Second))

	claims, _ := token.Claims.(*generator.Claims)

	return token, claims
}
