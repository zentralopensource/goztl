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

var mpkgListJSONResponse = `{
    "count": 1,
    "results": [
        {
            "id": "74d2bc35-7ebe-4d76-8e58-fb9a8b9adc11",
            "name": "Example",
            "description": "Example description",
            "type": "PKG",
            "source_uri": "https://example.com/example.pkg",
            "sha256": "0000000000000000000000000000000000000000000000000000000000000000",
            "size": 12345,
            "filename": "example.pkg",
            "product_id": "com.example.pkg",
            "product_version": "1.0",
            "bundles": [{"id": "com.example.pkg", "version": "1.0"}],
            "manifest": {"items": []},
            "created_at": "2022-07-22T01:02:03.444444",
            "updated_at": "2022-07-22T01:02:03.444444"
        }
    ]
}
`

var mpkgGetJSONResponse = `
{
    "id": "74d2bc35-7ebe-4d76-8e58-fb9a8b9adc11",
    "name": "Example",
    "description": "Example description",
    "type": "PKG",
    "source_uri": "https://example.com/example.pkg",
    "sha256": "0000000000000000000000000000000000000000000000000000000000000000",
    "size": 12345,
    "filename": "example.pkg",
    "product_id": "com.example.pkg",
    "product_version": "1.0",
    "bundles": [{"id": "com.example.pkg", "version": "1.0"}],
    "manifest": {"items": []},
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var mpkgCreateJSONResponse = `
{
    "id": "74d2bc35-7ebe-4d76-8e58-fb9a8b9adc11",
    "name": "Example",
    "description": "Example description",
    "type": "PKG",
    "source_uri": "https://example.com/example.pkg",
    "sha256": "0000000000000000000000000000000000000000000000000000000000000000",
    "size": 12345,
    "filename": "example.pkg",
    "product_id": "com.example.pkg",
    "product_version": "1.0",
    "bundles": [{"id": "com.example.pkg", "version": "1.0"}],
    "manifest": {"items": []},
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var mpkgUpdateJSONResponse = `
{
    "id": "74d2bc35-7ebe-4d76-8e58-fb9a8b9adc11",
    "name": "Renamed",
    "description": "Updated description",
    "type": "PKG",
    "source_uri": "https://example.com/example.pkg",
    "sha256": "0000000000000000000000000000000000000000000000000000000000000000",
    "size": 12345,
    "filename": "example.pkg",
    "product_id": "com.example.pkg",
    "product_version": "1.0",
    "bundles": [{"id": "com.example.pkg", "version": "1.0"}],
    "manifest": {"items": []},
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func wantMDMPackage(name, description string) MDMPackage {
	return MDMPackage{
		ID:             "74d2bc35-7ebe-4d76-8e58-fb9a8b9adc11",
		Name:           name,
		Description:    description,
		Type:           "PKG",
		SourceURI:      "https://example.com/example.pkg",
		SHA256:         "0000000000000000000000000000000000000000000000000000000000000000",
		Size:           12345,
		Filename:       "example.pkg",
		ProductID:      "com.example.pkg",
		ProductVersion: "1.0",
		Bundles:        []map[string]interface{}{{"id": "com.example.pkg", "version": "1.0"}},
		Manifest:       map[string]interface{}{"items": []interface{}{}},
		Created:        Timestamp{referenceTime},
		Updated:        Timestamp{referenceTime},
	}
}

func TestMDMPackagesService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/packages/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mpkgListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMPackages.List(ctx, nil)
	if err != nil {
		t.Errorf("MDMPackages.List returned error: %v", err)
	}

	want := []MDMPackage{wantMDMPackage("Example", "Example description")}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMPackages.List returned %+v, want %+v", got, want)
	}
}

func TestMDMPackagesService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/packages/74d2bc35-7ebe-4d76-8e58-fb9a8b9adc11/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mpkgGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMPackages.GetByID(ctx, "74d2bc35-7ebe-4d76-8e58-fb9a8b9adc11")
	if err != nil {
		t.Errorf("MDMPackages.GetByID returned error: %v", err)
	}

	w := wantMDMPackage("Example", "Example description")
	want := &w
	if !cmp.Equal(got, want) {
		t.Errorf("MDMPackages.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMDMPackagesService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/packages/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "Example")
		fmt.Fprint(w, mpkgListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMPackages.GetByName(ctx, "Example")
	if err != nil {
		t.Errorf("MDMPackages.GetByName returned error: %v", err)
	}

	want := []MDMPackage{wantMDMPackage("Example", "Example description")}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMPackages.GetByName returned %+v, want %+v", got, want)
	}
}

func TestMDMPackagesService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &MDMPackageCreateRequest{
		Name:        "Example",
		Description: "Example description",
		SourceURI:   "https://example.com/example.pkg",
		SHA256:      "0000000000000000000000000000000000000000000000000000000000000000",
	}

	mux.HandleFunc("/mdm/packages/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMPackageCreateRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, mpkgCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMPackages.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("MDMPackages.Create returned error: %v", err)
	}

	w := wantMDMPackage("Example", "Example description")
	want := &w
	if !cmp.Equal(got, want) {
		t.Errorf("MDMPackages.Create returned %+v, want %+v", got, want)
	}
}

func TestMDMPackagesService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &MDMPackageUpdateRequest{
		Name:        "Renamed",
		Description: "Updated description",
	}

	mux.HandleFunc("/mdm/packages/74d2bc35-7ebe-4d76-8e58-fb9a8b9adc11/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMPackageUpdateRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, mpkgUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMPackages.Update(ctx, "74d2bc35-7ebe-4d76-8e58-fb9a8b9adc11", updateRequest)
	if err != nil {
		t.Errorf("MDMPackages.Update returned error: %v", err)
	}

	w := wantMDMPackage("Renamed", "Updated description")
	want := &w
	if !cmp.Equal(got, want) {
		t.Errorf("MDMPackages.Update returned %+v, want %+v", got, want)
	}
}

func TestMDMPackagesService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/packages/74d2bc35-7ebe-4d76-8e58-fb9a8b9adc11/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MDMPackages.Delete(ctx, "74d2bc35-7ebe-4d76-8e58-fb9a8b9adc11")
	if err != nil {
		t.Errorf("MDMPackages.Delete returned error: %v", err)
	}
}
