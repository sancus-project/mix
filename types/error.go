package types

import (
	"net/http"
)

type ErrorHandler func(http.ResponseWriter, *http.Request, error)

// Error including HTTP Status and an optional wrapped error
type Error interface {
	Error() string
	Status() int
	Unwrap() error
}
