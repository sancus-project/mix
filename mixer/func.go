package mixer

import (
	"net/http"
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
