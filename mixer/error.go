package mixer

import (
	"net/http"

	"go.sancus.dev/web"
	"go.sancus.dev/web/errors"
)

func (m *Mixer) NotFound(w http.ResponseWriter, r *http.Request) {
	m.config.ErrorHandler(w, r, errors.ErrNotFound)
}

func (m *Mixer) SetErrorHandler(f web.ErrorHandlerFunc) error {
	if f == nil {
		f = errors.HandleError
	}
	m.config.ErrorHandler = f
	return nil
}

func SetErrorHandler(f web.ErrorHandlerFunc) MixerOption {
	return MixerOptionFunc(func(m *Mixer) error {
		return m.SetErrorHandler(f)
	})
}
