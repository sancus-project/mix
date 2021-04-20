package types

import (
	"net/http"
)

type RouterConstructor interface {
	Priority() int                           // Priority() defines Test order
	New(string, http.Handler) (Router, bool) // Attempts to create a Router
}

type Router interface {
	http.Handler
	Routes
}
