package adhoc

import (
	"net/http"

	"go.sancus.dev/web/errors"
)

type AdhocWrapper struct {
	h http.Handler
}

func (h *AdhocWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.h.ServeHTTP(w, r)
}

func TryHandler(w0 http.ResponseWriter, r *http.Request, h http.Handler, err_out *error) {
	var err error

	defer func() {
		if err = errors.Recover(); err != nil {
			if err_out == nil {
				panic(err)
			} else {
				*err_out = err
			}
		}
	}()

	w := Writer{w: w0}
	h.ServeHTTP(&w, r)

	if err = w.Error(); err == nil {
		err = w.Flush()
	}

	if err_out != nil {
		*err_out = err
	} else if err != nil {
		panic(errors.Panic(err))
	}
}

func (h *AdhocWrapper) TryServeHTTP(w http.ResponseWriter, r *http.Request) error {
	var err error
	TryHandler(w, r, h.h, &err)
	return err
}
