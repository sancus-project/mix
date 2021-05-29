package sancus

import (
	"log"

	"go.sancus.dev/mix/types"
	"go.sancus.dev/mix/wrapper/registry"
	"go.sancus.dev/web"
)

const ConstructorPriority = 9

type WrapperConstructor struct{}

func (_ WrapperConstructor) Priority() int {
	return ConstructorPriority
}

func (f *WrapperConstructor) New(pattern string, h interface{}) (types.Handler, bool) {
	if v, ok := h.(types.Handler); ok {
		return v, true
	} else if v, ok := h.(web.Handler); ok {
		r := &Wrapper{Handler: v}
		return r, true
	}
	return nil, false
}

func init() {
	f := &WrapperConstructor{}
	if err := registry.RegisterWrapperConstructor(f); err != nil {
		log.Fatal(err)
	}
}
