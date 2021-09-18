package kv

import "errors"

var ErrKeyFound = errors.New("key found")
var ErrKeyNotFound = errors.New("key not found")

// Kv instances store a set of key/value pairs.
// The keys are unique accross the store.
type Kv interface {
	// Check if the key is present in the store.
	Contains(key []byte) (bool, error)
	// Get the value for a given key or return an error if the
	// key is not found.
	Get(key []byte) ([]byte, error)
	// Set a value for a given key, overriding the current value
	// if the key is already present.
	Set(key []byte, value []byte) error
	// Update the value for a key or return an error if the key
	// is not present.
	Update(key []byte, value []byte) error
	// Insert a new key/value pair or return an error if the key
	// is already present.
	Insert(key []byte, value []byte) error
	// Delete a key or return an error if the key is not present.
	Delete(key []byte) error
}

// Check if two byte slices are equal.
func Eq(a []byte, b []byte) bool {
	if a == nil {
		return b == nil
	}
	if b == nil || len(a) != len(b) {
		return false
	}
	for k, va := range a {
		vb := b[k]
		if va != vb {
			return false
		}
	}
	return true
}
