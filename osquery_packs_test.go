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

var opListJSONResponse = `
[
    {
        "id": 4,
        "name": "Default",
	"slug": "default",
	"description": "Default query pack",
	"discovery_queries": [
	    "SELECT pid FROM processes WHERE name = 'ldap';"
	],
	"event_routing_key": "",
	"created_at": "2022-07-22T01:02:03.444444",
	"updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var opGetJSONResponse = `
{
    "id": 4,
    "name": "Default",
    "slug": "default",
    "description": "Default query pack",
    "discovery_queries": [
	"SELECT pid FROM processes WHERE name = 'ldap';"
    ],
    "shard": null,
    "event_routing_key": "",
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var opCreateJSONResponse = `
{
    "id": 4,
    "name": "Default",
    "slug": "default",
    "description": "Default query pack",
    "discovery_queries": [
	"SELECT pid FROM processes WHERE name = 'ldap';"
    ],
    "shard": 10,
    "event_routing_key": "important",
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var opUpdateJSONResponse = `
{
    "id": 4,
    "name": "Default",
    "slug": "default",
    "description": "Default query pack",
    "discovery_queries": [
	"SELECT pid FROM processes WHERE name = 'ldap';"
    ],
    "shard": 100,
    "event_routing_key": "important",
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestOsqueryPacksService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/packs/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, opListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryPacks.List(ctx, nil)
	if err != nil {
		t.Errorf("OsqueryPacks.List returned error: %v", err)
	}

	want := []OsqueryPack{
		{
			ID:               4,
			Name:             "Default",
			Slug:             "default",
			Description:      "Default query pack",
			DiscoveryQueries: []string{"SELECT pid FROM processes WHERE name = 'ldap';"},
			EventRoutingKey:  "",
			Created:          Timestamp{referenceTime},
			Updated:          Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryPacks.List returned %+v, want %+v", got, want)
	}
}

func TestOsqueryPacksService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/packs/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, opGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryPacks.GetByID(ctx, 1)
	if err != nil {
		t.Errorf("OsqueryPacks.GetByID returned error: %v", err)
	}

	want := &OsqueryPack{
		ID:               4,
		Name:             "Default",
		Slug:             "default",
		Description:      "Default query pack",
		DiscoveryQueries: []string{"SELECT pid FROM processes WHERE name = 'ldap';"},
		EventRoutingKey:  "",
		Created:          Timestamp{referenceTime},
		Updated:          Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryPacks.GetByID returned %+v, want %+v", got, want)
	}
}

func TestOsqueryPacksService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/packs/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "Default")
		fmt.Fprint(w, opListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryPacks.GetByName(ctx, "Default")
	if err != nil {
		t.Errorf("OsqueryPacks.GetByName returned error: %v", err)
	}

	want := &OsqueryPack{
		ID:               4,
		Name:             "Default",
		Slug:             "default",
		Description:      "Default query pack",
		DiscoveryQueries: []string{"SELECT pid FROM processes WHERE name = 'ldap';"},
		EventRoutingKey:  "",
		Created:          Timestamp{referenceTime},
		Updated:          Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryPacks.GetByName returned %+v, want %+v", got, want)
	}
}

func TestOsqueryPacksService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &OsqueryPackRequest{
		Name:             "Default",
		Description:      "Default query pack",
		DiscoveryQueries: []string{"SELECT pid FROM processes WHERE name = 'ldap';"},
		Shard:            Int(10),
		EventRoutingKey:  "important",
	}

	mux.HandleFunc("/osquery/packs/", func(w http.ResponseWriter, r *http.Request) {
		v := new(OsqueryPackRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, opCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryPacks.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("OsqueryPacks.Create returned error: %v", err)
	}

	want := &OsqueryPack{
		ID:               4,
		Name:             "Default",
		Slug:             "default",
		Description:      "Default query pack",
		DiscoveryQueries: []string{"SELECT pid FROM processes WHERE name = 'ldap';"},
		Shard:            Int(10),
		EventRoutingKey:  "important",
		Created:          Timestamp{referenceTime},
		Updated:          Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryPacks.Create returned %+v, want %+v", got, want)
	}
}

func TestOsqueryPacksService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &OsqueryPackRequest{
		Name:             "Default",
		Description:      "Default query pack",
		DiscoveryQueries: []string{"SELECT pid FROM processes WHERE name = 'ldap';"},
		Shard:            Int(100),
		EventRoutingKey:  "important",
	}

	mux.HandleFunc("/osquery/packs/1/", func(w http.ResponseWriter, r *http.Request) {
		v := new(OsqueryPackRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, opUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryPacks.Update(ctx, 1, updateRequest)
	if err != nil {
		t.Errorf("OsqueryPacks.Update returned error: %v", err)
	}

	want := &OsqueryPack{
		ID:               4,
		Name:             "Default",
		Slug:             "default",
		Description:      "Default query pack",
		DiscoveryQueries: []string{"SELECT pid FROM processes WHERE name = 'ldap';"},
		Shard:            Int(100),
		EventRoutingKey:  "important",
		Created:          Timestamp{referenceTime},
		Updated:          Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryPacks.Update returned %+v, want %+v", got, want)
	}
}

func TestOsqueryPacksService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/packs/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.OsqueryPacks.Delete(ctx, 1)
	if err != nil {
		t.Errorf("OsqueryPacks.Delete returned error: %v", err)
	}
}
