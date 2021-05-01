package registry

import (
	"errors"
	"net/http"
	"sort"
)

var list []RouterConstructor

func less(i, j int) bool {
	// inverted order
	return list[i].Priority() > list[j].Priority()
}

func RegisterRouterConstructor(f RouterConstructor) error {
	if f == nil {
		return errors.New("RegisterRouterConstructor called without function")
	}
	list = append(list, f)
	sort.Slice(list, less)
	return nil
}

func NewRouter(pattern string, h http.Handler) Router {
	for _, t := range list {
		if r, ok := t.New(pattern, h); ok {
			return r
		}
	}
	return nil
}
