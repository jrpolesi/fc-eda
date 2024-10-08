package webserver

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(webServerPort string) *WebServer {
	return &WebServer{
		Router: chi.NewRouter(),
		Handlers: map[string]map[string]http.HandlerFunc{
			"POST": {},
			"GET":  {},
		},
		WebServerPort: webServerPort,
	}
}

func (s *WebServer) AddHandler(path string, method string, handler http.HandlerFunc) {
	s.Handlers[method][path] = handler
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for path, handler := range s.Handlers["POST"] {
		s.Router.Post(path, handler)
	}

	for path, handler := range s.Handlers["GET"] {
		s.Router.Get(path, handler)
	}

	http.ListenAndServe(s.WebServerPort, s.Router)
}
