package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func startServer() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/users", func (writer http.ResponseWriter, req *http.Request) {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode("Hello")
	}).Methods("GET")

	println("Server starting")
	log.Fatal(http.ListenAndServe(":3000", router))
}


func main() {
	startServer()
}
