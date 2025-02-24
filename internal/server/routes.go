package server

import (
	"encoding/json"
	"log"
	"net/http"

	"teleport/cmd/web"

	"github.com/a-h/templ"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.Handle("/", templ.Handler(web.LongUrl()))
	mux.HandleFunc("/health", s.healthHandler)
	fileServer := http.FileServer(http.FS(web.Files))
	mux.Handle("/js/", fileServer)
	mux.HandleFunc("/short", func(w http.ResponseWriter, r *http.Request) {
		web.ShortUrlHandler(s.db, w, r)
	})
	mux.HandleFunc("GET /v1/{hash}", s.getLongUrl)
	return mux
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "server running well."

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) getLongUrl(w http.ResponseWriter, r *http.Request) {
	hash := r.PathValue("hash")
	longUrl := s.db.GetLongUrl(hash)
	if longUrl != "" {
		http.Redirect(w, r, longUrl, http.StatusMovedPermanently)
		return
	}
	http.NotFound(w, r)
}
