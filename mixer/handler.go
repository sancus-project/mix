package mixer

import (
	"net/http"

	"go.sancus.dev/web/errors"
)

func (m *Router) tryServeHTTP(w http.ResponseWriter, r *http.Request) error {
	return errors.ErrNotFound
}

func (m *Router) TryServeHTTP(w http.ResponseWriter, r *http.Request) error {
	// Server-Timing
	if t := m.GetServerTiming(r, "TryServeHTTP"); t != nil {
		defer t.Start().Stop()
	}

	return m.tryServeHTTP(w, r)
}

func (m *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Server-Timing
	if t := m.GetServerTiming(r, "ServeHTTP"); t != nil {
		defer t.Start().Stop()
	}

	if err := m.tryServeHTTP(w, r); err != nil {
		m.mixer.config.ErrorHandler(w, r, err)
	}
}
