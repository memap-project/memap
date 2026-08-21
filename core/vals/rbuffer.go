package vals

import (
	"sync"
	"time"
)

type intstr interface {
	int64 | string
}

type RingBuffer[T intstr] struct {
	mu        sync.RWMutex
	buf       []T
	cap       int64
	len       int64
	head      int64
	tail      int64
	expiresAt int64
}

func NewRingBuffer[T intstr](cap int64) *RingBuffer[T] {
	return &RingBuffer[T]{
		mu:   sync.RWMutex{},
		buf:  make([]T, cap),
		cap:  cap,
		head: 0,
		tail: 0,
		len:  0,
	}
}

// Push adds a value to the RingBuffer.
func (rb *RingBuffer[T]) Push(val T) {
	rb.mu.Lock()
	defer rb.mu.Unlock()
	rb.buf[rb.tail] = val
	rb.tail = (rb.tail + 1) % rb.cap
	if rb.len < rb.cap {
		rb.len++
	} else {
		rb.head = (rb.head + 1) % rb.cap
	}
}

// Pop removes and returns the value at the head of the RingBuffer.
// Returns false and zero value if the RingBuffer is empty.
func (rb *RingBuffer[T]) Pop() (T, bool) {
	rb.mu.Lock()
	defer rb.mu.Unlock()
	var zero T
	if rb.len == 0 {
		return zero, false
	}
	val := rb.buf[rb.head]
	rb.head = (rb.head + 1) % rb.cap
	rb.len--
	return val, true
}

// IsExpired returns true if the RingBuffer has an expiration time and is expired.
// Returns false if the RingBuffer has no expiration time or is not yet expired.
func (rb *RingBuffer[T]) IsExpired() bool {
	rb.mu.RLock()
	defer rb.mu.RUnlock()
	if rb.expiresAt == 0 {
		return false
	}
	return time.Now().Unix() > rb.expiresAt
}

// Expire sets the expiration time for the RingBuffer.
// Returns false if the RingBuffer is already expired or if ttl is negative.
func (rb *RingBuffer[T]) Expire(ttl int64) bool {
	rb.mu.Lock()
	defer rb.mu.Unlock()
	if rb.IsExpired() {
		return false
	}
	if ttl < 0 {
		return false
	}
	rb.expiresAt = time.Now().Unix() + ttl
	return true
}

// Ring retrieves a copy of the RingBuffer's contents as a slice.
func (rb *RingBuffer[T]) Ring() []T {
	rb.mu.RLock()
	defer rb.mu.RUnlock()
	result := make([]T, rb.len)
	if rb.len == 0 {
		return result
	}
	firstPart := rb.cap - rb.head
	if firstPart > rb.len {
		firstPart = rb.len
	}
	copy(result, rb.buf[rb.head:rb.head+firstPart])
	if secondPart := rb.len - firstPart; secondPart > 0 {
		copy(result[firstPart:], rb.buf[:secondPart])
	}
	return result
}

// Cap returns the capacity of the RingBuffer.
func (rb *RingBuffer[T]) Cap() int64 {
	rb.mu.RLock()
	defer rb.mu.RUnlock()
	return rb.cap
}

// Len returns the number of elements in the RingBuffer.
func (rb *RingBuffer[T]) Len() int64 {
	rb.mu.RLock()
	defer rb.mu.RUnlock()
	return rb.len
}

// Head returns the index of the head of the RingBuffer.
func (rb *RingBuffer[T]) Head() int64 {
	rb.mu.RLock()
	defer rb.mu.RUnlock()
	return rb.head
}

// Tail returns the index of the tail of the RingBuffer.
func (rb *RingBuffer[T]) Tail() int64 {
	rb.mu.RLock()
	defer rb.mu.RUnlock()
	return rb.tail
}
