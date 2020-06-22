package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var status HealthStatus

// HealthStatus defines status of current service
type HealthStatus struct {
	Status string `json:"status"`
}

func handleHealthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	status.Status = "OK"
	json.NewEncoder(w).Encode(status)
}

func handle(w http.ResponseWriter, r *http.Request) {
	if "/favicon.ico" == r.URL.Path {
		return
	}
	fmt.Print("request: " + r.URL.Path + " ")

	result := string("Hello World!")

	fmt.Print(result + "\n\r")
	fmt.Fprintf(w, result)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/droids", handle).Methods("GET")
	router.HandleFunc("/healthz", handleHealthz).Methods("GET")

	http.ListenAndServe(":8000", router)
}
