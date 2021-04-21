package mixer

import (
	"net/http"

	"go.sancus.dev/web/errors"
)

func (m *Mixer) Handle(w http.ResponseWriter, r *http.Request) error {
	return errors.ErrNotFound
}

func (m *Mixer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := m.Handle(w, r); err != nil {
		m.config.ErrorHandler(w, r, err)
	}
}
