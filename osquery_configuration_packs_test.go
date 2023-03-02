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

var ocpListJSONResponse = `
[
    {
        "id": 4,
	"configuration": 5,
	"pack": 6,
	"tags": [1, 2]
    }
]
`

var ocpGetJSONResponse = `
{
    "id": 4,
    "configuration": 5,
    "pack": 6,
    "tags": [1, 2]
}
`

var ocpCreateJSONResponse = `
{
    "id": 4,
    "configuration": 5,
    "pack": 6,
    "tags": [1, 2]
}
`

var ocpUpdateJSONResponse = `
{
    "id": 4,
    "configuration": 5,
    "pack": 6,
    "tags": [1, 2]
}
`

func TestOsqueryConfigurationPacksService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/configuration_packs/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, ocpListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryConfigurationPacks.List(ctx, nil)
	if err != nil {
		t.Errorf("OsqueryConfigurationPacks.List returned error: %v", err)
	}

	want := []OsqueryConfigurationPack{
		{
			ID:              4,
			ConfigurationID: 5,
			PackID:          6,
			TagIDs:          []int{1, 2},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryConfigurationPacks.List returned %+v, want %+v", got, want)
	}
}

func TestOsqueryConfigurationPacksService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/configuration_packs/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, ocpGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryConfigurationPacks.GetByID(ctx, 4)
	if err != nil {
		t.Errorf("OsqueryConfigurationPacks.GetByID returned error: %v", err)
	}

	want := &OsqueryConfigurationPack{
		ID:              4,
		ConfigurationID: 5,
		PackID:          6,
		TagIDs:          []int{1, 2},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryConfigurationPacks.GetByID returned %+v, want %+v", got, want)
	}
}

func TestOsqueryConfigurationPacksService_GetByConfigurationID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/configuration_packs/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "configuration_id", "5")
		fmt.Fprint(w, ocpListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryConfigurationPacks.GetByConfigurationID(ctx, 5)
	if err != nil {
		t.Errorf("OsqueryConfigurationPacks.GetByConfigurationID returned error: %v", err)
	}

	want := []OsqueryConfigurationPack{
		{
			ID:              4,
			ConfigurationID: 5,
			PackID:          6,
			TagIDs:          []int{1, 2},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryConfigurationPacks.GetByConfigurationID returned %+v, want %+v", got, want)
	}
}

func TestOsqueryConfigurationPacksService_GetByPackID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/configuration_packs/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "pack_id", "6")
		fmt.Fprint(w, ocpListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryConfigurationPacks.GetByPackID(ctx, 6)
	if err != nil {
		t.Errorf("OsqueryConfigurationPacks.GetByPackID returned error: %v", err)
	}

	want := []OsqueryConfigurationPack{
		{
			ID:              4,
			ConfigurationID: 5,
			PackID:          6,
			TagIDs:          []int{1, 2},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryConfigurationPacks.GetByPackID returned %+v, want %+v", got, want)
	}
}

func TestOsqueryConfigurationPacksService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &OsqueryConfigurationPackRequest{
		ConfigurationID: 5,
		PackID:          6,
		TagIDs:          []int{1, 2},
	}

	mux.HandleFunc("/osquery/configuration_packs/", func(w http.ResponseWriter, r *http.Request) {
		v := new(OsqueryConfigurationPackRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, ocpCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryConfigurationPacks.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("OsqueryConfigurationPacks.Create returned error: %v", err)
	}

	want := &OsqueryConfigurationPack{
		ID:              4,
		ConfigurationID: 5,
		PackID:          6,
		TagIDs:          []int{1, 2},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryConfigurationPacks.Create returned %+v, want %+v", got, want)
	}
}

func TestOsqueryConfigurationPacksService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &OsqueryConfigurationPackRequest{
		ConfigurationID: 5,
		PackID:          6,
		TagIDs:          []int{1, 2},
	}

	mux.HandleFunc("/osquery/configuration_packs/4/", func(w http.ResponseWriter, r *http.Request) {
		v := new(OsqueryConfigurationPackRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, ocpUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryConfigurationPacks.Update(ctx, 4, updateRequest)
	if err != nil {
		t.Errorf("OsqueryConfigurationPacks.Update returned error: %v", err)
	}

	want := &OsqueryConfigurationPack{
		ID:              4,
		ConfigurationID: 5,
		PackID:          6,
		TagIDs:          []int{1, 2},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryConfigurationPacks.Update returned %+v, want %+v", got, want)
	}
}

func TestOsqueryConfigurationPacksService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/configuration_packs/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.OsqueryConfigurationPacks.Delete(ctx, 4)
	if err != nil {
		t.Errorf("OsqueryConfigurationPacks.Delete returned error: %v", err)
	}
}
