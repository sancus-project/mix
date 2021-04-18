package mix

import (
	"net/http"

	"go.sancus.dev/mix/types"
)

func (m *Mixer) Handler(w http.ResponseWriter, r *http.Request) error {
	return types.ErrNotFound
}

func (m *Mixer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := m.Handler(w, r); err != nil {
		m.HandleError(w, r, err)
	}
}

func (m *Mixer) HandleError(w http.ResponseWriter, r *http.Request, err error) {
	panic(err)
}
