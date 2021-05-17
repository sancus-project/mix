package mixer

import (
	"context"
	"net/http"

	"go.sancus.dev/mix"
	"go.sancus.dev/web"
	"go.sancus.dev/web/errors"
)

// web.RouterPageInfo
func (m *router) PageInfo(r *http.Request) (interface{}, bool) {
	// Server-Timing
	if t := m.GetServerTiming(r, "PageInfo"); t != nil {
		defer t.Start().Stop()
	}

	ctx := r.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if page, rctx, ok := m.ResolveRequest(ctx, r); !ok {
		return nil, false
	} else if h, ok := page.(web.RouterPageInfo); ok {

		ctx = mix.WithRouteContext(ctx, rctx)
		r = r.WithContext(ctx)

		return h.PageInfo(r)
	} else {
		return page, true
	}
}

// web.Handler
func (m *router) tryServeHTTP(w http.ResponseWriter, r *http.Request) error {

	ctx := r.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if page, rctx, ok := m.ResolveRequest(ctx, r); ok {

		ctx = mix.WithRouteContext(ctx, rctx)
		r = r.WithContext(ctx)

		return page.TryServeHTTP(w, r)
	}

	return errors.ErrNotFound
}

func (m *router) TryServeHTTP(w http.ResponseWriter, r *http.Request) error {
	// Server-Timing
	if t := m.GetServerTiming(r, "TryServeHTTP"); t != nil {
		defer t.Start().Stop()
	}

	return m.tryServeHTTP(w, r)
}

// http.Handler
func (m *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Server-Timing
	if t := m.GetServerTiming(r, "ServeHTTP"); t != nil {
		defer t.Start().Stop()
	}

	if err := m.tryServeHTTP(w, r); err != nil {
		m.mixer.config.ErrorHandler(w, r, err)
	}
}
