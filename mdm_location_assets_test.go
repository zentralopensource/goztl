package goztl

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var mlaListJSONResponse = `
[
    {
        "id": 4,
	"location": 5,
	"asset": 6,
	"adam_id": "0123456789",
	"pricing_param": "STDQ",
	"assigned_count": 7,
	"available_count": 8,
	"retired_count": 9,
	"total_count": 24,
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

func TestMDMLocationAssetsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/location_assets/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mlaListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMLocationAssets.List(ctx, nil)
	if err != nil {
		t.Errorf("MDMLocationAssets.List returned error: %v", err)
	}

	want := []MDMLocationAsset{
		{
			ID:             4,
			LocationID:     5,
			AssetID:        6,
			AdamID:         "0123456789",
			PricingParam:   "STDQ",
			AssignedCount:  7,
			AvailableCount: 8,
			RetiredCount:   9,
			TotalCount:     24,
			Created:        Timestamp{referenceTime},
			Updated:        Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMLocationAssets.List returned %+v, want %+v", got, want)
	}
}

func TestMDMLocationAssetsService_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/location_assets/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "location_id", "5")
		testQueryArg(t, r, "adam_id", "0123456789")
		testQueryArg(t, r, "pricing_param", "STDQ")
		fmt.Fprint(w, mlaListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMLocationAssets.Get(ctx, 5, "0123456789", "STDQ")
	if err != nil {
		t.Errorf("MDMLocationAssets.Get returned error: %v", err)
	}

	want := &MDMLocationAsset{
		ID:             4,
		LocationID:     5,
		AssetID:        6,
		AdamID:         "0123456789",
		PricingParam:   "STDQ",
		AssignedCount:  7,
		AvailableCount: 8,
		RetiredCount:   9,
		TotalCount:     24,
		Created:        Timestamp{referenceTime},
		Updated:        Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMLocationAssets.Get returned %+v, want %+v", got, want)
	}
}
