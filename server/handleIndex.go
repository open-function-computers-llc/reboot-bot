package server

import "net/http"

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>🤖 server rebooter! 🤖</h1>"))
	}
}
