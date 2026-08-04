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

var tcListJSONResponse = `
{
    "count": 1,
    "next": null,
    "previous": null,
    "results": [
    {
        "id": "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
        "name": "Default",
        "description": "the default configuration",
        "collect_inventory": true,
        "inventory_interval": 86400,
        "default_check_interval": 86400,
        "config_refresh_interval": 600,
        "results_batch_size": 100,
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
    ]
}
`

var tcGetJSONResponse = `
{
    "id": "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
    "name": "Default",
    "description": "the default configuration",
    "collect_inventory": true,
    "inventory_interval": 86400,
    "default_check_interval": 86400,
    "config_refresh_interval": 600,
    "results_batch_size": 100,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var tcCreateJSONResponse = `
{
    "id": "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
    "name": "Default",
    "description": "the default configuration",
    "collect_inventory": false,
    "inventory_interval": 3600,
    "default_check_interval": 7200,
    "config_refresh_interval": 300,
    "results_batch_size": 50,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var tcUpdateJSONResponse = `
{
    "id": "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
    "name": "Default",
    "description": "the default configuration",
    "collect_inventory": false,
    "inventory_interval": 3600,
    "default_check_interval": 7200,
    "config_refresh_interval": 300,
    "results_batch_size": 50,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestTurboConfigurationsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/turbo/configurations/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, tcListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.TurboConfigurations.List(ctx, nil)
	if err != nil {
		t.Errorf("TurboConfigurations.List returned error: %v", err)
	}

	want := []TurboConfiguration{
		{
			ID:                    "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
			Name:                  "Default",
			Description:           "the default configuration",
			CollectInventory:      true,
			InventoryInterval:     86400,
			DefaultCheckInterval:  86400,
			ConfigRefreshInterval: 600,
			ResultsBatchSize:      100,
			Created:               Timestamp{referenceTime},
			Updated:               Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("TurboConfigurations.List returned %+v, want %+v", got, want)
	}
}

func TestTurboConfigurationsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/turbo/configurations/5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, tcGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.TurboConfigurations.GetByID(ctx, "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b")
	if err != nil {
		t.Errorf("TurboConfigurations.GetByID returned error: %v", err)
	}

	want := &TurboConfiguration{
		ID:                    "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
		Name:                  "Default",
		Description:           "the default configuration",
		CollectInventory:      true,
		InventoryInterval:     86400,
		DefaultCheckInterval:  86400,
		ConfigRefreshInterval: 600,
		ResultsBatchSize:      100,
		Created:               Timestamp{referenceTime},
		Updated:               Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("TurboConfigurations.GetByID returned %+v, want %+v", got, want)
	}
}

func TestTurboConfigurationsService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/turbo/configurations/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "Default")
		fmt.Fprint(w, tcListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.TurboConfigurations.GetByName(ctx, "Default")
	if err != nil {
		t.Errorf("TurboConfigurations.GetByName returned error: %v", err)
	}

	want := &TurboConfiguration{
		ID:                    "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
		Name:                  "Default",
		Description:           "the default configuration",
		CollectInventory:      true,
		InventoryInterval:     86400,
		DefaultCheckInterval:  86400,
		ConfigRefreshInterval: 600,
		ResultsBatchSize:      100,
		Created:               Timestamp{referenceTime},
		Updated:               Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("TurboConfigurations.GetByName returned %+v, want %+v", got, want)
	}
}

func TestTurboConfigurationsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &TurboConfigurationRequest{
		Name:                  "Default",
		Description:           "the default configuration",
		CollectInventory:      false,
		InventoryInterval:     3600,
		DefaultCheckInterval:  7200,
		ConfigRefreshInterval: 300,
		ResultsBatchSize:      50,
	}

	mux.HandleFunc("/turbo/configurations/", func(w http.ResponseWriter, r *http.Request) {
		v := new(TurboConfigurationRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, tcCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.TurboConfigurations.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("TurboConfigurations.Create returned error: %v", err)
	}

	want := &TurboConfiguration{
		ID:                    "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
		Name:                  "Default",
		Description:           "the default configuration",
		CollectInventory:      false,
		InventoryInterval:     3600,
		DefaultCheckInterval:  7200,
		ConfigRefreshInterval: 300,
		ResultsBatchSize:      50,
		Created:               Timestamp{referenceTime},
		Updated:               Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("TurboConfigurations.Create returned %+v, want %+v", got, want)
	}
}

func TestTurboConfigurationsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &TurboConfigurationRequest{
		Name:                  "Default",
		Description:           "the default configuration",
		CollectInventory:      false,
		InventoryInterval:     3600,
		DefaultCheckInterval:  7200,
		ConfigRefreshInterval: 300,
		ResultsBatchSize:      50,
	}

	mux.HandleFunc("/turbo/configurations/5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b/", func(w http.ResponseWriter, r *http.Request) {
		v := new(TurboConfigurationRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, tcUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.TurboConfigurations.Update(ctx, "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b", updateRequest)
	if err != nil {
		t.Errorf("TurboConfigurations.Update returned error: %v", err)
	}

	want := &TurboConfiguration{
		ID:                    "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
		Name:                  "Default",
		Description:           "the default configuration",
		CollectInventory:      false,
		InventoryInterval:     3600,
		DefaultCheckInterval:  7200,
		ConfigRefreshInterval: 300,
		ResultsBatchSize:      50,
		Created:               Timestamp{referenceTime},
		Updated:               Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("TurboConfigurations.Update returned %+v, want %+v", got, want)
	}
}

func TestTurboConfigurationsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/turbo/configurations/5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.TurboConfigurations.Delete(ctx, "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b")
	if err != nil {
		t.Errorf("TurboConfigurations.Delete returned error: %v", err)
	}
}
