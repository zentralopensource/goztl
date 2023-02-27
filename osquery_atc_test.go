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

var oaListJSONResponse = `
[
    {
        "id": 4,
        "name": "Santa rules",
	"description": "Access the Google Santa rules.db",
	"table_name": "santa_rules",
	"query": "SELECT * FROM rules;",
	"path": "/var/db/santa/rules.db",
	"columns": [
            "identifier",
	    "state",
	    "type",
	    "custommsg",
	    "timestamp"
	],
	"platforms": ["darwin"],
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var oaGetJSONResponse = `
{
    "id": 4,
    "name": "Santa rules",
    "description": "Access the Google Santa rules.db",
    "table_name": "santa_rules",
    "query": "SELECT * FROM rules;",
    "path": "/var/db/santa/rules.db",
    "columns": [
	"identifier",
	"state",
	"type",
	"custommsg",
	"timestamp"
    ],
    "platforms": ["darwin"],
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var oaCreateJSONResponse = `
{
    "id": 4,
    "name": "Santa rules",
    "description": "Access the Google Santa rules.db",
    "table_name": "santa_rules",
    "query": "SELECT * FROM rules;",
    "path": "/var/db/santa/rules.db",
    "columns": [
	"identifier",
	"state",
	"type",
	"custommsg",
	"timestamp"
    ],
    "platforms": ["darwin"],
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var oaUpdateJSONResponse = `
{
    "id": 4,
    "name": "Santa rules",
    "description": "Access the Google Santa rules.db",
    "table_name": "santa_rules",
    "query": "SELECT * FROM rules;",
    "path": "/var/db/santa/rules.db",
    "columns": [
	"identifier",
	"state",
	"type",
	"custommsg",
	"timestamp"
    ],
    "platforms": ["darwin"],
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestOsqueryATCService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/atcs/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, oaListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryATC.List(ctx, nil)
	if err != nil {
		t.Errorf("OsqueryATC.List returned error: %v", err)
	}

	want := []OsqueryATC{
		{
			ID:          4,
			Name:        "Santa rules",
			Description: "Access the Google Santa rules.db",
			TableName:   "santa_rules",
			Query:       "SELECT * FROM rules;",
			Path:        "/var/db/santa/rules.db",
			Columns:     []string{"identifier", "state", "type", "custommsg", "timestamp"},
			Platforms:   []string{"darwin"},
			Created:     Timestamp{referenceTime},
			Updated:     Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryATC.List returned %+v, want %+v", got, want)
	}
}

func TestOsqueryATCService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/atcs/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, oaGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryATC.GetByID(ctx, 1)
	if err != nil {
		t.Errorf("OsqueryATC.GetByID returned error: %v", err)
	}

	want := &OsqueryATC{
		ID:          4,
		Name:        "Santa rules",
		Description: "Access the Google Santa rules.db",
		TableName:   "santa_rules",
		Query:       "SELECT * FROM rules;",
		Path:        "/var/db/santa/rules.db",
		Columns:     []string{"identifier", "state", "type", "custommsg", "timestamp"},
		Platforms:   []string{"darwin"},
		Created:     Timestamp{referenceTime},
		Updated:     Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryATC.GetByID returned %+v, want %+v", got, want)
	}
}

func TestOsqueryATCService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/atcs/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "Santa rules")
		fmt.Fprint(w, oaListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryATC.GetByName(ctx, "Santa rules")
	if err != nil {
		t.Errorf("OsqueryATC.GetByName returned error: %v", err)
	}

	want := &OsqueryATC{
		ID:          4,
		Name:        "Santa rules",
		Description: "Access the Google Santa rules.db",
		TableName:   "santa_rules",
		Query:       "SELECT * FROM rules;",
		Path:        "/var/db/santa/rules.db",
		Columns:     []string{"identifier", "state", "type", "custommsg", "timestamp"},
		Platforms:   []string{"darwin"},
		Created:     Timestamp{referenceTime},
		Updated:     Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryATC.GetByName returned %+v, want %+v", got, want)
	}
}

func TestOsqueryATCService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &OsqueryATCRequest{
		Name:        "Santa rules",
		Description: "Access the Google Santa rules.db",
		TableName:   "santa_rules",
		Query:       "SELECT * FROM rules;",
		Path:        "/var/db/santa/rules.db",
		Columns:     []string{"identifier", "state", "type", "custommsg", "timestamp"},
		Platforms:   []string{"darwin"},
	}

	mux.HandleFunc("/osquery/atcs/", func(w http.ResponseWriter, r *http.Request) {
		v := new(OsqueryATCRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, oaCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryATC.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("OsqueryATC.Create returned error: %v", err)
	}

	want := &OsqueryATC{
		ID:          4,
		Name:        "Santa rules",
		Description: "Access the Google Santa rules.db",
		TableName:   "santa_rules",
		Query:       "SELECT * FROM rules;",
		Path:        "/var/db/santa/rules.db",
		Columns:     []string{"identifier", "state", "type", "custommsg", "timestamp"},
		Platforms:   []string{"darwin"},
		Created:     Timestamp{referenceTime},
		Updated:     Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryATC.Create returned %+v, want %+v", got, want)
	}
}

func TestOsqueryATCService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &OsqueryATCRequest{
		Name:        "Santa rules",
		Description: "Access the Google Santa rules.db",
		TableName:   "santa_rules",
		Query:       "SELECT * FROM rules;",
		Path:        "/var/db/santa/rules.db",
		Columns:     []string{"identifier", "state", "type", "custommsg", "timestamp"},
		Platforms:   []string{"darwin"},
	}

	mux.HandleFunc("/osquery/atcs/1/", func(w http.ResponseWriter, r *http.Request) {
		v := new(OsqueryATCRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, oaUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryATC.Update(ctx, 1, updateRequest)
	if err != nil {
		t.Errorf("OsqueryATC.Update returned error: %v", err)
	}

	want := &OsqueryATC{
		ID:          4,
		Name:        "Santa rules",
		Description: "Access the Google Santa rules.db",
		TableName:   "santa_rules",
		Query:       "SELECT * FROM rules;",
		Path:        "/var/db/santa/rules.db",
		Columns:     []string{"identifier", "state", "type", "custommsg", "timestamp"},
		Platforms:   []string{"darwin"},
		Created:     Timestamp{referenceTime},
		Updated:     Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryATC.Update returned %+v, want %+v", got, want)
	}
}

func TestOsqueryATCService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/atcs/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.OsqueryATC.Delete(ctx, 1)
	if err != nil {
		t.Errorf("OsqueryATC.Delete returned error: %v", err)
	}
}
