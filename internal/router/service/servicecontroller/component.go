package servicecontroller

import (
	"net/http"

	"github.com/devetek/go-core/render"
)

func ComponentServicePage(w http.ResponseWriter, r *http.Request) {
	rdr := render.Context(r.Context())

	rdr.Set("title", "Service - Golang Web App Boilerplate")
	rdr.Set("description", "Service golang web app boilerplate, will help you to create web app with Golang, HTMX and tailwind")

	rdr.RenderClean("router/service/view/component.html")
}
