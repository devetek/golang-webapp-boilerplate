package main

import (
	"net/http"

	"github.com/devetek/go-core/render"
	"github.com/devetek/golang-webapp-boilerplate/internal"
	"github.com/devetek/golang-webapp-boilerplate/internal/config"
	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := config.NewConfig()
	log := config.NewLogger(cfg)
	router := chi.NewRouter()
	view := render.NewEngine(
		internal.Template,
		render.WithDefaultLayout(cfg.GetString("view.default")),
		render.WithMinifier(render.MinifierOption{
			Enable: true,
			HTML: render.MinifierHTMLConfig{
				KeepComments:            false,
				KeepConditionalComments: false,
				KeepDefaultAttrVals:     false,
				KeepDocumentTags:        false,
				KeepEndTags:             false,
				KeepQuotes:              true,
				KeepWhitespace:          false,
			},
		}),
	)

	config.Bootstrap(&config.BootstrapConfig{
		Router: router,
		View:   view,
		Log:    log,
		Config: cfg,
	})

	// run app and make fatal omn failed
	log.Fatal(http.ListenAndServe(":"+cfg.GetString("application.port"), router))
}
