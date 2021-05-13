package mixer

import (
	"log"
	"net/http"

	"go.sancus.dev/mix"
	"go.sancus.dev/mix/types"
	"go.sancus.dev/web"
	"go.sancus.dev/web/errors"
)

type MixPage struct {
	h []types.Handler
}

func (m *MixPage) TryServeHTTP(w http.ResponseWriter, r *http.Request) error {
	log.Printf("%s: %# v", "TryServeHTTP", m)
	return m.h[0].TryServeHTTP(w, r)
}

func NewMixPage(h []types.Handler) web.Handler {
	if l := len(h); l == 0 {
		return nil
	} else if l == 1 {
		return h[0]
	} else {
		h1 := make([]types.Handler, l)
		copy(h1, h)
		return &MixPage{h: h1}
	}
}

func (m *Router) getPage(rctx *mix.Context) (web.Handler, *mix.Context, bool) {

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
