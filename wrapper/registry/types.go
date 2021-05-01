package registry

import (
	"net/http"

	"go.sancus.dev/mix/types"
)

type RouterConstructor interface {
	Priority() int                                  // Priority() defines Test order
	New(string, http.Handler) (types.Handler, bool) // Attempts to create a Router from a http.Handler
}
