package mixer

import (
	"net/http"

	"go.sancus.dev/web/errors"
)

func (m *Mixer) Sitemap(prefix string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := m.SitemapHandler(w, r, prefix); err != nil {
			m.config.ErrorHandler(w, r, err)
		}
	}
}

func (m *Mixer) SitemapHandler(w http.ResponseWriter, r *http.Request, prefix string) error {
	return errors.ErrNotFound
}
