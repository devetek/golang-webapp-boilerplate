package aboutcontroller

import (
	"net/http"

	"github.com/devetek/go-core/render"
)

func ComponentAboutPage(w http.ResponseWriter, r *http.Request) {
	rdr := render.Context(r.Context())

	rdr.RenderClean("router/about/view/index.html")
}
