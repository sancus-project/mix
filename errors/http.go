package errors

import (
	"fmt"
	"net/http"
)

var (
	ErrNotFound = &HandlerError{Code: http.StatusNotFound}
)

// Reference Handler error
type HandlerError struct {
	Code int
	Err  error
}

func (err HandlerError) Status() int {
	var code int

	if err.Code != 0 {
		code = err.Code
	} else if err.Err == nil {
		code = http.StatusOK
	} else {
		code = http.StatusInternalServerError
	}

	return code
}

func (err HandlerError) Unwrap() error {
	return err.Err
}

func ErrorText(code int) string {
	text := http.StatusText(code)

	if len(text) == 0 {
		text = fmt.Sprintf("Unknown Error %d", code)
	} else if code >= 400 {
		text = fmt.Sprintf("%s (Error %d)", text, code)
	}

	return text
}

func (err HandlerError) String() string {
	return ErrorText(err.Status())
}

func (err HandlerError) Error() string {
	return ErrorText(err.Status())
}

func (err HandlerError) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	if code := err.Status(); code == http.StatusOK {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(code)

		fmt.Fprintln(w, ErrorText(code))
	}
}
