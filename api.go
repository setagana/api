package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	var apirouter = mux.NewRouter()
	apirouter.HandleFunc("/health", health).Methods("GET")
	apirouter.HandleFunc("/message", handleQuery).Methods("GET")
	apirouter.HandleFunc("/m/{msg}", handleURL).Methods("GET")

	fmt.Println("API srv up and running")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS()(apirouter)))
}

// Health check. Polling /health outputs "ok"
func health(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("ok")
}

// Takes the GET parm from the URL and outputs json response
// /message?msg=this%20msg outputs {"message":"this msg"}
func handleQuery(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	message := vars.Get("msg")

	json.NewEncoder(w).Encode(map[string]string{"message": message})
}

// Takes URL path and outputs a json response
// /m/this%20msg outputs {"message":"this msg"}
func handleURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	message := vars["msg"]

	json.NewEncoder(w).Encode(map[string]string{"message": message})
}
