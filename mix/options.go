package mix

import (
	"go.sancus.dev/mix/types"
)

// Config
type MixerConfig struct {
	ErrorHandler types.ErrorHandler
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
	m.SetErrorHandler(m.config.ErrorHandler)
	return nil
}
