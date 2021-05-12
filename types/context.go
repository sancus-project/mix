package types

import (
	"context"
	"path/filepath"
	"strings"
)

type Context struct {
	RoutePrefix  string
	RoutePath    string
	RoutePattern string
}

// Clone() creates a copy of a routing Context object
func (rctx Context) Clone() *Context {
	return &rctx
}

func (rctx *Context) Init(ctx context.Context, prefix, path string) error {
	var pattern string

	if prefix == "" {
		prefix = "/"
	}

	if path == "" {
		pattern = prefix
	} else if n := strings.IndexRune(path[1:], '/'); n < 0 {
		pattern = filepath.Join(prefix, path)
	} else {
		pattern = filepath.Join(prefix, path[0:n+1]) + "/*"
	}

	*rctx = Context{
		RoutePrefix:  prefix,
		RoutePattern: pattern,
		RoutePath:    path,
	}

	return nil
}
