package types

import (
	"net/http"
)

type Mixer interface {
	http.Handler

	Mount(string, http.Handler)

	Close() error
	Reload() error

	Middleware(prefix string) MiddlewareHandlerFunc
	Sitemap(prefix string) http.HandlerFunc
}
