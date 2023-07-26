package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rafaelmgr12/jornada-milha-api/configs"
	"github.com/rafaelmgr12/jornada-milha-api/internal/infra/repository"
	"github.com/rafaelmgr12/jornada-milha-api/internal/infra/web"
	"github.com/rafaelmgr12/jornada-milha-api/internal/infra/web/webserver"
	"github.com/rafaelmgr12/jornada-milha-api/internal/usecase/testimonials"
)

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

	repo := repository.NewTestimonialRepository(conn)
	usecase := testimonials.NewTestimonialsUseCase(repo)

	webserver := webserver.NewWebServer(":" + configs.WebServerPort)
	webserverTestimonialHandler := web.NewWebTestimonialHandler(*usecase)

	webserver.AddHandlerWithMethod("/api/v1/depoimentos", http.MethodPost, webserverTestimonialHandler.CreateTestimonial)
	webserver.AddHandlerWithMethod("/api/v1/depoimentos", http.MethodGet, webserverTestimonialHandler.GetListTestimonials)
	webserver.AddHandlerWithMethod("/api/v1/depoimentos", http.MethodPut, webserverTestimonialHandler.UpdateTestimonial)
	webserver.AddHandlerWithMethod("/api/v1/depoimentos/{id}", http.MethodDelete, webserverTestimonialHandler.DeleteTestimonial)

	log.Println("Server running on port " + configs.WebServerPort)

	webserver.Start()
}
