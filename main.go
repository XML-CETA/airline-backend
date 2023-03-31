package main

import (
	"log"
	"main/auth/generator"
	authHandler "main/auth/handler"
	authService "main/auth/service"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

	"main/handler"
	"main/repo"
	"main/service"
	"net/http"
)

func startServer(userHandler *handler.UserHandler, authHandler *authHandler.AuthHandler,
	flightHandler *handler.FlightHandler) {
	router := mux.NewRouter().StrictSlash(true)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
		},
		AllowedHeaders: []string{
			"*",
		},
	})
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{username}", authHandler.Authorize(userHandler.GetOne, "Admin")).Methods("GET")
	router.HandleFunc("/login", authHandler.Login).Methods("POST")
	router.HandleFunc("/flights", flightHandler.CreateFlight).Methods("POST")
	router.HandleFunc("/flights", flightHandler.GetAll).Methods("GET")
	router.HandleFunc("/flights/upcoming", flightHandler.GetAllUpcoming).Methods("GET")
	router.HandleFunc("/flights/{id}", flightHandler.GetOne).Methods("GET")
	router.HandleFunc("/flights", flightHandler.UpdateFlight).Methods("PUT")
	router.HandleFunc("/flights/{id}", flightHandler.DeleteFlight).Methods("DELETE")

	println("Server starting")
	log.Fatal(http.ListenAndServe(":3000", c.Handler(router)))
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
	flightRepository := &repo.FlightRepository{}
	flightService := &service.FlightService{Repo: flightRepository}
	flightHandler := &handler.FlightHandler{Service: flightService}

	startServer(userHandler, authHandler, flightHandler)
}
