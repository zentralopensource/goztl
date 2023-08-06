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

var mrpcListJSONResponse = `
[
    {
        "id": 4,
        "name": "Default",
	"dynamic_password": true,
	"static_password": null,
	"rotation_interval_days": 90,
	"rotate_firmware_password": true,
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var mrpcGetJSONResponse = `
{
    "id": 4,
    "name": "Default",
    "dynamic_password": false,
    "static_password": "12345678",
    "rotation_interval_days": 0,
    "rotate_firmware_password": false,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var mrpcCreateJSONResponse = `
{
    "id": 4,
    "name": "Default",
    "dynamic_password": false,
    "static_password": "12345678",
    "rotation_interval_days": 0,
    "rotate_firmware_password": false,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var mrpcUpdateJSONResponse = `
{
    "id": 4,
    "name": "Default",
    "dynamic_password": true,
    "rotation_interval_days": 90,
    "rotate_firmware_password": true,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestMDMRecoveryPasswordConfigsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/recovery_password_configs/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mrpcListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMRecoveryPasswordConfigs.List(ctx, nil)
	if err != nil {
		t.Errorf("MDMRecoveryPasswordConfigs.List returned error: %v", err)
	}

	want := []MDMRecoveryPasswordConfig{
		{
			ID:                     4,
			Name:                   "Default",
			DynamicPassword:        true,
			RotationIntervalDays:   90,
			RotateFirmwarePassword: true,
			Created:                Timestamp{referenceTime},
			Updated:                Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMRecoveryPasswordConfigs.List returned %+v, want %+v", got, want)
	}
}

func TestMDMRecoveryPasswordConfigsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/recovery_password_configs/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mrpcGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMRecoveryPasswordConfigs.GetByID(ctx, 4)
	if err != nil {
		t.Errorf("MDMRecoveryPasswordConfigs.GetByID returned error: %v", err)
	}

	want := &MDMRecoveryPasswordConfig{
		ID:                     4,
		Name:                   "Default",
		DynamicPassword:        false,
		StaticPassword:         String("12345678"),
		RotationIntervalDays:   0,
		RotateFirmwarePassword: false,
		Created:                Timestamp{referenceTime},
		Updated:                Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMRecoveryPasswordConfigs.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMDMRecoveryPasswordConfigsService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/recovery_password_configs/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "Default")
		fmt.Fprint(w, mrpcListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMRecoveryPasswordConfigs.GetByName(ctx, "Default")
	if err != nil {
		t.Errorf("MDMRecoveryPasswordConfigs.GetByName returned error: %v", err)
	}

	want := &MDMRecoveryPasswordConfig{
		ID:                     4,
		Name:                   "Default",
		DynamicPassword:        true,
		RotationIntervalDays:   90,
		RotateFirmwarePassword: true,
		Created:                Timestamp{referenceTime},
		Updated:                Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMRecoveryPasswordConfigs.GetByName returned %+v, want %+v", got, want)
	}
}

func TestMDMRecoveryPasswordConfigsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &MDMRecoveryPasswordConfigRequest{
		Name:                   "Default",
		DynamicPassword:        false,
		StaticPassword:         String("12345678"),
		RotationIntervalDays:   0,
		RotateFirmwarePassword: false,
	}

	mux.HandleFunc("/mdm/recovery_password_configs/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMRecoveryPasswordConfigRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, mrpcCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMRecoveryPasswordConfigs.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("MDMRecoveryPasswordConfigs.Create returned error: %v", err)
	}

	want := &MDMRecoveryPasswordConfig{
		ID:                     4,
		Name:                   "Default",
		DynamicPassword:        false,
		StaticPassword:         String("12345678"),
		RotationIntervalDays:   0,
		RotateFirmwarePassword: false,
		Created:                Timestamp{referenceTime},
		Updated:                Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMRecoveryPasswordConfigs.Create returned %+v, want %+v", got, want)
	}
}

func TestMDMRecoveryPasswordConfigsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &MDMRecoveryPasswordConfigRequest{
		Name:                   "Default",
		DynamicPassword:        true,
		RotationIntervalDays:   90,
		RotateFirmwarePassword: true,
	}

	mux.HandleFunc("/mdm/recovery_password_configs/4/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMRecoveryPasswordConfigRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, mrpcUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMRecoveryPasswordConfigs.Update(ctx, 4, updateRequest)
	if err != nil {
		t.Errorf("MDMRecoveryPasswordConfigs.Update returned error: %v", err)
	}

	want := &MDMRecoveryPasswordConfig{
		ID:                     4,
		Name:                   "Default",
		DynamicPassword:        true,
		RotationIntervalDays:   90,
		RotateFirmwarePassword: true,
		Created:                Timestamp{referenceTime},
		Updated:                Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMRecoveryPasswordConfigs.Update returned %+v, want %+v", got, want)
	}
}

func TestMDMRecoveryPasswordConfigsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/recovery_password_configs/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MDMRecoveryPasswordConfigs.Delete(ctx, 4)
	if err != nil {
		t.Errorf("MDMRecoveryPasswordConfigs.Delete returned error: %v", err)
	}
}
