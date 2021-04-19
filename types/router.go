package types

import (
	"net/http"
)

type Router interface {
	http.Handler
	Routes
}
