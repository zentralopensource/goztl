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

var msueListJSONResponse = `
[
    {
        "id": 4,
        "name": "Default",
	"details_url": "https://www.example.com",
	"platforms": ["macOS"],
	"tags": [1, 2],
	"os_version": "14.1",
	"build_version": "23B74",
	"local_datetime": "2023-11-05T09:30:00",
	"max_os_version": "",
	"delay_days": null,
	"local_time": null,
	"rotate_firmware_password": true,
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var msueGetJSONResponse = `
{
    "id": 4,
    "name": "Default",
    "details_url": "https://www.example.com",
    "platforms": ["macOS"],
    "tags": [],
    "os_version": "",
    "build_version": "",
    "local_datetime": null,
    "max_os_version": "15",
    "delay_days": 7,
    "local_time": "09:30:00",
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var msueCreateJSONResponse = `
{
    "id": 4,
    "name": "Default",
    "details_url": "https://www.example.com",
    "platforms": ["macOS"],
    "tags": [1, 2],
    "os_version": "",
    "build_version": "",
    "local_datetime": null,
    "max_os_version": "15",
    "delay_days": 7,
    "local_time": "09:30:00",
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var msueUpdateJSONResponse = `
{
    "id": 4,
    "name": "Default",
    "details_url": "https://www.example.com",
    "platforms": ["macOS"],
    "tags": [1, 2],
    "os_version": "",
    "build_version": "",
    "local_datetime": null,
    "max_os_version": "15",
    "delay_days": 7,
    "local_time": "09:30:00",
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestMDMSoftwareUpdateEnforcementsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/software_update_enforcements/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, msueListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMSoftwareUpdateEnforcements.List(ctx, nil)
	if err != nil {
		t.Errorf("MDMSoftwareUpdateEnforcements.List returned error: %v", err)
	}

	want := []MDMSoftwareUpdateEnforcement{
		{
			ID:            4,
			Name:          "Default",
			DetailsURL:    "https://www.example.com",
			Platforms:     []string{"macOS"},
			TagIDs:        []int{1, 2},
			OSVersion:     "14.1",
			BuildVersion:  "23B74",
			LocalDateTime: String("2023-11-05T09:30:00"),
			Created:       Timestamp{referenceTime},
			Updated:       Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMSoftwareUpdateEnforcements.List returned %+v, want %+v", got, want)
	}
}

func TestMDMSoftwareUpdateEnforcementsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/software_update_enforcements/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, msueGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMSoftwareUpdateEnforcements.GetByID(ctx, 4)
	if err != nil {
		t.Errorf("MDMSoftwareUpdateEnforcements.GetByID returned error: %v", err)
	}

	want := &MDMSoftwareUpdateEnforcement{
		ID:           4,
		Name:         "Default",
		DetailsURL:   "https://www.example.com",
		Platforms:    []string{"macOS"},
		TagIDs:       []int{},
		MaxOSVersion: "15",
		DelayDays:    Int(7),
		LocalTime:    String("09:30:00"),
		Created:      Timestamp{referenceTime},
		Updated:      Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMSoftwareUpdateEnforcements.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMDMSoftwareUpdateEnforcementsService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/software_update_enforcements/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "Default")
		fmt.Fprint(w, msueListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMSoftwareUpdateEnforcements.GetByName(ctx, "Default")
	if err != nil {
		t.Errorf("MDMSoftwareUpdateEnforcements.GetByName returned error: %v", err)
	}

	want := &MDMSoftwareUpdateEnforcement{
		ID:            4,
		Name:          "Default",
		DetailsURL:    "https://www.example.com",
		Platforms:     []string{"macOS"},
		TagIDs:        []int{1, 2},
		OSVersion:     "14.1",
		BuildVersion:  "23B74",
		LocalDateTime: String("2023-11-05T09:30:00"),
		Created:       Timestamp{referenceTime},
		Updated:       Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMSoftwareUpdateEnforcements.GetByName returned %+v, want %+v", got, want)
	}
}

func TestMDMSoftwareUpdateEnforcementsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &MDMSoftwareUpdateEnforcementRequest{
		Name:         "Default",
		DetailsURL:   "https://www.example.com",
		Platforms:    []string{"macOS"},
		TagIDs:       []int{1, 2},
		MaxOSVersion: "15",
		DelayDays:    Int(7),
		LocalTime:    String("09:30:00"),
	}

	mux.HandleFunc("/mdm/software_update_enforcements/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMSoftwareUpdateEnforcementRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, msueCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMSoftwareUpdateEnforcements.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("MDMSoftwareUpdateEnforcements.Create returned error: %v", err)
	}

	want := &MDMSoftwareUpdateEnforcement{
		ID:           4,
		Name:         "Default",
		Platforms:    []string{"macOS"},
		TagIDs:       []int{1, 2},
		DetailsURL:   "https://www.example.com",
		MaxOSVersion: "15",
		DelayDays:    Int(7),
		LocalTime:    String("09:30:00"),
		Created:      Timestamp{referenceTime},
		Updated:      Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMSoftwareUpdateEnforcements.Create returned %+v, want %+v", got, want)
	}
}

func TestMDMSoftwareUpdateEnforcementsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &MDMSoftwareUpdateEnforcementRequest{
		Name:         "Default",
		Platforms:    []string{"macOS"},
		TagIDs:       []int{1, 2},
		DetailsURL:   "https://www.example.com",
		MaxOSVersion: "15",
		DelayDays:    Int(7),
		LocalTime:    String("09:30:00"),
	}

	mux.HandleFunc("/mdm/software_update_enforcements/4/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMSoftwareUpdateEnforcementRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, msueUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMSoftwareUpdateEnforcements.Update(ctx, 4, updateRequest)
	if err != nil {
		t.Errorf("MDMSoftwareUpdateEnforcements.Update returned error: %v", err)
	}

	want := &MDMSoftwareUpdateEnforcement{
		ID:           4,
		Name:         "Default",
		Platforms:    []string{"macOS"},
		TagIDs:       []int{1, 2},
		DetailsURL:   "https://www.example.com",
		MaxOSVersion: "15",
		DelayDays:    Int(7),
		LocalTime:    String("09:30:00"),
		Created:      Timestamp{referenceTime},
		Updated:      Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMSoftwareUpdateEnforcements.Update returned %+v, want %+v", got, want)
	}
}

func TestMDMSoftwareUpdateEnforcementsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/software_update_enforcements/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MDMSoftwareUpdateEnforcements.Delete(ctx, 4)
	if err != nil {
		t.Errorf("MDMSoftwareUpdateEnforcements.Delete returned error: %v", err)
	}
}
