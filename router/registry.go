package router

import (
	"net/http"

	router "go.sancus.dev/mix/router/registry"
)

func NewRouter(pattern string, h http.Handler) router.Router {
	return router.NewRouter(pattern, h)
}
