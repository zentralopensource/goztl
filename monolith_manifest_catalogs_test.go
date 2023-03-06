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

var mmcListJSONResponse = `
[
    {
        "id": 4,
	"manifest": 5,
	"catalog": 6,
	"tags": [7, 8]
    }
]
`

var mmcGetJSONResponse = `
{
    "id": 4,
    "manifest": 5,
    "catalog": 6,
    "tags": [7, 8]
}
`

var mmcCreateJSONResponse = `
{
    "id": 4,
    "manifest": 5,
    "catalog": 6,
    "tags": [7, 8]
}
`

var mmcUpdateJSONResponse = `
{
    "id": 4,
    "manifest": 5,
    "catalog": 6,
    "tags": [7, 8]
}
`

func TestMonolithManifestCatalogsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/manifest_catalogs/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mmcListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithManifestCatalogs.List(ctx, nil)
	if err != nil {
		t.Errorf("MonolithManifestCatalogs.List returned error: %v", err)
	}

	want := []MonolithManifestCatalog{
		{
			ID:         4,
			ManifestID: 5,
			CatalogID:  6,
			TagIDs:     []int{7, 8},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithManifestCatalogs.List returned %+v, want %+v", got, want)
	}
}

func TestMonolithManifestCatalogsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/manifest_catalogs/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mmcGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithManifestCatalogs.GetByID(ctx, 4)
	if err != nil {
		t.Errorf("MonolithManifestCatalogs.GetByID returned error: %v", err)
	}

	want := &MonolithManifestCatalog{
		ID:         4,
		ManifestID: 5,
		CatalogID:  6,
		TagIDs:     []int{7, 8},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithManifestCatalogs.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMonolithManifestCatalogsService_GetByCatalogID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/manifest_catalogs/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "catalog_id", "6")
		fmt.Fprint(w, mmcListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithManifestCatalogs.GetByCatalogID(ctx, 6)
	if err != nil {
		t.Errorf("MonolithManifestCatalogs.GetByCatalogID returned error: %v", err)
	}

	want := []MonolithManifestCatalog{
		{
			ID:         4,
			ManifestID: 5,
			CatalogID:  6,
			TagIDs:     []int{7, 8},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithManifestCatalogs.GetByCatalogID returned %+v, want %+v", got, want)
	}
}

func TestMonolithManifestCatalogsService_GetByManifestID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/manifest_catalogs/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "manifest_id", "5")
		fmt.Fprint(w, mmcListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithManifestCatalogs.GetByManifestID(ctx, 5)
	if err != nil {
		t.Errorf("MonolithManifestCatalogs.GetByManifestID returned error: %v", err)
	}

	want := []MonolithManifestCatalog{
		{
			ID:         4,
			ManifestID: 5,
			CatalogID:  6,
			TagIDs:     []int{7, 8},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithManifestCatalogs.GetByManifestID returned %+v, want %+v", got, want)
	}
}

func TestMonolithManifestCatalogsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &MonolithManifestCatalogRequest{
		ManifestID: 5,
		CatalogID:  6,
		TagIDs:     []int{7, 8},
	}

	mux.HandleFunc("/monolith/manifest_catalogs/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MonolithManifestCatalogRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, mmcCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithManifestCatalogs.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("MonolithManifestCatalogs.Create returned error: %v", err)
	}

	want := &MonolithManifestCatalog{
		ID:         4,
		ManifestID: 5,
		CatalogID:  6,
		TagIDs:     []int{7, 8},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithManifestCatalogs.Create returned %+v, want %+v", got, want)
	}
}

func TestMonolithManifestCatalogsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &MonolithManifestCatalogRequest{
		ManifestID: 5,
		CatalogID:  6,
		TagIDs:     []int{7, 8},
	}

	mux.HandleFunc("/monolith/manifest_catalogs/4/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MonolithManifestCatalogRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, mmcUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithManifestCatalogs.Update(ctx, 4, updateRequest)
	if err != nil {
		t.Errorf("MonolithManifestCatalogs.Update returned error: %v", err)
	}

	want := &MonolithManifestCatalog{
		ID:         4,
		ManifestID: 5,
		CatalogID:  6,
		TagIDs:     []int{7, 8},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithManifestCatalogs.Update returned %+v, want %+v", got, want)
	}
}

func TestMonolithManifestCatalogsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/manifest_catalogs/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MonolithManifestCatalogs.Delete(ctx, 4)
	if err != nil {
		t.Errorf("MonolithManifestCatalogs.Delete returned error: %v", err)
	}
}
