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

var trjListJSONResponse = `
{
    "count": 1,
    "next": null,
    "previous": null,
    "results": [
    {
        "id": "e1e2e3e4-0000-1111-2222-333344445555",
        "configuration": "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
        "job": "b0000000-1111-2222-3333-444455556666",
        "interval": null,
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

var trjGetJSONResponse = `
{
    "id": "e1e2e3e4-0000-1111-2222-333344445555",
    "configuration": "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
    "job": "b0000000-1111-2222-3333-444455556666",
    "interval": 3600,
    "tags": [6, 7],
    "excluded_tags": [8],
    "serial_numbers": ["huit", "neuf"],
    "excluded_serial_numbers": ["dix"],
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var trjCreateJSONResponse = `
{
    "id": "e1e2e3e4-0000-1111-2222-333344445555",
    "configuration": "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
    "job": "b0000000-1111-2222-3333-444455556666",
    "interval": 3600,
    "tags": [6, 7],
    "excluded_tags": [8],
    "serial_numbers": ["huit", "neuf"],
    "excluded_serial_numbers": ["dix"],
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var trjUpdateJSONResponse = `
{
    "id": "e1e2e3e4-0000-1111-2222-333344445555",
    "configuration": "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
    "job": "b0000000-1111-2222-3333-444455556666",
    "interval": 7200,
    "tags": [6, 7],
    "excluded_tags": [8],
    "serial_numbers": ["huit", "neuf"],
    "excluded_serial_numbers": ["dix"],
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestTurboRecurringJobsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/turbo/recurring_jobs/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, trjListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.TurboRecurringJobs.List(ctx, nil)
	if err != nil {
		t.Errorf("TurboRecurringJobs.List returned error: %v", err)
	}

	want := []TurboRecurringJob{
		{
			ID:                    "e1e2e3e4-0000-1111-2222-333344445555",
			ConfigurationID:       "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
			JobID:                 "b0000000-1111-2222-3333-444455556666",
			Interval:              nil,
			TagIDs:                []int{},
			ExcludedTagIDs:        []int{},
			SerialNumbers:         []string{},
			ExcludedSerialNumbers: []string{},
			Created:               Timestamp{referenceTime},
			Updated:               Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("TurboRecurringJobs.List returned %+v, want %+v", got, want)
	}
}

func TestTurboRecurringJobsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/turbo/recurring_jobs/e1e2e3e4-0000-1111-2222-333344445555/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, trjGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.TurboRecurringJobs.GetByID(ctx, "e1e2e3e4-0000-1111-2222-333344445555")
	if err != nil {
		t.Errorf("TurboRecurringJobs.GetByID returned error: %v", err)
	}

	want := &TurboRecurringJob{
		ID:                    "e1e2e3e4-0000-1111-2222-333344445555",
		ConfigurationID:       "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
		JobID:                 "b0000000-1111-2222-3333-444455556666",
		Interval:              Int(3600),
		TagIDs:                []int{6, 7},
		ExcludedTagIDs:        []int{8},
		SerialNumbers:         []string{"huit", "neuf"},
		ExcludedSerialNumbers: []string{"dix"},
		Created:               Timestamp{referenceTime},
		Updated:               Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("TurboRecurringJobs.GetByID returned %+v, want %+v", got, want)
	}
}

func TestTurboRecurringJobsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &TurboRecurringJobRequest{
		ConfigurationID:       "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
		JobID:                 "b0000000-1111-2222-3333-444455556666",
		Interval:              Int(3600),
		TagIDs:                []int{6, 7},
		ExcludedTagIDs:        []int{8},
		SerialNumbers:         []string{"huit", "neuf"},
		ExcludedSerialNumbers: []string{"dix"},
	}

	mux.HandleFunc("/turbo/recurring_jobs/", func(w http.ResponseWriter, r *http.Request) {
		v := new(TurboRecurringJobRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, trjCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.TurboRecurringJobs.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("TurboRecurringJobs.Create returned error: %v", err)
	}

	want := &TurboRecurringJob{
		ID:                    "e1e2e3e4-0000-1111-2222-333344445555",
		ConfigurationID:       "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
		JobID:                 "b0000000-1111-2222-3333-444455556666",
		Interval:              Int(3600),
		TagIDs:                []int{6, 7},
		ExcludedTagIDs:        []int{8},
		SerialNumbers:         []string{"huit", "neuf"},
		ExcludedSerialNumbers: []string{"dix"},
		Created:               Timestamp{referenceTime},
		Updated:               Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("TurboRecurringJobs.Create returned %+v, want %+v", got, want)
	}
}

func TestTurboRecurringJobsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &TurboRecurringJobRequest{
		ConfigurationID:       "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
		JobID:                 "b0000000-1111-2222-3333-444455556666",
		Interval:              Int(7200),
		TagIDs:                []int{6, 7},
		ExcludedTagIDs:        []int{8},
		SerialNumbers:         []string{"huit", "neuf"},
		ExcludedSerialNumbers: []string{"dix"},
	}

	mux.HandleFunc("/turbo/recurring_jobs/e1e2e3e4-0000-1111-2222-333344445555/", func(w http.ResponseWriter, r *http.Request) {
		v := new(TurboRecurringJobRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, trjUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.TurboRecurringJobs.Update(ctx, "e1e2e3e4-0000-1111-2222-333344445555", updateRequest)
	if err != nil {
		t.Errorf("TurboRecurringJobs.Update returned error: %v", err)
	}

	want := &TurboRecurringJob{
		ID:                    "e1e2e3e4-0000-1111-2222-333344445555",
		ConfigurationID:       "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
		JobID:                 "b0000000-1111-2222-3333-444455556666",
		Interval:              Int(7200),
		TagIDs:                []int{6, 7},
		ExcludedTagIDs:        []int{8},
		SerialNumbers:         []string{"huit", "neuf"},
		ExcludedSerialNumbers: []string{"dix"},
		Created:               Timestamp{referenceTime},
		Updated:               Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("TurboRecurringJobs.Update returned %+v, want %+v", got, want)
	}
}

func TestTurboRecurringJobsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/turbo/recurring_jobs/e1e2e3e4-0000-1111-2222-333344445555/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.TurboRecurringJobs.Delete(ctx, "e1e2e3e4-0000-1111-2222-333344445555")
	if err != nil {
		t.Errorf("TurboRecurringJobs.Delete returned error: %v", err)
	}
}
