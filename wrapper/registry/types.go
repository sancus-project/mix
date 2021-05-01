package registry

import (
	"net/http"

	"go.sancus.dev/mix/types"
)

type WrapperConstructor interface {
	Priority() int                                  // Priority() defines Test order
	New(string, http.Handler) (types.Handler, bool) // Attempts to create a Wrapper from a http.Handler
}
