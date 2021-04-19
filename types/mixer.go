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
}
