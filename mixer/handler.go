package mixer

import (
	"net/http"

	"go.sancus.dev/web/errors"
)

func (m *Mixer) TryServeHTTP(w http.ResponseWriter, r *http.Request) error {
	return errors.ErrNotFound
}

func (m *Mixer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := m.TryServeHTTP(w, r); err != nil {
		m.config.ErrorHandler(w, r, err)
	}
}
