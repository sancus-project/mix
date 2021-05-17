package mixer

import (
	"errors"
	"sync"

	"go.sancus.dev/mix/types"
	"go.sancus.dev/mix/wrapper"
)

type Mixer struct {
	RouterNode

	mu sync.RWMutex

	// singleton
	wrapper map[interface{}]types.Handler

	config      MixerConfig
	routerCount int
}

func NewMixer(options ...MixerOption) (types.Mixer, error) {

	// Initialise
	m := &Mixer{
		wrapper: make(map[interface{}]types.Handler),
	}

	m.initRouter(&m.RouterNode)

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

func (m *Mixer) newHandler(pattern string, h interface{}) (types.Handler, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	r, ok := m.wrapper[h]
	if ok {
		// known
	} else if r, ok = wrapper.NewWrapper(pattern, h); !ok {
		// unsupported
		err := errors.New("Handler not supported")
		return nil, err
	} else {
		// remember
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
