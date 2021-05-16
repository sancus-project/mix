package types

import (
	"net/http"

	"go.sancus.dev/web"
)

type Handler interface {
	web.Handler
	web.RouterPageInfo
}

type Router interface {
	http.Handler
	Handler

	Middleware(prefix string) web.MiddlewareHandlerFunc

	Route(pattern string, fn func(r Router)) Router
	Attach(h interface{}) error

	Mount(path string, h interface{}) error

	Resolve(*Context) (web.Handler, *Context, bool)
}

type Mixer interface {
	Router

	Close() error
	Reload() error

	Sitemap(prefix string) http.HandlerFunc
}
