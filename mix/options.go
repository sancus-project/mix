package mix

import (
	"net/http"
)

type MixerConfig struct {
	ErrorHandler func(w http.ResponseWriter, r *http.Request, err error)
}

type MixerOption interface {
	ApplyOption(m *Mixer) error
}

func (m *Mixer) SetDefaults() error {
	cfg := &m.config
	if cfg.ErrorHandler == nil {
		cfg.ErrorHandler = DefaultErrorHandler
	}
	return nil
}
