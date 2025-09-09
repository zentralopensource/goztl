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

var moeListJSONResponme = `
[
    {
        "id": 1,
	"name": "Yolo",
	"display_name": "Fomo",
	"blueprint": 2,
	"push_certificate": 3,
	"realm": "2217e326-5c12-406f-8c31-cc95fe9fea1b",
	"acme_issuer": null,
	"scep_issuer": "0a0281b1-6fc0-462b-9128-67d2c87a0f45",
	"enrollment_secret": {
	    "id": 6,
	    "secret": "SECRET",
	    "meta_business_unit": 7,
	    "tags": [8, 9],
	    "serial_numbers": ["dix", "onze"],
	    "udids": [],
	    "quota": null,
	    "request_count": 12
	},
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var moeGetJSONResponme = `
{
    "id": 1,
    "name": "Yolo",
    "display_name": "Fomo",
    "blueprint": null,
    "push_certificate": 3,
    "realm": null,
    "acme_issuer": "e55e9dca-1f90-47bb-851c-c28fbf9aa55a",
    "scep_issuer": "0a0281b1-6fc0-462b-9128-67d2c87a0f45",
    "enrollment_secret": {
	"id": 6,
	"secret": "SECRET",
	"meta_business_unit": 7,
	"tags": [8, 9],
	"serial_numbers": ["dix", "onze"],
	"udids": ["AF92DAAB-EC8A-42EB-A11A-60B0BD94CCC1"],
	"quota": 12,
	"request_count": 13
    },
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var moeCreateJSONResponme = `
{
    "id": 1,
    "name": "Yolo",
    "display_name": "Fomo",
    "blueprint": 2,
    "push_certificate": 3,
    "realm": "2217e326-5c12-406f-8c31-cc95fe9fea1b",
    "acme_issuer": null,
    "scep_issuer": "0a0281b1-6fc0-462b-9128-67d2c87a0f45",
    "enrollment_secret": {
	"id": 6,
	"secret": "SECRET",
	"meta_business_unit": 7,
	"tags": [8, 9],
	"serial_numbers": ["dix", "onze"],
	"udids": ["AF92DAAB-EC8A-42EB-A11A-60B0BD94CCC1"],
	"quota": 12,
	"request_count": 13
    },
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var moeUpdateJSONResponme = `
{
    "id": 1,
    "name": "Yolo",
    "display_name": "Fomo",
    "blueprint": 2,
    "push_certificate": 3,
    "realm": "2217e326-5c12-406f-8c31-cc95fe9fea1b",
    "acme_issuer": "e55e9dca-1f90-47bb-851c-c28fbf9aa55a",
    "scep_issuer": "0a0281b1-6fc0-462b-9128-67d2c87a0f45",
    "enrollment_secret": {
	"id": 6,
	"secret": "SECRET",
	"meta_business_unit": 7,
	"tags": [8, 9],
	"serial_numbers": ["dix", "onze"],
	"udids": ["AF92DAAB-EC8A-42EB-A11A-60B0BD94CCC1"],
	"quota": 12,
	"request_count": 13
    },
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestMDMOTAEnrollmentsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/ota_enrollments/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, moeListJSONResponme)
	})

	ctx := context.Background()
	got, _, err := client.MDMOTAEnrollments.List(ctx, nil)
	if err != nil {
		t.Errorf("MDMOTAEnrollments.List returned error: %v", err)
	}

	want := []MDMOTAEnrollment{
		{
			ID:                1,
			Name:              "Yolo",
			DisplayName:       "Fomo",
			BlueprintID:       Int(2),
			PushCertificateID: 3,
			RealmUUID:         String("2217e326-5c12-406f-8c31-cc95fe9fea1b"),
			SCEPIssuerUUID:    "0a0281b1-6fc0-462b-9128-67d2c87a0f45",
			Secret: EnrollmentSecret{
				ID:                 6,
				Secret:             "SECRET",
				MetaBusinessUnitID: 7,
				TagIDs:             []int{8, 9},
				SerialNumbers:      []string{"dix", "onze"},
				UDIDs:              []string{},
				RequestCount:       12,
			},
			Created: Timestamp{referenceTime},
			Updated: Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMOTAEnrollments.List returned %+v, want %+v", got, want)
	}
}

func TestMDMOTAEnrollmentsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/ota_enrollments/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, moeGetJSONResponme)
	})

	ctx := context.Background()
	got, _, err := client.MDMOTAEnrollments.GetByID(ctx, 1)
	if err != nil {
		t.Errorf("MDMOTAEnrollments.GetByID returned error: %v", err)
	}

	want := &MDMOTAEnrollment{
		ID:                1,
		Name:              "Yolo",
		DisplayName:       "Fomo",
		PushCertificateID: 3,
		ACMEIssuerUUID:    String("e55e9dca-1f90-47bb-851c-c28fbf9aa55a"),
		SCEPIssuerUUID:    "0a0281b1-6fc0-462b-9128-67d2c87a0f45",
		Secret: EnrollmentSecret{
			ID:                 6,
			Secret:             "SECRET",
			MetaBusinessUnitID: 7,
			TagIDs:             []int{8, 9},
			SerialNumbers:      []string{"dix", "onze"},
			UDIDs:              []string{"AF92DAAB-EC8A-42EB-A11A-60B0BD94CCC1"},
			Quota:              Int(12),
			RequestCount:       13,
		},
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMOTAEnrollments.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMDMOTAEnrollmentsService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/ota_enrollments/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "Yolo")
		fmt.Fprint(w, moeListJSONResponme)
	})

	ctx := context.Background()
	got, _, err := client.MDMOTAEnrollments.GetByName(ctx, "Yolo")
	if err != nil {
		t.Errorf("MDMOTAEnrollments.GetByName returned error: %v", err)
	}

	want := &MDMOTAEnrollment{
		ID:                1,
		Name:              "Yolo",
		DisplayName:       "Fomo",
		BlueprintID:       Int(2),
		PushCertificateID: 3,
		RealmUUID:         String("2217e326-5c12-406f-8c31-cc95fe9fea1b"),
		SCEPIssuerUUID:    "0a0281b1-6fc0-462b-9128-67d2c87a0f45",
		Secret: EnrollmentSecret{
			ID:                 6,
			Secret:             "SECRET",
			MetaBusinessUnitID: 7,
			TagIDs:             []int{8, 9},
			SerialNumbers:      []string{"dix", "onze"},
			UDIDs:              []string{},
			RequestCount:       12,
		},
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMOTAEnrollments.List returned %+v, want %+v", got, want)
	}
}

func TestMDMOTAEnrollmentsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &MDMOTAEnrollmentRequest{
		Name:              "Yolo",
		DisplayName:       String("Fomo"),
		BlueprintID:       Int(2),
		PushCertificateID: 3,
		RealmUUID:         String("2217e326-5c12-406f-8c31-cc95fe9fea1b"),
		SCEPIssuerUUID:    "0a0281b1-6fc0-462b-9128-67d2c87a0f45",
		Secret: EnrollmentSecretRequest{
			MetaBusinessUnitID: 7,
			TagIDs:             []int{8, 9},
			SerialNumbers:      []string{"dix", "onze"},
			UDIDs:              []string{},
			Quota:              Int(12),
		},
	}

	mux.HandleFunc("/mdm/ota_enrollments/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMOTAEnrollmentRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, moeCreateJSONResponme)
	})

	ctx := context.Background()
	got, _, err := client.MDMOTAEnrollments.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("MDMOTAEnrollments.Create returned error: %v", err)
	}

	want := &MDMOTAEnrollment{
		ID:                1,
		Name:              "Yolo",
		DisplayName:       "Fomo",
		BlueprintID:       Int(2),
		PushCertificateID: 3,
		RealmUUID:         String("2217e326-5c12-406f-8c31-cc95fe9fea1b"),
		SCEPIssuerUUID:    "0a0281b1-6fc0-462b-9128-67d2c87a0f45",
		Secret: EnrollmentSecret{
			ID:                 6,
			Secret:             "SECRET",
			MetaBusinessUnitID: 7,
			TagIDs:             []int{8, 9},
			SerialNumbers:      []string{"dix", "onze"},
			UDIDs:              []string{"AF92DAAB-EC8A-42EB-A11A-60B0BD94CCC1"},
			Quota:              Int(12),
			RequestCount:       13,
		},
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMOTAEnrollments.Create returned %+v, want %+v", got, want)
	}
}

func TestMDMOTAEnrollmentsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &MDMOTAEnrollmentRequest{
		Name:              "Yolo",
		DisplayName:       String("Fomo"),
		BlueprintID:       Int(2),
		PushCertificateID: 3,
		RealmUUID:         String("2217e326-5c12-406f-8c31-cc95fe9fea1b"),
		ACMEIssuerUUID:    String("e55e9dca-1f90-47bb-851c-c28fbf9aa55a"),
		SCEPIssuerUUID:    "0a0281b1-6fc0-462b-9128-67d2c87a0f45",
		Secret: EnrollmentSecretRequest{
			MetaBusinessUnitID: 7,
			TagIDs:             []int{8, 9},
			SerialNumbers:      []string{"dix", "onze"},
			UDIDs:              []string{},
			Quota:              Int(12),
		},
	}

	mux.HandleFunc("/mdm/ota_enrollments/1/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMOTAEnrollmentRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, moeUpdateJSONResponme)
	})

	ctx := context.Background()
	got, _, err := client.MDMOTAEnrollments.Update(ctx, 1, updateRequest)
	if err != nil {
		t.Errorf("MDMOTAEnrollments.Update returned error: %v", err)
	}

	want := &MDMOTAEnrollment{
		ID:                1,
		Name:              "Yolo",
		DisplayName:       "Fomo",
		BlueprintID:       Int(2),
		PushCertificateID: 3,
		RealmUUID:         String("2217e326-5c12-406f-8c31-cc95fe9fea1b"),
		ACMEIssuerUUID:    String("e55e9dca-1f90-47bb-851c-c28fbf9aa55a"),
		SCEPIssuerUUID:    "0a0281b1-6fc0-462b-9128-67d2c87a0f45",
		Secret: EnrollmentSecret{
			ID:                 6,
			Secret:             "SECRET",
			MetaBusinessUnitID: 7,
			TagIDs:             []int{8, 9},
			SerialNumbers:      []string{"dix", "onze"},
			UDIDs:              []string{"AF92DAAB-EC8A-42EB-A11A-60B0BD94CCC1"},
			Quota:              Int(12),
			RequestCount:       13,
		},
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMOTAEnrollments.Update returned %+v, want %+v", got, want)
	}
}

func TestMDMOTAEnrollmentsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/ota_enrollments/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MDMOTAEnrollments.Delete(ctx, 1)
	if err != nil {
		t.Errorf("MDMOTAEnrollments.Delete returned error: %v", err)
	}
}
