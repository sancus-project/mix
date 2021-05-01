package registry

import (
	"net/http"

	"go.sancus.dev/web"
)

type RouterConstructor interface {
	Priority() int                           // Priority() defines Test order
	New(string, http.Handler) (Router, bool) // Attempts to create a Router from a http.Handler
}

type Router interface {
	web.Handler
}
