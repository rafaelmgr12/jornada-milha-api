package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rafaelmgr12/jornada-milha-api/configs"
	_ "github.com/rafaelmgr12/jornada-milha-api/docs"
	"github.com/rafaelmgr12/jornada-milha-api/internal/infra/repository"
	"github.com/rafaelmgr12/jornada-milha-api/internal/infra/web"
	"github.com/rafaelmgr12/jornada-milha-api/internal/infra/web/webserver"
	"github.com/rafaelmgr12/jornada-milha-api/internal/usecase/destinations"
	"github.com/rafaelmgr12/jornada-milha-api/internal/usecase/testimonials"
)

// @title Jorna Milha API
// @version 1.0

// @contact.name Rafael Ribeiro
// @contact.email ribeirorafaelmatheus@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /v2

func main() {

	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	conn, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true",
		configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	repoTestimonials := repository.NewTestimonialRepository(conn)
	repoDestinations := repository.NewDestinationsRepository(conn)

	usecaseTestimonials := testimonials.NewTestimonialsUseCase(repoTestimonials)
	usecaseDestinations := destinations.NewDestinationsUseCase(repoDestinations)

	webserver := webserver.NewWebServer(":" + configs.WebServerPort)
	webserverTestimonialHandler := web.NewWebTestimonialHandler(*usecaseTestimonials)
	webserverDestinationsHandler := web.NewWebDestinationsHandler(*usecaseDestinations)

	webserver.AddHandlerWithMethod("/api/v1/depoimentos", http.MethodPost, webserverTestimonialHandler.CreateTestimonial)
	webserver.AddHandlerWithMethod("/api/v1/depoimentos", http.MethodGet, webserverTestimonialHandler.GetListTestimonials)
	webserver.AddHandlerWithMethod("/api/v1/depoimentos", http.MethodPut, webserverTestimonialHandler.UpdateTestimonial)
	webserver.AddHandlerWithMethod("/api/v1/depoimentos/{id}", http.MethodDelete, webserverTestimonialHandler.DeleteTestimonial)
	webserver.AddHandlerWithMethod("/api/v1/depoimentos-home", http.MethodGet, webserverTestimonialHandler.GetThreeRandonTestimonial)

	webserver.AddHandlerWithMethod("/api/v1/destinos", http.MethodPost, webserverDestinationsHandler.CreateDestinations)
	webserver.AddHandlerWithMethod("/api/v1/destinos", http.MethodGet, webserverDestinationsHandler.GetListDestinations)
	webserver.AddHandlerWithMethod("/api/v1/destinos", http.MethodPut, webserverDestinationsHandler.UpdateDestinations)
	webserver.AddHandlerWithMethod("/api/v1/destinos/{id}", http.MethodDelete, webserverDestinationsHandler.DeleteDestinations)

	log.Println("Server running on port " + configs.WebServerPort)

	webserver.Start()
}
