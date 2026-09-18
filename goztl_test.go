package goztl

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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

// setupMountedAPI is setup() for an API mounted under a path, which is how Zentral is
// deployed: the base URL the terraform provider documents ends in /api/.
func setupMountedAPI() (client *Client, mux *http.ServeMux, teardown func()) {
	mux = http.NewServeMux()
	server := httptest.NewServer(mux)

	client, _ = NewClient(nil, server.URL+"/api/", testToken)

	return client, mux, server.Close
}

// withOrigin fills the origin in a fixture. DRF builds the next link of a page from the
// incoming request, so a fixture cannot carry a host the client never talked to.
func withOrigin(body string, origin string) string {
	return strings.ReplaceAll(body, "$ORIGIN", origin)
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
	"next": "$ORIGIN/test/items/?page=2",
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
			fmt.Fprint(w, withOrigin(rapFirstPageJSONResponse, "http://"+r.Host))
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

func TestResolveAllPagesMountedUnderAPath(t *testing.T) {
	// the next link of a page is absolute, and its path carries the mount point. Resolving it
	// against a base URL that carries the mount point too asked for it twice: /api/api/...
	client, mux, teardown := setupMountedAPI()
	defer teardown()

	var paths []string

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		paths = append(paths, r.URL.Path)
		testMethod(t, r, "GET")

		if r.URL.Path != "/api/test/items/" {
			http.NotFound(w, r)
			return
		}

		if r.URL.Query().Get("page") == "" {
			fmt.Fprint(w, withOrigin(rapFirstPageJSONResponse, "http://"+r.Host+"/api"))
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

	wantPaths := []string{"/api/test/items/", "/api/test/items/"}
	if !cmp.Equal(paths, wantPaths) {
		t.Errorf("resolveAllPages requested %v, want %v", paths, wantPaths)
	}
}

func TestResolveAllPagesIgnoresTheHostOfTheNextLink(t *testing.T) {
	// only the path and the query of the link are read, so the requests stay on the base URL
	client, mux, teardown := setupMountedAPI()
	defer teardown()

	var hosts []string

	mux.HandleFunc("/api/test/items/", func(w http.ResponseWriter, r *http.Request) {
		hosts = append(hosts, r.Host)
		testMethod(t, r, "GET")

		if r.URL.Query().Get("page") == "" {
			fmt.Fprint(w, withOrigin(rapFirstPageJSONResponse, "https://other.example.com/api"))
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

	wantHosts := []string{client.BaseURL.Host, client.BaseURL.Host}
	if !cmp.Equal(hosts, wantHosts) {
		t.Errorf("resolveAllPages asked %v, want %v", hosts, wantHosts)
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
