package router

import (
	"github.com/ddrugeon/go-rest-api/internal/app"
	"github.com/ddrugeon/go-rest-api/internal/handlers"
	"github.com/ddrugeon/go-rest-api/internal/middlewares"
	"github.com/gorilla/mux"
)

func NewRouter(app *app.App) *mux.Router {
	router := mux.NewRouter()
	router.Use(middlewares.JSONContentTypeMiddleware)
	router.Handle("/healthz", handlers.Health(app)).Methods("GET")
	router.Handle("/droids", handlers.GetAllDroids(app)).Methods("GET")
	router.Handle("/droids/{id}", handlers.GetDroidByID(app)).Methods("GET")

	return router
}
