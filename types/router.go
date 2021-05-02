package types

import (
	"net/http"
)

type Router interface {
	http.Handler
	Handler

	Route(pattern string, fn func(r Router)) Router
	Attach(h interface{}) error

	Mount(path string, h interface{}) error
}
