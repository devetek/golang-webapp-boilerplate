package homecontroller

import (
	"net/http"

	"github.com/devetek/go-core/render"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	rdr := render.Context(r.Context())

	rdr.Set("name", "Golang Web App Boilerplate")

	rdr.Render("router/home/view/index.html")
}
