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

var mcoListJSONResponse = `
[
    {
        "id": 4,
        "name": "Laptop",
	"predicate": "machine_type == \"laptop\"",
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var mcoGetJSONResponse = `
{
    "id": 4,
    "name": "Laptop",
    "predicate": "machine_type == \"laptop\"",
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var mcoCreateJSONResponse = `
{
    "id": 4,
    "name": "Laptop",
    "predicate": "machine_type == \"laptop\"",
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var mcoUpdateJSONResponse = `
{
    "id": 4,
    "name": "Laptop",
    "predicate": "machine_type == \"laptop\"",
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestMonolithConditionsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/conditions/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mcoListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithConditions.List(ctx, nil)
	if err != nil {
		t.Errorf("MonolithConditions.List returned error: %v", err)
	}

	want := []MonolithCondition{
		{
			ID:        4,
			Name:      "Laptop",
			Predicate: "machine_type == \"laptop\"",
			Created:   Timestamp{referenceTime},
			Updated:   Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithConditions.List returned %+v, want %+v", got, want)
	}
}

func TestMonolithConditionsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/conditions/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mcoGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithConditions.GetByID(ctx, 4)
	if err != nil {
		t.Errorf("MonolithConditions.GetByID returned error: %v", err)
	}

	want := &MonolithCondition{
		ID:        4,
		Name:      "Laptop",
		Predicate: "machine_type == \"laptop\"",
		Created:   Timestamp{referenceTime},
		Updated:   Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithConditions.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMonolithConditionsService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/conditions/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "Laptop")
		fmt.Fprint(w, mcoListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithConditions.GetByName(ctx, "Laptop")
	if err != nil {
		t.Errorf("MonolithConditions.GetByName returned error: %v", err)
	}

	want := &MonolithCondition{
		ID:        4,
		Name:      "Laptop",
		Predicate: "machine_type == \"laptop\"",
		Created:   Timestamp{referenceTime},
		Updated:   Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithConditions.GetByName returned %+v, want %+v", got, want)
	}
}

func TestMonolithConditionsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &MonolithConditionRequest{
		Name:      "Laptop",
		Predicate: "machine_type == \"laptop\"",
	}

	mux.HandleFunc("/monolith/conditions/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MonolithConditionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, mcoCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithConditions.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("MonolithConditions.Create returned error: %v", err)
	}

	want := &MonolithCondition{
		ID:        4,
		Name:      "Laptop",
		Predicate: "machine_type == \"laptop\"",
		Created:   Timestamp{referenceTime},
		Updated:   Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithConditions.Create returned %+v, want %+v", got, want)
	}
}

func TestMonolithConditionsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &MonolithConditionRequest{
		Name:      "Laptop",
		Predicate: "machine_type == \"laptop\"",
	}

	mux.HandleFunc("/monolith/conditions/4/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MonolithConditionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, mcoUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithConditions.Update(ctx, 4, updateRequest)
	if err != nil {
		t.Errorf("MonolithConditions.Update returned error: %v", err)
	}

	want := &MonolithCondition{
		ID:        4,
		Name:      "Laptop",
		Predicate: "machine_type == \"laptop\"",
		Created:   Timestamp{referenceTime},
		Updated:   Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithConditions.Update returned %+v, want %+v", got, want)
	}
}

func TestMonolithConditionsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/conditions/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MonolithConditions.Delete(ctx, 4)
	if err != nil {
		t.Errorf("MonolithConditions.Delete returned error: %v", err)
	}
}
