package registry

import (
	"errors"
	"sort"

	"go.sancus.dev/mix/types"
)

var list []types.RouterConstructor

func less(i, j int) bool {
	return list[i].Priority() < list[j].Priority()
}

func RegisterRouterConstructor(f types.RouterConstructor) error {
	if f == nil {
		return errors.New("RegisterRouterConstructor called without function")
	}
	list = append(list, f)
	sort.Slice(list, less)
	return nil
}
