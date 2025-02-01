package mut

import "sync"

type Mutable[T any] struct {
	mu sync.Mutex
	V  *T
}

func New[T any](p *T) *Mutable[T] {
	return &Mutable[T]{
		V: p,
	}
}

func (m *Mutable[T]) Mut() *T {
	m.mu.Lock()

	return m.V
}

func (m *Mutable[T]) TryMut() (*T, bool) {
	return m.V, m.mu.TryLock()
}

func (m *Mutable[T]) Unmute() {
	m.mu.Unlock()
}
