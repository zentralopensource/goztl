package goztl

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var gwsConnectionListJSONResponse = `
[
    {
        "id": "2eb6636e-3266-46f4-8f88-a45e99dfdef9",
        "name": "Default Connection",
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var gwsConnectionGetByNameJSONResponse = `
[
    {
        "id": "2eb6636e-3266-46f4-8f88-a45e99dfdef9",
        "name": "Default Connection",
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var gwsConnectionGetJSONResponse = `
{
    "id": "2eb6636e-3266-46f4-8f88-a45e99dfdef9",
	"name": "Default Connection",
	"healthy": true,
	"created_at": "2022-07-22T01:02:03.444444",
	"updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestGWSConnectionsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/google_workspace/connections/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, gwsConnectionListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.GWSConnections.List(ctx, nil)
	if err != nil {
		t.Errorf("TestGWSConnections.List returned error: %v", err)
	}

	want := []GWSConnection{
		{
			ID:      "2eb6636e-3266-46f4-8f88-a45e99dfdef9",
			Name:    "Default Connection",
			Created: Timestamp{referenceTime},
			Updated: Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("TestGWSConnections.List returned %+v, want %+v", got, want)
	}
}

func TestGWSConnectionsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/google_workspace/connections/2eb6636e-3266-46f4-8f88-a45e99dfdef9/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, gwsConnectionGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.GWSConnections.GetByID(ctx, "2eb6636e-3266-46f4-8f88-a45e99dfdef9")
	if err != nil {
		t.Errorf("TestGWSConnections.GetByID returned error: %v", err)
	}

	want := &GWSConnection{
		ID:      "2eb6636e-3266-46f4-8f88-a45e99dfdef9",
		Name:    "Default Connection",
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("TestGWSConnections.GetByID returned %+v, want %+v", got, want)
	}
}

func TestGWSConnectionsService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/google_workspace/connections/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "Default Connection")
		fmt.Fprint(w, gwsConnectionGetByNameJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.GWSConnections.GetByName(ctx, "Default Connection")
	if err != nil {
		t.Errorf("TestGWSConnections.GetByName returned error: %v", err)
	}

	want := &GWSConnection{
		ID:      "2eb6636e-3266-46f4-8f88-a45e99dfdef9",
		Name:    "Default Connection",
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("TestGWSConnections.GetByName returned %+v, want %+v", got, want)
	}
}
