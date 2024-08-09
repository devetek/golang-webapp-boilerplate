package http

import (
	"net/http"

	"github.com/devetek/go-core/render"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type HomeController struct {
	log  *logrus.Logger
	view *render.Engine
}

func NewHomeController(
	db *gorm.DB,
	log *logrus.Logger,
	view *render.Engine,
	validate *validator.Validate,
) *HomeController {
	return &HomeController{
		log:  log,
		view: view,
	}
}

func (c *HomeController) setHeaderMeta() {
	c.view.Set("title", "Golang + EJS + Tailwind + HTMX")
	c.view.Set("description", "Golang web app boilerplate, help you to build web app with Golang, EJS, HTMX and tailwind")
	c.view.Set("keywords", "golang, ejs, htmx, tailwind, website, server-driven component")
}

func (c *HomeController) Home(w http.ResponseWriter, r *http.Request) {
	c.setHeaderMeta()

	// render page with template html (ejs)
	err := c.view.HTML(w).Render("views/pages/home/index.ejs")
	if err != nil {
		c.log.Warnf("Render error : %+v", err)
	}
}
