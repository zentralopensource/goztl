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
    "next": null,
    "previous": null,
    "results": [
        {
            "id": 1,
            "configuration": 2,
            "name": "Lockdown for the fleet",
            "description": "description",
            "client_mode": 2,
            "event_detail_source": "CUSTOM",
            "event_detail_url": "https://zentral.com/blocked/%file_identifier%",
            "event_detail_text": "More info",
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

var sscmGetJSONResponse = `
{
    "id": 1,
    "configuration": 2,
    "name": "Lockdown for the fleet",
    "description": "description",
    "client_mode": 2,
    "event_detail_source": "CUSTOM",
    "event_detail_url": "https://zentral.com/blocked/%file_identifier%",
    "event_detail_text": "More info",
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

var sscmCreateJSONResponse = `
{
    "id": 1,
    "configuration": 2,
    "name": "Lockdown for the fleet",
    "description": "description",
    "client_mode": 2,
    "event_detail_source": "CUSTOM",
    "event_detail_url": "https://zentral.com/blocked/%file_identifier%",
    "event_detail_text": "More info",
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

var sscmUpdateJSONResponse = `
{
    "id": 1,
    "configuration": 2,
    "name": "Monitor for the fleet",
    "description": "",
    "client_mode": 1,
    "event_detail_source": "INHERIT",
    "event_detail_url": "",
    "event_detail_text": "",
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

func TestSantaScopedClientModesService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/santa/scoped_client_modes/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "configuration_id", "2")
		fmt.Fprint(w, sscmListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.SantaScopedClientModes.List(ctx, 2, nil)
	if err != nil {
		t.Errorf("SantaScopedClientModes.List returned error: %v", err)
	}

	want := []SantaScopedClientMode{
		{
			ID:                    1,
			ConfigurationID:       2,
			Name:                  "Lockdown for the fleet",
			Description:           "description",
			ClientMode:            2,
			EventDetailSource:     "CUSTOM",
			EventDetailURL:        "https://zentral.com/blocked/%file_identifier%",
			EventDetailText:       "More info",
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
		t.Errorf("SantaScopedClientModes.List returned %+v, want %+v", got, want)
	}
}

func TestSantaScopedClientModesService_ListWithoutAConfiguration(t *testing.T) {
	// the endpoint answers the entries of one configuration and refuses a request without one,
	// so the client does not make it
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/santa/scoped_client_modes/", func(w http.ResponseWriter, r *http.Request) {
		t.Error("SantaScopedClientModes.List made a request without a configuration")
	})

	ctx := context.Background()
	_, _, err := client.SantaScopedClientModes.List(ctx, 0, nil)
	if err == nil {
		t.Error("SantaScopedClientModes.List did not return an error")
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

	ctx := context.Background()
	got, _, err := client.SantaScopedClientModes.GetByID(ctx, 1)
	if err != nil {
		t.Errorf("SantaScopedClientModes.GetByID returned error: %v", err)
	}

	want := &SantaScopedClientMode{
		ID:                    1,
		ConfigurationID:       2,
		Name:                  "Lockdown for the fleet",
		Description:           "description",
		ClientMode:            2,
		EventDetailSource:     "CUSTOM",
		EventDetailURL:        "https://zentral.com/blocked/%file_identifier%",
		EventDetailText:       "More info",
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
		t.Errorf("SantaScopedClientModes.GetByID returned %+v, want %+v", got, want)
	}
}

func TestSantaScopedClientModesService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &SantaScopedClientModeRequest{
		ConfigurationID:       2,
		Name:                  "Lockdown for the fleet",
		Description:           "description",
		ClientMode:            2,
		EventDetailSource:     "CUSTOM",
		EventDetailURL:        "https://zentral.com/blocked/%file_identifier%",
		EventDetailText:       "More info",
		PrimaryUsers:          []string{"un", "deux"},
		ExcludedPrimaryUsers:  []string{"trois", "quatre"},
		SerialNumbers:         []string{"cinq", "six"},
		ExcludedSerialNumbers: []string{"sept", "huit"},
		TagIDs:                []int{9, 10},
		ExcludedTagIDs:        []int{11, 12},
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

		fmt.Fprint(w, sscmCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.SantaScopedClientModes.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("SantaScopedClientModes.Create returned error: %v", err)
	}

	want := &SantaScopedClientMode{
		ID:                    1,
		ConfigurationID:       2,
		Name:                  "Lockdown for the fleet",
		Description:           "description",
		ClientMode:            2,
		EventDetailSource:     "CUSTOM",
		EventDetailURL:        "https://zentral.com/blocked/%file_identifier%",
		EventDetailText:       "More info",
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
		t.Errorf("SantaScopedClientModes.Create returned %+v, want %+v", got, want)
	}
}

func TestSantaScopedClientModesService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	// an update is a full update: the scope and the event detail are written even when they
	// carry nothing, because the server keeps the stored value of an attribute left out
	updateRequest := &SantaScopedClientModeRequest{
		ConfigurationID:       2,
		Name:                  "Monitor for the fleet",
		ClientMode:            1,
		EventDetailSource:     "INHERIT",
		PrimaryUsers:          []string{},
		ExcludedPrimaryUsers:  []string{},
		SerialNumbers:         []string{},
		ExcludedSerialNumbers: []string{},
		TagIDs:                []int{},
		ExcludedTagIDs:        []int{},
	}

	mux.HandleFunc("/santa/scoped_client_modes/1/", func(w http.ResponseWriter, r *http.Request) {
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

	ctx := context.Background()
	got, _, err := client.SantaScopedClientModes.Update(ctx, 1, updateRequest)
	if err != nil {
		t.Errorf("SantaScopedClientModes.Update returned error: %v", err)
	}

	want := &SantaScopedClientMode{
		ID:                    1,
		ConfigurationID:       2,
		Name:                  "Monitor for the fleet",
		Description:           "",
		ClientMode:            1,
		EventDetailSource:     "INHERIT",
		EventDetailURL:        "",
		EventDetailText:       "",
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
		t.Errorf("SantaScopedClientModes.Update returned %+v, want %+v", got, want)
	}
}

func TestSantaScopedClientModesService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/santa/scoped_client_modes/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.SantaScopedClientModes.Delete(ctx, 1)
	if err != nil {
		t.Errorf("SantaScopedClientModes.Delete returned error: %v", err)
	}
}
