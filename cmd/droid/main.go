package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

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

	http.ListenAndServe(":8000", router)
}
