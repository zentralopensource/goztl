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

var srListJSONResponse = `
[
    {
        "id": 1,
        "configuration": 2,
	"policy": 1,
	"target_type": "BINARY",
	"target_identifier": "311fe3feed16b9cd8df0f8b1517be5cb86048707df4889ba8dc37d4d68866d02",
	"description": "description",
	"custom_msg": "custom message",
	"ruleset": null,
	"primary_users": ["un", "deux"],
	"excluded_primary_users": ["trois", "quatre"],
	"serial_numbers": ["cinq", "six"],
	"excluded_serial_numbers": ["sept", "huit"],
	"tags": [9, 10],
	"excluded_tags": [11, 12],
	"version": 1,
	"created_at": "2022-07-22T01:02:03.444444",
	"updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var srGetJSONResponse = `
{
    "id": 1,
    "configuration": 2,
    "policy": 1,
    "target_type": "BINARY",
    "target_identifier": "311fe3feed16b9cd8df0f8b1517be5cb86048707df4889ba8dc37d4d68866d02",
    "description": "description",
    "custom_msg": "custom message",
    "ruleset": 1,
    "primary_users": ["un", "deux"],
    "excluded_primary_users": ["trois", "quatre"],
    "serial_numbers": ["cinq", "six"],
    "excluded_serial_numbers": ["sept", "huit"],
    "tags": [9, 10],
    "excluded_tags": [11, 12],
    "version": 1,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var srCreateJSONResponse = `
{
    "id": 1,
    "configuration": 2,
    "policy": 1,
    "target_type": "BINARY",
    "target_identifier": "311fe3feed16b9cd8df0f8b1517be5cb86048707df4889ba8dc37d4d68866d02",
    "description": "description",
    "custom_msg": "custom message",
    "ruleset": null,
    "primary_users": ["un", "deux"],
    "excluded_primary_users": ["trois", "quatre"],
    "serial_numbers": ["cinq", "six"],
    "excluded_serial_numbers": ["sept", "huit"],
    "tags": [9, 10],
    "excluded_tags": [11, 12],
    "version": 1,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var srUpdateJSONResponse = `
{
    "id": 1,
    "configuration": 2,
    "policy": 1,
    "target_type": "BINARY",
    "target_identifier": "311fe3feed16b9cd8df0f8b1517be5cb86048707df4889ba8dc37d4d68866d02",
    "description": "",
    "custom_msg": "",
    "ruleset": null,
    "primary_users": [],
    "excluded_primary_users": [],
    "serial_numbers": [],
    "excluded_serial_numbers": [],
    "tags": [],
    "excluded_tags": [],
    "version": 2,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestSantaRulesService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/santa/rules/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, srListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.SantaRules.List(ctx, nil)
	if err != nil {
		t.Errorf("SantaRules.List returned error: %v", err)
	}

	want := []SantaRule{
		{
			ID:                    1,
			ConfigurationID:       2,
			Policy:                1,
			TargetType:            "BINARY",
			TargetIdentifier:      "311fe3feed16b9cd8df0f8b1517be5cb86048707df4889ba8dc37d4d68866d02",
			Description:           "description",
			CustomMessage:         "custom message",
			RulesetID:             nil,
			PrimaryUsers:          []string{"un", "deux"},
			ExcludedPrimaryUsers:  []string{"trois", "quatre"},
			SerialNumbers:         []string{"cinq", "six"},
			ExcludedSerialNumbers: []string{"sept", "huit"},
			TagIDs:                []int{9, 10},
			ExcludedTagIDs:        []int{11, 12},
			Version:               1,
			Created:               Timestamp{referenceTime},
			Updated:               Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("SantaRules.List returned %+v, want %+v", got, want)
	}
}

func TestSantaRulesService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/santa/rules/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, srGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.SantaRules.GetByID(ctx, 1)
	if err != nil {
		t.Errorf("SantaRules.GetByID returned error: %v", err)
	}

	want := &SantaRule{
		ID:                    1,
		ConfigurationID:       2,
		Policy:                1,
		TargetType:            "BINARY",
		TargetIdentifier:      "311fe3feed16b9cd8df0f8b1517be5cb86048707df4889ba8dc37d4d68866d02",
		Description:           "description",
		CustomMessage:         "custom message",
		RulesetID:             Int(1),
		PrimaryUsers:          []string{"un", "deux"},
		ExcludedPrimaryUsers:  []string{"trois", "quatre"},
		SerialNumbers:         []string{"cinq", "six"},
		ExcludedSerialNumbers: []string{"sept", "huit"},
		TagIDs:                []int{9, 10},
		ExcludedTagIDs:        []int{11, 12},
		Version:               1,
		Created:               Timestamp{referenceTime},
		Updated:               Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("SantaRules.GetByID returned %+v, want %+v", got, want)
	}
}

func TestSantaRulesService_GetByConfigurationID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/santa/rules/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "configuration_id", "2")
		fmt.Fprint(w, srListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.SantaRules.GetByConfigurationID(ctx, 2)
	if err != nil {
		t.Errorf("SantaRules.GetByConfigurationID returned error: %v", err)
	}

	want := []SantaRule{
		{
			ID:                    1,
			ConfigurationID:       2,
			Policy:                1,
			TargetType:            "BINARY",
			TargetIdentifier:      "311fe3feed16b9cd8df0f8b1517be5cb86048707df4889ba8dc37d4d68866d02",
			Description:           "description",
			CustomMessage:         "custom message",
			RulesetID:             nil,
			PrimaryUsers:          []string{"un", "deux"},
			ExcludedPrimaryUsers:  []string{"trois", "quatre"},
			SerialNumbers:         []string{"cinq", "six"},
			ExcludedSerialNumbers: []string{"sept", "huit"},
			TagIDs:                []int{9, 10},
			ExcludedTagIDs:        []int{11, 12},
			Version:               1,
			Created:               Timestamp{referenceTime},
			Updated:               Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("SantaRules.GetByConfigurationID returned %+v, want %+v", got, want)
	}
}

func TestSantaRulesService_GetByTargetIdentifier(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/santa/rules/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "target_identifier", "311fe3feed16b9cd8df0f8b1517be5cb86048707df4889ba8dc37d4d68866d02")
		fmt.Fprint(w, srListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.SantaRules.GetByTargetIdentifier(ctx, "311fe3feed16b9cd8df0f8b1517be5cb86048707df4889ba8dc37d4d68866d02")
	if err != nil {
		t.Errorf("SantaRules.GetByTargetIdentifier returned error: %v", err)
	}

	want := []SantaRule{
		{
			ID:                    1,
			ConfigurationID:       2,
			Policy:                1,
			TargetType:            "BINARY",
			TargetIdentifier:      "311fe3feed16b9cd8df0f8b1517be5cb86048707df4889ba8dc37d4d68866d02",
			Description:           "description",
			CustomMessage:         "custom message",
			RulesetID:             nil,
			PrimaryUsers:          []string{"un", "deux"},
			ExcludedPrimaryUsers:  []string{"trois", "quatre"},
			SerialNumbers:         []string{"cinq", "six"},
			ExcludedSerialNumbers: []string{"sept", "huit"},
			TagIDs:                []int{9, 10},
			ExcludedTagIDs:        []int{11, 12},
			Version:               1,
			Created:               Timestamp{referenceTime},
			Updated:               Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("SantaRules.GetByTargetIdentifier returned %+v, want %+v", got, want)
	}
}

func TestSantaRulesService_GetByTargetType(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/santa/rules/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "target_type", "BINARY")
		fmt.Fprint(w, srListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.SantaRules.GetByTargetType(ctx, "BINARY")
	if err != nil {
		t.Errorf("SantaRules.GetByTargetType returned error: %v", err)
	}

	want := []SantaRule{
		{
			ID:                    1,
			ConfigurationID:       2,
			Policy:                1,
			TargetType:            "BINARY",
			TargetIdentifier:      "311fe3feed16b9cd8df0f8b1517be5cb86048707df4889ba8dc37d4d68866d02",
			Description:           "description",
			CustomMessage:         "custom message",
			RulesetID:             nil,
			PrimaryUsers:          []string{"un", "deux"},
			ExcludedPrimaryUsers:  []string{"trois", "quatre"},
			SerialNumbers:         []string{"cinq", "six"},
			ExcludedSerialNumbers: []string{"sept", "huit"},
			TagIDs:                []int{9, 10},
			ExcludedTagIDs:        []int{11, 12},
			Version:               1,
			Created:               Timestamp{referenceTime},
			Updated:               Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("SantaRules.GetByTargetType returned %+v, want %+v", got, want)
	}
}

func TestSantaRulesService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &SantaRuleRequest{
		ConfigurationID:       2,
		Policy:                1,
		TargetType:            "BINARY",
		TargetIdentifier:      "311fe3feed16b9cd8df0f8b1517be5cb86048707df4889ba8dc37d4d68866d02",
		Description:           "description",
		CustomMessage:         "custom message",
		PrimaryUsers:          []string{"un", "deux"},
		ExcludedPrimaryUsers:  []string{"trois", "quatre"},
		SerialNumbers:         []string{"cinq", "six"},
		ExcludedSerialNumbers: []string{"sept", "huit"},
		TagIDs:                []int{9, 10},
		ExcludedTagIDs:        []int{11, 12},
	}

	mux.HandleFunc("/santa/rules/", func(w http.ResponseWriter, r *http.Request) {
		v := new(SantaRuleRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, srCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.SantaRules.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("SantaRules.Create returned error: %v", err)
	}

	want := &SantaRule{
		ID:                    1,
		ConfigurationID:       2,
		Policy:                1,
		TargetType:            "BINARY",
		TargetIdentifier:      "311fe3feed16b9cd8df0f8b1517be5cb86048707df4889ba8dc37d4d68866d02",
		Description:           "description",
		CustomMessage:         "custom message",
		RulesetID:             nil,
		PrimaryUsers:          []string{"un", "deux"},
		ExcludedPrimaryUsers:  []string{"trois", "quatre"},
		SerialNumbers:         []string{"cinq", "six"},
		ExcludedSerialNumbers: []string{"sept", "huit"},
		TagIDs:                []int{9, 10},
		ExcludedTagIDs:        []int{11, 12},
		Version:               1,
		Created:               Timestamp{referenceTime},
		Updated:               Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("SantaRules.Create returned %+v, want %+v", got, want)
	}
}

func TestSantaRulesService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &SantaRuleRequest{
		ConfigurationID:  2,
		Policy:           1,
		TargetType:       "BINARY",
		TargetIdentifier: "311fe3feed16b9cd8df0f8b1517be5cb86048707df4889ba8dc37d4d68866d02",
	}

	mux.HandleFunc("/santa/rules/1/", func(w http.ResponseWriter, r *http.Request) {
		v := new(SantaRuleRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, srUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.SantaRules.Update(ctx, 1, updateRequest)
	if err != nil {
		t.Errorf("SantaRules.Update returned error: %v", err)
	}

	want := &SantaRule{
		ID:                    1,
		ConfigurationID:       2,
		Policy:                1,
		TargetType:            "BINARY",
		TargetIdentifier:      "311fe3feed16b9cd8df0f8b1517be5cb86048707df4889ba8dc37d4d68866d02",
		Description:           "",
		CustomMessage:         "",
		RulesetID:             nil,
		PrimaryUsers:          []string{},
		ExcludedPrimaryUsers:  []string{},
		SerialNumbers:         []string{},
		ExcludedSerialNumbers: []string{},
		TagIDs:                []int{},
		ExcludedTagIDs:        []int{},
		Version:               2,
		Created:               Timestamp{referenceTime},
		Updated:               Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("SantaRules.Update returned %+v, want %+v", got, want)
	}
}

func TestSantaRulesService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/santa/rules/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.SantaRules.Delete(ctx, 1)
	if err != nil {
		t.Errorf("SantaRules.Delete returned error: %v", err)
	}
}
