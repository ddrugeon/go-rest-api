package server

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ddrugeon/go-rest-api/internal/app"
	"github.com/ddrugeon/go-rest-api/internal/middlewares"
	"github.com/ddrugeon/go-rest-api/internal/router"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type Server struct {
	router *mux.Router
	server http.Server
}

func NewServer(app *app.App) Server {
	router.InitRoutes(app)
	return Server{
		router: router.NewRouter(app),
	}
}

// Initialize database connection
func (s *Server) Initialize() {
}

// Run application server
func (s *Server) Run(addr string) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt, os.Kill)

	go func() {
		n := negroni.Classic()
		n.Use(middlewares.NewJSONContentTypeMiddleware())
		n.UseHandler(s.router)
		log.Fatal(http.ListenAndServe(addr, n))

	}()

	log.Printf("Server is listning on http://%s\n", addr)
	sig := <-sigs
	log.Println("Signal: ", sig)
}
