package main

import (
	"log"
	"net/http"

	"github.com/devetek/golang-webapp-boilerplate/internal/middleware"
	"github.com/devetek/golang-webapp-boilerplate/internal/router"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	// setup all middlewares
	middleware.Setup(r)

	// setup all routes
	router.Setup(r)

	// log fatal on error app
	log.Fatal(http.ListenAndServe(":3000", r))
}
