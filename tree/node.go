package tree

import (
	"sync"
)

type Node struct {
	sync.RWMutex
}
