package router

import (
	"net/http"

	"github.com/ddrugeon/go-rest-api/internal/app"
	"github.com/ddrugeon/go-rest-api/internal/handlers"
	"github.com/gorilla/mux"
)

type route struct {
	name    string
	method  string
	path    string
	handler http.HandlerFunc
}

var routes []route

func InitRoutes(app *app.App) {
	routes = []route{
		{
			name:    "healthz",
			method:  http.MethodGet,
			path:    "/healthz",
			handler: handlers.Health(app),
		},
		{
			name:    "version",
			method:  http.MethodGet,
			path:    "/version",
			handler: handlers.Version(app),
		},
		{
			name:    "getAllDroids",
			method:  http.MethodGet,
			path:    "/droids",
			handler: handlers.GetAllDroids(app),
		},
		{
			name:    "getDroidByID",
			method:  http.MethodGet,
			path:    "/droids/{id}",
			handler: handlers.GetDroidByID(app),
		},
	}
}

func NewRouter(app *app.App) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.KeepContext = true

	for _, route := range routes {
		router.Methods(route.method).
			Path(route.path).
			Name(route.name).
			Handler(route.handler)
	}

	return router
}
