package adhoc

import (
	"net/http"

	"go.sancus.dev/web/intercept"
)

type AdhocWrapper struct {
	h http.Handler
}

func (h *AdhocWrapper) TryServeHTTP(w http.ResponseWriter, r *http.Request) error {
	f := intercept.Intercept(h.h)
	return f.TryServeHTTP(w, r)
}

func (h *AdhocWrapper) PageInfo(r *http.Request) (interface{}, bool) {
	w := &intercept.DummyWriter{}

	h.h.ServeHTTP(w, r)
	if err := w.Error(); err != nil {
		return err, false
	} else {
		return nil, true
	}
}
