package mix

import (
	"go.sancus.dev/mix/types"
)

// Mount
func (m *Mixer) Mount(pattern string, router types.Router) {
}

// Close
func (m *Mixer) Close() error {
	return nil
}

// Reload
func (m *Mixer) Reload() error {
	return nil
}
