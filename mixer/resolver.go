package mixer

import (
	"context"
	"log"
	"net/http"

	"go.sancus.dev/mix"
	"go.sancus.dev/web"
	"go.sancus.dev/web/errors"
)

func (m *Router) Resolve(rctx *mix.Context) (web.Handler, *mix.Context, bool) {

	path := rctx.RoutePath
	log.Printf("path:%q % #v", path, rctx)

	if path == "/" {
		if rctx.RoutePrefix == "/" {
			path = ""
		} else {
			h := errors.NewPermanentRedirect(rctx.RoutePrefix)
			return h, rctx, true
		}
	}

	if path == "" {
		// exact match

		m.mu.Lock()
		defer m.mu.Unlock()

		if len(m.handler) == 0 {
			goto fail
		} else {
			return NewMixPage(m.handler), rctx, true
		}

	}

fail:
	return nil, nil, false
}

func (m *Router) ResolvePath(ctx context.Context, prefix, path string) (web.Handler, *mix.Context, bool) {

	if rctx := mix.NewRouteContext(ctx, prefix, path); rctx != nil {
		return m.Resolve(rctx)
	}

	return nil, nil, false
}

func (m *Router) ResolveRequest(ctx context.Context, r *http.Request) (web.Handler, *mix.Context, bool) {

	path := m.mixer.config.GetRoutePath(r)

	return m.ResolvePath(ctx, "/", path)
}
