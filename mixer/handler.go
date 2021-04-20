package mix

import (
	"net/http"

	"go.sancus.dev/mix/types"
)

func (m *Mixer) Handler(w http.ResponseWriter, r *http.Request) error {
	return types.ErrNotFound
}

func (m *Mixer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := m.Handler(w, r); err != nil {
		m.config.ErrorHandler(w, r, err)
	}
}
