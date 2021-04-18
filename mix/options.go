package mix

type MixerConfig struct{}

type MixerOption interface {
	ApplyOption(m *Mixer) error
}

func (m *Mixer) SetDefaults() error {
	return nil
}
