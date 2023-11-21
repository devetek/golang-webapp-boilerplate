package router

import (
	"github.com/devetek/golang-webapp-boilerplate/internal/router/about/aboutcontroller"
	"github.com/devetek/golang-webapp-boilerplate/internal/router/home/homecontroller"
	"github.com/devetek/golang-webapp-boilerplate/internal/router/service/servicecontroller"
	"github.com/go-chi/chi/v5"
)

func Setup(r *chi.Mux) {
	r.Get("/", homecontroller.Handler)
	r.Get("/about", aboutcontroller.Handler)
	r.Get("/service", servicecontroller.Handler)

	// partial components
	r.Route("/api", func(r chi.Router) {
		r.Route("/home", func(r chi.Router) {
			r.Get("/", homecontroller.ComponentHomePage)
		})
		r.Route("/about", func(r chi.Router) {
			r.Get("/", aboutcontroller.ComponentAboutPage)
		})
		r.Route("/service", func(r chi.Router) {
			r.Get("/", servicecontroller.ComponentServicePage)
		})
	})
}
