package aboutcontroller

import (
	"net/http"

	"github.com/devetek/go-core/render"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	rdr := render.Context(r.Context())

	rdr.Set("title", "About - Golang Web App Boilerplate")
	rdr.Set("description", "About golang web app boilerplate, will help you to create web app with Golang, HTMX and tailwind")

	rdr.Render("router/about/view/index.html")
}
