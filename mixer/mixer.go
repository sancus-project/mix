package mixer

import (
	"errors"
	"net/http"
	"sync"

	"go.sancus.dev/mix/tree"
	"go.sancus.dev/mix/types"
	"go.sancus.dev/mix/wrapper"
)

type Mixer struct {
	sync.RWMutex
	Router

	// singleton
	wrapper map[http.Handler]types.Handler

	config MixerConfig
}

func NewMixer(options ...MixerOption) (types.Mixer, error) {

	// Initialise
	m := &Mixer{
		wrapper: make(map[http.Handler]types.Handler),
	}

	m.Router.mixer = m

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

// Preprocesses pattern and handler for Router.Route()/Router.Mount()
func (m *Mixer) NewHandler(pattern string, h http.Handler) (*tree.Path, types.Handler, error) {
	m.Lock()
	defer m.Unlock()

	p, err := tree.Compile(pattern)
	if err != nil {
		return nil, nil, err
	}

	r, ok := m.wrapper[h]
	if !ok {
		if r = wrapper.NewWrapper(pattern, h); r == nil {
			err = errors.New("Handler not supported")
			return nil, nil, err
		}

		m.wrapper[h] = r
	}

	return p, r, nil
}

// Close
func (m *Mixer) Close() error {
	return nil
}

// Reload
func (m *Mixer) Reload() error {
	return nil
}
