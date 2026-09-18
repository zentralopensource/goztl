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

var sscmListJSONResponse = `
{
    "count": 1,
    "results": [
        {
            "id": 1,
            "name": "Yolo",
            "description": "Description",
            "serial_numbers": ["un", "deux"],
            "excluded_serial_numbers": ["trois", "quatre"],
            "primary_users": ["cinq", "six"],
            "excluded_primary_users": ["sept", "huit"],
            "created_at": "2022-07-22T01:02:03.444444",
            "updated_at": "2022-07-22T01:02:03.444444",
            "client_mode": 2,
            "event_detail_source": "CUSTOM",
            "event_detail_url": "https://www.example.com/blocked/",
            "event_detail_text": "Request an exception",
            "configuration": 2,
            "tags": [9, 10],
            "excluded_tags": [11, 12]
        }
    ]
}
`

var sscmGetJSONResponse = `
{
    "id": 1,
    "name": "Yolo",
    "description": "Description",
    "serial_numbers": ["un", "deux"],
    "excluded_serial_numbers": ["trois", "quatre"],
    "primary_users": ["cinq", "six"],
    "excluded_primary_users": ["sept", "huit"],
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444",
    "client_mode": 2,
    "event_detail_source": "CUSTOM",
    "event_detail_url": "https://www.example.com/blocked/",
    "event_detail_text": "Request an exception",
    "configuration": 2,
    "tags": [9, 10],
    "excluded_tags": [11, 12]
}
`

var sscmCreateJSONResponse = `
{
    "id": 4,
    "name": "Fomo",
    "description": "",
    "serial_numbers": [],
    "excluded_serial_numbers": [],
    "primary_users": [],
    "excluded_primary_users": [],
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444",
    "client_mode": 1,
    "event_detail_source": "INHERIT",
    "event_detail_url": "",
    "event_detail_text": "",
    "configuration": 3,
    "tags": [],
    "excluded_tags": []
}
`

var sscmUpdateJSONResponse = `
{
    "id": 4,
    "name": "Fomo",
    "description": "Description",
    "serial_numbers": ["un"],
    "excluded_serial_numbers": [],
    "primary_users": [],
    "excluded_primary_users": [],
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444",
    "client_mode": 2,
    "event_detail_source": "NONE",
    "event_detail_url": "",
    "event_detail_text": "",
    "configuration": 3,
    "tags": [],
    "excluded_tags": []
}
`

func sscmWant() *SantaScopedClientMode {
	return &SantaScopedClientMode{
		ID:                    1,
		ConfigurationID:       2,
		Name:                  "Yolo",
		Description:           "Description",
		SerialNumbers:         []string{"un", "deux"},
		ExcludedSerialNumbers: []string{"trois", "quatre"},
		PrimaryUsers:          []string{"cinq", "six"},
		ExcludedPrimaryUsers:  []string{"sept", "huit"},
		TagIDs:                []int{9, 10},
		ExcludedTagIDs:        []int{11, 12},
		ClientMode:            2,
		EventDetailSource:     "CUSTOM",
		EventDetailURL:        "https://www.example.com/blocked/",
		EventDetailText:       "Request an exception",
		Created:               Timestamp{referenceTime},
		Updated:               Timestamp{referenceTime},
	}
}

func TestSantaScopedClientModesService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/santa/scoped_client_modes/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, sscmGetJSONResponse)
	})

	got, _, err := client.SantaScopedClientModes.GetByID(context.Background(), 1)
	if err != nil {
		t.Errorf("SantaScopedClientModes.GetByID returned error: %v", err)
	}

	if want := sscmWant(); !cmp.Equal(got, want) {
		t.Errorf("SantaScopedClientModes.GetByID returned %+v, want %+v", got, want)
	}
}

func TestSantaScopedClientModesService_GetByConfigurationID(t *testing.T) {
	// the configuration is required on the list endpoint: the server takes one authorization
	// decision per entry, so it never answers for every configuration at once
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/santa/scoped_client_modes/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "configuration_id", "2")
		fmt.Fprint(w, sscmListJSONResponse)
	})

	got, _, err := client.SantaScopedClientModes.GetByConfigurationID(context.Background(), 2)
	if err != nil {
		t.Errorf("SantaScopedClientModes.GetByConfigurationID returned error: %v", err)
	}

	if want := []SantaScopedClientMode{*sscmWant()}; !cmp.Equal(got, want) {
		t.Errorf("SantaScopedClientModes.GetByConfigurationID returned %+v, want %+v", got, want)
	}
}

