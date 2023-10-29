package main

import (
	"net/http"

	"github.com/danielzinhors/multithreading-go-expert/configs"
	"github.com/danielzinhors/multithreading-go-expert/infra/webservers/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Consulta CEP Go Expert API multithreading
// @version         1.0
// @description     Prodct API Consulta CEP
// @termsOfService  http://swagger.io/terms/

// @contact.name   Daniel Silveira
// @contact.url    http://danielsilveira.dev.br
// @contact.email  falecom@danielsilveira.dev.br

// @license.name   Daniel Silveira
// @license.url    http://danielsilveira.dev.br

// @host   localhost:8000
// @basePath  /
func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	cepHandler := handlers.NewCepHanbler()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Route("/cep", func(r chi.Router) {
		r.Get("/", cepHandler.BuscaCep)
	})
	portServer := ":" + configs.WebServerPort
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost"+portServer+"/docs/doc.json")))
	http.ListenAndServe(portServer, r)
}
