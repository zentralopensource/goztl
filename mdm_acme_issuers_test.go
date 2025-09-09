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

var macmeListJSONResponse = `
[
    {
        "id": "eaabf092-caed-4b0e-a8d5-851205b2fa56",
	"provisioning_uid": null,
        "name": "Default",
	"description": "Description",
	"directory_url": "https://www.example.com/acme/",
	"key_type": "ECSECPrimeRandom",
	"key_size": 384,
	"usage_flags": 1,
	"extended_key_usage": ["1.3.6.1.5.5.7.3.2"],
	"hardware_bound": true,
	"attest": true,
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

var macmeGetByNameJSONResponse = `
[
    {
        "id": "eaabf092-caed-4b0e-a8d5-851205b2fa56",
	"provisioning_uid": null,
        "name": "Default",
	"description": "Description",
	"directory_url": "https://www.example.com/acme/",
	"key_type": "RSA",
	"key_size": 2048,
	"usage_flags": 1,
	"extended_key_usage": [],
	"hardware_bound": false,
	"attest": false,
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

var macmeCreateJSONResponse = `
{
    "id": "eaabf092-caed-4b0e-a8d5-851205b2fa56",
    "provisioning_uid": null,
    "name": "Default",
    "description": "Description",
    "directory_url": "https://www.example.com/acme/",
    "key_type": "ECSECPrimeRandom",
    "key_size": 256,
    "usage_flags": 1,
    "extended_key_usage": ["1.3.6.1.5.5.7.3.2"],
    "hardware_bound": true,
    "attest": true,
    "backend": "OKTA_CA",
    "microsoft_ca_kwargs": null,
    "okta_ca_kwargs": {
	"url": "https://www.example.com/ndes/",
	"username": "Yolo",
	"password": "Fomo"
    },
    "static_challenge_kwargs": null,
    "version": 1,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var macmeUpdateJSONResponse = `
{
    "id": "eaabf092-caed-4b0e-a8d5-851205b2fa56",
    "provisioning_uid": null,
    "name": "Default",
    "description": "Description",
    "directory_url": "https://www.example.com/acme/",
    "key_type": "ECSECPrimeRandom",
    "key_size": 256,
    "usage_flags": 1,
    "extended_key_usage": ["1.3.6.1.5.5.7.3.2"],
    "hardware_bound": true,
    "attest": true,
    "backend": "IDENT",
    "ident_kwargs": {
        "url": "https://www.example.com/ident/",
	"bearer_token": "YoloFomo",
	"request_timeout": 123,
	"max_retries": 5
    },
    "version": 2,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var macmeGetJSONResponse = `
{
    "id": "eaabf092-caed-4b0e-a8d5-851205b2fa56",
    "provisioning_uid": "YoLoFoMo",
    "name": "Default",
    "description": "Description",
    "directory_url": "https://www.example.com/acme/",
    "key_type": "ECSECPrimeRandom",
    "key_size": 256,
    "usage_flags": 1,
    "extended_key_usage": ["1.3.6.1.5.5.7.3.2"],
    "hardware_bound": true,
    "attest": true,
    "backend": "STATIC_CHALLENGE",
    "static_challenge_kwargs": {
        "challenge": "fomo"
    },
    "version": 1,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestMDMACMEIssuersService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/acme_issuers/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, macmeListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMACMEIssuers.List(ctx, nil)
	if err != nil {
		t.Errorf("MDMACMEIssuers.List returned error: %v", err)
	}

	want := []MDMACMEIssuer{
		{
			ID:               "eaabf092-caed-4b0e-a8d5-851205b2fa56",
			Name:             "Default",
			Description:      "Description",
			DirectoryURL:     "https://www.example.com/acme/",
			KeyType:          "ECSECPrimeRandom",
			KeySize:          384,
			UsageFlags:       1,
			ExtendedKeyUsage: []string{"1.3.6.1.5.5.7.3.2"},
			HardwareBound:    true,
			Attest:           true,
			Backend:          String("MICROSOFT_CA"),
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
		t.Errorf("MDMACMEIssuers.List returned %+v, want %+v", got, want)
	}
}

func TestMDMACMEIssuersService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/acme_issuers/eaabf092-caed-4b0e-a8d5-851205b2fa56/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, macmeGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMACMEIssuers.GetByID(ctx, "eaabf092-caed-4b0e-a8d5-851205b2fa56")
	if err != nil {
		t.Errorf("MDMACMEIssuers.GetByID returned error: %v", err)
	}

	want := &MDMACMEIssuer{
		ID:               "eaabf092-caed-4b0e-a8d5-851205b2fa56",
		ProvisioningUID:  String("YoLoFoMo"),
		Name:             "Default",
		Description:      "Description",
		DirectoryURL:     "https://www.example.com/acme/",
		KeyType:          "ECSECPrimeRandom",
		KeySize:          256,
		UsageFlags:       1,
		ExtendedKeyUsage: []string{"1.3.6.1.5.5.7.3.2"},
		HardwareBound:    true,
		Attest:           true,
		Backend:          String("STATIC_CHALLENGE"),
		StaticChallenge: &StaticChallenge{
			Challenge: "fomo",
		},
		Version: 1,
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMACMEIssuers.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMDMACMEIssuersService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/acme_issuers/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "Default")
		fmt.Fprint(w, macmeGetByNameJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMACMEIssuers.GetByName(ctx, "Default")
	if err != nil {
		t.Errorf("MDMACMEIssuers.GetByName returned error: %v", err)
	}

	want := &MDMACMEIssuer{
		ID:               "eaabf092-caed-4b0e-a8d5-851205b2fa56",
		Name:             "Default",
		Description:      "Description",
		DirectoryURL:     "https://www.example.com/acme/",
		KeyType:          "RSA",
		KeySize:          2048,
		UsageFlags:       1,
		ExtendedKeyUsage: []string{},
		HardwareBound:    false,
		Attest:           false,
		Backend:          String("STATIC_CHALLENGE"),
		StaticChallenge: &StaticChallenge{
			Challenge: "YoloFomo",
		},
		Version: 1,
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMACMEIssuers.GetByName returned %+v, want %+v", got, want)
	}
}

func TestMDMACMEIssuersService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &MDMACMEIssuerRequest{
		Name:             "Default",
		Description:      "Description",
		DirectoryURL:     "https://www.example.com/acme/",
		KeyType:          "ECSECPrimeRandom",
		KeySize:          256,
		UsageFlags:       1,
		ExtendedKeyUsage: []string{"1.3.6.1.5.5.7.3.2"},
		HardwareBound:    true,
		Attest:           true,
		Backend:          "OKTA_CA",
		OktaCA: &MicrosoftCA{
			URL:      "https://www.example.com/ndes/",
			Username: "Yolo",
			Password: "Fomo",
		},
	}

	mux.HandleFunc("/mdm/acme_issuers/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMACMEIssuerRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, macmeCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMACMEIssuers.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("MDMACMEIssuers.Create returned error: %v", err)
	}

	want := &MDMACMEIssuer{
		ID:               "eaabf092-caed-4b0e-a8d5-851205b2fa56",
		Name:             "Default",
		Description:      "Description",
		DirectoryURL:     "https://www.example.com/acme/",
		KeyType:          "ECSECPrimeRandom",
		KeySize:          256,
		UsageFlags:       1,
		ExtendedKeyUsage: []string{"1.3.6.1.5.5.7.3.2"},
		HardwareBound:    true,
		Attest:           true,
		Backend:          String("OKTA_CA"),
		OktaCA: &MicrosoftCA{
			URL:      "https://www.example.com/ndes/",
			Username: "Yolo",
			Password: "Fomo",
		},
		Version: 1,
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMACMEIssuers.Create returned %+v, want %+v", got, want)
	}
}

func TestMDMACMEIssuersService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &MDMACMEIssuerRequest{
		Name:             "Default",
		Description:      "Description",
		DirectoryURL:     "https://www.example.com/acme/",
		KeyType:          "ECSECPrimeRandom",
		KeySize:          256,
		UsageFlags:       1,
		ExtendedKeyUsage: []string{"1.3.6.1.5.5.7.3.2"},
		HardwareBound:    true,
		Attest:           true,
		Backend:          "IDENT",
		IDent: &IDent{
			URL:            "https://www.example.com/ident/",
			BearerToken:    "YoloFomo",
			RequestTimeout: 123,
			MaxRetries:     5,
		},
	}

	mux.HandleFunc("/mdm/acme_issuers/eaabf092-caed-4b0e-a8d5-851205b2fa56/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMACMEIssuerRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, macmeUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMACMEIssuers.Update(ctx, "eaabf092-caed-4b0e-a8d5-851205b2fa56", updateRequest)
	if err != nil {
		t.Errorf("MDMACMEIssuers.Update returned error: %v", err)
	}

	want := &MDMACMEIssuer{
		ID:               "eaabf092-caed-4b0e-a8d5-851205b2fa56",
		Name:             "Default",
		Description:      "Description",
		DirectoryURL:     "https://www.example.com/acme/",
		KeyType:          "ECSECPrimeRandom",
		KeySize:          256,
		UsageFlags:       1,
		ExtendedKeyUsage: []string{"1.3.6.1.5.5.7.3.2"},
		HardwareBound:    true,
		Attest:           true,
		Backend:          String("IDENT"),
		IDent: &IDent{
			URL:            "https://www.example.com/ident/",
			BearerToken:    "YoloFomo",
			RequestTimeout: 123,
			MaxRetries:     5,
		},
		Version: 2,
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMACMEIssuers.Update returned %+v, want %+v", got, want)
	}
}

func TestMDMACMEIssuersService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/acme_issuers/eaabf092-caed-4b0e-a8d5-851205b2fa56/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MDMACMEIssuers.Delete(ctx, "eaabf092-caed-4b0e-a8d5-851205b2fa56")
	if err != nil {
		t.Errorf("MDMACMEIssuers.Delete returned error: %v", err)
	}
}
