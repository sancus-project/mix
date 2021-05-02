package wrapper

import (
	"go.sancus.dev/mix/types"
	"go.sancus.dev/mix/wrapper/registry"
)

func NewWrapper(pattern string, h interface{}) (types.Handler, bool) {
	return registry.NewWrapper(pattern, h)
}
