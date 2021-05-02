package types

import (
	"net/http"

	"go.sancus.dev/web"
)

type (
	ErrorHandlerFunc = web.ErrorHandlerFunc
)

type GetRoutePathFunc func(*http.Request) string
type SetRoutePathFunc func(*http.Request, string)
