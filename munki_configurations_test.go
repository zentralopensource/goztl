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

var mucListJSONResponse = `
[
    {
        "id": 1,
        "name": "Default",
	"description": "",
        "inventory_apps_full_info_shard": 100,
        "principal_user_detection_domains": [
            "example.com"
        ],
        "principal_user_detection_sources": [
            "logged_in_user"
        ],
        "collected_condition_keys": [
            "arch",
            "machine_type"
        ],
        "managed_installs_sync_interval_days": 7,
        "auto_reinstall_incidents": true,
        "auto_failed_install_incidents": false,
        "version": 5,
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var mucGetJSONResponse = `
{
  "id": 6,
  "name": "Default",
  "description": "Description",
  "inventory_apps_full_info_shard": 50,
  "principal_user_detection_sources": [
    "google_chrome",
    "company_portal"
  ],
  "principal_user_detection_domains": [
    "zentral.io"
  ],
  "collected_condition_keys": [
    "arch",
    "machine_type"
  ],
  "managed_installs_sync_interval_days": 1,
  "auto_reinstall_incidents": true,
  "auto_failed_install_incidents": true,
  "version": 5,
  "created_at": "2022-07-22T01:02:03.444444",
  "updated_at": "2022-07-22T01:02:03.444444"
}
`

var mucCreateJSONResponse = `
{
  "id": 6,
  "name": "Default",
  "description": "Description",
  "inventory_apps_full_info_shard": 50,
  "principal_user_detection_sources": [
    "google_chrome",
    "company_portal"
  ],
  "principal_user_detection_domains": [
    "zentral.io"
  ],
  "collected_condition_keys": [
    "arch",
    "machine_type"
  ],
  "managed_installs_sync_interval_days": 1,
  "auto_reinstall_incidents": true,
  "auto_failed_install_incidents": true,
  "version": 1,
  "created_at": "2022-07-22T01:02:03.444444",
  "updated_at": "2022-07-22T01:02:03.444444"
}
`

var mucUpdateJSONResponse = `
{
  "id": 6,
  "name": "Default",
  "description": "Description",
  "inventory_apps_full_info_shard": 50,
  "principal_user_detection_sources": [
    "google_chrome",
    "company_portal"
  ],
  "principal_user_detection_domains": [
    "zentral.io"
  ],
  "collected_condition_keys": [
    "arch",
    "machine_type"
  ],
  "managed_installs_sync_interval_days": 1,
  "auto_reinstall_incidents": true,
  "auto_failed_install_incidents": true,
  "version": 5,
  "created_at": "2022-07-22T01:02:03.444444",
  "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestMunkiConfigurationsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/munki/configurations/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mucListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MunkiConfigurations.List(ctx, nil)
	if err != nil {
		t.Errorf("MunkiConfigurations.List returned error: %v", err)
	}

	want := []MunkiConfiguration{
		{
			ID:                              1,
			Name:                            "Default",
			Description:                     "",
			InventoryAppsFullInfoShard:      100,
			PrincipalUserDetectionSources:   []string{"logged_in_user"},
			PrincipalUserDetectionDomains:   []string{"example.com"},
			CollectedConditionKeys:          []string{"arch", "machine_type"},
			ManagedInstallsSyncIntervalDays: 7,
			AutoReinstallIncidents:          true,
			AutoFailedInstallIncidents:      false,
			Version:                         5,
			Created:                         Timestamp{referenceTime},
			Updated:                         Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MunkiConfigurations.List returned %+v, want %+v", got, want)
	}
}

func TestMunkiConfigurationsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/munki/configurations/6/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mucGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MunkiConfigurations.GetByID(ctx, 6)
	if err != nil {
		t.Errorf("MunkiConfigurations.GetByID returned error: %v", err)
	}

	want := &MunkiConfiguration{
		ID:                              6,
		Name:                            "Default",
		Description:                     "Description",
		InventoryAppsFullInfoShard:      50,
		PrincipalUserDetectionSources:   []string{"google_chrome", "company_portal"},
		PrincipalUserDetectionDomains:   []string{"zentral.io"},
		CollectedConditionKeys:          []string{"arch", "machine_type"},
		ManagedInstallsSyncIntervalDays: 1,
		AutoReinstallIncidents:          true,
		AutoFailedInstallIncidents:      true,
		Version:                         5,
		Created:                         Timestamp{referenceTime},
		Updated:                         Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MunkiConfigurations.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMunkiConfigurationsService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/munki/configurations/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "Default")
		fmt.Fprint(w, mucListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MunkiConfigurations.GetByName(ctx, "Default")
	if err != nil {
		t.Errorf("MunkiConfigurations.GetByName returned error: %v", err)
	}

	want := &MunkiConfiguration{
		ID:                              1,
		Name:                            "Default",
		Description:                     "",
		InventoryAppsFullInfoShard:      100,
		PrincipalUserDetectionSources:   []string{"logged_in_user"},
		PrincipalUserDetectionDomains:   []string{"example.com"},
		CollectedConditionKeys:          []string{"arch", "machine_type"},
		ManagedInstallsSyncIntervalDays: 7,
		AutoReinstallIncidents:          true,
		AutoFailedInstallIncidents:      false,
		Version:                         5,
		Created:                         Timestamp{referenceTime},
		Updated:                         Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MunkiConfigurations.GetByName returned %+v, want %+v", got, want)
	}
}

func TestMunkiConfigurationsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &MunkiConfigurationRequest{
		Name:                            "Default",
		Description:                     "Description",
		InventoryAppsFullInfoShard:      50,
		PrincipalUserDetectionSources:   []string{"google_chrome", "company_portal"},
		PrincipalUserDetectionDomains:   []string{"zentral.io"},
		CollectedConditionKeys:          []string{"arch", "machine_type"},
		ManagedInstallsSyncIntervalDays: 1,
		AutoReinstallIncidents:          true,
		AutoFailedInstallIncidents:      true,
	}

	mux.HandleFunc("/munki/configurations/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MunkiConfigurationRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, mucCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MunkiConfigurations.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("MunkiConfigurations.Create returned error: %v", err)
	}

	want := &MunkiConfiguration{
		ID:                              6,
		Name:                            "Default",
		Description:                     "Description",
		InventoryAppsFullInfoShard:      50,
		PrincipalUserDetectionSources:   []string{"google_chrome", "company_portal"},
		PrincipalUserDetectionDomains:   []string{"zentral.io"},
		CollectedConditionKeys:          []string{"arch", "machine_type"},
		ManagedInstallsSyncIntervalDays: 1,
		AutoReinstallIncidents:          true,
		AutoFailedInstallIncidents:      true,
		Version:                         1,
		Created:                         Timestamp{referenceTime},
		Updated:                         Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MunkiConfigurations.Create returned %+v, want %+v", got, want)
	}
}

func TestMunkiConfigurationsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &MunkiConfigurationRequest{
		Name:                            "Default",
		Description:                     "Description",
		InventoryAppsFullInfoShard:      50,
		PrincipalUserDetectionSources:   []string{"google_chrome", "company_portal"},
		PrincipalUserDetectionDomains:   []string{"zentral.io"},
		CollectedConditionKeys:          []string{"arch", "machine_type"},
		ManagedInstallsSyncIntervalDays: 1,
		AutoReinstallIncidents:          true,
		AutoFailedInstallIncidents:      true,
	}

	mux.HandleFunc("/munki/configurations/6/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MunkiConfigurationRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, mucUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MunkiConfigurations.Update(ctx, 6, updateRequest)
	if err != nil {
		t.Errorf("MunkiConfigurations.Update returned error: %v", err)
	}

	want := &MunkiConfiguration{
		ID:                              6,
		Name:                            "Default",
		Description:                     "Description",
		InventoryAppsFullInfoShard:      50,
		PrincipalUserDetectionSources:   []string{"google_chrome", "company_portal"},
		PrincipalUserDetectionDomains:   []string{"zentral.io"},
		CollectedConditionKeys:          []string{"arch", "machine_type"},
		ManagedInstallsSyncIntervalDays: 1,
		AutoReinstallIncidents:          true,
		AutoFailedInstallIncidents:      true,
		Version:                         5,
		Created:                         Timestamp{referenceTime},
		Updated:                         Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MunkiConfigurations.Update returned %+v, want %+v", got, want)
	}
}

func TestMunkiConfigurationsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/munki/configurations/6/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MunkiConfigurations.Delete(ctx, 6)
	if err != nil {
		t.Errorf("MunkiConfigurations.Delete returned error: %v", err)
	}
}
