package homecontroller

import (
	"net/http"

	"github.com/devetek/go-core/render"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	rdr := render.Context(r.Context())

	rdr.Set("title", "Golang Web App Boilerplate")
	rdr.Set("description", "Welcome to Golang web app boilerplate, will help you to create web app with Golang, HTMX and tailwind")

	rdr.Render("router/home/view/index.html")
}
