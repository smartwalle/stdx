package syncx

import (
	"context"
	"errors"
	"sync"
)

var (
	ErrSignalClosed = errors.New("signal is closed")
)

type Signal[T any] struct {
	mu             sync.RWMutex
	subscribers    map[uint64]chan<- T
	nextSubscriber uint64
	closed         bool
}

func NewSignal[T any]() *Signal[T] {
	return &Signal[T]{
		subscribers: make(map[uint64]chan<- T),
	}
}

func (s *Signal[T]) Notify(ctx context.Context, payload T) error {
	if ctx == nil {
		ctx = context.Background()
	}
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.closed {
		return ErrSignalClosed
	}

	for _, subscriber := range s.subscribers {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case subscriber <- payload:
		default:
		}
	}
	return nil
}

func (s *Signal[T]) Listen(bufferSize int) (message <-chan T, cancel func()) {
	if bufferSize < 0 {
		bufferSize = 0
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.closed {
		var closeChan = make(chan T)
		close(closeChan)
		return closeChan, func() {}
	}

	s.nextSubscriber++
	var id = s.nextSubscriber
	var subscriber = make(chan T, bufferSize)
	s.subscribers[id] = subscriber

	message = subscriber
	cancel = func() {
		s.removeSubscriber(id)
	}

	return message, cancel
}

func (s *Signal[T]) removeSubscriber(id uint64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if subscriber, ok := s.subscribers[id]; ok {
		delete(s.subscribers, id)
		close(subscriber)
	}
}

func (s *Signal[T]) Closed() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.closed
}

func (s *Signal[T]) Close() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.closed {
		return
	}
	s.closed = true

	for _, subscriber := range s.subscribers {
		close(subscriber)
	}
	s.subscribers = nil
}
