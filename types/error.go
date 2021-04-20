package types

import (
	"net/http"
)

type ErrorHandler func(http.ResponseWriter, *http.Request, error)

// Error including HTTP Status
type Error interface {
	Error() string
	Status() int
}
