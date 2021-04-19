package types

import (
	"net/http"
)

type Mixer interface {
	Router

	Mount(string, Router)

	Close() error
	Reload() error

	Middleware(prefix string) MiddlewareHandlerFunc
	Sitemap(prefix string) http.HandlerFunc
}
