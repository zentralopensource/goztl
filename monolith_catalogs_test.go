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

var mcListJSONResponse = `
[
    {
        "id": 4,
	"repository": 3,
        "name": "Default",
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var mcGetJSONResponse = `
{
    "id": 4,
    "repository": 3,
    "name": "Default",
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444",
    "archived_at": "2022-07-22T01:02:03.444444"
}
`

var mcCreateJSONResponse = `
{
    "id": 4,
    "repository": 3,
    "name": "Default",
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444",
    "archived_at": null
}
`

var mcUpdateJSONResponse = `
{
    "id": 4,
    "repository": 3,
    "name": "Standard",
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444",
    "archived_at": null
}
`

func TestMonolithCatalogsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/catalogs/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mcListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithCatalogs.List(ctx, nil)
	if err != nil {
		t.Errorf("MonolithCatalogs.List returned error: %v", err)
	}

	want := []MonolithCatalog{
		{
			ID:           4,
			RepositoryID: 3,
			Name:         "Default",
			Created:      Timestamp{referenceTime},
			Updated:      Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithCatalogs.List returned %+v, want %+v", got, want)
	}
}

func TestMonolithCatalogsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/catalogs/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mcGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithCatalogs.GetByID(ctx, 4)
	if err != nil {
		t.Errorf("MonolithCatalogs.GetByID returned error: %v", err)
	}

	want := &MonolithCatalog{
		ID:           4,
		RepositoryID: 3,
		Name:         "Default",
		Created:      Timestamp{referenceTime},
		Updated:      Timestamp{referenceTime},
		ArchivedAt:   &Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithCatalogs.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMonolithCatalogsService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/catalogs/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "Default")
		fmt.Fprint(w, mcListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithCatalogs.GetByName(ctx, "Default")
	if err != nil {
		t.Errorf("MonolithCatalogs.GetByName returned error: %v", err)
	}

	want := &MonolithCatalog{
		ID:           4,
		RepositoryID: 3,
		Name:         "Default",
		Created:      Timestamp{referenceTime},
		Updated:      Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithCatalogs.GetByName returned %+v, want %+v", got, want)
	}
}

func TestMonolithCatalogsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &MonolithCatalogRequest{
		RepositoryID: 3,
		Name:         "Default",
	}

	mux.HandleFunc("/monolith/catalogs/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MonolithCatalogRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, mcCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithCatalogs.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("MonolithCatalogs.Create returned error: %v", err)
	}

	want := &MonolithCatalog{
		ID:           4,
		RepositoryID: 3,
		Name:         "Default",
		Created:      Timestamp{referenceTime},
		Updated:      Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithCatalogs.Create returned %+v, want %+v", got, want)
	}
}

func TestMonolithCatalogsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &MonolithCatalogRequest{
		RepositoryID: 3,
		Name:         "Standard",
	}

	mux.HandleFunc("/monolith/catalogs/4/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MonolithCatalogRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, mcUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithCatalogs.Update(ctx, 4, updateRequest)
	if err != nil {
		t.Errorf("MonolithCatalogs.Update returned error: %v", err)
	}

	want := &MonolithCatalog{
		ID:           4,
		RepositoryID: 3,
		Name:         "Standard",
		Created:      Timestamp{referenceTime},
		Updated:      Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithCatalogs.Update returned %+v, want %+v", got, want)
	}
}

func TestMonolithCatalogsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/catalogs/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MonolithCatalogs.Delete(ctx, 4)
	if err != nil {
		t.Errorf("MonolithCatalogs.Delete returned error: %v", err)
	}
}
