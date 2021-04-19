package mix

import (
	"net/http"

	"go.sancus.dev/mix/types"
)

// Middleware
func (m *Mixer) MiddlewareHandler(w http.ResponseWriter, r *http.Request, next http.Handler, prefix string) {

	m.NotFound(w, r)
}

func (m *Mixer) Middleware(prefix string) types.MiddlewareHandlerFunc {

	// Wrap MiddlewareHandler
	fn := func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			m.MiddlewareHandler(w, r, next, prefix)
		}

		return http.HandlerFunc(fn)
	}
	return fn
}
