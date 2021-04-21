package mixer

import (
	"go.sancus.dev/web"
)

// Config
type MixerConfig struct {
	ErrorHandler web.ErrorHandlerFunc
}

// Options
type MixerOption interface {
	ApplyOption(m *Mixer) error
}

type MixerOptionApplier func(*Mixer) error

func MixerOptionFunc(f MixerOptionApplier) MixerOption {
	return &mixerOption{apply: f}
}

type mixerOption struct {
	apply MixerOptionApplier
}

func (opt mixerOption) ApplyOption(m *Mixer) error {
	return opt.apply(m)
}

// Defaults
func (m *Mixer) SetDefaults() error {
	m.Router.mixer = m

	m.SetErrorHandler(m.config.ErrorHandler)
	return nil
}
