package tree

import (
	"go.sancus.dev/mix/mixer/segment"
)

type Path struct {
	s []segment.Segment
}

func Compile(pattern string) (*Path, error) {
	s, ok := segment.Compile(pattern)
	if !ok {
		return nil, InvalidPattern(pattern)
	}

	p := &Path{
		s: s,
	}
	return p, nil
}
