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

var tsListJSONResponse = `
{
    "count": 1,
    "next": null,
    "previous": null,
    "results": [
    {
        "id": "a1b2c3d4-0000-1111-2222-333344445555",
        "name": "Check something",
        "description": "a description",
        "source": "echo ok",
        "tag": null,
        "arch_amd64": true,
        "arch_arm64": true,
        "min_os_version": "",
        "max_os_version": "",
        "version": 1,
        "job_id": "b0000000-1111-2222-3333-444455556666",
        "compliance_check_enabled": false,
        "compliance_check_id": null,
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
    ]
}
`

var tsGetJSONResponse = `
{
    "id": "a1b2c3d4-0000-1111-2222-333344445555",
    "name": "Check something",
    "description": "a description",
    "source": "echo ok",
    "tag": 12,
    "arch_amd64": true,
    "arch_arm64": true,
    "min_os_version": "13.0",
    "max_os_version": "",
    "version": 1,
    "job_id": "b0000000-1111-2222-3333-444455556666",
    "compliance_check_enabled": true,
    "compliance_check_id": 42,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var tsCreateJSONResponse = `
{
    "id": "a1b2c3d4-0000-1111-2222-333344445555",
    "name": "Check something",
    "description": "a description",
    "source": "echo ok",
    "tag": 12,
    "arch_amd64": true,
    "arch_arm64": false,
    "min_os_version": "13.0",
    "max_os_version": "",
    "version": 1,
    "job_id": "b0000000-1111-2222-3333-444455556666",
    "compliance_check_enabled": true,
    "compliance_check_id": 42,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var tsUpdateJSONResponse = `
{
    "id": "a1b2c3d4-0000-1111-2222-333344445555",
    "name": "Check something",
    "description": "a description",
    "source": "echo ko",
    "tag": 12,
    "arch_amd64": true,
    "arch_arm64": false,
    "min_os_version": "13.0",
    "max_os_version": "",
    "version": 2,
    "job_id": "b0000000-1111-2222-3333-444455556666",
    "compliance_check_enabled": true,
    "compliance_check_id": 42,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestTurboScriptsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/turbo/scripts/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, tsListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.TurboScripts.List(ctx, nil)
	if err != nil {
		t.Errorf("TurboScripts.List returned error: %v", err)
	}

	want := []TurboScript{
		{
			ID:                     "a1b2c3d4-0000-1111-2222-333344445555",
			Name:                   "Check something",
			Description:            "a description",
			Source:                 "echo ok",
			TagID:                  nil,
			ArchAMD64:              true,
			ArchARM64:              true,
			MinOSVersion:           "",
			MaxOSVersion:           "",
			Version:                1,
			JobID:                  "b0000000-1111-2222-3333-444455556666",
			ComplianceCheckEnabled: false,
			ComplianceCheckID:      nil,
			Created:                Timestamp{referenceTime},
			Updated:                Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("TurboScripts.List returned %+v, want %+v", got, want)
	}
}

func TestTurboScriptsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/turbo/scripts/a1b2c3d4-0000-1111-2222-333344445555/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, tsGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.TurboScripts.GetByID(ctx, "a1b2c3d4-0000-1111-2222-333344445555")
	if err != nil {
		t.Errorf("TurboScripts.GetByID returned error: %v", err)
	}

	want := &TurboScript{
		ID:                     "a1b2c3d4-0000-1111-2222-333344445555",
		Name:                   "Check something",
		Description:            "a description",
		Source:                 "echo ok",
		TagID:                  Int(12),
		ArchAMD64:              true,
		ArchARM64:              true,
		MinOSVersion:           "13.0",
		MaxOSVersion:           "",
		Version:                1,
		JobID:                  "b0000000-1111-2222-3333-444455556666",
		ComplianceCheckEnabled: true,
		ComplianceCheckID:      Int(42),
		Created:                Timestamp{referenceTime},
		Updated:                Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("TurboScripts.GetByID returned %+v, want %+v", got, want)
	}
}

func TestTurboScriptsService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/turbo/scripts/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "Check something")
		fmt.Fprint(w, tsListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.TurboScripts.GetByName(ctx, "Check something")
	if err != nil {
		t.Errorf("TurboScripts.GetByName returned error: %v", err)
	}

	want := &TurboScript{
		ID:                     "a1b2c3d4-0000-1111-2222-333344445555",
		Name:                   "Check something",
		Description:            "a description",
		Source:                 "echo ok",
		TagID:                  nil,
		ArchAMD64:              true,
		ArchARM64:              true,
		MinOSVersion:           "",
		MaxOSVersion:           "",
		Version:                1,
		JobID:                  "b0000000-1111-2222-3333-444455556666",
		ComplianceCheckEnabled: false,
		ComplianceCheckID:      nil,
		Created:                Timestamp{referenceTime},
		Updated:                Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("TurboScripts.GetByName returned %+v, want %+v", got, want)
	}
}

func TestTurboScriptsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &TurboScriptRequest{
		Name:                   "Check something",
		Description:            "a description",
		Source:                 "echo ok",
		TagID:                  Int(12),
		ArchAMD64:              true,
		ArchARM64:              false,
		MinOSVersion:           "13.0",
		MaxOSVersion:           "",
		ComplianceCheckEnabled: true,
	}

	mux.HandleFunc("/turbo/scripts/", func(w http.ResponseWriter, r *http.Request) {
		v := new(TurboScriptRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, tsCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.TurboScripts.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("TurboScripts.Create returned error: %v", err)
	}

	want := &TurboScript{
		ID:                     "a1b2c3d4-0000-1111-2222-333344445555",
		Name:                   "Check something",
		Description:            "a description",
		Source:                 "echo ok",
		TagID:                  Int(12),
		ArchAMD64:              true,
		ArchARM64:              false,
		MinOSVersion:           "13.0",
		MaxOSVersion:           "",
		Version:                1,
		JobID:                  "b0000000-1111-2222-3333-444455556666",
		ComplianceCheckEnabled: true,
		ComplianceCheckID:      Int(42),
		Created:                Timestamp{referenceTime},
		Updated:                Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("TurboScripts.Create returned %+v, want %+v", got, want)
	}
}

func TestTurboScriptsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &TurboScriptRequest{
		Name:                   "Check something",
		Description:            "a description",
		Source:                 "echo ko",
		TagID:                  Int(12),
		ArchAMD64:              true,
		ArchARM64:              false,
		MinOSVersion:           "13.0",
		MaxOSVersion:           "",
		ComplianceCheckEnabled: true,
	}

	mux.HandleFunc("/turbo/scripts/a1b2c3d4-0000-1111-2222-333344445555/", func(w http.ResponseWriter, r *http.Request) {
		v := new(TurboScriptRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, tsUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.TurboScripts.Update(ctx, "a1b2c3d4-0000-1111-2222-333344445555", updateRequest)
	if err != nil {
		t.Errorf("TurboScripts.Update returned error: %v", err)
	}

	want := &TurboScript{
		ID:                     "a1b2c3d4-0000-1111-2222-333344445555",
		Name:                   "Check something",
		Description:            "a description",
		Source:                 "echo ko",
		TagID:                  Int(12),
		ArchAMD64:              true,
		ArchARM64:              false,
		MinOSVersion:           "13.0",
		MaxOSVersion:           "",
		Version:                2,
		JobID:                  "b0000000-1111-2222-3333-444455556666",
		ComplianceCheckEnabled: true,
		ComplianceCheckID:      Int(42),
		Created:                Timestamp{referenceTime},
		Updated:                Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("TurboScripts.Update returned %+v, want %+v", got, want)
	}
}

func TestTurboScriptsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/turbo/scripts/a1b2c3d4-0000-1111-2222-333344445555/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.TurboScripts.Delete(ctx, "a1b2c3d4-0000-1111-2222-333344445555")
	if err != nil {
		t.Errorf("TurboScripts.Delete returned error: %v", err)
	}
}
