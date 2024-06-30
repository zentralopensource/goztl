package goztl

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var mpcListJSONResponse = `
[
    {
        "id": 4,
	"provisioning_uid": null,
        "name": "Default",
	"topic": null,
	"not_before": null,
	"not_after": null,
	"certificate": null,
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var mpcGetJSONResponse = `
{
    "id": 4,
    "provisioning_uid": "YoLoFoMo",
    "name": "Default",
    "topic": "un-deux-trois",
    "not_before": "2022-07-22T01:02:03.444444",
    "not_after": "2022-07-22T01:02:03.444444",
    "certificate": "CERT_PEM",
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestMDMPushCertificatesService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/push_certificates/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mpcListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMPushCertificates.List(ctx, nil)
	if err != nil {
		t.Errorf("MDMPushCertificates.List returned error: %v", err)
	}

	want := []MDMPushCertificate{
		{
			ID:      4,
			Name:    "Default",
			Created: Timestamp{referenceTime},
			Updated: Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMPushCertificates.List returned %+v, want %+v", got, want)
	}
}

func TestMDMPushCertificatesService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/push_certificates/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mpcGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMPushCertificates.GetByID(ctx, 4)
	if err != nil {
		t.Errorf("MDMPushCertificates.GetByID returned error: %v", err)
	}

	want := &MDMPushCertificate{
		ID:              4,
		ProvisioningUID: String("YoLoFoMo"),
		Name:            "Default",
		Topic:           String("un-deux-trois"),
		Certificate:     String("CERT_PEM"),
		NotBefore:       &Timestamp{referenceTime},
		NotAfter:        &Timestamp{referenceTime},
		Created:         Timestamp{referenceTime},
		Updated:         Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMPushCertificates.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMDMPushCertificatesService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/push_certificates/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "Default")
		fmt.Fprint(w, mpcListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMPushCertificates.GetByName(ctx, "Default")
	if err != nil {
		t.Errorf("MDMPushCertificates.GetByName returned error: %v", err)
	}

	want := &MDMPushCertificate{
		ID:      4,
		Name:    "Default",
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMPushCertificates.GetByName returned %+v, want %+v", got, want)
	}
}
