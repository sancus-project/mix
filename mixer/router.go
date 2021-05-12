package mixer

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/armon/go-radix"

	"go.sancus.dev/mix"
	"go.sancus.dev/mix/mixer/tree"
	"go.sancus.dev/mix/types"
	"go.sancus.dev/web"
)

// Expression based routers
type RouterExp struct {
	Keys   []tree.Segment
	Values []types.Router

	trie *radix.Tree
}

// Adds a expression/router pair to the table
func (t *RouterExp) Append(key tree.Segment, value types.Router) {
	t.Keys = append(t.Keys, key)
	t.Values = append(t.Values, value)
	t.trie.Insert(key.String(), value)
}

// Returns all expression based routers that match a segment string
func (t *RouterExp) Match(s string) ([]tree.Match, []types.Router, bool) {
	var matches []tree.Match
	var routers []types.Router

	for i, k := range t.Keys {
		if v, ok := k.Match(s); ok {
			r := t.Values[i]

			matches = append(matches, v)
			routers = append(routers, r)
		}
	}

	if len(routers) > 0 {
		return matches, routers, true
	}

	return nil, nil, false
}

// Finds router by expression
func (t *RouterExp) Get(key tree.Segment) (r types.Router, ok bool) {
	if v, ok := t.trie.Get(key.String()); ok {
		r, ok = v.(types.Router)
	}
	return
}

// Router description at a specific segment
type Router struct {
	mixer *Mixer
	mu    sync.RWMutex

	trie *radix.Tree
	exps RouterExp

	handler []types.Handler

	ServerTimingPrefix string
}

func (m *Mixer) initRouter(r *Router) {
	n := m.routerCount
	m.routerCount++

	r.mixer = m
	r.trie = radix.New()
	r.exps.trie = radix.New()

	if len(m.config.ServerTiming) > 0 {
		m.ServerTimingPrefix = fmt.Sprintf("%s-%v", m.config.ServerTiming, n)
	}
}

// Match
func (m *Router) match(s string) ([]tree.Match, []types.Router, bool) {
	var matches []tree.Match
	var routers []types.Router

	m.mu.RLock()
	defer m.mu.RUnlock()

	// Expressions
	if v, r, ok := m.exps.Match(s); ok {
		matches = append(matches, v...)
		routers = append(routers, r...)
	}

	// Literal strings

	if len(routers) > 0 {
		return matches, routers, true
	}

	return nil, nil, false
}

// GetPage
func (m *Router) getpage(rctx *mix.Context) (web.Handler, *mix.Context, bool) {
	return nil, nil, false
}

func (m *Router) GetPageFromPath(ctx context.Context, prefix, path string) (web.Handler, *mix.Context, bool) {

	if rctx := mix.NewRouteContext(prefix, path); rctx != nil {
		return m.getpage(rctx)
	}

	return nil, nil, false
}

func (m *Router) GetPage(ctx context.Context, r *http.Request) (web.Handler, *mix.Context, bool) {

	path := m.mixer.config.GetRoutePath(r)

	return m.GetPageFromPath(ctx, "/", path)
}

// Gets Router for a given path pattern
func (m *Router) Route(pattern string, fn func(r types.Router)) types.Router {
	p, err := tree.Compile(pattern)
	if err != nil {
		log.Fatal(err)
	}

	r := m.route(p)
	if r == nil {
		return nil
	}

	if fn != nil {
		fn(r)
	}

	return r
}

func (m *Router) route(p []tree.Segment) *Router {
	var r *Router
	var p0 tree.Segment
	var s string

	m.mu.Lock()
	defer m.mu.Unlock()

	p0, p = p[0], p[1:]
	s = p0.String()

	if _, ok := p0.(tree.Literal); ok {
		// literal string
		v, ok := m.trie.Get(s)
		if ok {
			// reuse
			r = v.(*Router)
		} else {
			// new
			r = m.mixer.NewRouter()
			m.trie.Insert(s, r)
		}
	} else {
		// pattern
		v, ok := m.exps.Get(p0)
		if ok {
			// reuse
			r = v.(*Router)
		} else {
			// new
			r = m.mixer.NewRouter()
			m.exps.Append(p0, r)
		}
	}

	if len(p) == 0 {
		return r
	} else {
		return r.route(p)
	}
}

// Mounts handler at path
func (m *Router) Mount(path string, h interface{}) error {
	var pattern string

	if path == "" {
		pattern = "/*"
	} else if strings.HasSuffix(path, "/") {
		pattern = path + "*"
	} else {
		pattern = path
	}

	p, err := tree.Compile(pattern)
	if err != nil {
		return err
	}

	w, err := m.mixer.newHandler(pattern, h)
	if err != nil {
		return err
	}

	return m.route(p).attach(w)
}

// Attach handler to Router
func (m *Router) Attach(h interface{}) error {
	w, err := m.mixer.newHandler("/", h)
	if err != nil {
		return err
	}

	return m.attach(w)
}

func (m *Router) attach(w types.Handler) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.handler = append(m.handler, w)
	return nil
}
