package mixer

import (
	"net/http"

	"go.sancus.dev/web/errors"
)

func (m *Mixer) NotFound(w http.ResponseWriter, r *http.Request) {
	m.config.ErrorHandler(w, r, errors.ErrNotFound)
}
