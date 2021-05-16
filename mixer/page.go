package mixer

import (
	"log"
	"net/http"

	"go.sancus.dev/mix/types"
	"go.sancus.dev/web"
)

type MixPage struct {
	h []types.Handler
}

func (m *MixPage) TryServeHTTP(w http.ResponseWriter, r *http.Request) error {
	log.Printf("%s: %# v", "TryServeHTTP", m)
	return m.h[0].TryServeHTTP(w, r)
}

func NewMixPage(h []types.Handler) web.Handler {
	if l := len(h); l == 0 {
		return nil
	} else if l == 1 {
		return h[0]
	} else {
		h1 := make([]types.Handler, l)
		copy(h1, h)
		return &MixPage{h: h1}
	}
}
