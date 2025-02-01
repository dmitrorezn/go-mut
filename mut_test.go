package mut

import (
	"sync"
	"testing"
)

func TestMutable(t *testing.T) {
	val := 42
	m := New(&val)

	if m.V == nil {
		t.Fatal("expected non-nil pointer value")
	}

	// Test Mut method
	mutVal := m.Mut()
	if mutVal == nil || *mutVal != val {
		t.Fatalf("Mut returned incorrect value, got %v, want %v", mutVal, val)
	}

	if m.mu.TryLock() {
		t.Fatal("expected lock to be held by Mut")
	}

	// Unlock should release the lock
	m.Unmute()
	if !m.mu.TryLock() {
		t.Fatal("expected mutex to be available")
	}
	m.mu.Unlock()
}

func TestConcurrentAccess(t *testing.T) {
	val := struct {
		V int
	}{}
	m := New(&val)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		mutVal := m.Mut()
		mutVal.V += 1
		m.Unmute()
	}()

	go func() {
		defer wg.Done()
		mutVal := m.Mut()
		mutVal.V += 1
		m.Unmute()
	}()

	wg.Wait()

	if m.V.V != 2 {
		t.Fatalf("unexpected value: got %v, want 1 or 2", *m.V)
	}
}

func TestTryMutable(t *testing.T) {
	val := 42
	m := New(&val)

	if m.V == nil {
		t.Fatal("expected non-nil pointer value")
	}

	// Test Mut method
	mutVal, ok := m.TryMut()
	if mutVal == nil || !ok {
		t.Fatalf("Mut returned incorrect value, got %v, want %v", mutVal, val)
	}
	mutVal, ok = m.TryMut()
	if ok {
		t.Fatalf("Muted try mut should fail before unmuted want false, got true")
	}
	if m.mu.TryLock() {
		t.Fatal("expected lock to be held by Mut")
	}
	// Unmute should release the lock
	m.Unmute()
	if !m.mu.TryLock() {
		t.Fatal("expected mutex to be available")
	}
	m.mu.Unlock()
}

func TestTryConcurrentAccess(t *testing.T) {
	val := struct {
		V int
	}{}
	m := New(&val)

	if mutVal, ok := m.TryMut(); ok {
		mutVal.V++
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if mutVal, ok := m.TryMut(); ok {
				mutVal.V++
				m.Unmute()
				return
			}
		}
	}()

	if m.V.V != 1 {
		t.Fatalf("unexpected value: got %v, want 1", *m.V)
	}
	m.Unmute()

	wg.Wait()

	if m.V.V != 2 {
		t.Fatalf("unexpected value: got %v, want 1", *m.V)
	}
}
