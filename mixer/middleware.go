package mixer

import (
	"context"
	"net/http"
	"strings"

	"go.sancus.dev/mix"
	"go.sancus.dev/web"
	"go.sancus.dev/web/errors"
)

// Middleware
func (m *Router) middlewareTryServeHTTP(w http.ResponseWriter, r *http.Request, prefix string, path string) error {

	ctx := r.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if page, rctx, ok := m.ResolvePath(ctx, prefix, path); ok {

		ctx = mix.WithRouteContext(ctx, rctx)
		r = r.WithContext(ctx)

		return page.TryServeHTTP(w, r)
	}

	return errors.ErrNotFound
}

func (m *Router) MiddlewareHandler(w http.ResponseWriter, r *http.Request, next http.Handler, prefix string) {
	var err error

	path := m.mixer.config.GetRoutePath(r)

	if prefix == "/" || prefix == "" {
		// no prefix
		err = m.middlewareTryServeHTTP(w, r, "", path)

	} else if s := strings.TrimPrefix(path, prefix); s == path {
		// doesn't match prefix
		err = errors.ErrNotFound
	} else if s == "" || s[0] == '/' {
		// prefix match
		err = m.middlewareTryServeHTTP(w, r, prefix, s)
	} else {
		// doesn't match prefix
		err = errors.ErrNotFound
	}

	if next != nil {
		// skip 404
		if e, ok := err.(web.Error); ok && e.Status() == http.StatusNotFound {
			next.ServeHTTP(w, r)
			return
		}
	}

	// Failed
	if err != nil {
		m.mixer.config.ErrorHandler(w, r, err)
	}
}

func (m *Router) Middleware(prefix string) web.MiddlewareHandlerFunc {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if t := m.GetServerTiming(r, "Middleware"); t != nil {
				defer t.Start().Stop()
			}

			m.MiddlewareHandler(w, r, next, prefix)
		}

		return http.HandlerFunc(fn)
	}
}
