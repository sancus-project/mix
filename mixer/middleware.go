package mixer

import (
	"net/http"
	"strings"

	"go.sancus.dev/web"
)

// Middleware
func (m *Mixer) MiddlewareHandler(w http.ResponseWriter, r *http.Request, next http.Handler, prefix string) {
	var err error

	if prefix == "/" || prefix == "" {
		// no prefix

		if err = m.TryServeHTTP(w, r); err == nil {
			// Done.
			return
		}

	} else {
		// check prefix
		path := m.config.GetRoutePath(r)

		if s := strings.TrimPrefix(path, prefix); s != path {
			if s == "" {
				path = "/"
			} else {
				path = s
			}

			if path[0] == '/' {

				// Update RoutePath before Handling
				m.config.SetRoutePath(r, path)

				if err = m.TryServeHTTP(w, r); err == nil {
					// Done.
					return
				}
			}
		}
	}

	// Failed
	if err != nil {
		m.config.ErrorHandler(w, r, err)
	} else if next != nil {
		next.ServeHTTP(w, r)
	} else {
		m.NotFound(w, r)
	}
}

func (m *Mixer) Middleware(prefix string) web.MiddlewareHandlerFunc {

	// Wrap MiddlewareHandler
	fn := func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			m.MiddlewareHandler(w, r, next, prefix)
		}

		return http.HandlerFunc(fn)
	}
	return fn
}
