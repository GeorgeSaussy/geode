package kv

import (
	"sync"
)

type pair struct {
	key   []byte
	value []byte
}

type MemoryKv struct {
	pairs []pair
	mu    sync.RWMutex
}

func NewMemoryKv() Kv {
	return &MemoryKv{
		pairs: make([]pair, 0),
	}
}

// Check if the key is present in the store.
// This function will never fail.
func (m *MemoryKv) Contains(key []byte) (bool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.ki(key) != -1, nil
}

// Get the value for a given key or return an error if the
// key is not found.
func (m *MemoryKv) Get(key []byte) ([]byte, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	i := m.ki(key)
	if i == -1 {
		return nil, ErrKeyNotFound
	}
	return m.pairs[i].value, nil
}

// Set a value for a given key, overriding the current value
// if the key is already present.
// This function will never fail.
func (m *MemoryKv) Set(key []byte, value []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	i := m.ki(key)
	if i == -1 {
		p := pair{
			key:   key,
			value: value,
		}
		m.pairs = append(m.pairs, p)
		return nil
	}
	m.pairs[i].value = value
	return nil
}

// Update the value for a key or return an error if the key
// is not present.
func (m *MemoryKv) Update(key []byte, value []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	i := m.ki(key)
	if i == -1 {
		return ErrKeyNotFound
	}
	m.pairs[i].value = value
	return nil
}

// Insert a new key/value pair or return an error if the key
// is already present.
func (m *MemoryKv) Insert(key []byte, value []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	i := m.ki(key)
	if i != -1 {
		return ErrKeyFound
	}
	p := pair{
		key:   key,
		value: value,
	}
	m.pairs = append(m.pairs, p)
	return nil
}

// Delete a key or return an error if the key is not present.
func (m *MemoryKv) Delete(key []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	i := m.ki(key)
	if i == -1 {
		return ErrKeyNotFound
	}
	pl := len(m.pairs)
	m.pairs[i] = m.pairs[pl-1]
	m.pairs = m.pairs[:pl-1]
	return nil
}

// Get the index for a key within MemoryKv::pairs or -1 if it is
// not present. This implementation is not thread safe so the the
// MemoryKv instance should already be locked against mutations
// for the durration of the function call's execution.
func (m *MemoryKv) ki(key []byte) int {
	for k, p := range m.pairs {
		if Eq(p.key, key) {
			return k
		}
	}
	return -1
}
