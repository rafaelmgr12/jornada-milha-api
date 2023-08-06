package webserver

import (
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
)

var log = logrus.New()

func init() {
	logFile, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(logFile)
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.InfoLevel)
}

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		writer := &responseWriter{ResponseWriter: w}

		next.ServeHTTP(writer, r)

		// Ignora logs para o endpoint /metrics do Prometheus
		if r.URL.Path != "/metrics" {
			log.WithFields(logrus.Fields{
				"method":     r.Method,
				"url":        r.URL.Path,
				"remoteAddr": r.RemoteAddr,
				"userAgent":  r.UserAgent(),
				"status":     writer.status,
				"duration":   time.Since(start).String(),
			}).Info("Nova requisição")
		}
	})
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}

type WebServer struct {
	Router        chi.Router
	WebServerPort string
}

func NewWebServer(webServerPort string) *WebServer {

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
	}))

	router.Use(RequestLogger)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
	))

	router.Handle("/metrics", promhttp.Handler())

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
		log.Panic(err.Error())
	}
}
