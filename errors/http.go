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

func (err HandlerError) String() string {
	code := err.Status()
	text := http.StatusText(code)

	if len(text) == 0 {
		text = fmt.Sprintf("Unknown Error %d", code)
	} else if code >= 400 {
		text = fmt.Sprintf("%s (Error %d)", text, code)
	}

	return text
}

func (err HandlerError) Error() string {
	return err.String()
}

func (err HandlerError) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(err.Status())

	fmt.Fprintln(w, err)
}
