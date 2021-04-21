package types

import (
	"net/http"

	"go.sancus.dev/web"
)

type RouterConstructor interface {
	Priority() int                           // Priority() defines Test order
	New(string, http.Handler) (Router, bool) // Attempts to create a Router
}

type Router interface {
	web.Handler
}