func TestSantaScopedClientModesService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &SantaScopedClientModeRequest{
		ConfigurationID:       3,
		Name:                  "Fomo",
		SerialNumbers:         make([]string, 0),
		ExcludedSerialNumbers: make([]string, 0),
		PrimaryUsers:          make([]string, 0),
		ExcludedPrimaryUsers:  make([]string, 0),
		TagIDs:                make([]int, 0),
		ExcludedTagIDs:        make([]int, 0),
		ClientMode:            1,
		EventDetailSource:     "INHERIT",
	}

	mux.HandleFunc("/santa/scoped_client_modes/", func(w http.ResponseWriter, r *http.Request) {
		v := new(SantaScopedClientModeRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, sscmCreateJSONResponse)
	})

	got, _, err := client.SantaScopedClientModes.Create(context.Background(), createRequest)
	if err != nil {
		t.Errorf("SantaScopedClientModes.Create returned error: %v", err)
	}

	want := &SantaScopedClientMode{
		ID:                    4,
		ConfigurationID:       3,
		Name:                  "Fomo",
		SerialNumbers:         make([]string, 0),
		ExcludedSerialNumbers: make([]string, 0),
		PrimaryUsers:          make([]string, 0),
		ExcludedPrimaryUsers:  make([]string, 0),
		TagIDs:                make([]int, 0),
		ExcludedTagIDs:        make([]int, 0),
		ClientMode:            1,
		EventDetailSource:     "INHERIT",
		Created:               Timestamp{referenceTime},
		Updated:               Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("SantaScopedClientModes.Create returned %+v, want %+v", got, want)
	}
}

func TestSantaScopedClientModesService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &SantaScopedClientModeRequest{
		ConfigurationID:       3,
		Name:                  "Fomo",
		Description:           "Description",
		SerialNumbers:         []string{"un"},
		ExcludedSerialNumbers: make([]string, 0),
		PrimaryUsers:          make([]string, 0),
		ExcludedPrimaryUsers:  make([]string, 0),
		TagIDs:                make([]int, 0),
		ExcludedTagIDs:        make([]int, 0),
		ClientMode:            2,
		EventDetailSource:     "NONE",
	}

	mux.HandleFunc("/santa/scoped_client_modes/4/", func(w http.ResponseWriter, r *http.Request) {
		v := new(SantaScopedClientModeRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, sscmUpdateJSONResponse)
	})

	got, _, err := client.SantaScopedClientModes.Update(context.Background(), 4, updateRequest)
	if err != nil {
		t.Errorf("SantaScopedClientModes.Update returned error: %v", err)
	}

	want := &SantaScopedClientMode{
		ID:                    4,
		ConfigurationID:       3,
		Name:                  "Fomo",
		Description:           "Description",
		SerialNumbers:         []string{"un"},
		ExcludedSerialNumbers: make([]string, 0),
		PrimaryUsers:          make([]string, 0),
		ExcludedPrimaryUsers:  make([]string, 0),
		TagIDs:                make([]int, 0),
		ExcludedTagIDs:        make([]int, 0),
		ClientMode:            2,
		EventDetailSource:     "NONE",
		Created:               Timestamp{referenceTime},
		Updated:               Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("SantaScopedClientModes.Update returned %+v, want %+v", got, want)
	}
}

func TestSantaScopedClientModesService_UpdateSendsTheWholeScope(t *testing.T) {
	// every scope attribute is required on an update, and an empty one has to go out as an
	// empty list: a nil slice is null on the wire, which the server refuses
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &SantaScopedClientModeRequest{
		ConfigurationID:       3,
		Name:                  "Fomo",
		SerialNumbers:         make([]string, 0),
		ExcludedSerialNumbers: make([]string, 0),
		PrimaryUsers:          make([]string, 0),
		ExcludedPrimaryUsers:  make([]string, 0),
		TagIDs:                make([]int, 0),
		ExcludedTagIDs:        make([]int, 0),
		ClientMode:            2,
		EventDetailSource:     "INHERIT",
	}

	mux.HandleFunc("/santa/scoped_client_modes/4/", func(w http.ResponseWriter, r *http.Request) {
		testBody(t, r, `{"configuration":3,"name":"Fomo","description":"","serial_numbers":[],`+
			`"excluded_serial_numbers":[],"primary_users":[],"excluded_primary_users":[],`+
			`"tags":[],"excluded_tags":[],"client_mode":2,"event_detail_source":"INHERIT",`+
			`"event_detail_url":"","event_detail_text":""}`+"\n")
		fmt.Fprint(w, sscmUpdateJSONResponse)
	})

	_, _, err := client.SantaScopedClientModes.Update(context.Background(), 4, updateRequest)
	if err != nil {
		t.Errorf("SantaScopedClientModes.Update returned error: %v", err)
	}
}

func TestSantaScopedClientModesService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/santa/scoped_client_modes/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	_, err := client.SantaScopedClientModes.Delete(context.Background(), 1)
	if err != nil {
		t.Errorf("SantaScopedClientModes.Delete returned error: %v", err)
	}
}
