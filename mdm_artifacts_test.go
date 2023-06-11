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

var maListJSONResponse = `
[
    {
        "id": "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
        "name": "Default",
	"type": "Profile",
	"channel": "Device",
	"platforms": ["macOS"],
	"install_during_setup_assistant": false,
	"auto_update": true,
	"reinstall_interval": 1,
	"reinstall_on_os_update": "No",
	"requires": ["b89d21e8-76de-4ae5-948d-5627474ab8be"],
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var maGetJSONResponse = `
{
    "id": "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
    "name": "Default",
    "type": "Profile",
    "channel": "Device",
    "platforms": ["macOS"],
    "install_during_setup_assistant": false,
    "auto_update": true,
    "reinstall_interval": 1,
    "reinstall_on_os_update": "No",
    "requires": ["b89d21e8-76de-4ae5-948d-5627474ab8be"],
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var maCreateJSONResponse = `
{
    "id": "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
    "name": "Default",
    "type": "Profile",
    "channel": "Device",
    "platforms": ["macOS"],
    "install_during_setup_assistant": true,
    "auto_update": true,
    "reinstall_interval": 1,
    "reinstall_on_os_update": "No",
    "requires": ["b89d21e8-76de-4ae5-948d-5627474ab8be"],
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var maUpdateJSONResponse = `
{
    "id": "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
    "name": "Default",
    "type": "Profile",
    "channel": "Device",
    "platforms": ["macOS"],
    "install_during_setup_assistant": true,
    "auto_update": true,
    "reinstall_interval": 1,
    "reinstall_on_os_update": "No",
    "requires": ["b89d21e8-76de-4ae5-948d-5627474ab8be"],
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestMDMArtifactsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/artifacts/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, maListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMArtifacts.List(ctx, nil)
	if err != nil {
		t.Errorf("MDMArtifacts.List returned error: %v", err)
	}

	want := []MDMArtifact{
		{
			ID:                          "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
			Name:                        "Default",
			Type:                        "Profile",
			Channel:                     "Device",
			Platforms:                   []string{"macOS"},
			InstallDuringSetupAssistant: false,
			AutoUpdate:                  true,
			ReinstallInterval:           1,
			ReinstallOnOSUpdate:         "No",
			Requires:                    []string{"b89d21e8-76de-4ae5-948d-5627474ab8be"},
			Created:                     Timestamp{referenceTime},
			Updated:                     Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMArtifacts.List returned %+v, want %+v", got, want)
	}
}

func TestMDMArtifactsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/artifacts/526efd25-c1f7-498c-82b5-94ff0b39ba8e/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, maGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMArtifacts.GetByID(ctx, "526efd25-c1f7-498c-82b5-94ff0b39ba8e")
	if err != nil {
		t.Errorf("MDMArtifacts.GetByID returned error: %v", err)
	}

	want := &MDMArtifact{
		ID:                          "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
		Name:                        "Default",
		Type:                        "Profile",
		Channel:                     "Device",
		Platforms:                   []string{"macOS"},
		InstallDuringSetupAssistant: false,
		AutoUpdate:                  true,
		ReinstallInterval:           1,
		ReinstallOnOSUpdate:         "No",
		Requires:                    []string{"b89d21e8-76de-4ae5-948d-5627474ab8be"},
		Created:                     Timestamp{referenceTime},
		Updated:                     Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMArtifacts.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMDMArtifactsService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/artifacts/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "Default")
		fmt.Fprint(w, maListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMArtifacts.GetByName(ctx, "Default")
	if err != nil {
		t.Errorf("MDMArtifacts.GetByName returned error: %v", err)
	}

	want := &MDMArtifact{
		ID:                          "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
		Name:                        "Default",
		Type:                        "Profile",
		Channel:                     "Device",
		Platforms:                   []string{"macOS"},
		InstallDuringSetupAssistant: false,
		AutoUpdate:                  true,
		ReinstallInterval:           1,
		ReinstallOnOSUpdate:         "No",
		Requires:                    []string{"b89d21e8-76de-4ae5-948d-5627474ab8be"},
		Created:                     Timestamp{referenceTime},
		Updated:                     Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMArtifacts.GetByName returned %+v, want %+v", got, want)
	}
}

func TestMDMArtifactsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &MDMArtifactRequest{
		Name:                        "Default",
		Type:                        "Profile",
		Channel:                     "Device",
		Platforms:                   []string{"macOS"},
		InstallDuringSetupAssistant: true,
		AutoUpdate:                  true,
		ReinstallInterval:           1,
		ReinstallOnOSUpdate:         "No",
		Requires:                    []string{"b89d21e8-76de-4ae5-948d-5627474ab8be"},
	}

	mux.HandleFunc("/mdm/artifacts/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMArtifactRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, maCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMArtifacts.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("MDMArtifacts.Create returned error: %v", err)
	}

	want := &MDMArtifact{
		ID:                          "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
		Name:                        "Default",
		Type:                        "Profile",
		Channel:                     "Device",
		Platforms:                   []string{"macOS"},
		InstallDuringSetupAssistant: true,
		AutoUpdate:                  true,
		ReinstallInterval:           1,
		ReinstallOnOSUpdate:         "No",
		Requires:                    []string{"b89d21e8-76de-4ae5-948d-5627474ab8be"},
		Created:                     Timestamp{referenceTime},
		Updated:                     Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMArtifacts.Create returned %+v, want %+v", got, want)
	}
}

func TestMDMArtifactsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &MDMArtifactRequest{
		Name:                        "Default",
		Type:                        "Profile",
		Channel:                     "Device",
		Platforms:                   []string{"macOS"},
		InstallDuringSetupAssistant: true,
		AutoUpdate:                  true,
		ReinstallInterval:           1,
		ReinstallOnOSUpdate:         "No",
		Requires:                    []string{"b89d21e8-76de-4ae5-948d-5627474ab8be"},
	}

	mux.HandleFunc("/mdm/artifacts/526efd25-c1f7-498c-82b5-94ff0b39ba8e/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMArtifactRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, maUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMArtifacts.Update(ctx, "526efd25-c1f7-498c-82b5-94ff0b39ba8e", updateRequest)
	if err != nil {
		t.Errorf("MDMArtifacts.Update returned error: %v", err)
	}

	want := &MDMArtifact{
		ID:                          "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
		Name:                        "Default",
		Type:                        "Profile",
		Channel:                     "Device",
		Platforms:                   []string{"macOS"},
		InstallDuringSetupAssistant: true,
		AutoUpdate:                  true,
		ReinstallInterval:           1,
		ReinstallOnOSUpdate:         "No",
		Requires:                    []string{"b89d21e8-76de-4ae5-948d-5627474ab8be"},
		Created:                     Timestamp{referenceTime},
		Updated:                     Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMArtifacts.Update returned %+v, want %+v", got, want)
	}
}

func TestMDMArtifactsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/artifacts/526efd25-c1f7-498c-82b5-94ff0b39ba8e/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MDMArtifacts.Delete(ctx, "526efd25-c1f7-498c-82b5-94ff0b39ba8e")
	if err != nil {
		t.Errorf("MDMArtifacts.Delete returned error: %v", err)
	}
}
