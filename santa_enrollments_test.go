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

var seListJSONResponse = `
[
    {
        "id": 1,
	"configuration": 2,
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
	"configuration_profile_download_url": "/api/santa/enrollments/1/configuration_profile/",
	"plist_download_url": "/api/santa/enrollments/1/plist/",
	"version": 11,
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var seGetJSONResponse = `
{
    "id": 1,
    "configuration": 2,
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
    "configuration_profile_download_url": "/api/santa/enrollments/1/configuration_profile/",
    "plist_download_url": "/api/santa/enrollments/1/plist/",
    "version": 12,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var seCreateJSONResponse = `
{
    "id": 1,
    "configuration": 2,
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
    "configuration_profile_download_url": "/api/santa/enrollments/1/configuration_profile/",
    "plist_download_url": "/api/santa/enrollments/1/plist/",
    "version": 1,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var seUpdateJSONResponse = `
{
    "id": 1,
    "configuration": 2,
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
    "configuration_profile_download_url": "/api/santa/enrollments/1/configuration_profile/",
    "plist_download_url": "/api/santa/enrollments/1/plist/",
    "version": 12,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestSantaEnrollmentsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/santa/enrollments/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, seListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.SantaEnrollments.List(ctx, nil)
	if err != nil {
		t.Errorf("SantaEnrollments.List returned error: %v", err)
	}

	want := []SantaEnrollment{
		{
			ID:                    1,
			ConfigurationID:       2,
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
			ConfigProfileURL: "/api/santa/enrollments/1/configuration_profile/",
			PlistURL:         "/api/santa/enrollments/1/plist/",
			Version:          11,
			Created:          Timestamp{referenceTime},
			Updated:          Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("SantaEnrollments.List returned %+v, want %+v", got, want)
	}
}

func TestSantaEnrollmentsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/santa/enrollments/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, seGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.SantaEnrollments.GetByID(ctx, 1)
	if err != nil {
		t.Errorf("SantaEnrollments.GetByID returned error: %v", err)
	}

	want := &SantaEnrollment{
		ID:                    1,
		ConfigurationID:       2,
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
		ConfigProfileURL: "/api/santa/enrollments/1/configuration_profile/",
		PlistURL:         "/api/santa/enrollments/1/plist/",
		Version:          12,
		Created:          Timestamp{referenceTime},
		Updated:          Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("SantaEnrollments.GetByID returned %+v, want %+v", got, want)
	}
}

func TestSantaEnrollmentsService_GetByConfigurationID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/santa/enrollments/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "configuration_id", "2")
		fmt.Fprint(w, seListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.SantaEnrollments.GetByConfigurationID(ctx, 2)
	if err != nil {
		t.Errorf("SantaEnrollments.GetByConfigurationID returned error: %v", err)
	}

	want := &SantaEnrollment{
		ID:                    1,
		ConfigurationID:       2,
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
		ConfigProfileURL: "/api/santa/enrollments/1/configuration_profile/",
		PlistURL:         "/api/santa/enrollments/1/plist/",
		Version:          11,
		Created:          Timestamp{referenceTime},
		Updated:          Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("SantaEnrollments.GetByConfigurationID returned %+v, want %+v", got, want)
	}
}

func TestSantaEnrollmentsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &SantaEnrollmentRequest{
		ConfigurationID: 2,
		Secret: EnrollmentSecretRequest{
			MetaBusinessUnitID: 4,
			TagIDs:             []int{6, 7},
			SerialNumbers:      []string{"huit", "neuf"},
			UDIDs:              []string{},
		},
	}

	mux.HandleFunc("/santa/enrollments/", func(w http.ResponseWriter, r *http.Request) {
		v := new(SantaEnrollmentRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, seCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.SantaEnrollments.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("SantaEnrollments.Create returned error: %v", err)
	}

	want := &SantaEnrollment{
		ID:                    1,
		ConfigurationID:       2,
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
		ConfigProfileURL: "/api/santa/enrollments/1/configuration_profile/",
		PlistURL:         "/api/santa/enrollments/1/plist/",
		Version:          1,
		Created:          Timestamp{referenceTime},
		Updated:          Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("SantaEnrollments.Create returned %+v, want %+v", got, want)
	}
}

func TestSantaEnrollmentsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &SantaEnrollmentRequest{
		ConfigurationID: 2,
		Secret: EnrollmentSecretRequest{
			MetaBusinessUnitID: 4,
			TagIDs:             []int{6, 7},
			SerialNumbers:      []string{"huit", "neuf"},
			UDIDs:              []string{"AF92DAAB-EC8A-42EB-A11A-60B0BD94CCC1"},
			Quota:              Int(10),
		},
	}

	mux.HandleFunc("/santa/enrollments/1/", func(w http.ResponseWriter, r *http.Request) {
		v := new(SantaEnrollmentRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, seUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.SantaEnrollments.Update(ctx, 1, updateRequest)
	if err != nil {
		t.Errorf("SantaEnrollments.Update returned error: %v", err)
	}

	want := &SantaEnrollment{
		ID:                    1,
		ConfigurationID:       2,
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
		ConfigProfileURL: "/api/santa/enrollments/1/configuration_profile/",
		PlistURL:         "/api/santa/enrollments/1/plist/",
		Version:          12,
		Created:          Timestamp{referenceTime},
		Updated:          Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("SantaEnrollments.Update returned %+v, want %+v", got, want)
	}
}

func TestSantaEnrollmentsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/santa/enrollments/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.SantaEnrollments.Delete(ctx, 1)
	if err != nil {
		t.Errorf("SantaEnrollments.Delete returned error: %v", err)
	}
}
