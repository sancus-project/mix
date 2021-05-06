package types

import (
	"net/http"

	"go.sancus.dev/web"
)

type Handler interface {
	web.Handler
}

type Router interface {
	http.Handler
	Handler

	Route(pattern string, fn func(r Router)) Router
	Attach(h interface{}) error

	Mount(path string, h interface{}) error
}

type Mixer interface {
	Router

	Close() error
	Reload() error

	Middleware(prefix string) web.MiddlewareHandlerFunc
	Sitemap(prefix string) http.HandlerFunc
}
