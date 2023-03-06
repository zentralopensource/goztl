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

var msmListJSONResponse = `
[
    {
        "id": 4,
        "name": "Default",
	"description": "Description",
        "meta_business_unit": 5,
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var msmGetJSONResponse = `
{
    "id": 4,
    "name": "Default",
    "description": "Description",
    "meta_business_unit": 5,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var msmCreateJSONResponse = `
{
    "id": 4,
    "name": "Default",
    "description": "Description",
    "meta_business_unit": 5,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var msmUpdateJSONResponse = `
{
    "id": 4,
    "name": "Default",
    "description": "",
    "meta_business_unit": null,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestMonolithSubManifestsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/sub_manifests/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, msmListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithSubManifests.List(ctx, nil)
	if err != nil {
		t.Errorf("MonolithSubManifests.List returned error: %v", err)
	}

	want := []MonolithSubManifest{
		{
			ID:                 4,
			Name:               "Default",
			Description:        "Description",
			MetaBusinessUnitID: Int(5),
			Created:            Timestamp{referenceTime},
			Updated:            Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithSubManifests.List returned %+v, want %+v", got, want)
	}
}

func TestMonolithSubManifestsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/sub_manifests/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, msmGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithSubManifests.GetByID(ctx, 4)
	if err != nil {
		t.Errorf("MonolithSubManifests.GetByID returned error: %v", err)
	}

	want := &MonolithSubManifest{
		ID:                 4,
		Name:               "Default",
		Description:        "Description",
		MetaBusinessUnitID: Int(5),
		Created:            Timestamp{referenceTime},
		Updated:            Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithSubManifests.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMonolithSubManifestsService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/sub_manifests/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "Default")
		fmt.Fprint(w, msmListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithSubManifests.GetByName(ctx, "Default")
	if err != nil {
		t.Errorf("MonolithSubManifests.GetByName returned error: %v", err)
	}

	want := &MonolithSubManifest{
		ID:                 4,
		Name:               "Default",
		Description:        "Description",
		MetaBusinessUnitID: Int(5),
		Created:            Timestamp{referenceTime},
		Updated:            Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithSubManifests.GetByName returned %+v, want %+v", got, want)
	}
}

func TestMonolithSubManifestsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &MonolithSubManifestRequest{
		Name:               "Default",
		MetaBusinessUnitID: Int(5),
	}

	mux.HandleFunc("/monolith/sub_manifests/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MonolithSubManifestRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, msmCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithSubManifests.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("MonolithSubManifests.Create returned error: %v", err)
	}

	want := &MonolithSubManifest{
		ID:                 4,
		Name:               "Default",
		Description:        "Description",
		MetaBusinessUnitID: Int(5),
		Created:            Timestamp{referenceTime},
		Updated:            Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithSubManifests.Create returned %+v, want %+v", got, want)
	}
}

func TestMonolithSubManifestsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &MonolithSubManifestRequest{
		Name: "Default",
	}

	mux.HandleFunc("/monolith/sub_manifests/4/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MonolithSubManifestRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, msmUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithSubManifests.Update(ctx, 4, updateRequest)
	if err != nil {
		t.Errorf("MonolithSubManifests.Update returned error: %v", err)
	}

	want := &MonolithSubManifest{
		ID:          4,
		Name:        "Default",
		Description: "",
		Created:     Timestamp{referenceTime},
		Updated:     Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithSubManifests.Update returned %+v, want %+v", got, want)
	}
}

func TestMonolithSubManifestsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/sub_manifests/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MonolithSubManifests.Delete(ctx, 4)
	if err != nil {
		t.Errorf("MonolithSubManifests.Delete returned error: %v", err)
	}
}
