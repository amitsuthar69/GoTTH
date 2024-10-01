package server

import (
	"encoding/json"
	"net/http"

	"gotth/internal/middleware"
	"gotth/web"

	"github.com/a-h/templ"
)

type Middleware func(http.Handler) http.Handler

// wrapMiddleware applies middleware to a handler and returns the wrapped handler
func wrapMiddleware(h http.Handler, middlewares []Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	// Define routes with their handlers and middleware
	routes := map[string]struct {
		Handler     http.Handler
		Middlewares []Middleware
	}{
		"/": {
			Handler:     templ.Handler(web.HelloForm()),
			Middlewares: []Middleware{middleware.Logger}, // only logger middleware
		},
		"/home": {
			Handler:     http.HandlerFunc(s.homeHandler),
			Middlewares: []Middleware{middleware.Auth}, // only auth middleware
		},
		"/hello": {
			Handler:     http.HandlerFunc(web.HelloWebHandler),
			Middlewares: []Middleware{middleware.Logger, middleware.Auth}, // multiple middlewares
		},
		"/health": {
			Handler:     http.HandlerFunc(s.healthHandler),
			Middlewares: []Middleware{}, // No middleware
		},
	}

	// Register routes with pre-wrapped middleware
	for pattern, route := range routes {
		mux.Handle(pattern, wrapMiddleware(route.Handler, route.Middlewares))
	}

	// Static file server (no middleware)
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	return mux
}

func (s *Server) homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home"))
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}
