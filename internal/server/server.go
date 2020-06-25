package server

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ddrugeon/go-rest-api/internal/app"
	"github.com/ddrugeon/go-rest-api/internal/middlewares"
	"github.com/ddrugeon/go-rest-api/internal/router"
	"github.com/gorilla/mux"
	negronilogus "github.com/meatballhat/negroni-logrus"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

type Server struct {
	router *mux.Router
	server http.Server
	port   string
	logger *logrus.Logger
}

func NewServer(app *app.App) Server {
	router.InitRoutes(app)
	return Server{
		router: router.NewRouter(app),
		port:   app.Port,
		logger: app.Logger,
	}
}

// Initialize database connection
func (s *Server) Initialize() {
}

// Run application server
func (s *Server) Run() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt, os.Kill)

	go func() {
		n := negroni.New()
		n.Use(negronilogus.NewMiddlewareFromLogger(s.logger, "web"))
		n.Use(middlewares.NewJSONContentTypeMiddleware())
		n.UseHandler(s.router)
		s.logger.Fatal(http.ListenAndServe(s.port, n))

	}()

	s.logger.Printf("Server is listening on http://%s\n", s.port)
	sig := <-sigs
	s.logger.Println("Signal: ", sig)
}
