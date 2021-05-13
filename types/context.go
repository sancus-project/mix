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
		pattern = filepath.Join(prefix, "*")
	}

	*rctx = Context{
		RoutePrefix:  prefix,
		RoutePattern: pattern,
		RoutePath:    path,
	}

	return nil
}

func (rctx *Context) Next() (*Context, string) {

	path := rctx.RoutePath

	if len(path) > 1 {

		s := path[1:]
		prefix := rctx.RoutePrefix
		pattern := strings.TrimSuffix("/*", rctx.RoutePattern)

		if prefix == "/" {
			prefix = ""
		}

		if n := strings.IndexRune(s, '/'); n < 0 {
			prefix += path
			pattern += path
			path = ""

		} else {
			s = s[:n]
			prefix += path[:n+1]
			pattern += path[:n+1] + "/*"
			path = path[n+1:]
		}

		next := &Context{
			RoutePath: path,
			RoutePrefix: prefix,
			RoutePattern: pattern,
		}

		return next, s
	}

	return nil, ""
}
