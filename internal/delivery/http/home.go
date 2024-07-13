package http

import (
	"net/http"

	"github.com/devetek/go-core/render"
	"github.com/devetek/golang-webapp-boilerplate/internal/services/member"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type HomeController struct {
	log       *logrus.Logger
	myUsecase *member.UseCase
	view      *render.Engine
}

func NewHomeController(
	db *gorm.DB,
	log *logrus.Logger,
	view *render.Engine,
	validate *validator.Validate,
) *HomeController {
	// init module repositories
	myRepository := member.NewRepository(log)

	// init module usecase
	myUsecase := member.NewUseCase(db, log, validate, myRepository)

	return &HomeController{
		log:       log,
		myUsecase: myUsecase,
		view:      view,
	}
}

func (c *HomeController) setHeaderMeta() {
	c.view.Set("title", "Golang Web App Boilerplate")
	c.view.Set("description", "Welcome to Golang web app boilerplate, will help you to create web app with Golang, HTMX and tailwind")
}

func (c *HomeController) Home(w http.ResponseWriter, r *http.Request) {
	c.setHeaderMeta()

	// render page with template html (ejs)
	err := c.view.HTML(w).Render("views/pages/home/index.html")
	if err != nil {
		c.log.Warnf("Render error : %+v", err)
	}

}

func (c *HomeController) Component(w http.ResponseWriter, r *http.Request) {
	c.setHeaderMeta()

	err := c.view.HTML(w).RenderClean("views/pages/home/component.html")
	if err != nil {
		c.log.Warnf("RenderClean error : %+v", err)
	}
}
