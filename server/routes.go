package server

import "net/http"

func (s *server) bindRoutes() {
	openRoutes := map[string]http.HandlerFunc{
		"POST /reboot/{id}": s.handleReboot(),
		"GET /":             s.handleIndex(),
	}
	for path, handler := range openRoutes {
		s.router.Handle(path, s.LogRequest(handler))
	}
}
