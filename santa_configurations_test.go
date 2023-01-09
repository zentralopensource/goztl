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

var scListJSONResponse = `
[
    {
        "id": 4,
        "name": "Santa - Monitoring",
        "client_mode": 1,
        "client_certificate_auth": false,
        "batch_size": 50,
        "full_sync_interval": 600,
        "enable_bundles": false,
        "enable_transitive_rules": false,
        "allowed_path_regex": "",
        "blocked_path_regex": "",
        "block_usb_mount": false,
        "remount_usb_mode": [],
        "allow_unknown_shard": 100,
        "enable_all_event_upload_shard": 0,
        "sync_incident_severity": 0,
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var scGetJSONResponse = `
{
    "id": 4,
    "name": "Santa - Monitoring",
    "client_mode": 1,
    "client_certificate_auth": false,
    "batch_size": 50,
    "full_sync_interval": 600,
    "enable_bundles": false,
    "enable_transitive_rules": false,
    "allowed_path_regex": "",
    "blocked_path_regex": "",
    "block_usb_mount": false,
    "remount_usb_mode": [],
    "allow_unknown_shard": 100,
    "enable_all_event_upload_shard": 0,
    "sync_incident_severity": 0,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var scCreateJSONResponse = `
{
    "id": 4,
    "name": "Santa - Monitoring",
    "client_mode": 2,
    "client_certificate_auth": true,
    "batch_size": 49,
    "full_sync_interval": 601,
    "enable_bundles": true,
    "enable_transitive_rules": true,
    "allowed_path_regex": "un",
    "blocked_path_regex": "deux",
    "block_usb_mount": true,
    "remount_usb_mode": [],
    "allow_unknown_shard": 100,
    "enable_all_event_upload_shard": 1,
    "sync_incident_severity": 2,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var scUpdateJSONResponse = `
{
    "id": 4,
    "name": "Santa - Monitoring",
    "client_mode": 2,
    "client_certificate_auth": true,
    "batch_size": 49,
    "full_sync_interval": 601,
    "enable_bundles": true,
    "enable_transitive_rules": true,
    "allowed_path_regex": "un",
    "blocked_path_regex": "deux",
    "block_usb_mount": true,
    "remount_usb_mode": [],
    "allow_unknown_shard": 100,
    "enable_all_event_upload_shard": 1,
    "sync_incident_severity": 2,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestSantaConfigurationsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/santa/configurations/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, scListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.SantaConfigurations.List(ctx, nil)
	if err != nil {
		t.Errorf("SantaConfigurations.List returned error: %v", err)
	}

	want := []SantaConfiguration{
		{
			ID:                        4,
			Name:                      "Santa - Monitoring",
			ClientMode:                1,
			ClientCertificateAuth:     false,
			BatchSize:                 50,
			FullSyncInterval:          600,
			EnableBundles:             false,
			EnableTransitiveRules:     false,
			AllowedPathRegex:          "",
			BlockedPathRegex:          "",
			BlockUSBMount:             false,
			RemountUSBMode:            make([]string, 0),
			AllowUnknownShard:         100,
			EnableAllEventUploadShard: 0,
			SyncIncidentSeverity:      0,
			Created:                   Timestamp{referenceTime},
			Updated:                   Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("SantaConfigurations.List returned %+v, want %+v", got, want)
	}
}

func TestSantaConfigurationsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/santa/configurations/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, scGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.SantaConfigurations.GetByID(ctx, 1)
	if err != nil {
		t.Errorf("SantaConfigurations.GetByID returned error: %v", err)
	}

	want := &SantaConfiguration{
		ID:                        4,
		Name:                      "Santa - Monitoring",
		ClientMode:                1,
		ClientCertificateAuth:     false,
		BatchSize:                 50,
		FullSyncInterval:          600,
		EnableBundles:             false,
		EnableTransitiveRules:     false,
		AllowedPathRegex:          "",
		BlockedPathRegex:          "",
		BlockUSBMount:             false,
		RemountUSBMode:            make([]string, 0),
		AllowUnknownShard:         100,
		EnableAllEventUploadShard: 0,
		SyncIncidentSeverity:      0,
		Created:                   Timestamp{referenceTime},
		Updated:                   Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("SantaConfigurations.GetByID returned %+v, want %+v", got, want)
	}
}

func TestSantaConfigurationsService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/santa/configurations/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "Santa - Monitoring")
		fmt.Fprint(w, scListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.SantaConfigurations.GetByName(ctx, "Santa - Monitoring")
	if err != nil {
		t.Errorf("SantaConfigurations.GetByName returned error: %v", err)
	}

	want := &SantaConfiguration{
		ID:                        4,
		Name:                      "Santa - Monitoring",
		ClientMode:                1,
		ClientCertificateAuth:     false,
		BatchSize:                 50,
		FullSyncInterval:          600,
		EnableBundles:             false,
		EnableTransitiveRules:     false,
		AllowedPathRegex:          "",
		BlockedPathRegex:          "",
		BlockUSBMount:             false,
		RemountUSBMode:            make([]string, 0),
		AllowUnknownShard:         100,
		EnableAllEventUploadShard: 0,
		SyncIncidentSeverity:      0,
		Created:                   Timestamp{referenceTime},
		Updated:                   Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("SantaConfigurations.GetByName returned %+v, want %+v", got, want)
	}
}

func TestSantaConfigurationsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &SantaConfigurationRequest{
		Name:                      "Santa - Monitoring",
		ClientMode:                2,
		ClientCertificateAuth:     true,
		BatchSize:                 49,
		FullSyncInterval:          601,
		EnableBundles:             true,
		EnableTransitiveRules:     true,
		AllowedPathRegex:          "un",
		BlockedPathRegex:          "deux",
		BlockUSBMount:             true,
		RemountUSBMode:            make([]string, 0),
		AllowUnknownShard:         100,
		EnableAllEventUploadShard: 1,
		SyncIncidentSeverity:      2,
	}

	mux.HandleFunc("/santa/configurations/", func(w http.ResponseWriter, r *http.Request) {
		v := new(SantaConfigurationRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, scCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.SantaConfigurations.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("SantaConfigurations.Create returned error: %v", err)
	}

	want := &SantaConfiguration{
		ID:                        4,
		Name:                      "Santa - Monitoring",
		ClientMode:                2,
		ClientCertificateAuth:     true,
		BatchSize:                 49,
		FullSyncInterval:          601,
		EnableBundles:             true,
		EnableTransitiveRules:     true,
		AllowedPathRegex:          "un",
		BlockedPathRegex:          "deux",
		BlockUSBMount:             true,
		RemountUSBMode:            make([]string, 0),
		AllowUnknownShard:         100,
		EnableAllEventUploadShard: 1,
		SyncIncidentSeverity:      2,
		Created:                   Timestamp{referenceTime},
		Updated:                   Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("SantaConfigurations.Create returned %+v, want %+v", got, want)
	}
}

func TestSantaConfigurationsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &SantaConfigurationRequest{
		Name:                      "Santa - Monitoring",
		ClientMode:                2,
		ClientCertificateAuth:     true,
		BatchSize:                 49,
		FullSyncInterval:          601,
		EnableBundles:             true,
		EnableTransitiveRules:     true,
		AllowedPathRegex:          "un",
		BlockedPathRegex:          "deux",
		BlockUSBMount:             true,
		RemountUSBMode:            make([]string, 0),
		AllowUnknownShard:         100,
		EnableAllEventUploadShard: 1,
		SyncIncidentSeverity:      2,
	}

	mux.HandleFunc("/santa/configurations/1/", func(w http.ResponseWriter, r *http.Request) {
		v := new(SantaConfigurationRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, scUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.SantaConfigurations.Update(ctx, 1, updateRequest)
	if err != nil {
		t.Errorf("SantaConfigurations.Update returned error: %v", err)
	}

	want := &SantaConfiguration{
		ID:                        4,
		Name:                      "Santa - Monitoring",
		ClientMode:                2,
		ClientCertificateAuth:     true,
		BatchSize:                 49,
		FullSyncInterval:          601,
		EnableBundles:             true,
		EnableTransitiveRules:     true,
		AllowedPathRegex:          "un",
		BlockedPathRegex:          "deux",
		BlockUSBMount:             true,
		RemountUSBMode:            make([]string, 0),
		AllowUnknownShard:         100,
		EnableAllEventUploadShard: 1,
		SyncIncidentSeverity:      2,
		Created:                   Timestamp{referenceTime},
		Updated:                   Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("SantaConfigurations.Update returned %+v, want %+v", got, want)
	}
}

func TestSantaConfigurationsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/santa/configurations/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.SantaConfigurations.Delete(ctx, 1)
	if err != nil {
		t.Errorf("SantaConfigurations.Delete returned error: %v", err)
	}
}
