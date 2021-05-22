package mix

import (
	"context"

	mctx "go.sancus.dev/web/context"
)

type Context = mctx.Context

func RouteContext(ctx context.Context) *mctx.Context {
	return mctx.RouteContext(ctx)
}

func WithRouteContext(ctx context.Context, rctx *mctx.Context) context.Context {
	return mctx.WithRouteContext(ctx, rctx)
}

func NewRouteContext(ctx context.Context, prefix, path string) *mctx.Context {
	return mctx.NewRouteContext(prefix, path)
}
