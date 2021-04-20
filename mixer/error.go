package mixer

import (
	"net/http"

	"go.sancus.dev/mix/errors"
	"go.sancus.dev/mix/types"
)

func DefaultErrorHandler(w http.ResponseWriter, r *http.Request, err error) {

	// does the error know how to render itself?
	h, ok := err.(http.Handler)
	if !ok {
		var code int

		// but if it doesn't, wrap it in HandlerError{}
		if e, ok := err.(types.Error); ok {
			code = e.Status()
		} else {
			code = http.StatusInternalServerError
		}

		h = &errors.HandlerError{
			Code: code,
			Err:  err,
		}
	}

	h.ServeHTTP(w, r)
}

func (m *Mixer) NotFound(w http.ResponseWriter, r *http.Request) {
	m.config.ErrorHandler(w, r, errors.ErrNotFound)
}

func (m *Mixer) SetErrorHandler(f types.ErrorHandler) error {
	if f == nil {
		f = DefaultErrorHandler
	}
	m.config.ErrorHandler = f
	return nil
}

func SetErrorHandler(f types.ErrorHandler) MixerOption {
	return MixerOptionFunc(func(m *Mixer) error {
		return m.SetErrorHandler(f)
	})
}
