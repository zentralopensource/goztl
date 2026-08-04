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

var totjListJSONResponse = `
{
    "count": 1,
    "next": null,
    "previous": null,
    "results": [
    {
        "id": "f1f2f3f4-0000-1111-2222-333344445555",
        "configuration": "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
        "job": "b0000000-1111-2222-3333-444455556666",
        "not_before": null,
        "not_after": null,
        "tags": [],
        "excluded_tags": [],
        "serial_numbers": [],
        "excluded_serial_numbers": [],
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
    ]
}
`

var totjGetJSONResponse = `
{
    "id": "f1f2f3f4-0000-1111-2222-333344445555",
    "configuration": "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
    "job": "b0000000-1111-2222-3333-444455556666",
    "not_before": "2026-08-01T09:00:00Z",
    "not_after": "2026-08-31T09:00:00Z",
    "tags": [6, 7],
    "excluded_tags": [8],
    "serial_numbers": ["huit", "neuf"],
    "excluded_serial_numbers": ["dix"],
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var totjCreateJSONResponse = `
{
    "id": "f1f2f3f4-0000-1111-2222-333344445555",
    "configuration": "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
    "job": "b0000000-1111-2222-3333-444455556666",
    "not_before": "2026-08-01T09:00:00Z",
    "not_after": "2026-08-31T09:00:00Z",
    "tags": [6, 7],
    "excluded_tags": [8],
    "serial_numbers": ["huit", "neuf"],
    "excluded_serial_numbers": ["dix"],
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var totjUpdateJSONResponse = `
{
    "id": "f1f2f3f4-0000-1111-2222-333344445555",
    "configuration": "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
    "job": "b0000000-1111-2222-3333-444455556666",
    "not_before": null,
    "not_after": "2026-09-30T09:00:00Z",
    "tags": [6, 7],
    "excluded_tags": [8],
    "serial_numbers": ["huit", "neuf"],
    "excluded_serial_numbers": ["dix"],
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestTurboOneTimeJobsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/turbo/one_time_jobs/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, totjListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.TurboOneTimeJobs.List(ctx, nil)
	if err != nil {
		t.Errorf("TurboOneTimeJobs.List returned error: %v", err)
	}

	want := []TurboOneTimeJob{
		{
			ID:                    "f1f2f3f4-0000-1111-2222-333344445555",
			ConfigurationID:       "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
			JobID:                 "b0000000-1111-2222-3333-444455556666",
			NotBefore:             nil,
			NotAfter:              nil,
			TagIDs:                []int{},
			ExcludedTagIDs:        []int{},
			SerialNumbers:         []string{},
			ExcludedSerialNumbers: []string{},
			Created:               Timestamp{referenceTime},
			Updated:               Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("TurboOneTimeJobs.List returned %+v, want %+v", got, want)
	}
}

func TestTurboOneTimeJobsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/turbo/one_time_jobs/f1f2f3f4-0000-1111-2222-333344445555/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, totjGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.TurboOneTimeJobs.GetByID(ctx, "f1f2f3f4-0000-1111-2222-333344445555")
	if err != nil {
		t.Errorf("TurboOneTimeJobs.GetByID returned error: %v", err)
	}

	want := &TurboOneTimeJob{
		ID:                    "f1f2f3f4-0000-1111-2222-333344445555",
		ConfigurationID:       "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
		JobID:                 "b0000000-1111-2222-3333-444455556666",
		NotBefore:             String("2026-08-01T09:00:00Z"),
		NotAfter:              String("2026-08-31T09:00:00Z"),
		TagIDs:                []int{6, 7},
		ExcludedTagIDs:        []int{8},
		SerialNumbers:         []string{"huit", "neuf"},
		ExcludedSerialNumbers: []string{"dix"},
		Created:               Timestamp{referenceTime},
		Updated:               Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("TurboOneTimeJobs.GetByID returned %+v, want %+v", got, want)
	}
}

func TestTurboOneTimeJobsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &TurboOneTimeJobRequest{
		ConfigurationID:       "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
		JobID:                 "b0000000-1111-2222-3333-444455556666",
		NotBefore:             String("2026-08-01T09:00:00Z"),
		NotAfter:              String("2026-08-31T09:00:00Z"),
		TagIDs:                []int{6, 7},
		ExcludedTagIDs:        []int{8},
		SerialNumbers:         []string{"huit", "neuf"},
		ExcludedSerialNumbers: []string{"dix"},
	}

	mux.HandleFunc("/turbo/one_time_jobs/", func(w http.ResponseWriter, r *http.Request) {
		v := new(TurboOneTimeJobRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, totjCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.TurboOneTimeJobs.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("TurboOneTimeJobs.Create returned error: %v", err)
	}

	want := &TurboOneTimeJob{
		ID:                    "f1f2f3f4-0000-1111-2222-333344445555",
		ConfigurationID:       "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
		JobID:                 "b0000000-1111-2222-3333-444455556666",
		NotBefore:             String("2026-08-01T09:00:00Z"),
		NotAfter:              String("2026-08-31T09:00:00Z"),
		TagIDs:                []int{6, 7},
		ExcludedTagIDs:        []int{8},
		SerialNumbers:         []string{"huit", "neuf"},
		ExcludedSerialNumbers: []string{"dix"},
		Created:               Timestamp{referenceTime},
		Updated:               Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("TurboOneTimeJobs.Create returned %+v, want %+v", got, want)
	}
}

func TestTurboOneTimeJobsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &TurboOneTimeJobRequest{
		ConfigurationID:       "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
		JobID:                 "b0000000-1111-2222-3333-444455556666",
		NotBefore:             nil,
		NotAfter:              String("2026-09-30T09:00:00Z"),
		TagIDs:                []int{6, 7},
		ExcludedTagIDs:        []int{8},
		SerialNumbers:         []string{"huit", "neuf"},
		ExcludedSerialNumbers: []string{"dix"},
	}

	mux.HandleFunc("/turbo/one_time_jobs/f1f2f3f4-0000-1111-2222-333344445555/", func(w http.ResponseWriter, r *http.Request) {
		v := new(TurboOneTimeJobRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, totjUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.TurboOneTimeJobs.Update(ctx, "f1f2f3f4-0000-1111-2222-333344445555", updateRequest)
	if err != nil {
		t.Errorf("TurboOneTimeJobs.Update returned error: %v", err)
	}

	want := &TurboOneTimeJob{
		ID:                    "f1f2f3f4-0000-1111-2222-333344445555",
		ConfigurationID:       "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
		JobID:                 "b0000000-1111-2222-3333-444455556666",
		NotBefore:             nil,
		NotAfter:              String("2026-09-30T09:00:00Z"),
		TagIDs:                []int{6, 7},
		ExcludedTagIDs:        []int{8},
		SerialNumbers:         []string{"huit", "neuf"},
		ExcludedSerialNumbers: []string{"dix"},
		Created:               Timestamp{referenceTime},
		Updated:               Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("TurboOneTimeJobs.Update returned %+v, want %+v", got, want)
	}
}

func TestTurboOneTimeJobsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/turbo/one_time_jobs/f1f2f3f4-0000-1111-2222-333344445555/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.TurboOneTimeJobs.Delete(ctx, "f1f2f3f4-0000-1111-2222-333344445555")
	if err != nil {
		t.Errorf("TurboOneTimeJobs.Delete returned error: %v", err)
	}
}
