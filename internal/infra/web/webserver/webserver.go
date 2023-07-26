package webserver

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type WebServer struct {
	Router        chi.Router
	WebServerPort string
}

func NewWebServer(webServerPort string) *WebServer {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	return &WebServer{
		WebServerPort: webServerPort,
		Router:        router,
	}
}

func (s *WebServer) AddHandlerWithMethod(path string, method string, handler http.HandlerFunc) {
	switch method {
	case http.MethodGet:
		s.Router.Get(path, handler)
	case http.MethodPost:
		s.Router.Post(path, handler)
	case http.MethodPut:
		s.Router.Put(path, handler)
	case http.MethodDelete:
		s.Router.Delete(path, handler)
	default:
		panic("Invalid method")
	}
}

func (s *WebServer) Start() {
	if err := http.ListenAndServe(s.WebServerPort, s.Router); err != nil {
		panic(err.Error())
	}
}
