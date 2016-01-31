package filters

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/unrolled/render"
)

const (
	renderContextKey = "filters/render_setup"
)

type renderSetupHandler struct {
	next http.Handler
}

func (p *renderSetupHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	renderEngine := render.New(render.Options{
		IndentJSON: (r.Form.Get("pretty") == "true"),
	})

	context.Set(r, renderContextKey, renderEngine)
	defer context.Delete(r, renderContextKey)

	p.next.ServeHTTP(w, r)
}

func RenderSetupHandler(next http.Handler) http.Handler {

	return &renderSetupHandler{next}

}

func GetRenderer(r *http.Request) *render.Render {
	return context.Get(r, renderContextKey).(*render.Render)
}
