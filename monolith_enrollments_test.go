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

var meListJSONResponme = `
[
    {
        "id": 1,
	"manifest": 2,
	"enrolled_machines_count": 3,
	"secret": {
	    "id": 4,
	    "secret": "SECRET",
	    "meta_business_unit": 5,
	    "tags": [6, 7],
	    "serial_numbers": ["huit", "neuf"],
	    "udids": [],
	    "quota": null,
	    "request_count": 10
	},
	"configuration_profile_download_url": "/api/monolith/enrollments/1/configuration_profile/",
	"plist_download_url": "/api/monolith/enrollments/1/plist/",
	"version": 11,
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var meGetJSONResponme = `
{
    "id": 1,
    "manifest": 2,
    "enrolled_machines_count": 3,
    "secret": {
	"id": 4,
	"secret": "SECRET",
	"meta_business_unit": 5,
	"tags": [6, 7],
	"serial_numbers": ["huit", "neuf"],
	"udids": ["AF92DAAB-EC8A-42EB-A11A-60B0BD94CCC1"],
	"quota": 10,
	"request_count": 11
    },
    "configuration_profile_download_url": "/api/monolith/enrollments/1/configuration_profile/",
    "plist_download_url": "/api/monolith/enrollments/1/plist/",
    "version": 12,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var meCreateJSONResponme = `
{
    "id": 1,
    "manifest": 2,
    "enrolled_machines_count": 0,
    "secret": {
	"id": 4,
	"secret": "SECRET",
	"meta_business_unit": 5,
	"tags": [6, 7],
	"serial_numbers": ["huit", "neuf"],
	"udids": [],
	"quota": null,
	"request_count": 0
    },
    "configuration_profile_download_url": "/api/monolith/enrollments/1/configuration_profile/",
    "plist_download_url": "/api/monolith/enrollments/1/plist/",
    "version": 1,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var meUpdateJSONResponme = `
{
    "id": 1,
    "manifest": 2,
    "enrolled_machines_count": 3,
    "secret": {
	"id": 4,
	"secret": "SECRET",
	"meta_business_unit": 5,
	"tags": [6, 7],
	"serial_numbers": ["huit", "neuf"],
	"udids": ["AF92DAAB-EC8A-42EB-A11A-60B0BD94CCC1"],
	"quota": 10,
	"request_count": 11
    },
    "configuration_profile_download_url": "/api/monolith/enrollments/1/configuration_profile/",
    "plist_download_url": "/api/monolith/enrollments/1/plist/",
    "version": 12,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestMonolithEnrollmentsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/enrollments/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, meListJSONResponme)
	})

	ctx := context.Background()
	got, _, err := client.MonolithEnrollments.List(ctx, nil)
	if err != nil {
		t.Errorf("MonolithEnrollments.List returned error: %v", err)
	}

	want := []MonolithEnrollment{
		{
			ID:                    1,
			ManifestID:            2,
			EnrolledMachinesCount: 3,
			Secret: EnrollmentSecret{
				ID:                 4,
				Secret:             "SECRET",
				MetaBusinessUnitID: 5,
				TagIDs:             []int{6, 7},
				SerialNumbers:      []string{"huit", "neuf"},
				UDIDs:              []string{},
				RequestCount:       10,
			},
			ConfigProfileURL: "/api/monolith/enrollments/1/configuration_profile/",
			PlistURL:         "/api/monolith/enrollments/1/plist/",
			Version:          11,
			Created:          Timestamp{referenceTime},
			Updated:          Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithEnrollments.List returned %+v, want %+v", got, want)
	}
}

func TestMonolithEnrollmentsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/enrollments/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, meGetJSONResponme)
	})

	ctx := context.Background()
	got, _, err := client.MonolithEnrollments.GetByID(ctx, 1)
	if err != nil {
		t.Errorf("MonolithEnrollments.GetByID returned error: %v", err)
	}

	want := &MonolithEnrollment{
		ID:                    1,
		ManifestID:            2,
		EnrolledMachinesCount: 3,
		Secret: EnrollmentSecret{
			ID:                 4,
			Secret:             "SECRET",
			MetaBusinessUnitID: 5,
			TagIDs:             []int{6, 7},
			SerialNumbers:      []string{"huit", "neuf"},
			UDIDs:              []string{"AF92DAAB-EC8A-42EB-A11A-60B0BD94CCC1"},
			Quota:              Int(10),
			RequestCount:       11,
		},
		ConfigProfileURL: "/api/monolith/enrollments/1/configuration_profile/",
		PlistURL:         "/api/monolith/enrollments/1/plist/",
		Version:          12,
		Created:          Timestamp{referenceTime},
		Updated:          Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithEnrollments.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMonolithEnrollmentsService_GetByManifestID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/enrollments/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "manifest_id", "2")
		fmt.Fprint(w, meListJSONResponme)
	})

	ctx := context.Background()
	got, _, err := client.MonolithEnrollments.GetByManifestID(ctx, 2)
	if err != nil {
		t.Errorf("MonolithEnrollments.GetByManifestID returned error: %v", err)
	}

	want := []MonolithEnrollment{
		{
			ID:                    1,
			ManifestID:            2,
			EnrolledMachinesCount: 3,
			Secret: EnrollmentSecret{
				ID:                 4,
				Secret:             "SECRET",
				MetaBusinessUnitID: 5,
				TagIDs:             []int{6, 7},
				SerialNumbers:      []string{"huit", "neuf"},
				UDIDs:              []string{},
				RequestCount:       10,
			},
			ConfigProfileURL: "/api/monolith/enrollments/1/configuration_profile/",
			PlistURL:         "/api/monolith/enrollments/1/plist/",
			Version:          11,
			Created:          Timestamp{referenceTime},
			Updated:          Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithEnrollments.GetByManifestID returned %+v, want %+v", got, want)
	}
}

func TestMonolithEnrollmentsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &MonolithEnrollmentRequest{
		ManifestID: 2,
		Secret: EnrollmentSecretRequest{
			MetaBusinessUnitID: 4,
			TagIDs:             []int{6, 7},
			SerialNumbers:      []string{"huit", "neuf"},
			UDIDs:              []string{},
		},
	}

	mux.HandleFunc("/monolith/enrollments/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MonolithEnrollmentRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, meCreateJSONResponme)
	})

	ctx := context.Background()
	got, _, err := client.MonolithEnrollments.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("MonolithEnrollments.Create returned error: %v", err)
	}

	want := &MonolithEnrollment{
		ID:                    1,
		ManifestID:            2,
		EnrolledMachinesCount: 0,
		Secret: EnrollmentSecret{
			ID:                 4,
			Secret:             "SECRET",
			MetaBusinessUnitID: 5,
			TagIDs:             []int{6, 7},
			SerialNumbers:      []string{"huit", "neuf"},
			UDIDs:              []string{},
			RequestCount:       0,
		},
		ConfigProfileURL: "/api/monolith/enrollments/1/configuration_profile/",
		PlistURL:         "/api/monolith/enrollments/1/plist/",
		Version:          1,
		Created:          Timestamp{referenceTime},
		Updated:          Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithEnrollments.Create returned %+v, want %+v", got, want)
	}
}

func TestMonolithEnrollmentsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &MonolithEnrollmentRequest{
		ManifestID: 2,
		Secret: EnrollmentSecretRequest{
			MetaBusinessUnitID: 4,
			TagIDs:             []int{6, 7},
			SerialNumbers:      []string{"huit", "neuf"},
			UDIDs:              []string{"AF92DAAB-EC8A-42EB-A11A-60B0BD94CCC1"},
			Quota:              Int(10),
		},
	}

	mux.HandleFunc("/monolith/enrollments/1/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MonolithEnrollmentRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, meUpdateJSONResponme)
	})

	ctx := context.Background()
	got, _, err := client.MonolithEnrollments.Update(ctx, 1, updateRequest)
	if err != nil {
		t.Errorf("MonolithEnrollments.Update returned error: %v", err)
	}

	want := &MonolithEnrollment{
		ID:                    1,
		ManifestID:            2,
		EnrolledMachinesCount: 3,
		Secret: EnrollmentSecret{
			ID:                 4,
			Secret:             "SECRET",
			MetaBusinessUnitID: 5,
			TagIDs:             []int{6, 7},
			SerialNumbers:      []string{"huit", "neuf"},
			UDIDs:              []string{"AF92DAAB-EC8A-42EB-A11A-60B0BD94CCC1"},
			Quota:              Int(10),
			RequestCount:       11,
		},
		ConfigProfileURL: "/api/monolith/enrollments/1/configuration_profile/",
		PlistURL:         "/api/monolith/enrollments/1/plist/",
		Version:          12,
		Created:          Timestamp{referenceTime},
		Updated:          Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithEnrollments.Update returned %+v, want %+v", got, want)
	}
}

func TestMonolithEnrollmentsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/enrollments/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MonolithEnrollments.Delete(ctx, 1)
	if err != nil {
		t.Errorf("MonolithEnrollments.Delete returned error: %v", err)
	}
}
