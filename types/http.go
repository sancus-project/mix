package types

import (
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request) error

type MiddlewareHandlerFunc func(http.Handler) http.Handler

type Handler interface {
	Handle(http.ResponseWriter, *http.Request) error
}
