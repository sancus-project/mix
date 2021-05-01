package chi

import (
	"net/http"

	"go.sancus.dev/web/errors"
)

type ChiRouter struct {
	h ChiHandler
}

func (h *ChiRouter) TryServeHTTP(w http.ResponseWriter, r *http.Request) error {
	return errors.ErrNotFound
}
