package registry

import (
	"errors"
	"net/http"
	"sort"

	"go.sancus.dev/mix/types"
)

var list []WrapperConstructor

func less(i, j int) bool {
	// inverted order
	return list[i].Priority() > list[j].Priority()
}

func RegisterWrapperConstructor(f WrapperConstructor) error {
	if f == nil {
		return errors.New("RegisterWrapperConstructor called without function")
	}
	list = append(list, f)
	sort.Slice(list, less)
	return nil
}

func NewWrapper(pattern string, h http.Handler) types.Handler {
	for _, t := range list {
		if r, ok := t.New(pattern, h); ok {
			return r
		}
	}
	return nil
}
