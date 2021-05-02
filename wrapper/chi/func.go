package chi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetRoutePath(r *http.Request) string {
	rctx := chi.RouteContext(r.Context())
	if rctx != nil && rctx.RoutePath != "" {
		return rctx.RoutePath
	} else {
		return r.URL.Path
	}
}

func SetRoutePath(r *http.Request, path string) {
	rctx := chi.RouteContext(r.Context())
	if rctx != nil {
		rctx.RoutePath = path
	} else {
		r.URL.Path = path
	}
}
