package mixer

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"go.sancus.dev/mix/wrapper"
)

type Router struct {
	mixer *Mixer

	routes      []chi.Route
	middlewares chi.Middlewares
}

// Mount
func (m *Mixer) Mount(pattern string, h http.Handler) {
	wrapper.NewWrapper(pattern, h)
}

// Close
func (m *Mixer) Close() error {
	return nil
}

// Reload
func (m *Mixer) Reload() error {
	return nil
}

// types.Routes
func (r *Router) Routes() []chi.Route {
	return r.routes
}

func (r *Router) Middlewares() chi.Middlewares {
	return r.middlewares
}

func (r *Router) Match(rctx *chi.Context, method, path string) bool {
	return false
}
