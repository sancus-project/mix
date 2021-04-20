package mix

import (
	"go.sancus.dev/mix/types"
)

type Router struct {
	mixer *Mixer

	routes      []types.Route
	middlewares types.Middlewares
}

// Mount
func (m *Mixer) Mount(pattern string, router types.Router) {
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
func (r *Router) Routes() []types.Route {
	return r.routes
}

func (r *Router) Middlewares() types.Middlewares {
	return r.middlewares
}

func (r *Router) Match(rctx *types.Context, method, path string) bool {
	return false
}
