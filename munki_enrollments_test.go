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

var mueListJSONResponse = `
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
	"package_download_url": "/api/munki/enrollments/1/package/",
	"version": 11,
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var mueGetJSONResponse = `
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
    "package_download_url": "/api/munki/enrollments/1/package/",
    "version": 12,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var mueCreateJSONResponse = `
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
    "package_download_url": "/api/munki/enrollments/1/package/",
    "version": 1,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var mueUpdateJSONResponse = `
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
    "package_download_url": "/api/munki/enrollments/1/package/",
    "version": 12,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestMunkiEnrollmentsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/munki/enrollments/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mueListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MunkiEnrollments.List(ctx, nil)
	if err != nil {
		t.Errorf("MunkiEnrollments.List returned error: %v", err)
	}

	want := []MunkiEnrollment{
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
			PackageURL: "/api/munki/enrollments/1/package/",
			Version:    11,
			Created:    Timestamp{referenceTime},
			Updated:    Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MunkiEnrollments.List returned %+v, want %+v", got, want)
	}
}

func TestMunkiEnrollmentsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/munki/enrollments/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mueGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MunkiEnrollments.GetByID(ctx, 1)
	if err != nil {
		t.Errorf("MunkiEnrollments.GetByID returned error: %v", err)
	}

	want := &MunkiEnrollment{
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
		PackageURL: "/api/munki/enrollments/1/package/",
		Version:    12,
		Created:    Timestamp{referenceTime},
		Updated:    Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MunkiEnrollments.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMunkiEnrollmentsService_GetByConfigurationID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/munki/enrollments/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "configuration_id", "2")
		fmt.Fprint(w, mueListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MunkiEnrollments.GetByConfigurationID(ctx, 2)
	if err != nil {
		t.Errorf("MunkiEnrollments.GetByConfigurationID returned error: %v", err)
	}

	want := []MunkiEnrollment{
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
			PackageURL: "/api/munki/enrollments/1/package/",
			Version:    11,
			Created:    Timestamp{referenceTime},
			Updated:    Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MunkiEnrollments.GetByConfigurationID returned %+v, want %+v", got, want)
	}
}

func TestMunkiEnrollmentsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &MunkiEnrollmentRequest{
		ConfigurationID: 2,
		Secret: EnrollmentSecretRequest{
			MetaBusinessUnitID: 4,
			TagIDs:             []int{6, 7},
			SerialNumbers:      []string{"huit", "neuf"},
			UDIDs:              []string{},
		},
	}

	mux.HandleFunc("/munki/enrollments/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MunkiEnrollmentRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, mueCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MunkiEnrollments.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("MunkiEnrollments.Create returned error: %v", err)
	}

	want := &MunkiEnrollment{
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
		PackageURL: "/api/munki/enrollments/1/package/",
		Version:    1,
		Created:    Timestamp{referenceTime},
		Updated:    Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MunkiEnrollments.Create returned %+v, want %+v", got, want)
	}
}

func TestMunkiEnrollmentsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &MunkiEnrollmentRequest{
		ConfigurationID: 2,
		Secret: EnrollmentSecretRequest{
			MetaBusinessUnitID: 4,
			TagIDs:             []int{6, 7},
			SerialNumbers:      []string{"huit", "neuf"},
			UDIDs:              []string{"AF92DAAB-EC8A-42EB-A11A-60B0BD94CCC1"},
			Quota:              Int(10),
		},
	}

	mux.HandleFunc("/munki/enrollments/1/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MunkiEnrollmentRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, mueUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MunkiEnrollments.Update(ctx, 1, updateRequest)
	if err != nil {
		t.Errorf("MunkiEnrollments.Update returned error: %v", err)
	}

	want := &MunkiEnrollment{
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
		PackageURL: "/api/munki/enrollments/1/package/",
		Version:    12,
		Created:    Timestamp{referenceTime},
		Updated:    Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MunkiEnrollments.Update returned %+v, want %+v", got, want)
	}
}

func TestMunkiEnrollmentsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/munki/enrollments/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MunkiEnrollments.Delete(ctx, 1)
	if err != nil {
		t.Errorf("MunkiEnrollments.Delete returned error: %v", err)
	}
}
