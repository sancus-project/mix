package types

import (
	"net/http"

	"go.sancus.dev/web"
)

type Router interface {
	web.Handler

	Mount(pattern string, h http.Handler) error
}
