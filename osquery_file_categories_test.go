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

var ofcListJSONResponse = `
[
    {
        "id": 4,
        "name": "SSH",
	"slug": "ssh",
	"description": ".ssh folders",
	"file_paths": [
            "/root/.ssh/%%",
	    "/home/%/.ssh/%%"
	],
	"exclude_paths": [
	    "/home/not_to_monitor/.ssh/%%"
	],
	"file_paths_queries": [],
	"access_monitoring": true,
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var ofcGetJSONResponse = `
{
    "id": 4,
    "name": "SSH",
    "slug": "ssh",
    "description": ".ssh folders",
    "file_paths": [
	"/root/.ssh/%%",
	"/home/%/.ssh/%%"
    ],
    "exclude_paths": [
	"/home/not_to_monitor/.ssh/%%"
    ],
    "file_paths_queries": [],
    "access_monitoring": true,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var ofcCreateJSONResponse = `
{
    "id": 4,
    "name": "SSH",
    "slug": "ssh",
    "description": ".ssh folders",
    "file_paths": [
	"/root/.ssh/%%",
	"/home/%/.ssh/%%"
    ],
    "exclude_paths": [
	"/home/not_to_monitor/.ssh/%%"
    ],
    "file_paths_queries": [],
    "access_monitoring": true,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var ofcUpdateJSONResponse = `
{
    "id": 4,
    "name": "SSH",
    "slug": "ssh",
    "description": ".ssh folders",
    "file_paths": [
	"/root/.ssh/%%",
	"/home/%/.ssh/%%"
    ],
    "exclude_paths": [
	"/home/not_to_monitor/.ssh/%%"
    ],
    "file_paths_queries": [],
    "access_monitoring": true,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestOsqueryFileCategoriesService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/file_categories/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, ofcListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryFileCategories.List(ctx, nil)
	if err != nil {
		t.Errorf("OsqueryFileCategories.List returned error: %v", err)
	}

	want := []OsqueryFileCategory{
		{
			ID:               4,
			Name:             "SSH",
			Slug:             "ssh",
			Description:      ".ssh folders",
			FilePaths:        []string{"/root/.ssh/%%", "/home/%/.ssh/%%"},
			ExcludePaths:     []string{"/home/not_to_monitor/.ssh/%%"},
			FilePathsQueries: []string{},
			AccessMonitoring: true,
			Created:          Timestamp{referenceTime},
			Updated:          Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryFileCategories.List returned %+v, want %+v", got, want)
	}
}

func TestOsqueryFileCategoriesService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/file_categories/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, ofcGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryFileCategories.GetByID(ctx, 1)
	if err != nil {
		t.Errorf("OsqueryFileCategories.GetByID returned error: %v", err)
	}

	want := &OsqueryFileCategory{
		ID:               4,
		Name:             "SSH",
		Slug:             "ssh",
		Description:      ".ssh folders",
		FilePaths:        []string{"/root/.ssh/%%", "/home/%/.ssh/%%"},
		ExcludePaths:     []string{"/home/not_to_monitor/.ssh/%%"},
		FilePathsQueries: []string{},
		AccessMonitoring: true,
		Created:          Timestamp{referenceTime},
		Updated:          Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryFileCategories.GetByID returned %+v, want %+v", got, want)
	}
}

func TestOsqueryFileCategoriesService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/file_categories/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "SSH")
		fmt.Fprint(w, ofcListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryFileCategories.GetByName(ctx, "SSH")
	if err != nil {
		t.Errorf("OsqueryFileCategories.GetByName returned error: %v", err)
	}

	want := &OsqueryFileCategory{
		ID:               4,
		Name:             "SSH",
		Slug:             "ssh",
		Description:      ".ssh folders",
		FilePaths:        []string{"/root/.ssh/%%", "/home/%/.ssh/%%"},
		ExcludePaths:     []string{"/home/not_to_monitor/.ssh/%%"},
		FilePathsQueries: []string{},
		AccessMonitoring: true,
		Created:          Timestamp{referenceTime},
		Updated:          Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryFileCategories.GetByName returned %+v, want %+v", got, want)
	}
}

func TestOsqueryFileCategoriesService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &OsqueryFileCategoryRequest{
		Name:             "SSH",
		Description:      ".ssh folders",
		FilePaths:        []string{"/root/.ssh/%%", "/home/%/.ssh/%%"},
		ExcludePaths:     []string{"/home/not_to_monitor/.ssh/%%"},
		FilePathsQueries: []string{},
		AccessMonitoring: true,
	}

	mux.HandleFunc("/osquery/file_categories/", func(w http.ResponseWriter, r *http.Request) {
		v := new(OsqueryFileCategoryRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, ofcCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryFileCategories.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("OsqueryFileCategories.Create returned error: %v", err)
	}

	want := &OsqueryFileCategory{
		ID:               4,
		Name:             "SSH",
		Slug:             "ssh",
		Description:      ".ssh folders",
		FilePaths:        []string{"/root/.ssh/%%", "/home/%/.ssh/%%"},
		ExcludePaths:     []string{"/home/not_to_monitor/.ssh/%%"},
		FilePathsQueries: []string{},
		AccessMonitoring: true,
		Created:          Timestamp{referenceTime},
		Updated:          Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryFileCategories.Create returned %+v, want %+v", got, want)
	}
}

func TestOsqueryFileCategoriesService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &OsqueryFileCategoryRequest{
		Name:             "SSH",
		Description:      ".ssh folders",
		FilePaths:        []string{"/root/.ssh/%%", "/home/%/.ssh/%%"},
		ExcludePaths:     []string{"/home/not_to_monitor/.ssh/%%"},
		FilePathsQueries: []string{},
		AccessMonitoring: true,
	}

	mux.HandleFunc("/osquery/file_categories/1/", func(w http.ResponseWriter, r *http.Request) {
		v := new(OsqueryFileCategoryRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, ofcUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryFileCategories.Update(ctx, 1, updateRequest)
	if err != nil {
		t.Errorf("OsqueryFileCategories.Update returned error: %v", err)
	}

	want := &OsqueryFileCategory{
		ID:               4,
		Name:             "SSH",
		Slug:             "ssh",
		Description:      ".ssh folders",
		FilePaths:        []string{"/root/.ssh/%%", "/home/%/.ssh/%%"},
		ExcludePaths:     []string{"/home/not_to_monitor/.ssh/%%"},
		FilePathsQueries: []string{},
		AccessMonitoring: true,
		Created:          Timestamp{referenceTime},
		Updated:          Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryFileCategories.Update returned %+v, want %+v", got, want)
	}
}

func TestOsqueryFileCategoriesService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/file_categories/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.OsqueryFileCategories.Delete(ctx, 1)
	if err != nil {
		t.Errorf("OsqueryFileCategories.Delete returned error: %v", err)
	}
}
