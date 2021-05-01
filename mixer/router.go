package mixer

import (
	"errors"
	"net/http"
	"strings"
)

type Router struct {
	mixer *Mixer
}

// Mounts handler at path
func (m *Router) Mount(path string, h http.Handler) error {
	var pattern string

	if path == "" {
		pattern = "/*"
	} else if strings.HasSuffix(path, "/") {
		pattern = path + "*"
	} else {
		pattern = path
	}

	_, _, err := m.mixer.NewHandler(pattern, h)
	if err != nil {
		return err
	}

	return errors.New("Not Implemented")
}
