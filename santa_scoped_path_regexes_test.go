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
    "next": null,
    "previous": null,
    "results": [
        {
            "id": 1,
            "configuration": 2,
            "name": "Homebrew",
            "description": "description",
            "policy": "ALLOW",
            "regex": "/opt/homebrew/.+",
            "primary_users": ["un", "deux"],
            "excluded_primary_users": ["trois", "quatre"],
            "serial_numbers": ["cinq", "six"],
            "excluded_serial_numbers": ["sept", "huit"],
            "tags": [9, 10],
            "excluded_tags": [11, 12],
            "created_at": "2022-07-22T01:02:03.444444",
            "updated_at": "2022-07-22T01:02:03.444444"
        }
    ]
}
`

var ssprGetJSONResponse = `
{
    "id": 1,
    "configuration": 2,
    "name": "Homebrew",
    "description": "description",
    "policy": "ALLOW",
    "regex": "/opt/homebrew/.+",
    "primary_users": ["un", "deux"],
    "excluded_primary_users": ["trois", "quatre"],
    "serial_numbers": ["cinq", "six"],
    "excluded_serial_numbers": ["sept", "huit"],
    "tags": [9, 10],
    "excluded_tags": [11, 12],
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var ssprCreateJSONResponse = `
{
    "id": 1,
    "configuration": 2,
    "name": "Homebrew",
    "description": "description",
    "policy": "ALLOW",
    "regex": "/opt/homebrew/.+",
    "primary_users": ["un", "deux"],
    "excluded_primary_users": ["trois", "quatre"],
    "serial_numbers": ["cinq", "six"],
    "excluded_serial_numbers": ["sept", "huit"],
    "tags": [9, 10],
    "excluded_tags": [11, 12],
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var ssprUpdateJSONResponse = `
{
    "id": 1,
    "configuration": 2,
    "name": "Downloads",
    "description": "",
    "policy": "BLOCK",
    "regex": "/Users/.+/Downloads/.+",
    "primary_users": [],
    "excluded_primary_users": [],
    "serial_numbers": [],
    "excluded_serial_numbers": [],
    "tags": [],
    "excluded_tags": [],
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestSantaScopedPathRegexesService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/santa/scoped_path_regexes/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "configuration_id", "2")
		fmt.Fprint(w, ssprListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.SantaScopedPathRegexes.List(ctx, 2, nil)
	if err != nil {
		t.Errorf("SantaScopedPathRegexes.List returned error: %v", err)
	}

	want := []SantaScopedPathRegex{
		{
			ID:                    1,
			ConfigurationID:       2,
			Name:                  "Homebrew",
			Description:           "description",
			Policy:                "ALLOW",
			Regex:                 "/opt/homebrew/.+",
			PrimaryUsers:          []string{"un", "deux"},
			ExcludedPrimaryUsers:  []string{"trois", "quatre"},
			SerialNumbers:         []string{"cinq", "six"},
			ExcludedSerialNumbers: []string{"sept", "huit"},
			TagIDs:                []int{9, 10},
			ExcludedTagIDs:        []int{11, 12},
			Created:               Timestamp{referenceTime},
			Updated:               Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("SantaScopedPathRegexes.List returned %+v, want %+v", got, want)
	}
}

func TestSantaScopedPathRegexesService_ListWithoutAConfiguration(t *testing.T) {
	// the endpoint answers the entries of one configuration and refuses a request without one,
	// so the client does not make it
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/santa/scoped_path_regexes/", func(w http.ResponseWriter, r *http.Request) {
		t.Error("SantaScopedPathRegexes.List made a request without a configuration")
	})

	ctx := context.Background()
	_, _, err := client.SantaScopedPathRegexes.List(ctx, 0, nil)
	if err == nil {
		t.Error("SantaScopedPathRegexes.List did not return an error")
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

	ctx := context.Background()
	got, _, err := client.SantaScopedPathRegexes.GetByID(ctx, 1)
	if err != nil {
		t.Errorf("SantaScopedPathRegexes.GetByID returned error: %v", err)
	}

	want := &SantaScopedPathRegex{
		ID:                    1,
		ConfigurationID:       2,
		Name:                  "Homebrew",
		Description:           "description",
		Policy:                "ALLOW",
		Regex:                 "/opt/homebrew/.+",
		PrimaryUsers:          []string{"un", "deux"},
		ExcludedPrimaryUsers:  []string{"trois", "quatre"},
		SerialNumbers:         []string{"cinq", "six"},
		ExcludedSerialNumbers: []string{"sept", "huit"},
		TagIDs:                []int{9, 10},
		ExcludedTagIDs:        []int{11, 12},
		Created:               Timestamp{referenceTime},
		Updated:               Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("SantaScopedPathRegexes.GetByID returned %+v, want %+v", got, want)
	}
}

func TestSantaScopedPathRegexesService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &SantaScopedPathRegexRequest{
		ConfigurationID:       2,
		Name:                  "Homebrew",
		Description:           "description",
		Policy:                "ALLOW",
		Regex:                 "/opt/homebrew/.+",
		PrimaryUsers:          []string{"un", "deux"},
		ExcludedPrimaryUsers:  []string{"trois", "quatre"},
		SerialNumbers:         []string{"cinq", "six"},
		ExcludedSerialNumbers: []string{"sept", "huit"},
		TagIDs:                []int{9, 10},
		ExcludedTagIDs:        []int{11, 12},
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

		fmt.Fprint(w, ssprCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.SantaScopedPathRegexes.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("SantaScopedPathRegexes.Create returned error: %v", err)
	}

	want := &SantaScopedPathRegex{
		ID:                    1,
		ConfigurationID:       2,
		Name:                  "Homebrew",
		Description:           "description",
		Policy:                "ALLOW",
		Regex:                 "/opt/homebrew/.+",
		PrimaryUsers:          []string{"un", "deux"},
		ExcludedPrimaryUsers:  []string{"trois", "quatre"},
		SerialNumbers:         []string{"cinq", "six"},
		ExcludedSerialNumbers: []string{"sept", "huit"},
		TagIDs:                []int{9, 10},
		ExcludedTagIDs:        []int{11, 12},
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

	// an update is a full update: the scope is written even when it carries nothing, because
	// the server keeps the stored value of an attribute left out
	updateRequest := &SantaScopedPathRegexRequest{
		ConfigurationID:       2,
		Name:                  "Downloads",
		Policy:                "BLOCK",
		Regex:                 "/Users/.+/Downloads/.+",
		PrimaryUsers:          []string{},
		ExcludedPrimaryUsers:  []string{},
		SerialNumbers:         []string{},
		ExcludedSerialNumbers: []string{},
		TagIDs:                []int{},
		ExcludedTagIDs:        []int{},
	}

	mux.HandleFunc("/santa/scoped_path_regexes/1/", func(w http.ResponseWriter, r *http.Request) {
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

	ctx := context.Background()
	got, _, err := client.SantaScopedPathRegexes.Update(ctx, 1, updateRequest)
	if err != nil {
		t.Errorf("SantaScopedPathRegexes.Update returned error: %v", err)
	}

	want := &SantaScopedPathRegex{
		ID:                    1,
		ConfigurationID:       2,
		Name:                  "Downloads",
		Description:           "",
		Policy:                "BLOCK",
		Regex:                 "/Users/.+/Downloads/.+",
		PrimaryUsers:          []string{},
		ExcludedPrimaryUsers:  []string{},
		SerialNumbers:         []string{},
		ExcludedSerialNumbers: []string{},
		TagIDs:                []int{},
		ExcludedTagIDs:        []int{},
		Created:               Timestamp{referenceTime},
		Updated:               Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("SantaScopedPathRegexes.Update returned %+v, want %+v", got, want)
	}
}

func TestSantaScopedPathRegexesService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/santa/scoped_path_regexes/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.SantaScopedPathRegexes.Delete(ctx, 1)
	if err != nil {
		t.Errorf("SantaScopedPathRegexes.Delete returned error: %v", err)
	}
}
