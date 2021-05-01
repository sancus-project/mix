package chi

import (
	"log"
	"net/http"

	"go.sancus.dev/mix/types"
	"go.sancus.dev/mix/wrapper/registry"
)

const ConstructorPriority = 5

type ChiWrapperConstructor struct{}

func (_ ChiWrapperConstructor) Priority() int {
	return ConstructorPriority
}

func (f *ChiWrapperConstructor) New(pattern string, h http.Handler) (types.Handler, bool) {
	if v, ok := h.(ChiHandler); ok {
		r := &ChiWrapper{h: v}
		return r, true
	}
	return nil, false
}

func init() {
	f := &ChiWrapperConstructor{}
	if err := registry.RegisterWrapperConstructor(f); err != nil {
		log.Fatal(err)
	}
}
