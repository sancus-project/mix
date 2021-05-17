package mixer

import (
	"fmt"
	"net/http"

	"github.com/mitchellh/go-server-timing"
)

func (m *router) GetServerTiming(r *http.Request, name string) *servertiming.Metric {
	if len(m.ServerTimingPrefix) > 0 {
		if len(name) > 0 {
			name = fmt.Sprintf("%s.%s", m.ServerTimingPrefix, name)
		} else {
			name = m.ServerTimingPrefix
		}

		if t := servertiming.FromContext(r.Context()); t != nil {
			return t.NewMetric(name)
		}
	}
	return nil
}
