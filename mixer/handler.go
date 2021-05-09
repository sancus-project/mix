package mixer

import (
	"net/http"

	"go.sancus.dev/web/errors"
)

// web.RouterPageInfo
func (m *Router) pageinfo(r *http.Request) (interface{}, bool) {
	return nil, false
}

func (m *Router) PageInfo(r *http.Request) (interface{}, bool) {
	if t := m.GetServerTiming(r, "PageInfo"); t != nil {
		defer t.Start().Stop()
	}

	return m.pageinfo(r)
}

// web.Handler
func (m *Router) tryServeHTTP(w http.ResponseWriter, r *http.Request) error {
	_, ok := m.pageinfo(r)
	if !ok {
		return errors.ErrNotFound
	}

	return nil
}

func (m *Router) TryServeHTTP(w http.ResponseWriter, r *http.Request) error {
	// Server-Timing
	if t := m.GetServerTiming(r, "TryServeHTTP"); t != nil {
		defer t.Start().Stop()
	}

	return m.tryServeHTTP(w, r)
}

// http.Handler
func (m *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Server-Timing
	if t := m.GetServerTiming(r, "ServeHTTP"); t != nil {
		defer t.Start().Stop()
	}

	if err := m.tryServeHTTP(w, r); err != nil {
		m.mixer.config.ErrorHandler(w, r, err)
	}
}
