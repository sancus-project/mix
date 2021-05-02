package mixer

import (
	"net/http"

	"go.sancus.dev/web/errors"
)

// types.GetRoutePathFunc
func DefaultGetRoutePath(r *http.Request) string {
	return r.URL.Path
}

// types.SetRoutePathFunc
func DefaultSetRoutePath(r *http.Request, path string) {
	if path == "" {
		path = "/"
	}

	r.URL.Path = path
}

// types.ErrorHandlerFunc
var DefaultErrorHandler = errors.HandleError
