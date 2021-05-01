package tree

import (
	"sync"

	"github.com/armon/go-radix"
)

type Tree struct {
	sync.RWMutex

	trie *radix.Tree
}

func New() *Tree {
	t := &Tree{}
	return t.Init()
}

func (t *Tree) Init() *Tree {
	t.Lock()
	defer t.Unlock()

	if t.trie == nil {
		t.trie = radix.New()
	}
	return t
}

func (t *Tree) Len() int {
	t.RLock()
	defer t.RUnlock()

	return t.trie.Len()
}
