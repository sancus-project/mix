package chi

import (
	"log"
	"net/http"

	"go.sancus.dev/mix/types"
	"go.sancus.dev/mix/wrapper/registry"
)

const ConstructorPriority = 5

type ChiRouterConstructor struct{}

func (_ ChiRouterConstructor) Priority() int {
	return ConstructorPriority
}

func (f *ChiRouterConstructor) New(pattern string, h http.Handler) (types.Handler, bool) {
	if v, ok := h.(ChiHandler); ok {
		r := &ChiRouter{h: v}
		return r, true
	}
	return nil, false
}

func init() {
	f := &ChiRouterConstructor{}
	if err := registry.RegisterRouterConstructor(f); err != nil {
		log.Fatal(err)
	}
}
