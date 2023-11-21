package servicecontroller

import (
	"net/http"

	"github.com/devetek/go-core/render"
)

func ComponentServicePage(w http.ResponseWriter, r *http.Request) {
	rdr := render.Context(r.Context())

	rdr.RenderClean("router/service/view/index.html")
}
