package config

import (
	"github.com/devetek/go-core/render"
	"github.com/devetek/golang-webapp-boilerplate/internal/delivery/http"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type BootstrapConfig struct {
	Config *viper.Viper
	Router *chi.Mux
	View   *render.Engine
	Log    *logrus.Logger
}

func Bootstrap(config *BootstrapConfig) {
	homeController := http.NewHomeController(config.View, config.Log)

	route := &http.RouteConfig{
		Router:         config.Router,
		HomeController: homeController,
	}

	// init registered router
	route.Setup()
}
