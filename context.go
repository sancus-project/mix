package mix

import (
	"context"

	"go.sancus.dev/mix/types"
)

type (
	Context = types.Context
)

// RouteContext returns mix's routing Context object from a
// http.Request Context.
func RouteContext(ctx context.Context) *Context {
	if rctx, ok := ctx.Value(RouteCtxKey).(*Context); ok {
		return rctx
	}
	return nil
}

// NewRouteContext returns a new routing Context object.
func NewRouteContext() *Context {
	return &Context{}
}

// WithRouteContext returns a new http.Request Context with
// a given mix routing Context object connected to it, so it
// can later be extracted using RouteContext()
func WithRouteContext(ctx context.Context, rctx *Context) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	if rctx == nil {
		rctx = NewRouteContext()
	}
	return context.WithValue(ctx, RouteCtxKey, rctx)
}

var (
	// RouteCtxKey is the context.Context key to store the request context.
	RouteCtxKey = &contextKey{"RouteContext"}
)

// contextKey is a value for use with context.WithValue
type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "mix context value " + k.name
}
