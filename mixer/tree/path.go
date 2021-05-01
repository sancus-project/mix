package tree

import (
	"go.sancus.dev/mix/mixer/segment"
)

func Compile(pattern string) (Path, error) {
	s, ok := segment.Compile(pattern)
	if !ok {
		return nil, InvalidPattern(pattern)
	}

	return Path(s), nil
}
