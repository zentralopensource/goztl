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

var mmListJSONResponse = `
[
    {
        "id": 4,
        "name": "Default",
        "meta_business_unit": 5,
	"version": 6,
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var mmGetJSONResponse = `
{
    "id": 4,
    "name": "Default",
    "meta_business_unit": 5,
    "version": 6,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var mmCreateJSONResponse = `
{
    "id": 4,
    "name": "Default",
    "meta_business_unit": 5,
    "version": 1,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var mmUpdateJSONResponse = `
{
    "id": 4,
    "name": "Default",
    "meta_business_unit": 5,
    "version": 6,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestMonolithManifestsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/manifests/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mmListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithManifests.List(ctx, nil)
	if err != nil {
		t.Errorf("MonolithManifests.List returned error: %v", err)
	}

	want := []MonolithManifest{
		{
			ID:                 4,
			Name:               "Default",
			MetaBusinessUnitID: 5,
			Version:            6,
			Created:            Timestamp{referenceTime},
			Updated:            Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithManifests.List returned %+v, want %+v", got, want)
	}
}

func TestMonolithManifestsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/manifests/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mmGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithManifests.GetByID(ctx, 4)
	if err != nil {
		t.Errorf("MonolithManifests.GetByID returned error: %v", err)
	}

	want := &MonolithManifest{
		ID:                 4,
		Name:               "Default",
		MetaBusinessUnitID: 5,
		Version:            6,
		Created:            Timestamp{referenceTime},
		Updated:            Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithManifests.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMonolithManifestsService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/manifests/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "Default")
		fmt.Fprint(w, mmListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithManifests.GetByName(ctx, "Default")
	if err != nil {
		t.Errorf("MonolithManifests.GetByName returned error: %v", err)
	}

	want := &MonolithManifest{
		ID:                 4,
		Name:               "Default",
		MetaBusinessUnitID: 5,
		Version:            6,
		Created:            Timestamp{referenceTime},
		Updated:            Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithManifests.GetByName returned %+v, want %+v", got, want)
	}
}

func TestMonolithManifestsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &MonolithManifestRequest{
		Name:               "Default",
		MetaBusinessUnitID: 5,
	}

	mux.HandleFunc("/monolith/manifests/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MonolithManifestRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, mmCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithManifests.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("MonolithManifests.Create returned error: %v", err)
	}

	want := &MonolithManifest{
		ID:                 4,
		Name:               "Default",
		MetaBusinessUnitID: 5,
		Version:            1,
		Created:            Timestamp{referenceTime},
		Updated:            Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithManifests.Create returned %+v, want %+v", got, want)
	}
}

func TestMonolithManifestsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &MonolithManifestRequest{
		Name:               "Default",
		MetaBusinessUnitID: 5,
	}

	mux.HandleFunc("/monolith/manifests/4/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MonolithManifestRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, mmUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithManifests.Update(ctx, 4, updateRequest)
	if err != nil {
		t.Errorf("MonolithManifests.Update returned error: %v", err)
	}

	want := &MonolithManifest{
		ID:                 4,
		Name:               "Default",
		MetaBusinessUnitID: 5,
		Version:            6,
		Created:            Timestamp{referenceTime},
		Updated:            Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithManifests.Update returned %+v, want %+v", got, want)
	}
}

func TestMonolithManifestsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/manifests/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MonolithManifests.Delete(ctx, 4)
	if err != nil {
		t.Errorf("MonolithManifests.Delete returned error: %v", err)
	}
}
