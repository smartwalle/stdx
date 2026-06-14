package syncx

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

var (
	ErrRoutineClosed    = errors.New("routine is closed")
	ErrRoutineQueueFull = errors.New("routine queue is full")
	ErrRoutineBadTask   = errors.New("routine task is nil")
)

type Routine struct {
	tasks chan func()

	mu         sync.Mutex
	cond       *sync.Cond
	workers    int
	submitters int

	closed    chan struct{}
	closeOnce sync.Once

	workerWg sync.WaitGroup

	workerSize  int
	idleTimeout time.Duration

	panicHandler atomic.Value
}

type PanicHandler func(any)

func NewRoutine(workerSize int, queueSize int) *Routine {
	if workerSize < 1 {
		workerSize = 1
	}
	if queueSize < 0 {
		queueSize = 0
	}

	var routine = &Routine{
		tasks:       make(chan func(), queueSize),
		closed:      make(chan struct{}),
		workerSize:  workerSize,
		idleTimeout: time.Second,
	}
	routine.cond = sync.NewCond(&routine.mu)

	return routine
}

func (r *Routine) Submit(ctx context.Context, fn func()) error {
	if fn == nil {
		return ErrRoutineBadTask
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if !r.beginSubmit() {
		return ErrRoutineClosed
	}
	defer r.endSubmit()

	select {
	case r.tasks <- fn:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	case <-r.closed:
		return ErrRoutineClosed
	}
}

func (r *Routine) TrySubmit(fn func()) error {
	if fn == nil {
		return ErrRoutineBadTask
	}
	if !r.beginSubmit() {
		return ErrRoutineClosed
	}
	defer r.endSubmit()

	select {
	case r.tasks <- fn:
		return nil
	case <-r.closed:
		return ErrRoutineClosed
	default:
		return ErrRoutineQueueFull
	}
}

func (r *Routine) Close() {
	r.closeOnce.Do(func() {
		r.mu.Lock()
		close(r.closed)
		for r.submitters > 0 {
			r.cond.Wait()
		}
		close(r.tasks)
		r.mu.Unlock()

		r.workerWg.Wait()
	})
}

func (r *Routine) HandlePanic(handler PanicHandler) {
	if handler == nil {
		return
	}
	r.panicHandler.Store(handler)
}

func (r *Routine) Closed() bool {
	select {
	case <-r.closed:
		return true
	default:
		return false
	}
}

func (r *Routine) beginSubmit() bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	select {
	case <-r.closed:
		return false
	default:
	}
	r.submitters++

	if r.workers < r.workerSize {
		r.workers++
		r.workerWg.Add(1)
		go r.worker()
	}

	return true
}

func (r *Routine) endSubmit() {
	r.mu.Lock()
	r.submitters--
	if r.submitters == 0 {
		r.cond.Broadcast()
	}
	r.mu.Unlock()
}

func (r *Routine) worker() {
	var stopped = false
	defer func() {
		if !stopped {
			r.stopWorker()
		}
		r.workerWg.Done()
	}()

	var timer = time.NewTimer(r.idleTimeout)
	defer timer.Stop()

	for {
		select {
		case fn, ok := <-r.tasks:
			if !ok {
				return
			}
			if !timer.Stop() {
				select {
				case <-timer.C:
				default:
				}
			}
			r.run(fn)
			timer.Reset(r.idleTimeout)
		case <-timer.C:
			select {
			case fn, ok := <-r.tasks:
				if !ok {
					return
				}
				r.run(fn)
				timer.Reset(r.idleTimeout)
			default:
				r.mu.Lock()
				if r.submitters > 0 {
					r.mu.Unlock()
					timer.Reset(r.idleTimeout)
					continue
				}
				r.workers--
				stopped = true
				r.mu.Unlock()
				return
			}
		}
	}
}

func (r *Routine) stopWorker() {
	r.mu.Lock()
	r.workers--
	r.mu.Unlock()
}

func (r *Routine) run(fn func()) {
	defer func() {
		if x := recover(); x != nil {
			r.handlePanic(x)
		}
	}()
	fn()
}

func (r *Routine) handlePanic(x any) {
	var handler = r.panicHandler.Load()
	if handler == nil {
		return
	}

	defer func() {
		_ = recover()
	}()
	handler.(PanicHandler)(x)
}
