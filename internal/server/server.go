package server

import (
	"net/http"

	"github.com/ddrugeon/go-rest-api/internal/app"
	"github.com/ddrugeon/go-rest-api/internal/router"
	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

func NewServer(app *app.App) *Server {
	return &Server{
		router: router.NewRouter(app),
	}
}

// Initialize database connection
func (s *Server) Initialize() {
}

// Run application server
func (s *Server) Run(addr string) {
	http.ListenAndServe(addr, s.router)
}
