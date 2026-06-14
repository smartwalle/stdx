package syncx

import (
	"context"
	"sync"
)

type Signal[T any] struct {
	mu             sync.RWMutex
	subscribers    map[uint64]chan<- T
	nextSubscriber uint64
}

func NewSignal[T any]() *Signal[T] {
	return &Signal[T]{
		subscribers: make(map[uint64]chan<- T),
	}
}

func (s *Signal[T]) Notify(ctx context.Context, payload T) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, subscriber := range s.subscribers {
		select {
		case <-ctx.Done():
			return
		default:
		}

		select {
		case <-ctx.Done():
			return
		case subscriber <- payload:
		default:
		}
	}
}

func (s *Signal[T]) Listen(bufferSize int) (message <-chan T, cancel func()) {
	s.mu.Lock()
	defer s.mu.Unlock()

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
