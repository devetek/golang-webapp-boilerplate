package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// register your controller here!
type RouteConfig struct {
	Router         *chi.Mux
	HomeController *HomeController
}

func (c *RouteConfig) Setup() {
	c.SetupMiddleware()
	c.SetupStaticFileServing()
	c.SetupGuestRoute()
}

// setup global middleware
func (c *RouteConfig) SetupMiddleware() {
	c.Router.Use(middleware.Logger)
}

func (c *RouteConfig) SetupStaticFileServing() {
	var fs = http.FileServer(http.Dir("public"))

	c.Router.Handle("/static/*", http.StripPrefix("/static/", fs))
}

func (c *RouteConfig) SetupGuestRoute() {
	c.Router.Get("/", c.HomeController.Home)
}
