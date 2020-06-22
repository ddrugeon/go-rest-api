package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ddrugeon/go-rest-api/internal/healthz"
	"github.com/gorilla/mux"
)

var healthzService healthz.Service

func handle(w http.ResponseWriter, r *http.Request) {
	if "/favicon.ico" == r.URL.Path {
		return
	}
	fmt.Print("request: " + r.URL.Path + " ")

	result := string("Hello World!")

	fmt.Print(result + "\n\r")
	fmt.Fprintf(w, result)
}

func handleHealthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(healthzService.Health())
}

func main() {
	router := mux.NewRouter()
	healthzService = healthz.NewService()

	router.HandleFunc("/droids", handle).Methods("GET")
	router.HandleFunc("/healthz", handleHealthz).Methods("GET")

	http.ListenAndServe(":8000", router)
}
