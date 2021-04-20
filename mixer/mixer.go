package mixer

import (
	"go.sancus.dev/mix/types"
)

type Mixer struct {
	Router

	config MixerConfig
}

func NewMixer(options ...MixerOption) (types.Mixer, error) {
	m := &Mixer{}

	for _, opt := range options {
		if err := opt.ApplyOption(m); err != nil {
			return nil, err
		}
	}

	if err := m.SetDefaults(); err != nil {
		return nil, err
	}

	return m, nil
}
