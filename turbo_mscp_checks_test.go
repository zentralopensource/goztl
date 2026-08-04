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

var tmcListJSONResponse = `
{
    "count": 1,
    "next": null,
    "previous": null,
    "results": [
    {
        "id": "c1c2c3c4-0000-1111-2222-333344445555",
        "rule_id": "audit_acls_files_configure",
        "baseline": "",
        "odv_int": null,
        "odv_string": null,
        "odv_bool": null,
        "version": 1,
        "job_id": "d0000000-1111-2222-3333-444455556666",
        "compliance_check_id": 42,
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
    ]
}
`

var tmcGetJSONResponse = `
{
    "id": "c1c2c3c4-0000-1111-2222-333344445555",
    "rule_id": "audit_acls_files_configure",
    "baseline": "",
    "odv_int": 10,
    "odv_string": null,
    "odv_bool": null,
    "version": 1,
    "job_id": "d0000000-1111-2222-3333-444455556666",
    "compliance_check_id": 42,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var tmcCreateJSONResponse = `
{
    "id": "c1c2c3c4-0000-1111-2222-333344445555",
    "rule_id": "audit_acls_files_configure",
    "baseline": "",
    "odv_int": 10,
    "odv_string": null,
    "odv_bool": null,
    "version": 1,
    "job_id": "d0000000-1111-2222-3333-444455556666",
    "compliance_check_id": 42,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var tmcUpdateJSONResponse = `
{
    "id": "c1c2c3c4-0000-1111-2222-333344445555",
    "rule_id": "audit_acls_files_configure",
    "baseline": "cis_lvl1",
    "odv_int": null,
    "odv_string": null,
    "odv_bool": null,
    "version": 2,
    "job_id": "d0000000-1111-2222-3333-444455556666",
    "compliance_check_id": 42,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestTurboMSCPChecksService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/turbo/mscp_checks/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, tmcListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.TurboMSCPChecks.List(ctx, nil)
	if err != nil {
		t.Errorf("TurboMSCPChecks.List returned error: %v", err)
	}

	want := []TurboMSCPCheck{
		{
			ID:                "c1c2c3c4-0000-1111-2222-333344445555",
			RuleID:            "audit_acls_files_configure",
			Baseline:          "",
			ODVInt:            nil,
			ODVString:         nil,
			ODVBool:           nil,
			Version:           1,
			JobID:             "d0000000-1111-2222-3333-444455556666",
			ComplianceCheckID: 42,
			Created:           Timestamp{referenceTime},
			Updated:           Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("TurboMSCPChecks.List returned %+v, want %+v", got, want)
	}
}

func TestTurboMSCPChecksService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/turbo/mscp_checks/c1c2c3c4-0000-1111-2222-333344445555/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, tmcGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.TurboMSCPChecks.GetByID(ctx, "c1c2c3c4-0000-1111-2222-333344445555")
	if err != nil {
		t.Errorf("TurboMSCPChecks.GetByID returned error: %v", err)
	}

	want := &TurboMSCPCheck{
		ID:                "c1c2c3c4-0000-1111-2222-333344445555",
		RuleID:            "audit_acls_files_configure",
		Baseline:          "",
		ODVInt:            Int(10),
		ODVString:         nil,
		ODVBool:           nil,
		Version:           1,
		JobID:             "d0000000-1111-2222-3333-444455556666",
		ComplianceCheckID: 42,
		Created:           Timestamp{referenceTime},
		Updated:           Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("TurboMSCPChecks.GetByID returned %+v, want %+v", got, want)
	}
}

func TestTurboMSCPChecksService_GetByRuleID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/turbo/mscp_checks/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "rule_id", "audit_acls_files_configure")
		fmt.Fprint(w, tmcListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.TurboMSCPChecks.GetByRuleID(ctx, "audit_acls_files_configure")
	if err != nil {
		t.Errorf("TurboMSCPChecks.GetByRuleID returned error: %v", err)
	}

	want := []TurboMSCPCheck{
		{
			ID:                "c1c2c3c4-0000-1111-2222-333344445555",
			RuleID:            "audit_acls_files_configure",
			Baseline:          "",
			ODVInt:            nil,
			ODVString:         nil,
			ODVBool:           nil,
			Version:           1,
			JobID:             "d0000000-1111-2222-3333-444455556666",
			ComplianceCheckID: 42,
			Created:           Timestamp{referenceTime},
			Updated:           Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("TurboMSCPChecks.GetByRuleID returned %+v, want %+v", got, want)
	}
}

func TestTurboMSCPChecksService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &TurboMSCPCheckRequest{
		RuleID:    "audit_acls_files_configure",
		Baseline:  "",
		ODVInt:    Int(10),
		ODVString: nil,
		ODVBool:   nil,
	}

	mux.HandleFunc("/turbo/mscp_checks/", func(w http.ResponseWriter, r *http.Request) {
		v := new(TurboMSCPCheckRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, tmcCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.TurboMSCPChecks.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("TurboMSCPChecks.Create returned error: %v", err)
	}

	want := &TurboMSCPCheck{
		ID:                "c1c2c3c4-0000-1111-2222-333344445555",
		RuleID:            "audit_acls_files_configure",
		Baseline:          "",
		ODVInt:            Int(10),
		ODVString:         nil,
		ODVBool:           nil,
		Version:           1,
		JobID:             "d0000000-1111-2222-3333-444455556666",
		ComplianceCheckID: 42,
		Created:           Timestamp{referenceTime},
		Updated:           Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("TurboMSCPChecks.Create returned %+v, want %+v", got, want)
	}
}

func TestTurboMSCPChecksService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &TurboMSCPCheckRequest{
		RuleID:    "audit_acls_files_configure",
		Baseline:  "cis_lvl1",
		ODVInt:    nil,
		ODVString: nil,
		ODVBool:   nil,
	}

	mux.HandleFunc("/turbo/mscp_checks/c1c2c3c4-0000-1111-2222-333344445555/", func(w http.ResponseWriter, r *http.Request) {
		v := new(TurboMSCPCheckRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, tmcUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.TurboMSCPChecks.Update(ctx, "c1c2c3c4-0000-1111-2222-333344445555", updateRequest)
	if err != nil {
		t.Errorf("TurboMSCPChecks.Update returned error: %v", err)
	}

	want := &TurboMSCPCheck{
		ID:                "c1c2c3c4-0000-1111-2222-333344445555",
		RuleID:            "audit_acls_files_configure",
		Baseline:          "cis_lvl1",
		ODVInt:            nil,
		ODVString:         nil,
		ODVBool:           nil,
		Version:           2,
		JobID:             "d0000000-1111-2222-3333-444455556666",
		ComplianceCheckID: 42,
		Created:           Timestamp{referenceTime},
		Updated:           Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("TurboMSCPChecks.Update returned %+v, want %+v", got, want)
	}
}

func TestTurboMSCPChecksService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/turbo/mscp_checks/c1c2c3c4-0000-1111-2222-333344445555/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.TurboMSCPChecks.Delete(ctx, "c1c2c3c4-0000-1111-2222-333344445555")
	if err != nil {
		t.Errorf("TurboMSCPChecks.Delete returned error: %v", err)
	}
}
