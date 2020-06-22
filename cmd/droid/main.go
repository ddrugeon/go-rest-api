package main

import (
	"fmt"
	"net/http"
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
	fmt.Print("Serving at 0.0.0.0:9090... ")
	http.HandleFunc("/", handle)
	http.ListenAndServe("0.0.0.0:9090", nil)
}
