package goztl

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var depVirtualServerGetJsonResponse = `{
	"id": 39480, 
	"name": "CUvRnNufiprz", 
	"uuid": "8265be4b-f097-4e03-8093-172e01056866", 
	"created_at": "2022-07-22T01:02:03.444444", 
	"updated_at": "2022-07-22T01:02:03.444444"
}`

var depVirtualServerListJsonResponse = `{
	"count": 1,
	"results": [
		{
			"id": 39480, 
			"name": "CUvRnNufiprz", 
			"uuid": "8265be4b-f097-4e03-8093-172e01056866", 
			"created_at": "2022-07-22T01:02:03.444444", 
			"updated_at": "2022-07-22T01:02:03.444444"
		}
	]
}`

var depVirtualServerListFirstPageJsonResponse = `{
	"count": 2,
	"next": "http://example.com/mdm/dep/virtual_servers/?page=2",
	"results": [
		{
			"id": 39480, 
			"name": "CUvRnNufiprz", 
			"uuid": "8265be4b-f097-4e03-8093-172e01056866", 
			"created_at": "2022-07-22T01:02:03.444444", 
			"updated_at": "2022-07-22T01:02:03.444444"
		}
	]
}`

var depVirtualServerListNextPageJsonResponse = `{
	"count": 2,
	"results": [
		{
			"id": 39481, 
			"name": "DIbTmMigOatu", 
			"uuid": "7fdb067b-cda4-4d3a-aa0f-881667e61a13", 
			"created_at": "2022-07-22T01:02:03.444444", 
			"updated_at": "2022-07-22T01:02:03.444444"
		}
	]
}`

func TestMDMDEPVirtualServersService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/dep/virtual_servers/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")

		if r.URL.Query().Get("page") == "" {
			fmt.Fprint(w, depVirtualServerListFirstPageJsonResponse)
			return
		}

		testQueryArg(t, r, "page", "2")
		fmt.Fprint(w, depVirtualServerListNextPageJsonResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMDEPVirtualServers.List(ctx, nil)
	if err != nil {
		t.Errorf("MDMDEPVirtualServers.List returned error: %v", err)
	}

	want := []MDMDEPVirtualServer{
		{
			ID:      39480,
			Name:    "CUvRnNufiprz",
			UUID:    "8265be4b-f097-4e03-8093-172e01056866",
			Created: Timestamp{referenceTime},
			Updated: Timestamp{referenceTime},
		},
		{
			ID:      39481,
			Name:    "DIbTmMigOatu",
			UUID:    "7fdb067b-cda4-4d3a-aa0f-881667e61a13",
			Created: Timestamp{referenceTime},
			Updated: Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMDEPVirtualServers.List returned %+v, want %+v", got, want)
	}
}

func TestMDMDEPVirtualServersService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/dep/virtual_servers/39480/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, depVirtualServerGetJsonResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMDEPVirtualServers.GetByID(ctx, 39480)
	if err != nil {
		t.Errorf("MDMDEPVirtualServers.GetByID returned error: %v", err)
	}

	want := &MDMDEPVirtualServer{
		ID:      39480,
		Name:    "CUvRnNufiprz",
		UUID:    "8265be4b-f097-4e03-8093-172e01056866",
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMDEPVirtualServers.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMDMDEPVirtualServersService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/dep/virtual_servers/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "CUvRnNufiprz")
		fmt.Fprint(w, depVirtualServerListJsonResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMDEPVirtualServers.GetByName(ctx, "CUvRnNufiprz")
	if err != nil {
		t.Errorf("MDMDEPVirtualServers.GetByName returned error: %v", err)
	}
	if !cmp.Equal(len(got), 1) {
		t.Errorf("MDMDEPVirtualServers.GetByName returned not unique result.")
	}
	first := &got[0]

	want := &MDMDEPVirtualServer{
		ID:      39480,
		Name:    "CUvRnNufiprz",
		UUID:    "8265be4b-f097-4e03-8093-172e01056866",
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}

	if !cmp.Equal(first, want) {
		t.Errorf("MDMDEPVirtualServers.GetByName returned %+v, want %+v", first, want)
	}

}
