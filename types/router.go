package types

import (
	"net/http"

	"go.sancus.dev/web"
)

type Router interface {
	web.Handler

	Route(pattern string, fn func(r Router)) Router

	Mount(path string, h http.Handler) error

	Attach(h http.Handler) error
}
