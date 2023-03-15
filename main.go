package main

import (
	"log"
	"main/handler"
	"main/repo"
	"main/service"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func startServer(handler *handler.UserHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/users", handler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{username}", handler.GetOne).Methods("GET")

	println("Server starting")
	log.Fatal(http.ListenAndServe(":3000", router))
}


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	userRepository := &repo.UserRepository{}
	userService := &service.UserService{Repo: userRepository}
	userHandler := &handler.UserHandler{Service: userService}

	startServer(userHandler)
}
