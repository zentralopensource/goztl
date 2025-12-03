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

var gwsGroupTagMappingListJSONResponse = `
[
    {
        "id": "42ac3eed-884c-4abf-8845-e304f5c10122",
		"group_email": "no-reply@zentral.com",
		"connection": "e4681fa5-5eb8-4e38-bc71-8a0c091468b1",
		"tags": [1],
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var gwsGroupTagMappingGetByGroupEmailJSONResponse = `
[
    {
        "id": "42ac3eed-884c-4abf-8845-e304f5c10122",
		"group_email": "no-reply@zentral.com",
		"connection": "e4681fa5-5eb8-4e38-bc71-8a0c091468b1",
		"tags": [1],
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var gwsGroupTagMappingGetByConnectionIdJSONResponse = `
[
    {
        "id": "42ac3eed-884c-4abf-8845-e304f5c10122",
		"group_email": "no-reply@zentral.com",
		"connection": "e4681fa5-5eb8-4e38-bc71-8a0c091468b1",
		"tags": [1],
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var gwsGroupTagMappingCreateJSONResponse = `
{
	"id": "42ac3eed-884c-4abf-8845-e304f5c10122",
	"group_email": "no-reply@zentral.com",
	"connection": "e4681fa5-5eb8-4e38-bc71-8a0c091468b1",
	"tags": [1],
	"created_at": "2022-07-22T01:02:03.444444",
	"updated_at": "2022-07-22T01:02:03.444444"
}
`

var gwsGroupTagMappingUpdateJSONResponse = `
{
	"id": "42ac3eed-884c-4abf-8845-e304f5c10122",
	"group_email": "no-reply@zentral.com",
	"connection": "e4681fa5-5eb8-4e38-bc71-8a0c091468b1",
	"tags": [1],
	"created_at": "2022-07-22T01:02:03.444444",
	"updated_at": "2022-07-22T01:02:03.444444"
}
`

var gwsGroupTagMappingGetJSONResponse = `
{
	"id": "42ac3eed-884c-4abf-8845-e304f5c10122",
	"group_email": "no-reply@zentral.com",
	"connection": "e4681fa5-5eb8-4e38-bc71-8a0c091468b1",
	"tags": [1],
	"created_at": "2022-07-22T01:02:03.444444",
	"updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestGWSGroupTagMappingsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/google_workspace/group_tag_mappings/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, gwsGroupTagMappingListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.GWSGroupTagMappings.List(ctx, nil)
	if err != nil {
		t.Errorf("GWSGroupTagMappings.List returned error: %v", err)
	}

	want := []GWSGroupTagMapping{
		{
			ID:           "42ac3eed-884c-4abf-8845-e304f5c10122",
			GroupEmail:   "no-reply@zentral.com",
			ConnectionID: "e4681fa5-5eb8-4e38-bc71-8a0c091468b1",
			TagIDs:       []int{1},
			Created:      Timestamp{referenceTime},
			Updated:      Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("GWSGroupTagMappings.List returned %+v, want %+v", got, want)
	}
}

func TestGWSGroupTagMappingsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/google_workspace/group_tag_mappings/42ac3eed-884c-4abf-8845-e304f5c10122/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, gwsGroupTagMappingGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.GWSGroupTagMappings.GetByID(ctx, "42ac3eed-884c-4abf-8845-e304f5c10122")
	if err != nil {
		t.Errorf("GWSGroupTagMappings.GetByID returned error: %v", err)
	}

	want := &GWSGroupTagMapping{
		ID:           "42ac3eed-884c-4abf-8845-e304f5c10122",
		GroupEmail:   "no-reply@zentral.com",
		ConnectionID: "e4681fa5-5eb8-4e38-bc71-8a0c091468b1",
		TagIDs:       []int{1},
		Created:      Timestamp{referenceTime},
		Updated:      Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("GWSGroupTagMappings.GetByID returned %+v, want %+v", got, want)
	}
}

func TestGWSGroupTagMappingsService_GetByGroupEmail(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/google_workspace/group_tag_mappings/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "group_email", "no-reply@zentral.com")
		fmt.Fprint(w, gwsGroupTagMappingGetByGroupEmailJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.GWSGroupTagMappings.GetByGroupEmail(ctx, "no-reply@zentral.com")
	if err != nil {
		t.Errorf("GWSGroupTagMappings.GetByGroupEmail returned error: %v", err)
	}

	want := []GWSGroupTagMapping{
		{
			ID:           "42ac3eed-884c-4abf-8845-e304f5c10122",
			GroupEmail:   "no-reply@zentral.com",
			ConnectionID: "e4681fa5-5eb8-4e38-bc71-8a0c091468b1",
			TagIDs:       []int{1},
			Created:      Timestamp{referenceTime},
			Updated:      Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("GWSGroupTagMappings.GetByGroupEmail returned %+v, want %+v", got, want)
	}
}

func TestGWSGroupTagMappingsService_GetConnectionID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/google_workspace/group_tag_mappings/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "connection_id", "e4681fa5-5eb8-4e38-bc71-8a0c091468b1")
		fmt.Fprint(w, gwsGroupTagMappingGetByConnectionIdJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.GWSGroupTagMappings.GetByConnectionID(ctx, "e4681fa5-5eb8-4e38-bc71-8a0c091468b1")
	if err != nil {
		t.Errorf("GWSGroupTagMappings.GetByConnectionID returned error: %v", err)
	}

	want := []GWSGroupTagMapping{
		{
			ID:           "42ac3eed-884c-4abf-8845-e304f5c10122",
			GroupEmail:   "no-reply@zentral.com",
			ConnectionID: "e4681fa5-5eb8-4e38-bc71-8a0c091468b1",
			TagIDs:       []int{1},
			Created:      Timestamp{referenceTime},
			Updated:      Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("GWSGroupTagMappings.GetByConnectionID returned %+v, want %+v", got, want)
	}
}

func TestGWSGroupTagMappingsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &GWSGroupTagMappingRequest{
		GroupEmail:   "no-reply@zentral.com",
		ConnectionID: "e4681fa5-5eb8-4e38-bc71-8a0c091468b1",
		TagIDs:       []int{1},
	}

	mux.HandleFunc("/google_workspace/group_tag_mappings/", func(w http.ResponseWriter, r *http.Request) {
		v := new(GWSGroupTagMappingRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, gwsGroupTagMappingCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.GWSGroupTagMappings.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("GWSGroupTagMappings.Create returned error: %v", err)
	}

	want := &GWSGroupTagMapping{
		ID:           "42ac3eed-884c-4abf-8845-e304f5c10122",
		GroupEmail:   "no-reply@zentral.com",
		ConnectionID: "e4681fa5-5eb8-4e38-bc71-8a0c091468b1",
		TagIDs:       []int{1},
		Created:      Timestamp{referenceTime},
		Updated:      Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("GWSGroupTagMappings.Create returned %+v, want %+v", got, want)
	}
}

func TestGWSGroupTagMappingsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &GWSGroupTagMappingRequest{
		GroupEmail:   "no-reply@zentral.com",
		ConnectionID: "e4681fa5-5eb8-4e38-bc71-8a0c091468b1",
		TagIDs:       []int{1},
	}

	mux.HandleFunc("/google_workspace/group_tag_mappings/42ac3eed-884c-4abf-8845-e304f5c10122/", func(w http.ResponseWriter, r *http.Request) {
		v := new(GWSGroupTagMappingRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, gwsGroupTagMappingUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.GWSGroupTagMappings.Update(ctx, "42ac3eed-884c-4abf-8845-e304f5c10122", updateRequest)
	if err != nil {
		t.Errorf("GWSGroupTagMappings.Update returned error: %v", err)
	}

	want := &GWSGroupTagMapping{
		ID:           "42ac3eed-884c-4abf-8845-e304f5c10122",
		GroupEmail:   "no-reply@zentral.com",
		ConnectionID: "e4681fa5-5eb8-4e38-bc71-8a0c091468b1",
		TagIDs:       []int{1},
		Created:      Timestamp{referenceTime},
		Updated:      Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("GWSGroupTagMappings.Update returned %+v, want %+v", got, want)
	}
}

func TestGWSGroupTagMappingsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/google_workspace/group_tag_mappings/42ac3eed-884c-4abf-8845-e304f5c10122/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.GWSGroupTagMappings.Delete(ctx, "42ac3eed-884c-4abf-8845-e304f5c10122")
	if err != nil {
		t.Errorf("GWSGroupTagMappings.Delete returned error: %v", err)
	}
}
