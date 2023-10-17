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

var oqListJSONResponse = `
[
    {
        "id": 4,
        "name": "Users",
	"sql": "SELECT * FROM users;",
	"platforms": ["darwin", "linux", "windows"],
	"description": "List all users",
	"value": "A list of user attributes",
	"version": 1,
	"scheduling": {
	    "can_be_denylisted": true,
	    "interval": 60,
	    "log_removed_actions": false,
	    "pack": 2,
	    "shard": 10,
	    "snapshot_mode": false
	},
	"compliance_check_enabled": false,
	"tag": null,
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var oqGetJSONResponse = `
{
    "id": 4,
    "name": "Users",
    "sql": "SELECT * FROM users;",
    "platforms": ["darwin", "linux", "windows"],
    "minimum_osquery_version": null,
    "description": "List all users",
    "value": "A list of user attributes",
    "version": 1,
    "compliance_check_enabled": false,
    "tag": 17,
    "scheduling": null,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var oqCreateJSONResponse = `
{
    "id": 4,
    "name": "Users",
    "sql": "SELECT * FROM users;",
    "platforms": ["darwin", "linux", "windows"],
    "minimum_osquery_version": "0.1.0",
    "description": "List all users",
    "value": "A list of user attributes",
    "version": 1,
    "scheduling": null,
    "compliance_check_enabled": false,
    "tag": 17,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var oqUpdateJSONResponse = `
{
    "id": 4,
    "name": "Users",
    "sql": "SELECT * FROM users;",
    "platforms": ["darwin", "linux", "windows"],
    "minimum_osquery_version": "0.1.0",
    "description": "List all users",
    "value": "A list of user attributes",
    "version": 1,
    "compliance_check_enabled": false,
    "tag": null,
    "scheduling": {
	"can_be_denylisted": true,
	"interval": 161,
	"log_removed_actions": true,
	"pack": 2,
	"shard": null,
	"snapshot_mode": false
    },
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestOsqueryQueriesService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/queries/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, oqListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryQueries.List(ctx, nil)
	if err != nil {
		t.Errorf("OsqueryQueries.List returned error: %v", err)
	}

	want := []OsqueryQuery{
		{
			ID:                     4,
			Name:                   "Users",
			SQL:                    "SELECT * FROM users;",
			Platforms:              []string{"darwin", "linux", "windows"},
			Description:            "List all users",
			Value:                  "A list of user attributes",
			Version:                1,
			ComplianceCheckEnabled: false,
			Scheduling: &OsqueryQueryScheduling{
				CanBeDenyListed:   true,
				Interval:          60,
				LogRemovedActions: false,
				PackID:            2,
				Shard:             Int(10),
				SnapshotMode:      false,
			},
			Created: Timestamp{referenceTime},
			Updated: Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryQueries.List returned %+v, want %+v", got, want)
	}
}

func TestOsqueryQueriesService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/queries/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, oqGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryQueries.GetByID(ctx, 1)
	if err != nil {
		t.Errorf("OsqueryQueries.GetByID returned error: %v", err)
	}

	want := &OsqueryQuery{
		ID:                     4,
		Name:                   "Users",
		SQL:                    "SELECT * FROM users;",
		Platforms:              []string{"darwin", "linux", "windows"},
		Description:            "List all users",
		Value:                  "A list of user attributes",
		Version:                1,
		ComplianceCheckEnabled: false,
		TagID:                  Int(17),
		Created:                Timestamp{referenceTime},
		Updated:                Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryQueries.GetByID returned %+v, want %+v", got, want)
	}
}

func TestOsqueryQueriesService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/queries/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "Users")
		fmt.Fprint(w, oqListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryQueries.GetByName(ctx, "Users")
	if err != nil {
		t.Errorf("OsqueryQueries.GetByName returned error: %v", err)
	}

	want := &OsqueryQuery{
		ID:                     4,
		Name:                   "Users",
		SQL:                    "SELECT * FROM users;",
		Platforms:              []string{"darwin", "linux", "windows"},
		Description:            "List all users",
		Value:                  "A list of user attributes",
		Version:                1,
		ComplianceCheckEnabled: false,
		Scheduling: &OsqueryQueryScheduling{
			CanBeDenyListed:   true,
			Interval:          60,
			LogRemovedActions: false,
			PackID:            2,
			Shard:             Int(10),
			SnapshotMode:      false,
		},
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryQueries.GetByName returned %+v, want %+v", got, want)
	}
}

