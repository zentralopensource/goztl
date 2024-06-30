package goztl

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var mscepListJSONResponse = `
[
    {
        "id": 4,
	"provisioning_uid": null,
        "name": "Default",
	"url": "https://www.example.com/scep/",
	"key_usage": 1,
	"key_is_extractable": false,
	"key_size": 2048,
	"allow_all_apps_access": false,
	"challenge_type": "MICROSOFT_CA",
	"microsoft_ca_challenge_kwargs" : {
	    "url": "https://www.example.com/ndes/",
	    "username": "Yolo",
	    "password": "Fomo"
	},
	"okta_ca_challenge_kwargs": null,
	"static_challenge_kwargs": null,
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var mscepGetByNameJSONResponse = `
[
    {
        "id": 4,
	"provisioning_uid": null,
        "name": "Default",
	"url": "https://www.example.com/scep/",
	"key_usage": 1,
	"key_is_extractable": false,
	"key_size": 2048,
	"allow_all_apps_access": false,
	"challenge_type": "STATIC",
	"microsoft_ca_challenge_kwargs" : null,
	"okta_ca_challenge_kwargs": null,
	"static_challenge_kwargs": {
	    "challenge": "YoloFomo"
	},
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var mscepGetJSONResponse = `
{
    "id": 4,
    "provisioning_uid": "YoLoFoMo",
    "name": "Default",
    "url": "https://www.example.com/scep/",
    "key_usage": 1,
    "key_is_extractable": true,
    "key_size": 2048,
    "allow_all_apps_access": true,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestMDMSCEPConfigsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/scep_configs/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mscepListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMSCEPConfigs.List(ctx, nil)
	if err != nil {
		t.Errorf("MDMSCEPConfigs.List returned error: %v", err)
	}

	want := []MDMSCEPConfig{
		{
			ID:                 4,
			Name:               "Default",
			URL:                "https://www.example.com/scep/",
			KeyUsage:           1,
			KeyIsExtractable:   false,
			KeySize:            2048,
			AllowAllAppsAccess: false,
			ChallengeType:      String("MICROSOFT_CA"),
			MicrosoftCAChallenge: &MicrosoftCASCEPChallenge{
				URL:      "https://www.example.com/ndes/",
				Username: "Yolo",
				Password: "Fomo",
			},
			Created: Timestamp{referenceTime},
			Updated: Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMSCEPConfigs.List returned %+v, want %+v", got, want)
	}
}

func TestMDMSCEPConfigsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/scep_configs/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mscepGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMSCEPConfigs.GetByID(ctx, 4)
	if err != nil {
		t.Errorf("MDMSCEPConfigs.GetByID returned error: %v", err)
	}

	want := &MDMSCEPConfig{
		ID:                 4,
		ProvisioningUID:    String("YoLoFoMo"),
		Name:               "Default",
		URL:                "https://www.example.com/scep/",
		KeyUsage:           1,
		KeyIsExtractable:   true,
		KeySize:            2048,
		AllowAllAppsAccess: true,
		Created:            Timestamp{referenceTime},
		Updated:            Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMSCEPConfigs.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMDMSCEPConfigsService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/scep_configs/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "Default")
		fmt.Fprint(w, mscepGetByNameJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMSCEPConfigs.GetByName(ctx, "Default")
	if err != nil {
		t.Errorf("MDMSCEPConfigs.GetByName returned error: %v", err)
	}

	want := &MDMSCEPConfig{
		ID:                 4,
		Name:               "Default",
		URL:                "https://www.example.com/scep/",
		KeyUsage:           1,
		KeyIsExtractable:   false,
		KeySize:            2048,
		AllowAllAppsAccess: false,
		ChallengeType:      String("STATIC"),
		StaticChallenge: &StaticSCEPChallenge{
			Challenge: "YoloFomo",
		},
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMSCEPConfigs.GetByName returned %+v, want %+v", got, want)
	}
}
