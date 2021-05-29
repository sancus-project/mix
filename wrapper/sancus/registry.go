package sancus

import (
	"log"

	"go.sancus.dev/mix/types"
	"go.sancus.dev/mix/wrapper/registry"
)

const ConstructorPriority = 9

type WrapperConstructor struct{}

func (_ WrapperConstructor) Priority() int {
	return ConstructorPriority
}

func (f *WrapperConstructor) New(pattern string, h interface{}) (types.Handler, bool) {
	if v, ok := h.(Handler); ok {
		return v, true
	}
	return nil, false
}

func init() {
	f := &WrapperConstructor{}
	if err := registry.RegisterWrapperConstructor(f); err != nil {
		log.Fatal(err)
	}
}
