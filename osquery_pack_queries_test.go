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

var opqListJSONResponse = `
[
    {
        "id": 4,
	"pack": 5,
	"query": 6,
	"slug": "users",
	"interval": 60,
	"log_removed_actions": true,
	"snapshot_mode": false,
	"can_be_denylisted": true,
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var opqGetJSONResponse = `
{
    "id": 4,
    "pack": 5,
    "query": 6,
    "slug": "users",
    "interval": 60,
    "log_removed_actions": true,
    "snapshot_mode": false,
    "shard": 100,
    "can_be_denylisted": true,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var opqCreateJSONResponse = `
{
    "id": 4,
    "pack": 5,
    "query": 6,
    "slug": "users",
    "interval": 60,
    "log_removed_actions": true,
    "snapshot_mode": false,
    "shard": 100,
    "can_be_denylisted": true,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var opqUpdateJSONResponse = `
{
    "id": 4,
    "pack": 5,
    "query": 6,
    "slug": "users",
    "interval": 1200,
    "log_removed_actions": false,
    "snapshot_mode": true,
    "shard": 100,
    "can_be_denylisted": true,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestOsqueryPackQueriesService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/pack_queries/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, opqListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryPackQueries.List(ctx, nil)
	if err != nil {
		t.Errorf("OsqueryPackQueries.List returned error: %v", err)
	}

	want := []OsqueryPackQuery{
		{
			ID:                4,
			PackID:            5,
			QueryID:           6,
			Slug:              "users",
			Interval:          60,
			LogRemovedActions: true,
			SnapshotMode:      false,
			CanBeDenyListed:   true,
			Created:           Timestamp{referenceTime},
			Updated:           Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryPackQueries.List returned %+v, want %+v", got, want)
	}
}

func TestOsqueryPackQueriesService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/pack_queries/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, opqGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryPackQueries.GetByID(ctx, 4)
	if err != nil {
		t.Errorf("OsqueryPackQueries.GetByID returned error: %v", err)
	}

	want := &OsqueryPackQuery{
		ID:                4,
		PackID:            5,
		QueryID:           6,
		Slug:              "users",
		Interval:          60,
		LogRemovedActions: true,
		SnapshotMode:      false,
		Shard:             Int(100),
		CanBeDenyListed:   true,
		Created:           Timestamp{referenceTime},
		Updated:           Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryPackQueries.GetByID returned %+v, want %+v", got, want)
	}
}

func TestOsqueryPackQueriesService_GetByPackID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/pack_queries/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "pack_id", "4")
		fmt.Fprint(w, opqListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryPackQueries.GetByPackID(ctx, 4)
	if err != nil {
		t.Errorf("OsqueryPackQueries.GetByPackID returned error: %v", err)
	}

	want := []OsqueryPackQuery{
		{
			ID:                4,
			PackID:            5,
			QueryID:           6,
			Slug:              "users",
			Interval:          60,
			LogRemovedActions: true,
			SnapshotMode:      false,
			CanBeDenyListed:   true,
			Created:           Timestamp{referenceTime},
			Updated:           Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryPackQueries.GetByPackID returned %+v, want %+v", got, want)
	}
}

func TestOsqueryPackQueriesService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &OsqueryPackQueryRequest{
		PackID:            5,
		QueryID:           6,
		Interval:          60,
		LogRemovedActions: true,
		SnapshotMode:      false,
		Shard:             Int(100),
		CanBeDenyListed:   true,
	}

	mux.HandleFunc("/osquery/pack_queries/", func(w http.ResponseWriter, r *http.Request) {
		v := new(OsqueryPackQueryRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, opqCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryPackQueries.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("OsqueryPackQueries.Create returned error: %v", err)
	}

	want := &OsqueryPackQuery{
		ID:                4,
		PackID:            5,
		QueryID:           6,
		Slug:              "users",
		Interval:          60,
		LogRemovedActions: true,
		SnapshotMode:      false,
		Shard:             Int(100),
		CanBeDenyListed:   true,
		Created:           Timestamp{referenceTime},
		Updated:           Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryPackQueries.Create returned %+v, want %+v", got, want)
	}
}

func TestOsqueryPackQueriesService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &OsqueryPackQueryRequest{
		PackID:            5,
		QueryID:           6,
		Interval:          1200,
		LogRemovedActions: false,
		SnapshotMode:      true,
		Shard:             Int(100),
		CanBeDenyListed:   true,
	}

	mux.HandleFunc("/osquery/pack_queries/4/", func(w http.ResponseWriter, r *http.Request) {
		v := new(OsqueryPackQueryRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, opqUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryPackQueries.Update(ctx, 4, updateRequest)
	if err != nil {
		t.Errorf("OsqueryPackQueries.Update returned error: %v", err)
	}

	want := &OsqueryPackQuery{
		ID:                4,
		PackID:            5,
		QueryID:           6,
		Slug:              "users",
		Interval:          1200,
		LogRemovedActions: false,
		SnapshotMode:      true,
		Shard:             Int(100),
		CanBeDenyListed:   true,
		Created:           Timestamp{referenceTime},
		Updated:           Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryPackQueries.Update returned %+v, want %+v", got, want)
	}
}

func TestOsqueryPackQueriesService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/pack_queries/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.OsqueryPackQueries.Delete(ctx, 4)
	if err != nil {
		t.Errorf("OsqueryPackQueries.Delete returned error: %v", err)
	}
}
