package adhoc

import (
	"log"
	"net/http"

	"go.sancus.dev/mix/types"
	"go.sancus.dev/mix/wrapper/registry"
)

const ConstructorPriority = -1

type AdhocWrapperConstructor struct{}

func (_ AdhocWrapperConstructor) Priority() int {
	return ConstructorPriority
}

func (f *AdhocWrapperConstructor) New(pattern string, h interface{}) (types.Handler, bool) {
	if v, ok := h.(http.Handler); ok {
		r := &AdhocWrapper{h: v}
		return r, true
	}
	return nil, false
}

func init() {
	f := &AdhocWrapperConstructor{}
	if err := registry.RegisterWrapperConstructor(f); err != nil {
		log.Fatal(err)
	}
}
