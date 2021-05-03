package adhoc

import (
	"bytes"
	"net/http"

	"go.sancus.dev/web/errors"
)

type Writer struct {
	buf    bytes.Buffer
	w      http.ResponseWriter
	code   int
	header http.Header
}

func (w *Writer) Header() http.Header {
	if w.header != nil {

	} else if w.w == nil {
		w.header = make(map[string][]string)
	} else {
		w.header = w.w.Header()
	}
	return w.header
}

func (w *Writer) Write(b []byte) (int, error) {
	return w.buf.Write(b)
}

func (w *Writer) WriteHeader(statusCode int) {
	w.code = statusCode
}

func (w *Writer) Status() int {
	if w.code == 0 {
		return http.StatusOK
	} else if w.code < 100 {
		return http.StatusInternalServerError
	} else {
		return w.code
	}
}

func (w *Writer) Error() error {
	code := w.Status()

	if code < http.StatusBadRequest {
		return nil // OK
	}

	// Error
	return &errors.HandlerError{Code: code}
}

func (w *Writer) Flush() error {
	if w.w != nil {
		w.w.WriteHeader(w.Status())

		if _, err := w.buf.WriteTo(w.w); err != nil {
			return err
		}
	}
	return nil
}
