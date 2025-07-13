package webserver

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	s.Handlers[path] = handler
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for path, handler := range s.Handlers {
		parts := strings.Split(path, " ")
		if len(parts) > 1 {
			// Format: "METHOD /path"
			method := parts[0]
			route := parts[1]
			s.Router.Method(method, route, handler)
		} else {
			// Legacy format: just "/path"
			s.Router.Handle(path, handler)
		}
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
