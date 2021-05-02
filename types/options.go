package types

import (
	"net/http"
)

type GetRoutePathFunc func(*http.Request) string
type SetRoutePathFunc func(*http.Request, string)
