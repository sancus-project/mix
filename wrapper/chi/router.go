package chi

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	sancus "go.sancus.dev/web/context"
	"go.sancus.dev/web/intercept"
)

type ChiWrapper struct {
	h ChiHandler
}

func (cw *ChiWrapper) TryServeHTTP(w http.ResponseWriter, r *http.Request) error {
	cc := RouteContext(r)
	ctx := WithRouteContext(r, cc)
	h := intercept.Intercept(cw.h)

	return h.TryServeHTTP(w, r.WithContext(ctx))
}

func (cw *ChiWrapper) PageInfo(r *http.Request) (interface{}, bool) {

	cc := RouteContext(r)
	found := cw.h.Match(cc, r.Method, cc.RoutePath)

	return nil, found
}

func RouteContext(r *http.Request) *chi.Context {
	rctx := sancus.RouteContext(r.Context())

	return &chi.Context{
		RoutePath:   rctx.RoutePath,
		RouteMethod: r.Method,
	}
}

func WithRouteContext(r *http.Request, cc *chi.Context) context.Context {
	return context.WithValue(r.Context(), chi.RouteCtxKey, cc)
}
