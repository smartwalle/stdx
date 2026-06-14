package syncx

import (
	"context"
	"sync"
	"sync/atomic"
)

type Group struct {
	ctx    context.Context
	cancel func(error)

	wg sync.WaitGroup

	sem chan struct{}

	errOnce sync.Once
	err     error

	panicHandler atomic.Value
}

func NewGroup(ctx context.Context) *Group {
	if ctx == nil {
		ctx = context.Background()
	}
	ctx, cancel := context.WithCancelCause(ctx)
	return &Group{ctx: ctx, cancel: cancel}
}

func (g *Group) Wait() error {
	g.wg.Wait()
	if g.cancel != nil {
		g.cancel(g.err)
	}
	return g.err
}

func (g *Group) done() {
	defer g.wg.Done()

	if x := recover(); x != nil {
		g.handlePanic(x)
	}

	if g.sem != nil {
		<-g.sem
	}
}

func (g *Group) HandlePanic(handler PanicHandler) {
	if handler == nil {
		return
	}
	g.panicHandler.Store(handler)
}

func (g *Group) handlePanic(x any) {
	var handler = g.panicHandler.Load()
	if handler == nil {
		return
	}

	defer func() {
		_ = recover()
	}()
	handler.(PanicHandler)(x)
}

func (g *Group) Go(fn func(context.Context) error) {
	select {
	case <-g.ctx.Done():
		return
	default:
	}

	if g.sem != nil {
		select {
		case <-g.ctx.Done():
			return
		case g.sem <- struct{}{}:
		}
	}

	g.wg.Add(1)
	go func() {
		defer g.done()
		if err := fn(g.ctx); err != nil {
			g.errOnce.Do(func() {
				g.err = err
				if g.cancel != nil {
					g.cancel(g.err)
				}
			})
		}
	}()
}

func (g *Group) Run(fn func(ctx context.Context) error) {
	select {
	case <-g.ctx.Done():
		return
	default:
	}

	if g.sem != nil {
		select {
		case <-g.ctx.Done():
			return
		case g.sem <- struct{}{}:
		}
	}

	g.wg.Add(1)
	go func() {
		defer g.done()
		if err := fn(g.ctx); err != nil {
			g.errOnce.Do(func() {
				g.err = err
			})
		}
	}()
}

func (g *Group) TryGo(fn func(context.Context) error) bool {
	select {
	case <-g.ctx.Done():
		return false
	default:
	}

	if g.sem != nil {
		select {
		case <-g.ctx.Done():
			return false
		case g.sem <- struct{}{}:
		default:
			return false
		}
	}

	g.wg.Add(1)
	go func() {
		defer g.done()
		if err := fn(g.ctx); err != nil {
			g.errOnce.Do(func() {
				g.err = err
				if g.cancel != nil {
					g.cancel(g.err)
				}
			})
		}
	}()
	return true
}

func (g *Group) Limit(n int) bool {
	if n < 1 || g.sem != nil {
		return false
	}
	g.sem = make(chan struct{}, n)
	return true
}
