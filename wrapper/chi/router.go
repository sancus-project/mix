package chi

import (
	"net/http"

	"go.sancus.dev/web/errors"
)

type ChiWrapper struct {
	h ChiHandler
}

func (h *ChiWrapper) TryServeHTTP(w http.ResponseWriter, r *http.Request) error {
	return errors.ErrNotFound
}
