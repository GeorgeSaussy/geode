package test_server

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"testing"
	"time"
)

type Server interface {
	StartInBackground(t *testing.T) error
	BlockUntilStarted(t *testing.T) error
	Stop(t *testing.T) error
}

// ServerWrapper can be used to wrap an http.Server for testing.
type ServerWrapper struct {
	// The server must respond with a 200 status code for GET requests to the path "/".
	Server *http.Server
	// The port will be used by the server instance.
	Port int
}

func (s *ServerWrapper) StartInBackground(t *testing.T) error {
	s.Server.Addr = "localhost:" + strconv.Itoa(s.Port)
	go func() {
		if err := s.Server.ListenAndServe(); err != nil && err.Error() != "http: Server closed" {
			t.Error(err)
		}
	}()
	return nil
}

func (s *ServerWrapper) BlockUntilStarted(t *testing.T) error {
	ok := false
	for i := 0; i < 10; i += 1 {
		// heatbeat helloworld server until it starts
		if resp, err := http.Get("http://localhost:" + strconv.Itoa(s.Port) + "/"); err == nil {
			if resp.StatusCode == 200 {
				ok = true
				break
			}
		}
		time.Sleep(time.Second)
	}
	if !ok {
		err := errors.New("server did not start")
		t.Error(err)
		return err
	}
	return nil
}

func (s *ServerWrapper) Stop(t *testing.T) error {
	if err := s.Server.Shutdown(context.Background()); err != nil {
		t.Error(err)
		return err
	}
	return nil
}