func TestOsqueryQueriesService_GetByPackID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/queries/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "pack_id", "2")
		fmt.Fprint(w, oqListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryQueries.GetByPackID(ctx, 2)
	if err != nil {
		t.Errorf("OsqueryQueries.GetByPackID returned error: %v", err)
	}

	want := []OsqueryQuery{
		{
			ID:                     4,
			Name:                   "Users",
			SQL:                    "SELECT * FROM users;",
			Platforms:              []string{"darwin", "linux", "windows"},
			Description:            "List all users",
			Value:                  "A list of user attributes",
			Version:                1,
			ComplianceCheckEnabled: false,
			Scheduling: &OsqueryQueryScheduling{
				CanBeDenyListed:   true,
				Interval:          60,
				LogRemovedActions: false,
				PackID:            2,
				Shard:             Int(10),
				SnapshotMode:      false,
			},
			Created: Timestamp{referenceTime},
			Updated: Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryQueries.GetByName returned %+v, want %+v", got, want)
	}
}

func TestOsqueryQueriesService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &OsqueryQueryRequest{
		Name:                   "Users",
		SQL:                    "SELECT * FROM users;",
		Platforms:              []string{"darwin", "linux", "windows"},
		MinOsqueryVersion:      String("0.1.0"),
		Description:            "List all users",
		Value:                  "A list of user attributes",
		ComplianceCheckEnabled: false,
	}

	mux.HandleFunc("/osquery/queries/", func(w http.ResponseWriter, r *http.Request) {
		v := new(OsqueryQueryRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, oqCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryQueries.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("OsqueryQueries.Create returned error: %v", err)
	}

	want := &OsqueryQuery{
		ID:                     4,
		Name:                   "Users",
		SQL:                    "SELECT * FROM users;",
		Platforms:              []string{"darwin", "linux", "windows"},
		MinOsqueryVersion:      String("0.1.0"),
		Description:            "List all users",
		Value:                  "A list of user attributes",
		Version:                1,
		ComplianceCheckEnabled: false,
		TagID:                  Int(17),
		Created:                Timestamp{referenceTime},
		Updated:                Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryQueries.Create returned %+v, want %+v", got, want)
	}
}

func TestOsqueryQueriesService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &OsqueryQueryRequest{
		Name:                   "Users",
		SQL:                    "SELECT * FROM users;",
		Platforms:              []string{"darwin", "linux", "windows"},
		MinOsqueryVersion:      String("0.1.0"),
		Description:            "List all users",
		Value:                  "A list of user attributes",
		ComplianceCheckEnabled: false,
		Scheduling: &OsqueryQuerySchedulingRequest{
			Interval: 161,
			PackID:   2,
		},
	}

	mux.HandleFunc("/osquery/queries/1/", func(w http.ResponseWriter, r *http.Request) {
		v := new(OsqueryQueryRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, oqUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryQueries.Update(ctx, 1, updateRequest)
	if err != nil {
		t.Errorf("OsqueryQueries.Update returned error: %v", err)
	}

	want := &OsqueryQuery{
		ID:                     4,
		Name:                   "Users",
		SQL:                    "SELECT * FROM users;",
		Platforms:              []string{"darwin", "linux", "windows"},
		MinOsqueryVersion:      String("0.1.0"),
		Description:            "List all users",
		Value:                  "A list of user attributes",
		Version:                1,
		ComplianceCheckEnabled: false,
		Scheduling: &OsqueryQueryScheduling{
			CanBeDenyListed:   true,
			Interval:          161,
			LogRemovedActions: true,
			PackID:            2,
			SnapshotMode:      false,
		},
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryQueries.Update returned %+v, want %+v", got, want)
	}
}

func TestOsqueryQueriesService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/queries/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.OsqueryQueries.Delete(ctx, 1)
	if err != nil {
		t.Errorf("OsqueryQueries.Delete returned error: %v", err)
	}
}
