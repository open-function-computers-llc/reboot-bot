package server

import (
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

type server struct {
	port   int
	apiKey string
	router *http.ServeMux
	logger *logrus.Logger
}

func New(port int, apiKey string) (server, error) {
	s := server{
		port:   port,
		apiKey: apiKey,
		logger: logrus.New(),
		router: http.NewServeMux(),
	}

	s.bindRoutes()

	return s, nil
}

func (s *server) Serve() error {
	return http.ListenAndServe(":"+strconv.Itoa(s.port), s.router)
}
