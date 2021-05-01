package segment

import (
	"strings"
)

func compile(chunks ...string) ([]Segment, bool) {
	var list []Segment

	last := len(chunks) - 1

	for i, chunk := range chunks {

		// empty chunks only if it's the last
		if len(chunk) > 0 {

			s, ok := NewSegment(chunk)
			if !ok {
				goto fail
			}

			list = append(list, s)

		} else if i == last {
			goto done
		} else {
			goto fail
		}
	}

done:
	return list, true
fail:
	return nil, false
}

func Compile(pattern string) ([]Segment, bool) {
	chunks := strings.Split(pattern, "/")
	if len(chunks) > 1 && len(chunks[0]) == 0 {
		return compile(chunks[1:]...)
	}

	return nil, false
}
