package mixer

import (
	"log"
	"net/http"
	"strings"
	"sync"

	"go.sancus.dev/mix/mixer/tree"
	"go.sancus.dev/mix/types"
)

type Router struct {
	mixer *Mixer
	mu    sync.Mutex

	handler []types.Handler
}

// Gets Router for a given path pattern
func (m *Router) Route(pattern string, fn func(r types.Router)) types.Router {
	p, err := tree.Compile(pattern)
	if err != nil {
		log.Fatal(err)
	}

	r := m.route(p)
	if r == nil {
		return nil
	}

	if fn != nil {
		fn(r)
	}

	return r
}

func (m *Router) route(p *tree.Path) *Router {
	return nil
}

// Mounts handler at path
func (m *Router) Mount(path string, h http.Handler) error {
	var pattern string

	if path == "" {
		pattern = "/*"
	} else if strings.HasSuffix(path, "/") {
		pattern = path + "*"
	} else {
		pattern = path
	}

	p, err := tree.Compile(pattern)
	if err != nil {
		return err
	}

	w, err := m.mixer.NewHandler(pattern, h)
	if err != nil {
		return err
	}

	return m.route(p).attach(w)
}

// Attach handler to Router
func (m *Router) Attach(h http.Handler) error {
	w, err := m.mixer.NewHandler("/", h)
	if err != nil {
		return err
	}

	return m.attach(w)
}

func (m *Router) attach(w types.Handler) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.handler = append(m.handler, w)
	return nil
}
