package goztl

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	testToken = "TOKEN"
)

func setup() (client *Client, mux *http.ServeMux, teardown func()) {
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(mux)

	// client is the Zentral client being tested and is configured
	// to use the test server.
	client, _ = NewClient(nil, server.URL, testToken)

	return client, mux, server.Close
}

func testBody(t *testing.T, r *http.Request, want string) {
	t.Helper()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		t.Errorf("Error reading request body: %v", err)
	}
	if got := string(b); got != want {
		t.Errorf("request Body is %s, want %s", got, want)
	}
}

func testHeader(t *testing.T, r *http.Request, header string, want string) {
	t.Helper()
	if got := r.Header.Get(header); got != want {
		t.Errorf("Header.Get(%q) returned %q, want %q", header, got, want)
	}
}

func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testQueryArg(t *testing.T, r *http.Request, arg string, want string) {
	t.Helper()
	if got := r.URL.Query().Get(arg); got != want {
		t.Errorf("Request query arg %q: value %q, want %q", arg, got, want)
	}
}
