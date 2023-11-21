package middleware

import (
	"github.com/devetek/go-core/render"
	"github.com/devetek/golang-webapp-boilerplate/internal"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Setup(r *chi.Mux) {

	r.Use(middleware.Logger)

	// create template engine with global default value
	r.Use(render.Middleware(
		render.NewEngine(
			internal.Template,
			render.WithValues(
				map[string]any{
					"keywords": "golang, HTMX, Web Platform, SPA, Tailwind",
				}),
		)),
	)
}
