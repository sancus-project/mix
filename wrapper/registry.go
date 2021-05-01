package wrapper

import (
	"net/http"

	"go.sancus.dev/mix/types"
	"go.sancus.dev/mix/wrapper/registry"
)

func NewWrapper(pattern string, h http.Handler) types.Handler {
	return registry.NewWrapper(pattern, h)
}
