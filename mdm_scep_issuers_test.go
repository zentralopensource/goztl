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

var mscepListJSONResponse = `
[
    {
        "id": "eaabf092-caed-4b0e-a8d5-851205b2fa56",
	"provisioning_uid": null,
        "name": "Default",
	"description": "Description",
	"url": "https://www.example.com/scep/",
	"key_usage": 1,
	"key_size": 2048,
	"backend": "MICROSOFT_CA",
	"microsoft_ca_kwargs" : {
	    "url": "https://www.example.com/ndes/",
	    "username": "Yolo",
	    "password": "Fomo"
	},
	"okta_ca_kwargs": null,
	"static_challenge_kwargs": null,
	"version": 1,
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var mscepGetByNameJSONResponse = `
[
    {
        "id": "eaabf092-caed-4b0e-a8d5-851205b2fa56",
	"provisioning_uid": null,
        "name": "Default",
	"description": "Description",
	"url": "https://www.example.com/scep/",
	"key_usage": 1,
	"key_size": 2048,
	"backend": "STATIC_CHALLENGE",
	"microsoft_ca_kwargs" : null,
	"okta_ca_kwargs": null,
	"static_challenge_kwargs": {
	    "challenge": "YoloFomo"
	},
	"version": 1,
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var mscepCreateJSONResponse = `
{
    "id": "eaabf092-caed-4b0e-a8d5-851205b2fa56",
    "provisioning_uid": null,
    "name": "Default",
    "description": "Description",
    "url": "https://www.example.com/scep/",
    "key_size": 2048,
    "key_usage": 1,
    "backend": "IDENT",
    "ident_kwargs": {
	"url": "https://www.example.com/ident/",
	"bearer_token": "YoloFomo",
	"request_timeout": 123,
	"max_retries": 5
    },
    "okta_ca_kwargs": null,
    "static_challenge_kwargs": null,
    "version": 1,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var mscepUpdateJSONResponse = `
{
    "id": "eaabf092-caed-4b0e-a8d5-851205b2fa56",
    "provisioning_uid": null,
    "name": "Default",
    "description": "Description",
    "url": "https://www.example.com/scep/",
    "key_size": 2048,
    "key_usage": 1,
    "backend": "STATIC_CHALLENGE",
    "microsoft_ca_kwargs": null,
    "okta_ca_kwargs": null,
    "static_challenge_kwargs": {
        "challenge": "fomo"
    },
    "version": 2,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var mscepGetJSONResponse = `
{
    "id": "eaabf092-caed-4b0e-a8d5-851205b2fa56",
    "provisioning_uid": "YoLoFoMo",
    "name": "Default",
    "description": "Description",
    "url": "https://www.example.com/scep/",
    "key_usage": 1,
    "key_size": 2048,
    "backend": "STATIC_CHALLENGE",
    "static_challenge_kwargs": {
        "challenge": "fomo"
    },
    "version": 1,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var mscepGetJSONResponse2 = `
{
    "id": "62ed7c87-dc8b-4367-84d9-0491ece2196d",
    "provisioning_uid": "YoLoFoMo",
    "name": "Default",
    "description": "Description",
    "url": "https://www.example.com/scep/",
    "key_usage": 1,
    "key_size": 2048,
    "backend": "DIGICERT",
    "digicert_kwargs": {
        "api_base_url": "https://one.digicert.com/mpki/api/",
	"api_token": "secret",
	"profile_guid": "60a3ce98-b05f-4f1b-83b0-200d82723134",
	"business_unit_guid": "34f0d9a5-4603-4d07-baf3-2071f6e5b874",
	"seat_type": "DEVICE_SEAT",
	"seat_id_mapping": "common_name",
	"default_seat_email": "yolo@example.com"
    },
    "version": 1,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestMDMSCEPIssuersService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/scep_issuers/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mscepListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMSCEPIssuers.List(ctx, nil)
	if err != nil {
		t.Errorf("MDMSCEPIssuers.List returned error: %v", err)
	}

	want := []MDMSCEPIssuer{
		{
			ID:          "eaabf092-caed-4b0e-a8d5-851205b2fa56",
			Name:        "Default",
			Description: "Description",
			URL:         "https://www.example.com/scep/",
			KeyUsage:    1,
			KeySize:     2048,
			Backend:     String("MICROSOFT_CA"),
			MicrosoftCA: &MicrosoftCA{
				URL:      "https://www.example.com/ndes/",
				Username: "Yolo",
				Password: "Fomo",
			},
			Version: 1,
			Created: Timestamp{referenceTime},
			Updated: Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMSCEPIssuers.List returned %+v, want %+v", got, want)
	}
}

func TestMDMSCEPIssuersService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/scep_issuers/eaabf092-caed-4b0e-a8d5-851205b2fa56/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mscepGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMSCEPIssuers.GetByID(ctx, "eaabf092-caed-4b0e-a8d5-851205b2fa56")
	if err != nil {
		t.Errorf("MDMSCEPIssuers.GetByID returned error: %v", err)
	}

	want := &MDMSCEPIssuer{
		ID:              "eaabf092-caed-4b0e-a8d5-851205b2fa56",
		ProvisioningUID: String("YoLoFoMo"),
		Name:            "Default",
		Description:     "Description",
		URL:             "https://www.example.com/scep/",
		KeyUsage:        1,
		KeySize:         2048,
		Backend:         String("STATIC_CHALLENGE"),
		StaticChallenge: &StaticChallenge{
			Challenge: "fomo",
		},
		Version: 1,
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMSCEPIssuers.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMDMSCEPIssuersService_GetByID2(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/scep_issuers/62ed7c87-dc8b-4367-84d9-0491ece2196d/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mscepGetJSONResponse2)
	})

	ctx := context.Background()
	got, _, err := client.MDMSCEPIssuers.GetByID(ctx, "62ed7c87-dc8b-4367-84d9-0491ece2196d")
	if err != nil {
		t.Errorf("MDMSCEPIssuers.GetByID returned error: %v", err)
	}

	want := &MDMSCEPIssuer{
		ID:              "62ed7c87-dc8b-4367-84d9-0491ece2196d",
		ProvisioningUID: String("YoLoFoMo"),
		Name:            "Default",
		Description:     "Description",
		URL:             "https://www.example.com/scep/",
		KeyUsage:        1,
		KeySize:         2048,
		Backend:         String("DIGICERT"),
		Digicert: &Digicert{
			APIBaseURL:       "https://one.digicert.com/mpki/api/",
			APIToken:         "secret",
			ProfileGUID:      "60a3ce98-b05f-4f1b-83b0-200d82723134",
			BusinessUnitGUID: "34f0d9a5-4603-4d07-baf3-2071f6e5b874",
			SeatType:         "DEVICE_SEAT",
			SeatIDMapping:    "common_name",
			DefaultSeatEmail: "yolo@example.com",
		},
		Version: 1,
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMSCEPIssuers.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMDMSCEPIssuersService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/scep_issuers/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "Default")
		fmt.Fprint(w, mscepGetByNameJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMSCEPIssuers.GetByName(ctx, "Default")
	if err != nil {
		t.Errorf("MDMSCEPIssuers.GetByName returned error: %v", err)
	}

	want := &MDMSCEPIssuer{
		ID:          "eaabf092-caed-4b0e-a8d5-851205b2fa56",
		Name:        "Default",
		Description: "Description",
		URL:         "https://www.example.com/scep/",
		KeyUsage:    1,
		KeySize:     2048,
		Backend:     String("STATIC_CHALLENGE"),
		StaticChallenge: &StaticChallenge{
			Challenge: "YoloFomo",
		},
		Version: 1,
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMSCEPIssuers.GetByName returned %+v, want %+v", got, want)
	}
}

func TestMDMSCEPIssuersService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &MDMSCEPIssuerRequest{
		Name:        "Default",
		Description: "Description",
		URL:         "https://www.example.com/scep/",
		KeySize:     2048,
		KeyUsage:    1,
		Backend:     "IDENT",
		IDent: &IDent{
			URL:            "https://www.example.com/ident/",
			BearerToken:    "YoloFomo",
			RequestTimeout: 123,
			MaxRetries:     5,
		},
	}

	mux.HandleFunc("/mdm/scep_issuers/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMSCEPIssuerRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, mscepCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMSCEPIssuers.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("MDMSCEPIssuers.Create returned error: %v", err)
	}

	want := &MDMSCEPIssuer{
		ID:          "eaabf092-caed-4b0e-a8d5-851205b2fa56",
		Name:        "Default",
		Description: "Description",
		URL:         "https://www.example.com/scep/",
		KeySize:     2048,
		KeyUsage:    1,
		Backend:     String("IDENT"),
		IDent: &IDent{
			URL:            "https://www.example.com/ident/",
			BearerToken:    "YoloFomo",
			RequestTimeout: 123,
			MaxRetries:     5,
		},
		Version: 1,
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMSCEPIssuers.Create returned %+v, want %+v", got, want)
	}
}

func TestMDMSCEPIssuersService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &MDMSCEPIssuerRequest{
		Name:        "Default",
		Description: "Description",
		URL:         "https://www.example.com/scep/",
		KeySize:     2048,
		KeyUsage:    1,
		Backend:     "STATIC_CHALLENGE",
		StaticChallenge: &StaticChallenge{
			Challenge: "fomo",
		},
	}

	mux.HandleFunc("/mdm/scep_issuers/eaabf092-caed-4b0e-a8d5-851205b2fa56/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMSCEPIssuerRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, mscepUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMSCEPIssuers.Update(ctx, "eaabf092-caed-4b0e-a8d5-851205b2fa56", updateRequest)
	if err != nil {
		t.Errorf("MDMSCEPIssuers.Update returned error: %v", err)
	}

	want := &MDMSCEPIssuer{
		ID:          "eaabf092-caed-4b0e-a8d5-851205b2fa56",
		Name:        "Default",
		Description: "Description",
		URL:         "https://www.example.com/scep/",
		KeySize:     2048,
		KeyUsage:    1,
		Backend:     String("STATIC_CHALLENGE"),
		StaticChallenge: &StaticChallenge{
			Challenge: "fomo",
		},
		Version: 2,
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMSCEPIssuers.Update returned %+v, want %+v", got, want)
	}
}

func TestMDMSCEPIssuersService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/scep_issuers/eaabf092-caed-4b0e-a8d5-851205b2fa56/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MDMSCEPIssuers.Delete(ctx, "eaabf092-caed-4b0e-a8d5-851205b2fa56")
	if err != nil {
		t.Errorf("MDMSCEPIssuers.Delete returned error: %v", err)
	}
}
