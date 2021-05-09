package mixer

import (
	"context"
	"net/http"
	"path/filepath"
	"strings"

	"go.sancus.dev/mix"
	"go.sancus.dev/web"
	"go.sancus.dev/web/errors"
)

func (m *Router) NewContext(r *http.Request, prefix, path string) context.Context {

	ctx := r.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	// New routing Context object
	rctx := mix.RouteContext(ctx)
	if rctx == nil {
		rctx = mix.NewRouteContext()
	} else {
		rctx = rctx.Clone()
	}

	if rctx.RoutePath == "" {
		// New routing Context
		var pattern string

		if prefix == "" {
			prefix = "/"
		}

		if path == "" {
			path, pattern = prefix, prefix
		} else if n := strings.IndexRune(path[1:], '/'); n < 0 {
			pattern = filepath.Join(prefix, path)
		} else {
			pattern = filepath.Join(prefix, path[0:n+1]) + "/*"
		}

		rctx.RoutePrefix = prefix
		rctx.RoutePattern = pattern
		rctx.RoutePath = path
	} else {
		// Update
		panic(errors.New("Nested Mixer not implemented"))
	}

	return mix.WithRouteContext(ctx, rctx)
}

// web.RouterPageInfo
func (m *Router) pageinfo(r *http.Request) (web.Handler, bool) {
	return nil, false
}

func (m *Router) PageInfo(r *http.Request) (interface{}, bool) {
	if t := m.GetServerTiming(r, "PageInfo"); t != nil {
		defer t.Start().Stop()
	}

	return m.pageinfo(r)
}

// web.Handler
func (m *Router) tryServeHTTP(w http.ResponseWriter, r *http.Request) error {

	path := m.mixer.config.GetRoutePath(r)

	// New http.Request Context including our routing Context inside
	ctx := m.NewContext(r, "", path)

	// And new http.Request with it
	r = r.WithContext(ctx)

	if page, ok := m.pageinfo(r); ok {
		return page.TryServeHTTP(w, r)
	}

	return errors.ErrNotFound
}

func (m *Router) TryServeHTTP(w http.ResponseWriter, r *http.Request) error {
	// Server-Timing
	if t := m.GetServerTiming(r, "TryServeHTTP"); t != nil {
		defer t.Start().Stop()
	}

	return m.tryServeHTTP(w, r)
}

// http.Handler
func (m *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Server-Timing
	if t := m.GetServerTiming(r, "ServeHTTP"); t != nil {
		defer t.Start().Stop()
	}

	if err := m.tryServeHTTP(w, r); err != nil {
		m.mixer.config.ErrorHandler(w, r, err)
	}
}
