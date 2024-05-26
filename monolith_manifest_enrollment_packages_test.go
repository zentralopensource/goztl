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

var mmepListJSONResponse = `
[
    {
        "id": 4,
	"manifest": 5,
	"builder": "zentral.contrib.munki.osx_package.builder.MunkiZentralEnrollPkgBuilder",
	"enrollment_pk": 6,
	"version": 7,
	"tags": [7, 8],
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var mmepGetJSONResponse = `
{
    "id": 4,
    "manifest": 5,
    "builder": "zentral.contrib.munki.osx_package.builder.MunkiZentralEnrollPkgBuilder",
    "enrollment_pk": 6,
    "version": 7,
    "tags": [7, 8],
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var mmepCreateJSONResponse = `
{
    "id": 4,
    "manifest": 5,
    "builder": "zentral.contrib.munki.osx_package.builder.MunkiZentralEnrollPkgBuilder",
    "enrollment_pk": 6,
    "version": 7,
    "tags": [7, 8],
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var mmepUpdateJSONResponse = `
{
    "id": 4,
    "manifest": 5,
    "builder": "zentral.contrib.munki.osx_package.builder.MunkiZentralEnrollPkgBuilder",
    "enrollment_pk": 6,
    "version": 7,
    "tags": [7, 8],
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestMonolithManifestEnrollmentPackagesService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/manifest_enrollment_packages/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mmepListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithManifestEnrollmentPackages.List(ctx, nil)
	if err != nil {
		t.Errorf("MonolithManifestEnrollmentPackages.List returned error: %v", err)
	}

	want := []MonolithManifestEnrollmentPackage{
		{
			ID:           4,
			ManifestID:   5,
			Builder:      "zentral.contrib.munki.osx_package.builder.MunkiZentralEnrollPkgBuilder",
			EnrollmentID: 6,
			Version:      7,
			TagIDs:       []int{7, 8},
			Created:      Timestamp{referenceTime},
			Updated:      Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithManifestEnrollmentPackages.List returned %+v, want %+v", got, want)
	}
}

func TestMonolithManifestEnrollmentPackagesService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/manifest_enrollment_packages/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mmepGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithManifestEnrollmentPackages.GetByID(ctx, 4)
	if err != nil {
		t.Errorf("MonolithManifestEnrollmentPackages.GetByID returned error: %v", err)
	}

	want := &MonolithManifestEnrollmentPackage{
		ID:           4,
		ManifestID:   5,
		Builder:      "zentral.contrib.munki.osx_package.builder.MunkiZentralEnrollPkgBuilder",
		EnrollmentID: 6,
		Version:      7,
		TagIDs:       []int{7, 8},
		Created:      Timestamp{referenceTime},
		Updated:      Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithManifestEnrollmentPackages.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMonolithManifestEnrollmentPackagesService_GetByManifestID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/manifest_enrollment_packages/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "manifest_id", "5")
		fmt.Fprint(w, mmepListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithManifestEnrollmentPackages.GetByManifestID(ctx, 5)
	if err != nil {
		t.Errorf("MonolithManifestEnrollmentPackages.GetByManifestID returned error: %v", err)
	}

	want := []MonolithManifestEnrollmentPackage{
		{
			ID:           4,
			ManifestID:   5,
			Builder:      "zentral.contrib.munki.osx_package.builder.MunkiZentralEnrollPkgBuilder",
			EnrollmentID: 6,
			Version:      7,
			TagIDs:       []int{7, 8},
			Created:      Timestamp{referenceTime},
			Updated:      Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithManifestEnrollmentPackages.GetByManifestID returned %+v, want %+v", got, want)
	}
}

func TestMonolithManifestEnrollmentPackagesService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &MonolithManifestEnrollmentPackageRequest{
		ManifestID:   5,
		Builder:      "zentral.contrib.munki.osx_package.builder.MunkiZentralEnrollPkgBuilder",
		EnrollmentID: 6,
		TagIDs:       []int{7, 8},
	}

	mux.HandleFunc("/monolith/manifest_enrollment_packages/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MonolithManifestEnrollmentPackageRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, mmepCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithManifestEnrollmentPackages.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("MonolithManifestEnrollmentPackages.Create returned error: %v", err)
	}

	want := &MonolithManifestEnrollmentPackage{
		ID:           4,
		ManifestID:   5,
		Builder:      "zentral.contrib.munki.osx_package.builder.MunkiZentralEnrollPkgBuilder",
		EnrollmentID: 6,
		Version:      7,
		TagIDs:       []int{7, 8},
		Created:      Timestamp{referenceTime},
		Updated:      Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithManifestEnrollmentPackages.Create returned %+v, want %+v", got, want)
	}
}

func TestMonolithManifestEnrollmentPackagesService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &MonolithManifestEnrollmentPackageRequest{
		ManifestID:   5,
		Builder:      "zentral.contrib.munki.osx_package.builder.MunkiZentralEnrollPkgBuilder",
		EnrollmentID: 6,
		TagIDs:       []int{7, 8},
	}

	mux.HandleFunc("/monolith/manifest_enrollment_packages/4/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MonolithManifestEnrollmentPackageRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, mmepUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithManifestEnrollmentPackages.Update(ctx, 4, updateRequest)
	if err != nil {
		t.Errorf("MonolithManifestEnrollmentPackages.Update returned error: %v", err)
	}

	want := &MonolithManifestEnrollmentPackage{
		ID:           4,
		ManifestID:   5,
		Builder:      "zentral.contrib.munki.osx_package.builder.MunkiZentralEnrollPkgBuilder",
		EnrollmentID: 6,
		Version:      7,
		TagIDs:       []int{7, 8},
		Created:      Timestamp{referenceTime},
		Updated:      Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithManifestEnrollmentPackages.Update returned %+v, want %+v", got, want)
	}
}

func TestMonolithManifestEnrollmentPackagesService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/manifest_enrollment_packages/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MonolithManifestEnrollmentPackages.Delete(ctx, 4)
	if err != nil {
		t.Errorf("MonolithManifestEnrollmentPackages.Delete returned error: %v", err)
	}
}
