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

var teListJSONResponse = `
{
    "count": 1,
    "next": null,
    "previous": null,
    "results": [
    {
        "id": 1,
        "configuration": "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
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
        "configuration_profile_download_url": "/api/turbo/enrollments/1/configuration_profile/",
        "plist_download_url": "/api/turbo/enrollments/1/plist/",
        "version": 11,
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
    ]
}
`

var teGetJSONResponse = `
{
    "id": 1,
    "configuration": "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
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
    "configuration_profile_download_url": "/api/turbo/enrollments/1/configuration_profile/",
    "plist_download_url": "/api/turbo/enrollments/1/plist/",
    "version": 12,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var teCreateJSONResponse = `
{
    "id": 1,
    "configuration": "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
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
    "configuration_profile_download_url": "/api/turbo/enrollments/1/configuration_profile/",
    "plist_download_url": "/api/turbo/enrollments/1/plist/",
    "version": 1,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var teUpdateJSONResponse = `
{
    "id": 1,
    "configuration": "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
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
    "configuration_profile_download_url": "/api/turbo/enrollments/1/configuration_profile/",
    "plist_download_url": "/api/turbo/enrollments/1/plist/",
    "version": 12,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestTurboEnrollmentsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/turbo/enrollments/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, teListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.TurboEnrollments.List(ctx, nil)
	if err != nil {
		t.Errorf("TurboEnrollments.List returned error: %v", err)
	}

	want := []TurboEnrollment{
		{
			ID:                    1,
			ConfigurationID:       "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
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
			ConfigProfileURL: "/api/turbo/enrollments/1/configuration_profile/",
			PlistURL:         "/api/turbo/enrollments/1/plist/",
			Version:          11,
			Created:          Timestamp{referenceTime},
			Updated:          Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("TurboEnrollments.List returned %+v, want %+v", got, want)
	}
}

func TestTurboEnrollmentsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/turbo/enrollments/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, teGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.TurboEnrollments.GetByID(ctx, 1)
	if err != nil {
		t.Errorf("TurboEnrollments.GetByID returned error: %v", err)
	}

	want := &TurboEnrollment{
		ID:                    1,
		ConfigurationID:       "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
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
		ConfigProfileURL: "/api/turbo/enrollments/1/configuration_profile/",
		PlistURL:         "/api/turbo/enrollments/1/plist/",
		Version:          12,
		Created:          Timestamp{referenceTime},
		Updated:          Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("TurboEnrollments.GetByID returned %+v, want %+v", got, want)
	}
}

func TestTurboEnrollmentsService_GetByConfigurationID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/turbo/enrollments/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "configuration", "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b")
		fmt.Fprint(w, teListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.TurboEnrollments.GetByConfigurationID(ctx, "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b")
	if err != nil {
		t.Errorf("TurboEnrollments.GetByConfigurationID returned error: %v", err)
	}

	want := []TurboEnrollment{
		{
			ID:                    1,
			ConfigurationID:       "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
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
			ConfigProfileURL: "/api/turbo/enrollments/1/configuration_profile/",
			PlistURL:         "/api/turbo/enrollments/1/plist/",
			Version:          11,
			Created:          Timestamp{referenceTime},
			Updated:          Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("TurboEnrollments.GetByConfigurationID returned %+v, want %+v", got, want)
	}
}

func TestTurboEnrollmentsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &TurboEnrollmentRequest{
		ConfigurationID: "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
		Secret: EnrollmentSecretRequest{
			MetaBusinessUnitID: 4,
			TagIDs:             []int{6, 7},
			SerialNumbers:      []string{"huit", "neuf"},
			UDIDs:              []string{},
		},
	}

	mux.HandleFunc("/turbo/enrollments/", func(w http.ResponseWriter, r *http.Request) {
		v := new(TurboEnrollmentRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, teCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.TurboEnrollments.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("TurboEnrollments.Create returned error: %v", err)
	}

	want := &TurboEnrollment{
		ID:                    1,
		ConfigurationID:       "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
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
		ConfigProfileURL: "/api/turbo/enrollments/1/configuration_profile/",
		PlistURL:         "/api/turbo/enrollments/1/plist/",
		Version:          1,
		Created:          Timestamp{referenceTime},
		Updated:          Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("TurboEnrollments.Create returned %+v, want %+v", got, want)
	}
}

func TestTurboEnrollmentsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &TurboEnrollmentRequest{
		ConfigurationID: "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
		Secret: EnrollmentSecretRequest{
			MetaBusinessUnitID: 4,
			TagIDs:             []int{6, 7},
			SerialNumbers:      []string{"huit", "neuf"},
			UDIDs:              []string{"AF92DAAB-EC8A-42EB-A11A-60B0BD94CCC1"},
			Quota:              Int(10),
		},
	}

	mux.HandleFunc("/turbo/enrollments/1/", func(w http.ResponseWriter, r *http.Request) {
		v := new(TurboEnrollmentRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, teUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.TurboEnrollments.Update(ctx, 1, updateRequest)
	if err != nil {
		t.Errorf("TurboEnrollments.Update returned error: %v", err)
	}

	want := &TurboEnrollment{
		ID:                    1,
		ConfigurationID:       "5f7f0f8a-1d2e-4b3c-8a9b-0c1d2e3f4a5b",
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
		ConfigProfileURL: "/api/turbo/enrollments/1/configuration_profile/",
		PlistURL:         "/api/turbo/enrollments/1/plist/",
		Version:          12,
		Created:          Timestamp{referenceTime},
		Updated:          Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("TurboEnrollments.Update returned %+v, want %+v", got, want)
	}
}

func TestTurboEnrollmentsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/turbo/enrollments/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.TurboEnrollments.Delete(ctx, 1)
	if err != nil {
		t.Errorf("TurboEnrollments.Delete returned error: %v", err)
	}
}
