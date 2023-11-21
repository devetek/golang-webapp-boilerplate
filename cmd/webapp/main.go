package main

import (
	"net/http"

	"github.com/devetek/golang-webapp-boilerplate/internal/middleware"
	"github.com/devetek/golang-webapp-boilerplate/internal/router/home/homecontroller"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	// setup all middlewares
	middleware.Setup(r)

	r.Get("/", homecontroller.Handler)

	http.ListenAndServe(":3000", r)
}
