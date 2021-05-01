package mixer

import (
	"errors"
	"net/http"
	"sync"

	"go.sancus.dev/mix/types"
	"go.sancus.dev/mix/wrapper"
)

type Mixer struct {
	Router

	mu sync.RWMutex

	// singleton
	wrapper map[http.Handler]types.Handler

	config MixerConfig
}

func NewMixer(options ...MixerOption) (types.Mixer, error) {

	// Initialise
	m := &Mixer{
		wrapper: make(map[http.Handler]types.Handler),
	}

	m.initRouter(&m.Router)

	// Configure
	for _, opt := range options {
		if err := opt.ApplyOption(m); err != nil {
			return nil, err
		}
	}

	// Finish
	if err := m.SetDefaults(); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Mixer) NewRouter() *Router {
	r := &Router{}
	m.initRouter(r)
	return r
}

func (m *Mixer) NewHandler(pattern string, h http.Handler) (types.Handler, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	r, ok := m.wrapper[h]
	if !ok {
		if r = wrapper.NewWrapper(pattern, h); r == nil {
			err := errors.New("Handler not supported")
			return nil, err
		}

		m.wrapper[h] = r
	}

	return r, nil
}

// Close
func (m *Mixer) Close() error {
	return nil
}

// Reload
func (m *Mixer) Reload() error {
	return nil
}
