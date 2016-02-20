package main

import (
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRepro(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello world")
	}))
	defer server.Close()

	addr := server.Listener.Addr().String()
	t.Logf("Listening on %v", addr)

	for d := time.Second; d > 0; d *= 2 {
		t.Logf("Attempting timeout of %v", d)
		c, err := net.DialTimeout("tcp", addr, d)
		if err != nil {
			t.Fatalf("Connect failed with timeout of %v (i64 value: %v): %v", d, int64(d), err)
		}
		c.Close()
	}
}
