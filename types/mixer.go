package types

import (
	"net/http"

	"go.sancus.dev/web"
)

type Mixer interface {
	Router

	Close() error
	Reload() error

	Middleware(prefix string) web.MiddlewareHandlerFunc
	Sitemap(prefix string) http.HandlerFunc
}
