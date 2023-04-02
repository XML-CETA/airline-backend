package main

import (
	"log"
	"main/handler"
	"main/repo"
	"main/service"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func startServer(userHandler *handler.UserHandler, authHandler *handler.AuthHandler,
	flightHandler *handler.FlightHandler, ticketHandler *handler.TicketHandler) {

	router := mux.NewRouter().StrictSlash(true)

	cors := cors.New(cors.Options{
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
	router.HandleFunc("/tickets", ticketHandler.CreateTicket).Methods("POST")
	router.HandleFunc("/tickets/{id}", ticketHandler.GetOne).Methods("GET")
	router.HandleFunc("/tickets", ticketHandler.GetAll).Methods("GET")
	router.HandleFunc("/flights/search", flightHandler.SearchFlights).Methods("POST")

	println("Server starting")
	log.Fatal(http.ListenAndServe(":3000", cors.Handler(router)))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	userRepository := &repo.UserRepository{}
	userService := &service.UserService{Repo: userRepository}
	authService := &service.AuthService{Repo: userRepository}
	userHandler := &handler.UserHandler{Service: userService}
	authHandler := &handler.AuthHandler{AuthService: authService}
	flightRepository := &repo.FlightRepository{}
	flightService := &service.FlightService{Repo: flightRepository}
	flightHandler := &handler.FlightHandler{Service: flightService}
	ticketRepository := &repo.TicketRepository{}
	ticketService := &service.TicketService{Repo: ticketRepository, FlightRepo: flightRepository, UserRepo: userRepository}
	ticketHandler := &handler.TicketHandler{Service: ticketService, Auth: authHandler}
	startServer(userHandler, authHandler, flightHandler, ticketHandler)
}
