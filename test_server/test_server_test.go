package test_server

import (
	"errors"
	"testing"
)

type BadTestServer struct{}

func (bts BadTestServer) StartInBackground(t *testing.T) error {
	err := errors.New("starting in the background failed")
	t.Error(err)
	return err
}
func (bts BadTestServer) BlockUntilStarted(t *testing.T) error {
	err := errors.New("server never started")
	t.Error(err)
	return err

}
func (bts BadTestServer) Stop(t *testing.T) error {
	err := errors.New("server never stopped")
	t.Error(err)
	return err
}

func TestBadTestServer(t *testing.T) {
	tp := &testing.T{}
	if tp.Failed() {
		t.Errorf("new T should not have failed")
	}
	bts := BadTestServer{}
	bts.StartInBackground(tp)
	if !tp.Failed() {
		t.Errorf("server should not have started in the background")
	}
	tp = &testing.T{}
	bts.BlockUntilStarted(tp)
	if !tp.Failed() {
		t.Errorf("server should never have started")
	}
	tp = &testing.T{}
	bts.Stop(tp)
	if !tp.Failed() {
		t.Errorf("server should not have stopped")
	}
}
