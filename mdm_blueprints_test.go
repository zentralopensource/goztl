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

var mbListJSONResponse = `
[
    {
        "id": 4,
        "name": "Default",
	"inventory_interval": 77777,
	"collect_apps": 0,
	"collect_certificates": 1,
	"collect_profiles": 2,
	"filevault_config": null,
	"recovery_password_config": null,
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var mbGetJSONResponse = `
{
    "id": 4,
    "name": "Default",
    "inventory_interval": 77777,
    "collect_apps": 0,
    "collect_certificates": 1,
    "collect_profiles": 2,
    "filevault_config": 3,
    "recovery_password_config": 4,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var mbCreateJSONResponse = `
{
    "id": 4,
    "name": "Default",
    "inventory_interval": 77777,
    "collect_apps": 0,
    "collect_certificates": 1,
    "collect_profiles": 2,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var mbUpdateJSONResponse = `
{
    "id": 4,
    "name": "Default",
    "inventory_interval": 77777,
    "collect_apps": 0,
    "collect_certificates": 1,
    "collect_profiles": 2,
    "filevault_config": 3,
    "recovery_password_config": 4,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestMDMBlueprintsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/blueprints/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mbListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMBlueprints.List(ctx, nil)
	if err != nil {
		t.Errorf("MDMBlueprints.List returned error: %v", err)
	}

	want := []MDMBlueprint{
		{
			ID:                  4,
			Name:                "Default",
			InventoryInterval:   77777,
			CollectApps:         0,
			CollectCertificates: 1,
			CollectProfiles:     2,
			Created:             Timestamp{referenceTime},
			Updated:             Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMBlueprints.List returned %+v, want %+v", got, want)
	}
}

func TestMDMBlueprintsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/blueprints/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mbGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMBlueprints.GetByID(ctx, 4)
	if err != nil {
		t.Errorf("MDMBlueprints.GetByID returned error: %v", err)
	}

	want := &MDMBlueprint{
		ID:                       4,
		Name:                     "Default",
		InventoryInterval:        77777,
		CollectApps:              0,
		CollectCertificates:      1,
		CollectProfiles:          2,
		FileVaultConfigID:        Int(3),
		RecoveryPasswordConfigID: Int(4),
		Created:                  Timestamp{referenceTime},
		Updated:                  Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMBlueprints.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMDMBlueprintsService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/blueprints/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "Default")
		fmt.Fprint(w, mbListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMBlueprints.GetByName(ctx, "Default")
	if err != nil {
		t.Errorf("MDMBlueprints.GetByName returned error: %v", err)
	}

	want := &MDMBlueprint{
		ID:                  4,
		Name:                "Default",
		InventoryInterval:   77777,
		CollectApps:         0,
		CollectCertificates: 1,
		CollectProfiles:     2,
		Created:             Timestamp{referenceTime},
		Updated:             Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMBlueprints.GetByName returned %+v, want %+v", got, want)
	}
}

func TestMDMBlueprintsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &MDMBlueprintRequest{
		Name:                "Default",
		InventoryInterval:   77777,
		CollectApps:         0,
		CollectCertificates: 1,
		CollectProfiles:     2,
	}

	mux.HandleFunc("/mdm/blueprints/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMBlueprintRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, mbCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMBlueprints.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("MDMBlueprints.Create returned error: %v", err)
	}

	want := &MDMBlueprint{
		ID:                  4,
		Name:                "Default",
		InventoryInterval:   77777,
		CollectApps:         0,
		CollectCertificates: 1,
		CollectProfiles:     2,
		Created:             Timestamp{referenceTime},
		Updated:             Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMBlueprints.Create returned %+v, want %+v", got, want)
	}
}

func TestMDMBlueprintsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &MDMBlueprintRequest{
		Name:                     "Default",
		InventoryInterval:        77777,
		CollectApps:              0,
		CollectCertificates:      1,
		CollectProfiles:          2,
		FileVaultConfigID:        Int(3),
		RecoveryPasswordConfigID: Int(4),
	}

	mux.HandleFunc("/mdm/blueprints/4/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMBlueprintRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, mbUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMBlueprints.Update(ctx, 4, updateRequest)
	if err != nil {
		t.Errorf("MDMBlueprints.Update returned error: %v", err)
	}

	want := &MDMBlueprint{
		ID:                       4,
		Name:                     "Default",
		InventoryInterval:        77777,
		CollectApps:              0,
		CollectCertificates:      1,
		CollectProfiles:          2,
		FileVaultConfigID:        Int(3),
		RecoveryPasswordConfigID: Int(4),
		Created:                  Timestamp{referenceTime},
		Updated:                  Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMBlueprints.Update returned %+v, want %+v", got, want)
	}
}

func TestMDMBlueprintsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/blueprints/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MDMBlueprints.Delete(ctx, 4)
	if err != nil {
		t.Errorf("MDMBlueprints.Delete returned error: %v", err)
	}
}
