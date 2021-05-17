package mixer

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/armon/go-radix"

	"go.sancus.dev/mix/mixer/tree"
	"go.sancus.dev/mix/types"
)

type Router interface {
	types.Router

	Init(m *Mixer, servertiming string)
	Mixer() *Mixer

	RouteSegments(p []tree.Segment) Router
	AttachHandler(w types.Handler) error
}

type router struct {
	mixer   *Mixer
	mu      sync.RWMutex
	top     Router
	handler []types.Handler

	ServerTimingPrefix string
}

func (r *router) init(m *Mixer, top Router, servertiming string) {
	r.mixer = m
	r.top = top
	r.ServerTimingPrefix = servertiming
}

func (r *router) Mixer() *Mixer {
	if r != nil {
		return r.mixer
	}
	return nil
}

// CatchAll
type RouterCatchAll struct {
	router
}

func newCatchAll(m *Mixer) Router {
	r := &RouterCatchAll{}
	m.initRouter(r)
	return r
}

func (r *RouterCatchAll) Init(m *Mixer, servertiming string) {
	r.router.init(m, r, servertiming)
}

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
type RouterNode struct {
	router

	trie *radix.Tree
	exps RouterExp

	catchall Router
}

func newRouter(m *Mixer) Router {
	r := &RouterNode{}
	m.initRouter(r)
	return r
}

func (r *RouterNode) Init(m *Mixer, servertiming string) {
	r.router.init(m, r, servertiming)

	r.trie = radix.New()
	r.exps.trie = radix.New()
}

func (m *Mixer) initRouter(r Router) {
	var servicetiming string

	if len(m.config.ServerTiming) > 0 {
		n := m.routerCount
		servicetiming = fmt.Sprintf("%s-%v", m.config.ServerTiming, n)
	}

	m.routerCount++

	r.Init(m, servicetiming)
}

// Gets Router for a given path pattern
func (m *router) Route(pattern string, fn func(r types.Router)) types.Router {

	p, err := tree.Compile(pattern)
	if err != nil {
		log.Fatal(err)
	}

	r := m.routeSegments(p)
	if r == nil {
		return nil
	}

	if fn != nil {
		fn(r)
	}

	return r
}

func (m *router) routeSegments(p []tree.Segment) Router {
	if m != nil && m.top != nil {
		return m.top.RouteSegments(p)
	} else {
		return nil
	}
}

func (m *RouterNode) RouteSegments(p []tree.Segment) Router {
	var r Router
	var p0 tree.Segment

	m.mu.Lock()
	defer m.mu.Unlock()

	p0, p = p[0], p[1:]

	if _, ok := p0.(tree.CatchAll); ok {
		r = m.catchall
		if r == nil {
			r = newCatchAll(m.mixer)
			m.catchall = r
		}
	} else if _, ok := p0.(tree.Literal); ok {
		// literal string
		s := p0.String()
		v, ok := m.trie.Get(s)
		if ok {
			// reuse
			r = v.(Router)
		} else {
			// new
			r = newRouter(m.mixer)
			m.trie.Insert(s, r)
		}
	} else {
		// pattern
		v, ok := m.exps.Get(p0)
		if ok {
			// reuse
			r = v.(Router)
		} else {
			// new
			r = newRouter(m.mixer)
			m.exps.Append(p0, r)
		}
	}

	if len(p) == 0 {
		return r
	} else {
		return r.RouteSegments(p)
	}
}

func (m *RouterCatchAll) RouteSegments(p []tree.Segment) Router {
	if len(p) == 0 {
		return m
	}
	return nil
}

// Mounts handler at path
func (m *router) Mount(path string, h interface{}) error {
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

	r := m.routeSegments(p)
	if r == nil {
		return tree.InvalidPattern(pattern)
	}

	w, err := m.mixer.newHandler(pattern, h)
	if err != nil {
		return err
	}

	return r.AttachHandler(w)
}

// Attach handler to Router
func (m *router) Attach(h interface{}) error {
	w, err := m.mixer.newHandler("/", h)
	if err != nil {
		return err
	}

	return m.AttachHandler(w)
}

func (m *router) AttachHandler(w types.Handler) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.handler = append(m.handler, w)
	return nil
}
