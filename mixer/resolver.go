package mixer

import (
	"bytes"
	"context"
	"net/http"
	"sync"

	"go.sancus.dev/mix"
	"go.sancus.dev/mix/mixer/tree"
	"go.sancus.dev/mix/types"
	"go.sancus.dev/web"
)

// Router.Resolve()
type RouterResolver struct {
	mu     sync.Mutex
	closed bool

	h    web.Handler
	rctx *mix.Context
}

func (rr *RouterResolver) NewRequest(rctx *mix.Context, method string) *http.Request {
	var buf bytes.Buffer

	ctx := context.Background()
	ctx = mix.WithRouteContext(ctx, rctx)

	url := rctx.Path()
	req, err := http.NewRequestWithContext(ctx, method, url, &buf)
	if err != nil {
		panic(err)
	}
	return req
}

func (rr *RouterResolver) Report(h web.Handler, rctx *mix.Context) {
	rr.mu.Lock()
	defer rr.mu.Unlock()

	if !rr.closed {
		rr.h = h
		rr.rctx = rctx
		rr.closed = true
	}
}

func (rr *RouterResolver) Result() (web.Handler, *mix.Context, bool) {
	rr.mu.Lock()
	defer rr.mu.Unlock()

	if !rr.closed {
		rr.closed = true
	}

	if rr.h != nil {
		return rr.h, rr.rctx, true
	}

	return nil, nil, false
}

type RouterMatch struct {
	m    tree.Match
	h    types.Handler
	rctx *mix.Context
}

func (v *RouterMatch) Resolve(rr *RouterResolver) {
	if m, ok := v.h.(Router); !ok {
		// types.Handler

		req := rr.NewRequest(v.rctx, "HEAD")
		if p, ok := v.h.PageInfo(req); ok {

			// Page found, did we get a direct handler?
			if h, ok := p.(web.Handler); ok {
				rr.Report(h, v.rctx)
			} else {
				rr.Report(v.h, v.rctx)
			}
		}

	} else if h, rctx, ok := m.Resolve(v.rctx); ok {
		// Handler Validated by another Router
		rr.Report(h, rctx)
	}
}

func resolve(matches []RouterMatch) (web.Handler, *mix.Context, bool) {
	var wg sync.WaitGroup
	var rr RouterResolver

	for i, _ := range matches {
		m := matches[i]

		wg.Add(1)
		go func() {
			defer wg.Done()
			m.Resolve(&rr)
		}()
	}
	wg.Wait()

	return rr.Result()
}

func (m *RouterNode) match(rctx *mix.Context, s string) ([]RouterMatch, bool) {
	var matches []RouterMatch

	// Literal strings
	if v, ok := m.trie.Get(s); ok {
		if h, ok := v.(Router); ok {
			matches = append(matches, RouterMatch{
				m:    s,
				h:    h,
				rctx: rctx,
			})
		}
	}

	// Expressions
	if v, r, ok := m.exps.Match(s); ok {

		for i, h := range r {
			matches = append(matches, RouterMatch{
				m:    v[i],
				h:    h,
				rctx: rctx,
			})
		}

	}

	if len(matches) > 0 {
		return matches, true
	}

	return nil, false
}

func (m *RouterNode) Resolve(rctx *mix.Context) (web.Handler, *mix.Context, bool) {
	var matches []RouterMatch

	path := rctx.RoutePath

	m.mu.Lock()
	defer m.mu.Unlock()

	if path == "" {
		// exact match
		for _, h := range m.handler {
			matches = append(matches, RouterMatch{
				m:    path,
				h:    h,
				rctx: rctx,
			})

		}
	} else {
		if m.catchall != nil {
			matches = append(matches, RouterMatch{
				m:    &tree.CatchAll{},
				h:    m.catchall,
				rctx: rctx,
			})
		}

		if rctx, s := rctx.Next(); s != "" {
			if v, ok := m.match(rctx, s); ok {
				matches = append(matches, v...)
			}
		}
	}

	if len(matches) > 0 {
		return resolve(matches)
	}

	return nil, nil, false
}

func (m *RouterCatchAll) Resolve(rctx *mix.Context) (web.Handler, *mix.Context, bool) {
	var matches []RouterMatch

	m.mu.Lock()
	defer m.mu.Unlock()

	for _, h := range m.handler {
		matches = append(matches, RouterMatch{
			m:    &tree.CatchAll{},
			h:    h,
			rctx: rctx,
		})
	}

	if len(matches) > 0 {
		return resolve(matches)
	}

	return nil, nil, false
}

func (m *router) resolve(rctx *mix.Context) (web.Handler, *mix.Context, bool) {
	if m != nil && m.top != nil {
		return m.top.Resolve(rctx)
	}

	return nil, nil, false
}

// Router.ResolvePath()
func (m *router) ResolvePath(ctx context.Context, prefix, path string) (web.Handler, *mix.Context, bool) {
	if rctx := mix.NewRouteContext(ctx, prefix, path); rctx != nil {
		return m.resolve(rctx)
	}

	return nil, nil, false
}

// Router.ResolveRequest()
func (m *router) ResolveRequest(ctx context.Context, r *http.Request) (web.Handler, *mix.Context, bool) {
	path := m.mixer.config.GetRoutePath(r)

	return m.ResolvePath(ctx, "/", path)
}
