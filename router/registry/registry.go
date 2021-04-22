package registry

import (
	"errors"
	"net/http"
	"sort"

	"go.sancus.dev/mix/types"
)

var list []types.RouterConstructor

func less(i, j int) bool {
	// inverted order
	return list[i].Priority() > list[j].Priority()
}

func RegisterRouterConstructor(f types.RouterConstructor) error {
	if f == nil {
		return errors.New("RegisterRouterConstructor called without function")
	}
	list = append(list, f)
	sort.Slice(list, less)
	return nil
}

func NewRouter(pattern string, h http.Handler) types.Router {
	for _, t := range list {
		if r, ok := t.New(pattern, h); ok {
			return r
		}
	}
	return nil
}
