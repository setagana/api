package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var apirouter = mux.NewRouter()
	apirouter.HandleFunc("/health", health).Methods("GET")

	fmt.Println("API srv up and running")
	log.Fatal(http.ListenAndServe(":3000", apirouter))
}

func health(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("alive")
}
