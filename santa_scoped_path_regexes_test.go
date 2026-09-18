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

var ssprListJSONResponse = `
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
            "policy": "ALLOW",
            "regex": "/Library/Example/",
            "configuration": 2,
            "tags": [9, 10],
            "excluded_tags": [11, 12]
        }
    ]
}
`

var ssprGetJSONResponse = `
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
    "policy": "ALLOW",
    "regex": "/Library/Example/",
    "configuration": 2,
    "tags": [9, 10],
    "excluded_tags": [11, 12]
}
`

var ssprCreateJSONResponse = `
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
    "policy": "BLOCK",
    "regex": ".*/Downloads/",
    "configuration": 3,
    "tags": [],
    "excluded_tags": []
}
`

var ssprUpdateJSONResponse = `
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
    "policy": "ALLOW",
    "regex": "/Library/Example/",
    "configuration": 3,
    "tags": [],
    "excluded_tags": []
}
`

func ssprWant() *SantaScopedPathRegex {
	return &SantaScopedPathRegex{
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
		Policy:                "ALLOW",
		Regex:                 "/Library/Example/",
		Created:               Timestamp{referenceTime},
		Updated:               Timestamp{referenceTime},
	}
}

func TestSantaScopedPathRegexesService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/santa/scoped_path_regexes/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, ssprGetJSONResponse)
	})

	got, _, err := client.SantaScopedPathRegexes.GetByID(context.Background(), 1)
	if err != nil {
		t.Errorf("SantaScopedPathRegexes.GetByID returned error: %v", err)
	}

	if want := ssprWant(); !cmp.Equal(got, want) {
		t.Errorf("SantaScopedPathRegexes.GetByID returned %+v, want %+v", got, want)
	}
}

func TestSantaScopedPathRegexesService_GetByConfigurationID(t *testing.T) {
	// the configuration is required on the list endpoint: the server takes one authorization
	// decision per entry, so it never answers for every configuration at once
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/santa/scoped_path_regexes/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "configuration_id", "2")
		fmt.Fprint(w, ssprListJSONResponse)
	})

	got, _, err := client.SantaScopedPathRegexes.GetByConfigurationID(context.Background(), 2)
	if err != nil {
		t.Errorf("SantaScopedPathRegexes.GetByConfigurationID returned error: %v", err)
	}

	if want := []SantaScopedPathRegex{*ssprWant()}; !cmp.Equal(got, want) {
		t.Errorf("SantaScopedPathRegexes.GetByConfigurationID returned %+v, want %+v", got, want)
	}
}

func TestSantaScopedPathRegexesService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &SantaScopedPathRegexRequest{
		ConfigurationID:       3,
		Name:                  "Fomo",
		SerialNumbers:         make([]string, 0),
		ExcludedSerialNumbers: make([]string, 0),
		PrimaryUsers:          make([]string, 0),
		ExcludedPrimaryUsers:  make([]string, 0),
		TagIDs:                make([]int, 0),
		ExcludedTagIDs:        make([]int, 0),
		Policy:                "BLOCK",
		Regex:                 ".*/Downloads/",
	}

	mux.HandleFunc("/santa/scoped_path_regexes/", func(w http.ResponseWriter, r *http.Request) {
		v := new(SantaScopedPathRegexRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, ssprCreateJSONResponse)
	})

	got, _, err := client.SantaScopedPathRegexes.Create(context.Background(), createRequest)
	if err != nil {
		t.Errorf("SantaScopedPathRegexes.Create returned error: %v", err)
	}

	want := &SantaScopedPathRegex{
		ID:                    4,
		ConfigurationID:       3,
		Name:                  "Fomo",
		SerialNumbers:         make([]string, 0),
		ExcludedSerialNumbers: make([]string, 0),
		PrimaryUsers:          make([]string, 0),
		ExcludedPrimaryUsers:  make([]string, 0),
		TagIDs:                make([]int, 0),
		ExcludedTagIDs:        make([]int, 0),
		Policy:                "BLOCK",
		Regex:                 ".*/Downloads/",
		Created:               Timestamp{referenceTime},
		Updated:               Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("SantaScopedPathRegexes.Create returned %+v, want %+v", got, want)
	}
}

func TestSantaScopedPathRegexesService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &SantaScopedPathRegexRequest{
		ConfigurationID:       3,
		Name:                  "Fomo",
		Description:           "Description",
		SerialNumbers:         []string{"un"},
		ExcludedSerialNumbers: make([]string, 0),
		PrimaryUsers:          make([]string, 0),
		ExcludedPrimaryUsers:  make([]string, 0),
		TagIDs:                make([]int, 0),
		ExcludedTagIDs:        make([]int, 0),
		Policy:                "ALLOW",
		Regex:                 "/Library/Example/",
	}

	mux.HandleFunc("/santa/scoped_path_regexes/4/", func(w http.ResponseWriter, r *http.Request) {
		v := new(SantaScopedPathRegexRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, ssprUpdateJSONResponse)
	})

	got, _, err := client.SantaScopedPathRegexes.Update(context.Background(), 4, updateRequest)
	if err != nil {
		t.Errorf("SantaScopedPathRegexes.Update returned error: %v", err)
	}

	want := &SantaScopedPathRegex{
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
		Policy:                "ALLOW",
		Regex:                 "/Library/Example/",
		Created:               Timestamp{referenceTime},
		Updated:               Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("SantaScopedPathRegexes.Update returned %+v, want %+v", got, want)
	}
}

func TestSantaScopedPathRegexesService_UpdateSendsTheWholeScope(t *testing.T) {
	// every scope attribute is required on an update, and an empty one has to go out as an
	// empty list: a nil slice is null on the wire, which the server refuses
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &SantaScopedPathRegexRequest{
		ConfigurationID:       3,
		Name:                  "Fomo",
		SerialNumbers:         make([]string, 0),
		ExcludedSerialNumbers: make([]string, 0),
		PrimaryUsers:          make([]string, 0),
		ExcludedPrimaryUsers:  make([]string, 0),
		TagIDs:                make([]int, 0),
		ExcludedTagIDs:        make([]int, 0),
		Policy:                "ALLOW",
		Regex:                 "/Library/Example/",
	}

	mux.HandleFunc("/santa/scoped_path_regexes/4/", func(w http.ResponseWriter, r *http.Request) {
		testBody(t, r, `{"configuration":3,"name":"Fomo","description":"","serial_numbers":[],`+
			`"excluded_serial_numbers":[],"primary_users":[],"excluded_primary_users":[],`+
			`"tags":[],"excluded_tags":[],"policy":"ALLOW","regex":"/Library/Example/"}`+"\n")
		fmt.Fprint(w, ssprUpdateJSONResponse)
	})

	_, _, err := client.SantaScopedPathRegexes.Update(context.Background(), 4, updateRequest)
	if err != nil {
		t.Errorf("SantaScopedPathRegexes.Update returned error: %v", err)
	}
}

func TestSantaScopedPathRegexesService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/santa/scoped_path_regexes/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	_, err := client.SantaScopedPathRegexes.Delete(context.Background(), 1)
	if err != nil {
		t.Errorf("SantaScopedPathRegexes.Delete returned error: %v", err)
	}
}
