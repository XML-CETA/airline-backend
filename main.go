package main

import (
	"context"
	"fmt"
	"log"
	"main/handler"
	"net/http"
	"main/utils"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/joho/godotenv"
)

func initDatabase() {
	client, cancel := utils.GetConn()
	defer cancel()

	databases, _ := client.ListDatabaseNames(context.TODO(), bson.M{})
	fmt.Println(databases)
}


func startServer(handler *handler.UserHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/users", handler.CreateUser).Methods("GET")

	println("Server starting")
	log.Fatal(http.ListenAndServe(":3000", router))
}


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	initDatabase()
	userHandler := &handler.UserHandler{}

	startServer(userHandler)
}
