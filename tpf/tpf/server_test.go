package tpf

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestPortCheck(t *testing.T) {
	log.Println("checking bad port handling")
	if _, err := NewServer(-1, 5000); err == nil {
		t.Errorf("expected an error for a negative listening port")
	}
	if _, err := NewServer(80, -1); err == nil {
		t.Errorf("expected an error for a negative proxied port")
	}
	if _, err := NewServer(80, 5000); err != nil {
		t.Error(err)
	}
}

func helloWorldServer() *http.Server {
	s := &http.Server{}
	s.Addr = "localhost:4000"
	s.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	return s
}

func waitUntilServing(s string) {
	for {
		// heatbeat helloworld server until it starts
		if resp, err := http.Get(s); err == nil {
			if resp.StatusCode == 200 {
				break
			}
		}
		time.Sleep(time.Second)
		log.Println("server did not start, heartbeating again")
	}
}

func laserServer(s string, t *testing.T) time.Duration {
	log.Printf("running speed test for %s\n", s)
	now := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count := 0
			errCount := 0
			for count < 1000 {
				if resp, err := http.Get(s); err == nil {
					if resp.StatusCode == 200 {
						count += 1
						continue
					}
				}
				errCount += 1
			}
			if errCount > 100 {
				t.Errorf("error rate is higher than .1 for %s", s)
			}
		}()
	}
	wg.Wait()
	return time.Since(now)
}

func TestServer(t *testing.T) {
	log.Println("checking HTTP forwarding")
	h := helloWorldServer()
	go func() {
		if err := h.ListenAndServe(); err != nil && err.Error() != "http: Server closed" {
			t.Error(err)
		}
	}()
	defer func() {
		if err := h.Shutdown(context.Background()); err != nil {
			t.Error(err)
		}
	}()
	be := "http://localhost:4000/"
	waitUntilServing(be)
	d0 := laserServer(be, t)
	log.Printf("4000 requests sent to helloworld backend on 4 threads in %s\n", d0.String())
	f, err := NewServer(4001, 4000)
	if err != nil {
		t.Error(err)
	}
	go func() {
		if err := f.Serve(); err != nil {
			t.Error(err)
		}
	}()
	ll := "http://localhost:4001/"
	waitUntilServing(ll)
	d1 := laserServer(ll, t)
	log.Printf("4000 requests sent to helloworld backend on 4 threads in %s\n", d1.String())
	log.Printf("proxied requests performed %f x slower than unproxied requests", d1.Seconds()/d0.Seconds())
	if err := f.Shutdown(); err != nil {
		t.Error(err)
	}
}

func TestBadConnections(t *testing.T) {
	f, err := NewServer(80, 5000)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("trying to serve on protected port")
	if err := f.Serve(); err == nil {
		t.Error("expected permissions error for port 80")
	}
	f, err = NewServer(4005, 4006)
	go func() {
		if err := f.Serve(); err != nil {
			t.Error(err)
		}
	}()
	for try := 0; try < 10; try++ {
		// should never respond
		if _, err := http.Get("http://localhost:4005"); err == nil {
			t.Error("there should be no response")
		}
		time.Sleep(250 * time.Millisecond)
	}
	defer func() {
		if err := f.Shutdown(); err != nil {
			t.Error(err)
		}
	}()
}
