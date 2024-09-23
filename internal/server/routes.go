package server

import (
	"encoding/json"
	"net/http"

	"gotth/internal/middleware"
	"gotth/web"

	"github.com/a-h/templ"
)

type Middleware func(http.Handler) http.Handler

func Use(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	// Register the static file handlers
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// db health
	mux.HandleFunc("/health", s.healthHandler)

	// Register other routes
	mux.Handle("/", templ.Handler(web.HelloForm()))
	mux.HandleFunc("/hello", web.HelloWebHandler)

	return Use(mux, middleware.Logger, middleware.Auth)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}
