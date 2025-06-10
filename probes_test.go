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

var probeListJSONResponse = `
[
  {
    "id": 17,
    "name": "Default",
    "slug": "default",
    "description": "Description",
    "inventory_filters": [{"tag_ids": [42]}],
    "active": true,
    "actions": ["3dab941d-8e44-46ba-848f-98dfb9797664"],
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
  }
]
`

var probeGetJSONResponse = `
{
  "id": 17,
  "name": "Default",
  "slug": "default",
  "description": "Description",
  "inventory_filters": [],
  "metadata_filters": [],
  "payload_filters": [[{"attribute": "decision", "operator": "IN", "values": ["BLOCK_YOLO", "BLOCK_FOMO"]}]],
  "incident_severity": 100,
  "active": false,
  "actions": [],
  "created_at": "2022-07-22T01:02:03.444444",
  "updated_at": "2022-07-22T01:02:03.444444"
}
`

var probeCreateJSONResponse = `
{
  "id": 17,
  "name": "Default",
  "slug": "default",
  "description": "Description",
  "inventory_filters": [],
  "metadata_filters": [{"event_tags": ["un"]}],
  "payload_filters": [[{"attribute": "decision", "operator": "IN", "values": ["BLOCK_YOLO", "BLOCK_FOMO"]}]],
  "active": true,
  "actions": ["3dab941d-8e44-46ba-848f-98dfb9797664"],
  "created_at": "2022-07-22T01:02:03.444444",
  "updated_at": "2022-07-22T01:02:03.444444"
}
`

var probeUpdateJSONResponse = `
{
  "id": 17,
  "name": "Default",
  "slug": "default",
  "description": "Description",
  "inventory_filters": [{"platforms": ["macOS"]}],
  "metadata_filters": [{"event_types": ["zentral_login"]}],
  "payload_filters": [],
  "active": false,
  "actions": ["3dab941d-8e44-46ba-848f-98dfb9797664"],
  "created_at": "2022-07-22T01:02:03.444444",
  "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestProbesService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/probes/probes/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, probeListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.Probes.List(ctx, nil)
	if err != nil {
		t.Errorf("Probes.List returned error: %v", err)
	}

	want := []Probe{
		{
			ID:               17,
			Name:             "Default",
			Slug:             "default",
			Description:      "Description",
			InventoryFilters: []InventoryFilter{{TagIDs: []int{42}}},
			Active:           true,
			ActionIDs:        []string{"3dab941d-8e44-46ba-848f-98dfb9797664"},
			Created:          Timestamp{referenceTime},
			Updated:          Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("Probes.List returned %+v, want %+v", got, want)
	}
}

func TestProbesService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/probes/probes/17/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, probeGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.Probes.GetByID(ctx, 17)
	if err != nil {
		t.Errorf("Probes.GetByID returned error: %v", err)
	}

	want := &Probe{
		ID:               17,
		Name:             "Default",
		Slug:             "default",
		Description:      "Description",
		InventoryFilters: []InventoryFilter{},
		MetadataFilters:  []MetadataFilter{},
		PayloadFilters:   [][]PayloadFilterItem{{{Attribute: "decision", Operator: "IN", Values: []string{"BLOCK_YOLO", "BLOCK_FOMO"}}}},
		IncidentSeverity: Int(100),
		Active:           false,
		ActionIDs:        []string{},
		Created:          Timestamp{referenceTime},
		Updated:          Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("Probes.GetByID returned %+v, want %+v", got, want)
	}
}

func TestProbesService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/probes/probes/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "Default")
		fmt.Fprint(w, probeListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.Probes.GetByName(ctx, "Default")
	if err != nil {
		t.Errorf("Probes.GetByName returned error: %v", err)
	}

	want := &Probe{
		ID:               17,
		Name:             "Default",
		Slug:             "default",
		Description:      "Description",
		InventoryFilters: []InventoryFilter{{TagIDs: []int{42}}},
		Active:           true,
		ActionIDs:        []string{"3dab941d-8e44-46ba-848f-98dfb9797664"},
		Created:          Timestamp{referenceTime},
		Updated:          Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("Probes.GetByName returned %+v, want %+v", got, want)
	}
}

func TestProbesService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &ProbeRequest{
		Name:             "Default",
		Description:      "Description",
		InventoryFilters: []InventoryFilter{},
		MetadataFilters:  []MetadataFilter{{EventTags: []string{"un"}}},
		PayloadFilters:   [][]PayloadFilterItem{{{Attribute: "decision", Operator: "IN", Values: []string{"BLOCK_YOLO", "BLOCK_FOMO"}}}},
		Active:           true,
		ActionIDs:        []string{"3dab941d-8e44-46ba-848f-98dfb9797664"},
	}

	mux.HandleFunc("/probes/probes/", func(w http.ResponseWriter, r *http.Request) {
		v := new(ProbeRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, probeCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.Probes.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("Probes.Create returned error: %v", err)
	}

	want := &Probe{
		ID:               17,
		Name:             "Default",
		Slug:             "default",
		Description:      "Description",
		InventoryFilters: []InventoryFilter{},
		MetadataFilters:  []MetadataFilter{{EventTags: []string{"un"}}},
		PayloadFilters:   [][]PayloadFilterItem{{{Attribute: "decision", Operator: "IN", Values: []string{"BLOCK_YOLO", "BLOCK_FOMO"}}}},
		Active:           true,
		ActionIDs:        []string{"3dab941d-8e44-46ba-848f-98dfb9797664"},
		Created:          Timestamp{referenceTime},
		Updated:          Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("Probes.Create returned %+v, want %+v", got, want)
	}
}

func TestProbesService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &ProbeRequest{
		Name:             "Default",
		Description:      "Description",
		InventoryFilters: []InventoryFilter{{Platforms: []string{"macOS"}}},
		MetadataFilters:  []MetadataFilter{{EventTypes: []string{"zentral_login"}}},
		PayloadFilters:   [][]PayloadFilterItem{},
		Active:           false,
		ActionIDs:        []string{"3dab941d-8e44-46ba-848f-98dfb9797664"},
	}

	mux.HandleFunc("/probes/probes/17/", func(w http.ResponseWriter, r *http.Request) {
		v := new(ProbeRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, probeUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.Probes.Update(ctx, 17, updateRequest)
	if err != nil {
		t.Errorf("Probes.Update returned error: %v", err)
	}

	want := &Probe{
		ID:               17,
		Name:             "Default",
		Slug:             "default",
		Description:      "Description",
		InventoryFilters: []InventoryFilter{{Platforms: []string{"macOS"}}},
		MetadataFilters:  []MetadataFilter{{EventTypes: []string{"zentral_login"}}},
		PayloadFilters:   [][]PayloadFilterItem{},
		Active:           false,
		ActionIDs:        []string{"3dab941d-8e44-46ba-848f-98dfb9797664"},
		Created:          Timestamp{referenceTime},
		Updated:          Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("Probes.Update returned %+v, want %+v", got, want)
	}
}

func TestProbesService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/probes/probes/17/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.Probes.Delete(ctx, 17)
	if err != nil {
		t.Errorf("Probes.Delete returned error: %v", err)
	}
}
