package test_server

import (
	"io"
	"net/http"
	"strconv"
	"testing"
)

func TestHelloServer(t *testing.T) {
	s := MakeHelloServer(4100)
	s.StartInBackground(t)
	s.BlockUntilStarted(t)
	res, err := http.Get("http://localhost:4100/")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Error code was %d, not 200 as expected.", res.StatusCode)
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	if string(b) != "hello world" {
		t.Errorf("Messages was %s, not 'hello world' as expected.", strconv.Quote(string(b)))
	}
	s.Stop(t)
}
