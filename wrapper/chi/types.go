package chi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ChiHandler interface {
	http.Handler
	chi.Routes
}
