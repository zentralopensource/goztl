package goztl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

var mmsmListJSONResponse = `
[
    {
        "id": 4,
	"manifest": 5,
	"sub_manifest": 6,
	"tags": [7, 8]
    }
]
`

var mmsmGetJSONResponse = `
{
    "id": 4,
    "manifest": 5,
    "sub_manifest": 6,
    "tags": [7, 8]
}
`

var mmsmCreateJSONResponse = `
{
    "id": 4,
    "manifest": 5,
    "sub_manifest": 6,
    "tags": [7, 8]
}
`

var mmsmUpdateJSONResponse = `
{
    "id": 4,
    "manifest": 5,
    "sub_manifest": 6,
    "tags": [7, 8]
}
`

func TestMonolithManifestSubManifestsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/manifest_sub_manifests/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mmsmListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithManifestSubManifests.List(ctx, nil)
	if err != nil {
		t.Errorf("MonolithManifestSubManifests.List returned error: %v", err)
	}

	want := []MonolithManifestSubManifest{
		{
			ID:            4,
			ManifestID:    5,
			SubManifestID: 6,
			TagIDs:        []int{7, 8},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithManifestSubManifests.List returned %+v, want %+v", got, want)
	}
}

func TestMonolithManifestSubManifestsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/manifest_sub_manifests/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mmsmGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithManifestSubManifests.GetByID(ctx, 4)
	if err != nil {
		t.Errorf("MonolithManifestSubManifests.GetByID returned error: %v", err)
	}

	want := &MonolithManifestSubManifest{
		ID:            4,
		ManifestID:    5,
		SubManifestID: 6,
		TagIDs:        []int{7, 8},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithManifestSubManifests.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMonolithManifestSubManifestsService_GetBySubManifestID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/manifest_sub_manifests/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "sub_manifest_id", "6")
		fmt.Fprint(w, mmsmListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithManifestSubManifests.GetBySubManifestID(ctx, 6)
	if err != nil {
		t.Errorf("MonolithManifestSubManifests.GetBySubManifestID returned error: %v", err)
	}

	want := []MonolithManifestSubManifest{
		{
			ID:            4,
			ManifestID:    5,
			SubManifestID: 6,
			TagIDs:        []int{7, 8},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithManifestSubManifests.GetBySubManifestID returned %+v, want %+v", got, want)
	}
}

func TestMonolithManifestSubManifestsService_GetByManifestID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/manifest_sub_manifests/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "manifest_id", "5")
		fmt.Fprint(w, mmsmListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithManifestSubManifests.GetByManifestID(ctx, 5)
	if err != nil {
		t.Errorf("MonolithManifestSubManifests.GetByManifestID returned error: %v", err)
	}

	want := []MonolithManifestSubManifest{
		{
			ID:            4,
			ManifestID:    5,
			SubManifestID: 6,
			TagIDs:        []int{7, 8},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithManifestSubManifests.GetByManifestID returned %+v, want %+v", got, want)
	}
}

func TestMonolithManifestSubManifestsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &MonolithManifestSubManifestRequest{
		ManifestID:    5,
		SubManifestID: 6,
		TagIDs:        []int{7, 8},
	}

	mux.HandleFunc("/monolith/manifest_sub_manifests/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MonolithManifestSubManifestRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, mmsmCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithManifestSubManifests.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("MonolithManifestSubManifests.Create returned error: %v", err)
	}

	want := &MonolithManifestSubManifest{
		ID:            4,
		ManifestID:    5,
		SubManifestID: 6,
		TagIDs:        []int{7, 8},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithManifestSubManifests.Create returned %+v, want %+v", got, want)
	}
}

func TestMonolithManifestSubManifestsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &MonolithManifestSubManifestRequest{
		ManifestID:    5,
		SubManifestID: 6,
		TagIDs:        []int{7, 8},
	}

	mux.HandleFunc("/monolith/manifest_sub_manifests/4/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MonolithManifestSubManifestRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, mmsmUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithManifestSubManifests.Update(ctx, 4, updateRequest)
	if err != nil {
		t.Errorf("MonolithManifestSubManifests.Update returned error: %v", err)
	}

	want := &MonolithManifestSubManifest{
		ID:            4,
		ManifestID:    5,
		SubManifestID: 6,
		TagIDs:        []int{7, 8},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithManifestSubManifests.Update returned %+v, want %+v", got, want)
	}
}

func TestMonolithManifestSubManifestsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/manifest_sub_manifests/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MonolithManifestSubManifests.Delete(ctx, 4)
	if err != nil {
		t.Errorf("MonolithManifestSubManifests.Delete returned error: %v", err)
	}
}
