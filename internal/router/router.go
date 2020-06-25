package router

import (
	"github.com/ddrugeon/go-rest-api/internal/app"
	"github.com/ddrugeon/go-rest-api/internal/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(app *app.App) *mux.Router {
	router := mux.NewRouter()
	router.Use(handlers.JSONContentTypeMiddleware)
	router.Handle("/healthz", handlers.Health(app)).Methods("GET")
	router.Handle("/droids", handlers.GetAllDroids(app)).Methods("GET")
	router.Handle("/droid/{id}", handlers.GetDroidByID(app)).Methods("GET")

	return router
}