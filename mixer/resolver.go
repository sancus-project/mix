package mixer

import (
	"context"
	"log"
	"net/http"

	"go.sancus.dev/mix"
	"go.sancus.dev/web"
)

// Router.Resolve()
func (m *RouterNode) Resolve(rctx *mix.Context) (web.Handler, *mix.Context, bool) {

	path := rctx.RoutePath
	log.Printf("path:%q % #v", path, rctx)

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

func (m *router) resolve(rctx *mix.Context) (web.Handler, *mix.Context, bool) {
	if m != nil && m.top != nil {
		return m.top.Resolve(rctx)
	}

	return nil, nil, false
}

// Router.ResolvePath()
func (m *router) ResolvePath(ctx context.Context, prefix, path string) (web.Handler, *mix.Context, bool) {
	if rctx := mix.NewRouteContext(ctx, prefix, path); rctx != nil {
		return m.resolve(rctx)
	}

	return nil, nil, false
}

// Router.ResolveRequest()
func (m *router) ResolveRequest(ctx context.Context, r *http.Request) (web.Handler, *mix.Context, bool) {
	path := m.mixer.config.GetRoutePath(r)

	return m.ResolvePath(ctx, "/", path)
}
