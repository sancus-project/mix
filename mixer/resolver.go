package mixer

import (
	"context"
	"log"
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

func (rr *RouterResolver) Report(h web.Handler, rctx *mix.Context) {
	rr.mu.Lock()
	defer rr.mu.Unlock()

	if !rr.closed {
		log.Printf("%s:%s h:%#v rctx:%#v", "Report", "", h, rctx)

		rr.h = h
		rr.rctx = rctx
		rr.closed = true
	} else {
		log.Printf("%s:%s h:%#v rctx:%#v", "Report", "<CLOSED>", h, rctx)
	}
}

func (rr *RouterResolver) Result() (web.Handler, *mix.Context, bool) {
	rr.mu.Lock()
	defer rr.mu.Unlock()

	if !rr.closed {
		rr.closed = true
	}

	if rr.h != nil {
		log.Printf("%s:%s h:%#v rctx:%#v", "Result", "", rr.h, rr.rctx)
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
	log.Printf("%T.%s: m:%#v h:%T ctx:%#v", v, "Resolve", v.m, v.h, v.rctx)

	if r, ok := v.h.(Router); ok {
		// Router
		if h, rctx, ok := r.Resolve(v.rctx); ok {
			rr.Report(h, rctx)
		}
	} else {
		// Handler
		rr.Report(v.h, v.rctx)
	}
}

func resolve(matches []RouterMatch) (web.Handler, *mix.Context, bool) {
	var wg sync.WaitGroup
	var rr RouterResolver

	for _, m := range matches {
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
	log.Printf("%T.%s: path:%q % #v", m, "Resolve", path, rctx)

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

	path := rctx.RoutePath
	log.Printf("%T.%s: path:%q % #v", m, "Resolve", path, rctx)

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
