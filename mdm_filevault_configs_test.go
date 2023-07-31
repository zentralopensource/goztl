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

var mfcListJSONResponse = `
[
    {
        "id": 4,
        "name": "Default",
	"escrow_location_display_name": "Escrow",
	"at_login_only": true,
	"bypass_attempts": 1,
	"show_recovery_key": true,
	"destroy_key_on_standby": true,
	"prk_rotation_interval_days": 90,
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var mfcGetJSONResponse = `
{
    "id": 4,
    "name": "Default",
    "escrow_location_display_name": "Escrow",
    "at_login_only": true,
    "bypass_attempts": 1,
    "show_recovery_key": true,
    "destroy_key_on_standby": true,
    "prk_rotation_interval_days": 90,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var mfcCreateJSONResponse = `
{
    "id": 4,
    "name": "Default",
    "escrow_location_display_name": "Escrow",
    "at_login_only": true,
    "bypass_attempts": 1,
    "show_recovery_key": true,
    "destroy_key_on_standby": true,
    "prk_rotation_interval_days": 90,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var mfcUpdateJSONResponse = `
{
    "id": 4,
    "name": "Default",
    "escrow_location_display_name": "Escrow",
    "at_login_only": true,
    "bypass_attempts": 1,
    "show_recovery_key": true,
    "destroy_key_on_standby": true,
    "prk_rotation_interval_days": 90,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestMDMFileVaultConfigsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/filevault_configs/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mfcListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMFileVaultConfigs.List(ctx, nil)
	if err != nil {
		t.Errorf("MDMFileVaultConfigs.List returned error: %v", err)
	}

	want := []MDMFileVaultConfig{
		{
			ID:                        4,
			Name:                      "Default",
			EscrowLocationDisplayName: "Escrow",
			AtLoginOnly:               true,
			BypassAttempts:            1,
			ShowRecoveryKey:           true,
			DestroyKeyOnStandby:       true,
			PRKRotationIntervalDays:   90,
			Created:                   Timestamp{referenceTime},
			Updated:                   Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMFileVaultConfigs.List returned %+v, want %+v", got, want)
	}
}

func TestMDMFileVaultConfigsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/filevault_configs/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mfcGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMFileVaultConfigs.GetByID(ctx, 4)
	if err != nil {
		t.Errorf("MDMFileVaultConfigs.GetByID returned error: %v", err)
	}

	want := &MDMFileVaultConfig{
		ID:                        4,
		Name:                      "Default",
		EscrowLocationDisplayName: "Escrow",
		AtLoginOnly:               true,
		BypassAttempts:            1,
		ShowRecoveryKey:           true,
		DestroyKeyOnStandby:       true,
		PRKRotationIntervalDays:   90,
		Created:                   Timestamp{referenceTime},
		Updated:                   Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMFileVaultConfigs.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMDMFileVaultConfigsService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/filevault_configs/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "Default")
		fmt.Fprint(w, mfcListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMFileVaultConfigs.GetByName(ctx, "Default")
	if err != nil {
		t.Errorf("MDMFileVaultConfigs.GetByName returned error: %v", err)
	}

	want := &MDMFileVaultConfig{
		ID:                        4,
		Name:                      "Default",
		EscrowLocationDisplayName: "Escrow",
		AtLoginOnly:               true,
		BypassAttempts:            1,
		ShowRecoveryKey:           true,
		DestroyKeyOnStandby:       true,
		PRKRotationIntervalDays:   90,
		Created:                   Timestamp{referenceTime},
		Updated:                   Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMFileVaultConfigs.GetByName returned %+v, want %+v", got, want)
	}
}

func TestMDMFileVaultConfigsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &MDMFileVaultConfigRequest{
		Name:                      "Default",
		EscrowLocationDisplayName: "Escrow",
		AtLoginOnly:               true,
		BypassAttempts:            1,
		ShowRecoveryKey:           true,
		DestroyKeyOnStandby:       true,
		PRKRotationIntervalDays:   90,
	}

	mux.HandleFunc("/mdm/filevault_configs/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMFileVaultConfigRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, mfcCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMFileVaultConfigs.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("MDMFileVaultConfigs.Create returned error: %v", err)
	}

	want := &MDMFileVaultConfig{
		ID:                        4,
		Name:                      "Default",
		EscrowLocationDisplayName: "Escrow",
		AtLoginOnly:               true,
		BypassAttempts:            1,
		ShowRecoveryKey:           true,
		DestroyKeyOnStandby:       true,
		PRKRotationIntervalDays:   90,
		Created:                   Timestamp{referenceTime},
		Updated:                   Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMFileVaultConfigs.Create returned %+v, want %+v", got, want)
	}
}

func TestMDMFileVaultConfigsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &MDMFileVaultConfigRequest{
		Name:                      "Default",
		EscrowLocationDisplayName: "Escrow",
		AtLoginOnly:               true,
		BypassAttempts:            1,
		ShowRecoveryKey:           true,
		DestroyKeyOnStandby:       true,
		PRKRotationIntervalDays:   90,
	}

	mux.HandleFunc("/mdm/filevault_configs/4/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMFileVaultConfigRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, mfcUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMFileVaultConfigs.Update(ctx, 4, updateRequest)
	if err != nil {
		t.Errorf("MDMFileVaultConfigs.Update returned error: %v", err)
	}

	want := &MDMFileVaultConfig{
		ID:                        4,
		Name:                      "Default",
		EscrowLocationDisplayName: "Escrow",
		AtLoginOnly:               true,
		BypassAttempts:            1,
		ShowRecoveryKey:           true,
		DestroyKeyOnStandby:       true,
		PRKRotationIntervalDays:   90,
		Created:                   Timestamp{referenceTime},
		Updated:                   Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMFileVaultConfigs.Update returned %+v, want %+v", got, want)
	}
}

func TestMDMFileVaultConfigsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/filevault_configs/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MDMFileVaultConfigs.Delete(ctx, 4)
	if err != nil {
		t.Errorf("MDMFileVaultConfigs.Delete returned error: %v", err)
	}
}
