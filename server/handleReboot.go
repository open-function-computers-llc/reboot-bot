package server

import (
	"bytes"
	"io"
	"net/http"
)

func (s *server) handleReboot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		body := bytes.NewBuffer([]byte(""))
		client := &http.Client{}
		req, err := http.NewRequest("POST", "https://api.vultr.com/v2/instances/"+id+"/reboot", body)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		req.Header.Add("Authorization", "Bearer "+s.apiKey)

		resp, err := client.Do(req)
		if err != nil {
			w.Write([]byte(err.Error() + id))
			return
		}
		defer resp.Body.Close()

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte("<h1>🤖 Rebooting " + id + "! 🤖</h1>"))
		w.Write(respBody)
	}
}
