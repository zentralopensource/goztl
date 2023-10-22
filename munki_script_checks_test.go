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

var mscListJSONResponse = `
[
    {
        "id": 1,
        "name": "Default",
	"description": "Description",
	"type": "ZTL_INT",
	"source": "echo 10",
	"expected_result": "10",
	"arch_amd64": false,
	"arch_arm64": true,
	"min_os_version": "14",
	"max_os_version": "15",
	"tags": [2, 3, 4],
        "version": 5,
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var mscGetJSONResponse = `
{
    "id": 1,
    "name": "Default",
    "description": "Description",
    "type": "ZTL_INT",
    "source": "echo 10",
    "expected_result": "10",
    "arch_amd64": false,
    "arch_arm64": true,
    "min_os_version": "14",
    "max_os_version": "15",
    "tags": [2, 3, 4],
    "version": 5,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var mscCreateJSONResponse = `
{
    "id": 1,
    "name": "Default",
    "description": "Description",
    "type": "ZTL_INT",
    "source": "echo 10",
    "expected_result": "10",
    "arch_amd64": false,
    "arch_arm64": true,
    "min_os_version": "14",
    "max_os_version": "15",
    "tags": [2, 3, 4],
    "version": 5,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var mscUpdateJSONResponse = `
{
    "id": 1,
    "name": "Default",
    "description": "Description",
    "type": "ZTL_INT",
    "source": "echo 10",
    "expected_result": "10",
    "arch_amd64": false,
    "arch_arm64": true,
    "min_os_version": "14",
    "max_os_version": "15",
    "tags": [2, 3, 4],
    "version": 5,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestMunkiScriptChecksService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/munki/script_checks/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mscListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MunkiScriptChecks.List(ctx, nil)
	if err != nil {
		t.Errorf("MunkiScriptChecks.List returned error: %v", err)
	}

	want := []MunkiScriptCheck{
		{
			ID:             1,
			Name:           "Default",
			Description:    "Description",
			Type:           "ZTL_INT",
			Source:         "echo 10",
			ExpectedResult: "10",
			ArchAMD64:      false,
			ArchARM64:      true,
			MinOSVersion:   "14",
			MaxOSVersion:   "15",
			TagIDs:         []int{2, 3, 4},
			Version:        5,
			Created:        Timestamp{referenceTime},
			Updated:        Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MunkiScriptChecks.List returned %+v, want %+v", got, want)
	}
}

func TestMunkiScriptChecksService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/munki/script_checks/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mscGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MunkiScriptChecks.GetByID(ctx, 1)
	if err != nil {
		t.Errorf("MunkiScriptChecks.GetByID returned error: %v", err)
	}

	want := &MunkiScriptCheck{
		ID:             1,
		Name:           "Default",
		Description:    "Description",
		Type:           "ZTL_INT",
		Source:         "echo 10",
		ExpectedResult: "10",
		ArchAMD64:      false,
		ArchARM64:      true,
		MinOSVersion:   "14",
		MaxOSVersion:   "15",
		TagIDs:         []int{2, 3, 4},
		Version:        5,
		Created:        Timestamp{referenceTime},
		Updated:        Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MunkiScriptChecks.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMunkiScriptChecksService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/munki/script_checks/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "Default")
		fmt.Fprint(w, mscListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MunkiScriptChecks.GetByName(ctx, "Default")
	if err != nil {
		t.Errorf("MunkiScriptChecks.GetByName returned error: %v", err)
	}

	want := &MunkiScriptCheck{
		ID:             1,
		Name:           "Default",
		Description:    "Description",
		Type:           "ZTL_INT",
		Source:         "echo 10",
		ExpectedResult: "10",
		ArchAMD64:      false,
		ArchARM64:      true,
		MinOSVersion:   "14",
		MaxOSVersion:   "15",
		TagIDs:         []int{2, 3, 4},
		Version:        5,
		Created:        Timestamp{referenceTime},
		Updated:        Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MunkiScriptChecks.GetByName returned %+v, want %+v", got, want)
	}
}

func TestMunkiScriptChecksService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &MunkiScriptCheckRequest{
		Name:           "Default",
		Description:    "Description",
		Type:           "ZTL_INT",
		Source:         "echo 10",
		ExpectedResult: "10",
		ArchAMD64:      false,
		ArchARM64:      true,
		MinOSVersion:   "14",
		MaxOSVersion:   "15",
		TagIDs:         []int{2, 3, 4},
	}

	mux.HandleFunc("/munki/script_checks/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MunkiScriptCheckRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, mscCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MunkiScriptChecks.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("MunkiScriptChecks.Create returned error: %v", err)
	}

	want := &MunkiScriptCheck{
		ID:             1,
		Name:           "Default",
		Description:    "Description",
		Type:           "ZTL_INT",
		Source:         "echo 10",
		ExpectedResult: "10",
		ArchAMD64:      false,
		ArchARM64:      true,
		MinOSVersion:   "14",
		MaxOSVersion:   "15",
		TagIDs:         []int{2, 3, 4},
		Version:        5,
		Created:        Timestamp{referenceTime},
		Updated:        Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MunkiScriptChecks.Create returned %+v, want %+v", got, want)
	}
}

func TestMunkiScriptChecksService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &MunkiScriptCheckRequest{
		Name:           "Default",
		Description:    "Description",
		Type:           "ZTL_INT",
		Source:         "echo 10",
		ExpectedResult: "10",
		ArchAMD64:      false,
		ArchARM64:      true,
		MinOSVersion:   "14",
		MaxOSVersion:   "15",
		TagIDs:         []int{2, 3, 4},
	}

	mux.HandleFunc("/munki/script_checks/1/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MunkiScriptCheckRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, mscUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MunkiScriptChecks.Update(ctx, 1, updateRequest)
	if err != nil {
		t.Errorf("MunkiScriptChecks.Update returned error: %v", err)
	}

	want := &MunkiScriptCheck{
		ID:             1,
		Name:           "Default",
		Description:    "Description",
		Type:           "ZTL_INT",
		Source:         "echo 10",
		ExpectedResult: "10",
		ArchAMD64:      false,
		ArchARM64:      true,
		MinOSVersion:   "14",
		MaxOSVersion:   "15",
		TagIDs:         []int{2, 3, 4},
		Version:        5,
		Created:        Timestamp{referenceTime},
		Updated:        Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MunkiScriptChecks.Update returned %+v, want %+v", got, want)
	}
}

func TestMunkiScriptChecksService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/munki/script_checks/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MunkiScriptChecks.Delete(ctx, 1)
	if err != nil {
		t.Errorf("MunkiScriptChecks.Delete returned error: %v", err)
	}
}
