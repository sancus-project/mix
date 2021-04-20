package tree

import (
	"testing"
)

func TestTreeNew(t *testing.T) {
	t1 := New()
	t2 := &Tree{}
	t3 := t2.Init()

	if t2 != t3 {
		t.Fatalf("`Tree.Init()` wrong return (%T)", t3)
	}

	if t1.Len() != t2.Len() {
		t.Fatalf("`tree.New()` wrong return (%T)", t1)
	}
}
