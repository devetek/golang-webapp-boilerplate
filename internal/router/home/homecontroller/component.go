package homecontroller

import (
	"net/http"

	"github.com/devetek/go-core/render"
)

func ComponentHomePage(w http.ResponseWriter, r *http.Request) {
	rdr := render.Context(r.Context())

	rdr.RenderClean("router/home/view/index.html")
}
