package main

import (
	"log"
	"main/auth/generator"
	authHandler "main/auth/handler"
	authService "main/auth/service"
	"main/handler"
	"main/repo"
	"main/service"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func startServer(userHandler *handler.UserHandler, authHandler *authHandler.AuthHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{username}", authHandler.Authorize(userHandler.GetOne, "Admin")).Methods("GET")
	router.HandleFunc("/login", authHandler.Login).Methods("POST")
	println("Server starting")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	userRepository := &repo.UserRepository{}
	jwtGenerator := &generator.JwtGenerator{}
	userService := &service.UserService{Repo: userRepository}
	authService := &authService.AuthService{Repo: userRepository, JwtGenerator: jwtGenerator}
	userHandler := &handler.UserHandler{Service: userService}
	authHandler := &authHandler.AuthHandler{AuthService: authService}

	startServer(userHandler, authHandler)
}
