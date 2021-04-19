package types

import (
	"net/http"
)

type Router interface {
	http.Handler
}

type Mixer interface {
	Router
	Mount(string, Router)

	Middleware(prefix string) MiddlewareHandlerFunc
	Sitemap(prefix string) http.HandlerFunc
}
