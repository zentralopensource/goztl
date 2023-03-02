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

var oeListJSONResponse = `
[
    {
        "id": 1,
	"configuration": 2,
	"osquery_release": "",
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
	"package_download_url": "/api/osquery/enrollments/1/package/",
	"script_download_url": "/api/osquery/enrollments/1/script/",
	"powershell_script_download_url": "/api/osquery/enrollments/1/powershell_script_script/",
	"version": 11,
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var oeGetJSONResponse = `
{
    "id": 1,
    "configuration": 2,
    "osquery_release": "",
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
    "package_download_url": "/api/osquery/enrollments/1/package/",
    "script_download_url": "/api/osquery/enrollments/1/script/",
    "powershell_script_download_url": "/api/osquery/enrollments/1/powershell_script_script/",
    "version": 12,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var oeCreateJSONResponse = `
{
    "id": 1,
    "configuration": 2,
    "osquery_release": "",
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
    "package_download_url": "/api/osquery/enrollments/1/package/",
    "script_download_url": "/api/osquery/enrollments/1/script/",
    "powershell_script_download_url": "/api/osquery/enrollments/1/powershell_script_script/",
    "version": 1,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var oeUpdateJSONResponse = `
{
    "id": 1,
    "configuration": 2,
    "osquery_release": "5.7.0",
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
    "package_download_url": "/api/osquery/enrollments/1/package/",
    "script_download_url": "/api/osquery/enrollments/1/script/",
    "powershell_script_download_url": "/api/osquery/enrollments/1/powershell_script_script/",
    "version": 12,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestOsqueryEnrollmentsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/enrollments/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, oeListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryEnrollments.List(ctx, nil)
	if err != nil {
		t.Errorf("OsqueryEnrollments.List returned error: %v", err)
	}

	want := []OsqueryEnrollment{
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
			PackageURL:          "/api/osquery/enrollments/1/package/",
			ScriptURL:           "/api/osquery/enrollments/1/script/",
			PowershellScriptURL: "/api/osquery/enrollments/1/powershell_script_script/",
			Version:             11,
			Created:             Timestamp{referenceTime},
			Updated:             Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryEnrollments.List returned %+v, want %+v", got, want)
	}
}

func TestOsqueryEnrollmentsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/enrollments/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, oeGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryEnrollments.GetByID(ctx, 1)
	if err != nil {
		t.Errorf("OsqueryEnrollments.GetByID returned error: %v", err)
	}

	want := &OsqueryEnrollment{
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
		PackageURL:          "/api/osquery/enrollments/1/package/",
		ScriptURL:           "/api/osquery/enrollments/1/script/",
		PowershellScriptURL: "/api/osquery/enrollments/1/powershell_script_script/",
		Version:             12,
		Created:             Timestamp{referenceTime},
		Updated:             Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryEnrollments.GetByID returned %+v, want %+v", got, want)
	}
}

func TestOsqueryEnrollmentsService_GetByConfigurationID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/enrollments/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "configuration_id", "2")
		fmt.Fprint(w, oeListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryEnrollments.GetByConfigurationID(ctx, 2)
	if err != nil {
		t.Errorf("OsqueryEnrollments.GetByConfigurationID returned error: %v", err)
	}

	want := []OsqueryEnrollment{
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
			PackageURL:          "/api/osquery/enrollments/1/package/",
			ScriptURL:           "/api/osquery/enrollments/1/script/",
			PowershellScriptURL: "/api/osquery/enrollments/1/powershell_script_script/",
			Version:             11,
			Created:             Timestamp{referenceTime},
			Updated:             Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryEnrollments.GetByConfigurationID returned %+v, want %+v", got, want)
	}
}

func TestOsqueryEnrollmentsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &OsqueryEnrollmentRequest{
		ConfigurationID: 2,
		Secret: EnrollmentSecretRequest{
			MetaBusinessUnitID: 4,
			TagIDs:             []int{6, 7},
			SerialNumbers:      []string{"huit", "neuf"},
			UDIDs:              []string{},
		},
	}

	mux.HandleFunc("/osquery/enrollments/", func(w http.ResponseWriter, r *http.Request) {
		v := new(OsqueryEnrollmentRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, oeCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryEnrollments.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("OsqueryEnrollments.Create returned error: %v", err)
	}

	want := &OsqueryEnrollment{
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
		PackageURL:          "/api/osquery/enrollments/1/package/",
		ScriptURL:           "/api/osquery/enrollments/1/script/",
		PowershellScriptURL: "/api/osquery/enrollments/1/powershell_script_script/",
		Version:             1,
		Created:             Timestamp{referenceTime},
		Updated:             Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryEnrollments.Create returned %+v, want %+v", got, want)
	}
}

func TestOsqueryEnrollmentsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &OsqueryEnrollmentRequest{
		ConfigurationID: 2,
		OsqueryRelease:  "5.7.0",
		Secret: EnrollmentSecretRequest{
			MetaBusinessUnitID: 4,
			TagIDs:             []int{6, 7},
			SerialNumbers:      []string{"huit", "neuf"},
			UDIDs:              []string{"AF92DAAB-EC8A-42EB-A11A-60B0BD94CCC1"},
			Quota:              Int(10),
		},
	}

	mux.HandleFunc("/osquery/enrollments/1/", func(w http.ResponseWriter, r *http.Request) {
		v := new(OsqueryEnrollmentRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, oeUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.OsqueryEnrollments.Update(ctx, 1, updateRequest)
	if err != nil {
		t.Errorf("OsqueryEnrollments.Update returned error: %v", err)
	}

	want := &OsqueryEnrollment{
		ID:                    1,
		ConfigurationID:       2,
		OsqueryRelease:        "5.7.0",
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
		PackageURL:          "/api/osquery/enrollments/1/package/",
		ScriptURL:           "/api/osquery/enrollments/1/script/",
		PowershellScriptURL: "/api/osquery/enrollments/1/powershell_script_script/",
		Version:             12,
		Created:             Timestamp{referenceTime},
		Updated:             Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("OsqueryEnrollments.Update returned %+v, want %+v", got, want)
	}
}

func TestOsqueryEnrollmentsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/osquery/enrollments/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.OsqueryEnrollments.Delete(ctx, 1)
	if err != nil {
		t.Errorf("OsqueryEnrollments.Delete returned error: %v", err)
	}
}
