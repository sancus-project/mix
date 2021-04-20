package tree

import (
	"testing"
)

func TestLock(t *testing.T) {
	n := Node{}
	n.Lock()
	n.Unlock()
}
