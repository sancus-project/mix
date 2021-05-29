package sancus

import (
	"net/http"

	"go.sancus.dev/web"
	"go.sancus.dev/web/intercept"
)

type Wrapper struct {
	web.Handler
}

func (h *Wrapper) PageInfo(r *http.Request) (interface{}, bool) {
	w := &intercept.DummyWriter{}

	if err := h.TryServeHTTP(w, r); err != nil {
		return err, false
	} else if err := w.Error(); err != nil {
		return err, false
	} else {
		return h, true
	}
}
