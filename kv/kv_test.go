package kv

import "testing"

func TestEq(t *testing.T) {
	if !Eq(nil, nil) {
		t.Fatal("nil/nil comparison failed")
	}
	bs := []byte("some bytes")
	if Eq(nil, bs) || Eq(bs, nil) {
		t.Fatal("nil/non-nil comparison failed")
	}
}
