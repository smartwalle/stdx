package syncx

import (
	"context"
	"sync"
)

type Group struct {
	ctx    context.Context
	cancel func(error)

	wg sync.WaitGroup

	routine *Routine

	errOnce sync.Once
	err     error
}

func NewGroup(ctx context.Context, maxConcurrency int, queueCapacity int) *Group {
	if ctx == nil {
		ctx = context.Background()
	}
	ctx, cancel := context.WithCancelCause(ctx)
	if maxConcurrency < 1 {
		maxConcurrency = 1
	}
	if queueCapacity < 0 {
		queueCapacity = 0
	}
	return &Group{
		ctx:     ctx,
		cancel:  cancel,
		routine: NewRoutine(maxConcurrency, queueCapacity),
	}
}

func (g *Group) Wait() error {
	g.wg.Wait()
	if g.routine != nil {
		g.routine.Close()
	}
	cause := context.Cause(g.ctx)
	if g.cancel != nil {
		g.cancel(g.err)
	}
	if g.err != nil {
		return g.err
	}
	return cause
}

func (g *Group) OnPanic(handler PanicHandler) {
	g.routine.OnPanic(handler)
}

func (g *Group) Go(fn func(context.Context) error) {
	select {
	case <-g.ctx.Done():
		return
	default:
	}

	g.wg.Add(1)
	if g.routine.Go(g.ctx, g.makeTask(fn, true)) != nil {
		g.wg.Done()
	}
}

func (g *Group) Run(fn func(ctx context.Context) error) {
	select {
	case <-g.ctx.Done():
		return
	default:
	}

	g.wg.Add(1)
	if g.routine.Go(g.ctx, g.makeTask(fn, false)) != nil {
		g.wg.Done()
	}
}

func (g *Group) TryGo(fn func(context.Context) error) bool {
	select {
	case <-g.ctx.Done():
		return false
	default:
	}

	g.wg.Add(1)
	if g.routine.TryGo(g.ctx, g.makeTask(fn, true)) != nil {
		g.wg.Done()
		return false
	}
	return true
}

func (g *Group) makeTask(fn func(context.Context) error, cancelOnError bool) func() {
	return func() {
		defer g.wg.Done()
		if err := fn(g.ctx); err != nil {
			g.errOnce.Do(func() {
				g.err = err
				if cancelOnError && g.cancel != nil {
					g.cancel(g.err)
				}
			})
		}
	}
}
