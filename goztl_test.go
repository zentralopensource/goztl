package goztl

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
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

type rapTestItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var rapFirstPageJSONResponse = `{
	"count": 2,
	"next": "http://example.com/test/items/?page=2",
	"results": [
		{"id": 1, "name": "un"}
	]
}`

var rapNextPageJSONResponse = `{
	"count": 2,
	"results": [
		{"id": 2, "name": "deux"}
	]
}`

var rapUnpaginatedJSONResponse = `[
	{"id": 1, "name": "un"},
	{"id": 2, "name": "deux"}
]`

func TestResolveAllPages(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/test/items/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		if r.URL.Query().Get("page") == "" {
			fmt.Fprint(w, rapFirstPageJSONResponse)
			return
		}

		testQueryArg(t, r, "page", "2")
		fmt.Fprint(w, rapNextPageJSONResponse)
	})

	ctx := context.Background()
	items, _, err := resolveAllPages[rapTestItem](ctx, client, "test/items/")
	if err != nil {
		t.Errorf("resolveAllPages returned error: %v", err)
	}

	want := []rapTestItem{{ID: 1, Name: "un"}, {ID: 2, Name: "deux"}}
	if !cmp.Equal(items, want) {
		t.Errorf("resolveAllPages returned %+v, want %+v", items, want)
	}
}

func TestResolveAllPagesUnpaginatedEndpoint(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	var requests int

	mux.HandleFunc("/test/items/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		requests++
		fmt.Fprint(w, rapUnpaginatedJSONResponse)
	})

	ctx := context.Background()
	items, _, err := resolveAllPages[rapTestItem](ctx, client, "test/items/")
	if err != nil {
		t.Errorf("resolveAllPages returned error: %v", err)
	}

	want := []rapTestItem{{ID: 1, Name: "un"}, {ID: 2, Name: "deux"}}
	if !cmp.Equal(items, want) {
		t.Errorf("resolveAllPages returned %+v, want %+v", items, want)
	}

	if requests != 1 {
		t.Errorf("resolveAllPages made %d requests, want 1", requests)
	}
}
