package mixer

import (
	"go.sancus.dev/mix/types"
	"go.sancus.dev/web"
)

// Config
type MixerConfig struct {
	GetRoutePath types.GetRoutePathFunc // Extracts RoutePath from http.Request
	SetRoutePath types.SetRoutePathFunc // Sets RoutePath to http.Request
	ErrorHandler web.ErrorHandlerFunc   // Error Handler

	ServerTiming string // Server-Timing metric name
}

// Options
type MixerOption interface {
	ApplyOption(m *Mixer) error
}

type MixerOptionFunc func(*Mixer) error

func (f MixerOptionFunc) ApplyOption(m *Mixer) error {
	return f(m)
}

// GetRoutePath
func (m *Mixer) SetGetRoutePath(f types.GetRoutePathFunc) error {
	if f == nil {
		f = DefaultGetRoutePath
	}
	m.config.GetRoutePath = f
	return nil
}

func SetGetRoutePath(f types.GetRoutePathFunc) MixerOption {
	return MixerOptionFunc(func(m *Mixer) error {
		return m.SetGetRoutePath(f)
	})
}

// SetRoutePath
func (m *Mixer) SetSetRoutePath(f types.SetRoutePathFunc) error {
	if f == nil {
		f = DefaultSetRoutePath
	}
	m.config.SetRoutePath = f
	return nil
}

func SetSetRoutePath(f types.SetRoutePathFunc) MixerOption {
	return MixerOptionFunc(func(m *Mixer) error {
		return m.SetSetRoutePath(f)
	})
}

// ErrorHandler
func (m *Mixer) SetErrorHandler(f types.ErrorHandlerFunc) error {
	if f == nil {
		f = DefaultErrorHandler
	}
	m.config.ErrorHandler = f
	return nil
}

func SetErrorHandler(f web.ErrorHandlerFunc) MixerOption {
	return MixerOptionFunc(func(m *Mixer) error {
		return m.SetErrorHandler(f)
	})
}

// ServerTiming
func (m *Mixer) SetServerTiming(name string) error {
	m.config.ServerTiming = name
	return nil
}

func SetServerTiming(name string) MixerOption {
	return MixerOptionFunc(func(m *Mixer) error {
		return m.SetServerTiming(name)
	})
}

// Defaults
func (m *Mixer) SetDefaults() error {
	m.SetGetRoutePath(m.config.GetRoutePath)
	m.SetSetRoutePath(m.config.SetRoutePath)
	m.SetErrorHandler(m.config.ErrorHandler)
	return nil
}
