package goztl

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var mlListJSONResponse = `
[
    {
        "id": 4,
	"name": "Yolo",
	"organization_name": "Fomo",
	"country_code": "DE",
	"library_uid": "01234578910",
	"mdm_info_id": "0dcc02e3-5802-47cd-9ee4-fcf977277ff0",
	"platform": "enterprisestore",
	"server_token_expiration_date": "2022-07-22T01:02:03.444444",
	"website_url": "https://business.apple.com",
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var mlGetJSONResponse = `
{
    "id": 4,
    "name": "Yolo",
    "organization_name": "Fomo",
    "country_code": "DE",
    "library_uid": "01234578910",
    "mdm_info_id": "0dcc02e3-5802-47cd-9ee4-fcf977277ff0",
    "platform": "enterprisestore",
    "server_token_expiration_date": "2022-07-22T01:02:03.444444",
    "website_url": "https://business.apple.com",
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestMDMLocationsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/locations/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mlListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMLocations.List(ctx, nil)
	if err != nil {
		t.Errorf("MDMLocations.List returned error: %v", err)
	}

	want := []MDMLocation{
		{
			ID:                        4,
			Name:                      "Yolo",
			OrganizationName:          "Fomo",
			CountryCode:               "DE",
			LibraryUID:                "01234578910",
			MDMInfoID:                 "0dcc02e3-5802-47cd-9ee4-fcf977277ff0",
			Platform:                  "enterprisestore",
			ServerTokenExpirationDate: Timestamp{referenceTime},
			WebsiteURL:                "https://business.apple.com",
			Created:                   Timestamp{referenceTime},
			Updated:                   Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMLocations.List returned %+v, want %+v", got, want)
	}
}

func TestMDMLocationsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/locations/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mlGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMLocations.GetByID(ctx, 4)
	if err != nil {
		t.Errorf("MDMLocations.GetByID returned error: %v", err)
	}

	want := &MDMLocation{
		ID:                        4,
		Name:                      "Yolo",
		OrganizationName:          "Fomo",
		CountryCode:               "DE",
		LibraryUID:                "01234578910",
		MDMInfoID:                 "0dcc02e3-5802-47cd-9ee4-fcf977277ff0",
		Platform:                  "enterprisestore",
		ServerTokenExpirationDate: Timestamp{referenceTime},
		WebsiteURL:                "https://business.apple.com",
		Created:                   Timestamp{referenceTime},
		Updated:                   Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMLocations.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMDMLocationsService_GetByMDMInfoID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/locations/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "mdm_info_id", "0dcc02e3-5802-47cd-9ee4-fcf977277ff0")
		fmt.Fprint(w, mlListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMLocations.GetByMDMInfoID(ctx, "0dcc02e3-5802-47cd-9ee4-fcf977277ff0")
	if err != nil {
		t.Errorf("MDMLocations.GetByMDMInfoID returned error: %v", err)
	}

	want := &MDMLocation{
		ID:                        4,
		Name:                      "Yolo",
		OrganizationName:          "Fomo",
		CountryCode:               "DE",
		LibraryUID:                "01234578910",
		MDMInfoID:                 "0dcc02e3-5802-47cd-9ee4-fcf977277ff0",
		Platform:                  "enterprisestore",
		ServerTokenExpirationDate: Timestamp{referenceTime},
		WebsiteURL:                "https://business.apple.com",
		Created:                   Timestamp{referenceTime},
		Updated:                   Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMLocations.GetByMDMInfoID returned %+v, want %+v", got, want)
	}
}

func TestMDMLocationsService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/locations/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "Yolo")
		fmt.Fprint(w, mlListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMLocations.GetByName(ctx, "Yolo")
	if err != nil {
		t.Errorf("MDMLocations.GetByName returned error: %v", err)
	}

	want := &MDMLocation{
		ID:                        4,
		Name:                      "Yolo",
		OrganizationName:          "Fomo",
		CountryCode:               "DE",
		LibraryUID:                "01234578910",
		MDMInfoID:                 "0dcc02e3-5802-47cd-9ee4-fcf977277ff0",
		Platform:                  "enterprisestore",
		ServerTokenExpirationDate: Timestamp{referenceTime},
		WebsiteURL:                "https://business.apple.com",
		Created:                   Timestamp{referenceTime},
		Updated:                   Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMLocations.GetByName returned %+v, want %+v", got, want)
	}
}
