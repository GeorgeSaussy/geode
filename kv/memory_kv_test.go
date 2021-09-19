package kv

import "testing"

func TestNewMemoryKv(t *testing.T) {
	kv := NewMemoryKv()
	ks := [][]byte{
		[]byte("key0"),
		[]byte("key1"),
		[]byte("key2"),
		[]byte("key3"),
	}
	vs := [][]byte{
		[]byte("value0"),
		[]byte("value1"),
		[]byte("value2"),
		[]byte("value3"),
		[]byte("value4"),
	}
	if b, err := kv.Contains(ks[0]); err != nil {
		t.Fatal(err)
	} else if b {
		t.Fatal("value present")
	}
	if err := kv.Set(ks[0], vs[0]); err != nil {
		t.Fatal(err)
	}
	if b, err := kv.Contains(ks[0]); err != nil {
		t.Fatal(err)
	} else if !b {
		t.Fatal("value not present")
	}
	if v, err := kv.Get(ks[0]); err != nil {
		t.Fatal(err)
	} else if !Eq(v, vs[0]) {
		t.Fatal("incorrect value")
	}
	if _, err := kv.Get(ks[1]); err == nil {
		t.Fatal("value not present")
	}
	if err := kv.Insert(ks[0], vs[1]); err == nil {
		t.Fatal("error expected")
	}
	if err := kv.Insert(ks[1], vs[1]); err != nil {
		t.Fatal(err)
	}
	if err := kv.Update(ks[2], vs[2]); err == nil {
		t.Fatal("error expected")
	}
	if err := kv.Insert(ks[2], vs[2]); err != nil {
		t.Fatal(err)
	}
	if err := kv.Update(ks[2], vs[3]); err != nil {
		t.Fatal(err)
	}
	if v, err := kv.Get(ks[2]); err != nil {
		t.Fatal(err)
	} else if !Eq(v, vs[3]) {
		t.Fatal("incorrect value")
	}
	if err := kv.Delete(ks[2]); err != nil {
		t.Fatal(err)
	}
	if b, err := kv.Contains(ks[2]); err != nil {
		t.Fatal(err)
	} else if b {
		t.Fatal("value present")
	}
	if err := kv.Delete(ks[2]); err == nil {
		t.Fatal("error expected")
	}
	if err := kv.Set(ks[3], vs[3]); err != nil {
		t.Fatal(err)
	}
	if v, err := kv.Get(ks[3]); err != nil {
		t.Fatal(err)
	} else if !Eq(v, vs[3]) {
		t.Fatal("incorrect value")
	}
	if err := kv.Set(ks[3], vs[4]); err != nil {
		t.Fatal(err)
	}
	if v, err := kv.Get(ks[3]); err != nil {
		t.Fatal(err)
	} else if !Eq(v, vs[4]) {
		t.Fatal("incorrect value")
	}
}
