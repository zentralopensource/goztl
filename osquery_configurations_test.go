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

var ocListJSONResponse = `
[
    {
        "id": 4,
        "name": "Default",
	"description": "Description",
	"inventory": true,
	"inventory_apps": true,
	"inventory_ec2": false,
	"inventory_interval": 600,
	"options": {"config_refresh": 120},
	"automatic_table_constructions": [],
	"file_categories": [],
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var ocGetJSONResponse = `
{
    "id": 4,
    "name": "Default",
    "description": "Description",
    "inventory": true,
    "inventory_apps": true,
    "inventory_ec2": false,
    "inventory_interval": 600,
    "options": {"config_refresh": 120},
    "automatic_table_constructions": [1],
    "file_categories": [1],
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var ocCreateJSONResponse = `
{
    "id": 4,
    "name": "Default",
    "description": "Description",
    "inventory": true,
    "inventory_apps": true,
    "inventory_ec2": false,
    "inventory_interval": 600,
    "options": {"config_refresh": 120},
    "automatic_table_constructions": [1, 2],
    "file_categories": [1, 2],
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var ocUpdateJSONResponse = `
{
    "id": 4,
    "name": "Default",
    "description": "Description",
    "inventory": true,
    "inventory_apps": true,
    "inventory_ec2": false,
    "inventory_interval": 600,
    "options": {"config_refresh": 120},
    "automatic_table_constructions": [1, 2, 3],
    "file_categories": [1, 2, 3],
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestOsqueryConfigurationsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/configurations/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, ocListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryConfigurations.List(ctx, nil)
	if err != nil {
		t.Errorf("OsqueryConfigurations.List returned error: %v", err)
	}

	want := []OsqueryConfiguration{
		{
			ID:                4,
			Name:              "Default",
			Description:       "Description",
			Inventory:         true,
			InventoryApps:     true,
			InventoryEC2:      false,
			InventoryInterval: 600,
			Options:           map[string]interface{}{"config_refresh": 120.0},
			ATCs:              []int{},
			FileCategories:    []int{},
			Created:           Timestamp{referenceTime},
			Updated:           Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryConfigurations.List returned %+v, want %+v", got, want)
	}
}

func TestOsqueryConfigurationsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/configurations/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, ocGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryConfigurations.GetByID(ctx, 1)
	if err != nil {
		t.Errorf("OsqueryConfigurations.GetByID returned error: %v", err)
	}

	want := &OsqueryConfiguration{
		ID:                4,
		Name:              "Default",
		Description:       "Description",
		Inventory:         true,
		InventoryApps:     true,
		InventoryEC2:      false,
		InventoryInterval: 600,
		Options:           map[string]interface{}{"config_refresh": 120.0},
		ATCs:              []int{1},
		FileCategories:    []int{1},
		Created:           Timestamp{referenceTime},
		Updated:           Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryConfigurations.GetByID returned %+v, want %+v", got, want)
	}
}

func TestOsqueryConfigurationsService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/configurations/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "Default")
		fmt.Fprint(w, ocListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryConfigurations.GetByName(ctx, "Default")
	if err != nil {
		t.Errorf("OsqueryConfigurations.GetByName returned error: %v", err)
	}

	want := &OsqueryConfiguration{
		ID:                4,
		Name:              "Default",
		Description:       "Description",
		Inventory:         true,
		InventoryApps:     true,
		InventoryEC2:      false,
		InventoryInterval: 600,
		Options:           map[string]interface{}{"config_refresh": 120.0},
		ATCs:              []int{},
		FileCategories:    []int{},
		Created:           Timestamp{referenceTime},
		Updated:           Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryConfigurations.GetByName returned %+v, want %+v", got, want)
	}
}

func TestOsqueryConfigurationsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &OsqueryConfigurationRequest{
		Name:              "Default",
		Description:       "Description",
		Inventory:         true,
		InventoryApps:     true,
		InventoryEC2:      false,
		InventoryInterval: 600,
		Options:           map[string]interface{}{"config_refresh": 120.0},
		ATCs:              []int{1, 2},
		FileCategories:    []int{1, 2},
	}

	mux.HandleFunc("/osquery/configurations/", func(w http.ResponseWriter, r *http.Request) {
		v := new(OsqueryConfigurationRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, ocCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryConfigurations.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("OsqueryConfigurations.Create returned error: %v", err)
	}

	want := &OsqueryConfiguration{
		ID:                4,
		Name:              "Default",
		Description:       "Description",
		Inventory:         true,
		InventoryApps:     true,
		InventoryEC2:      false,
		InventoryInterval: 600,
		Options:           map[string]interface{}{"config_refresh": 120.0},
		ATCs:              []int{1, 2},
		FileCategories:    []int{1, 2},
		Created:           Timestamp{referenceTime},
		Updated:           Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryConfigurations.Create returned %+v, want %+v", got, want)
	}
}

func TestOsqueryConfigurationsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &OsqueryConfigurationRequest{
		Name:              "Default",
		Description:       "Description",
		Inventory:         true,
		InventoryApps:     true,
		InventoryEC2:      false,
		InventoryInterval: 600,
		ATCs:              []int{1, 2, 3},
		FileCategories:    []int{1, 2, 3},
		Options:           map[string]interface{}{"config_refresh": 120.0},
	}

	mux.HandleFunc("/osquery/configurations/1/", func(w http.ResponseWriter, r *http.Request) {
		v := new(OsqueryConfigurationRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, ocUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryConfigurations.Update(ctx, 1, updateRequest)
	if err != nil {
		t.Errorf("OsqueryConfigurations.Update returned error: %v", err)
	}

	want := &OsqueryConfiguration{
		ID:                4,
		Name:              "Default",
		Description:       "Description",
		Inventory:         true,
		InventoryApps:     true,
		InventoryEC2:      false,
		InventoryInterval: 600,
		Options:           map[string]interface{}{"config_refresh": 120.0},
		ATCs:              []int{1, 2, 3},
		FileCategories:    []int{1, 2, 3},
		Created:           Timestamp{referenceTime},
		Updated:           Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryConfigurations.Update returned %+v, want %+v", got, want)
	}
}

func TestOsqueryConfigurationsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/configurations/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.OsqueryConfigurations.Delete(ctx, 1)
	if err != nil {
		t.Errorf("OsqueryConfigurations.Delete returned error: %v", err)
	}
}
