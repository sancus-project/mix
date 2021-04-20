package mixer

import (
	"net/http"

	"go.sancus.dev/mix/errors"
)

func (m *Mixer) Handler(w http.ResponseWriter, r *http.Request) error {
	return errors.ErrNotFound
}

func (m *Mixer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := m.Handler(w, r); err != nil {
		m.config.ErrorHandler(w, r, err)
	}
}
