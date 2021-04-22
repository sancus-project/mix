package router

import (
	"net/http"

	"go.sancus.dev/mix/router/registry"

	"go.sancus.dev/mix/types"
)

func NewRouter(pattern string, h http.Handler) types.Router {
	return registry.NewRouter(pattern, h)
}
