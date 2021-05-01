package tree

import (
	"fmt"
)

func InvalidPattern(pattern string) error {
	return fmt.Errorf("%q: Invalid Pattern", pattern)
}
