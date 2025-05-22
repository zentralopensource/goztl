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

var mdListJSONResponse = `
[
    {
        "id": "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
	"artifact": "b89d21e8-76de-4ae5-948d-5627474ab8be",
        "source": {
	    "Type": "com.apple.configuration.passcode.settings",
	    "Identifier": "com.example.zentral.pcs",
	    "ServerToken": "8cbb059c-326a-4ad8-8ffc-ea6c72e368a1",
	    "Payload": {"MinimumLength": 10}
	},
	"ios": true,
	"ios_max_version": "14",
	"ios_min_version": "13",
	"ipados": true,
	"ipados_max_version": "16",
	"ipados_min_version": "15",
	"macos": true,
	"macos_max_version": "18",
	"macos_min_version": "17",
	"tvos": true,
	"tvos_max_version": "20",
	"tvos_min_version": "19",
	"default_shard": 17,
	"shard_modulo": 35,
	"excluded_tags": [1],
	"tag_shards": [
	  {"tag": 2, "shard": 11}
        ],
	"version": 42,
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var mdGetJSONResponse = `
{
    "id": "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
    "artifact": "b89d21e8-76de-4ae5-948d-5627474ab8be",
    "source": {
        "Type": "com.apple.configuration.passcode.settings",
        "Identifier": "com.example.zentral.pcs",
        "ServerToken": "8cbb059c-326a-4ad8-8ffc-ea6c72e368a1",
        "Payload": {"MinimumLength": 10}
    },
    "ios": true,
    "ios_max_version": "14",
    "ios_min_version": "13",
    "ipados": true,
    "ipados_max_version": "16",
    "ipados_min_version": "15",
    "macos": true,
    "macos_max_version": "18",
    "macos_min_version": "17",
    "tvos": true,
    "tvos_max_version": "20",
    "tvos_min_version": "19",
    "default_shard": 17,
    "shard_modulo": 35,
    "excluded_tags": [1],
    "tag_shards": [
      {"tag": 2, "shard": 11}
    ],
    "version": 42,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var mdCreateJSONResponse = `
{
    "id": "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
    "artifact": "b89d21e8-76de-4ae5-948d-5627474ab8be",
    "source": {
        "Type": "com.apple.configuration.passcode.settings",
        "Identifier": "com.example.zentral.pcs",
        "ServerToken": "8cbb059c-326a-4ad8-8ffc-ea6c72e368a1",
        "Payload": {"MinimumLength": 10}
    },
    "ios": true,
    "ios_max_version": "14",
    "ios_min_version": "13",
    "ipados": true,
    "ipados_max_version": "16",
    "ipados_min_version": "15",
    "macos": true,
    "macos_max_version": "18",
    "macos_min_version": "17",
    "tvos": true,
    "tvos_max_version": "20",
    "tvos_min_version": "19",
    "default_shard": 17,
    "shard_modulo": 35,
    "excluded_tags": [1],
    "tag_shards": [
      {"tag": 2, "shard": 11}
    ],
    "version": 42,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var mdUpdateJSONResponse = `
{
    "id": "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
    "artifact": "b89d21e8-76de-4ae5-948d-5627474ab8be",
    "source": {
        "Type": "com.apple.configuration.passcode.settings",
        "Identifier": "com.example.zentral.pcs",
        "ServerToken": "8cbb059c-326a-4ad8-8ffc-ea6c72e368a1",
        "Payload": {"MinimumLength": 10}
    },
    "ios": true,
    "ios_max_version": "14",
    "ios_min_version": "13",
    "ipados": true,
    "ipados_max_version": "16",
    "ipados_min_version": "15",
    "macos": true,
    "macos_max_version": "18",
    "macos_min_version": "17",
    "tvos": true,
    "tvos_max_version": "20",
    "tvos_min_version": "19",
    "default_shard": 17,
    "shard_modulo": 35,
    "excluded_tags": [1],
    "tag_shards": [
      {"tag": 2, "shard": 11}
    ],
    "version": 42,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestMDMDeclarationsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/declarations/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mdListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMDeclarations.List(ctx, nil)
	if err != nil {
		t.Errorf("MDMDeclarations.List returned error: %v", err)
	}

	want := []MDMDeclaration{
		{
			ID: "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
			Source: MDMDeclarationSource{
				Type:        "com.apple.configuration.passcode.settings",
				Identifier:  "com.example.zentral.pcs",
				ServerToken: "8cbb059c-326a-4ad8-8ffc-ea6c72e368a1",
				Payload:     map[string]interface{}{"MinimumLength": 10.0},
			},
			MDMArtifactVersion: MDMArtifactVersion{
				ArtifactID:       "b89d21e8-76de-4ae5-948d-5627474ab8be",
				IOS:              true,
				IOSMaxVersion:    "14",
				IOSMinVersion:    "13",
				IPadOS:           true,
				IPadOSMaxVersion: "16",
				IPadOSMinVersion: "15",
				MacOS:            true,
				MacOSMaxVersion:  "18",
				MacOSMinVersion:  "17",
				TVOS:             true,
				TVOSMaxVersion:   "20",
				TVOSMinVersion:   "19",
				DefaultShard:     17,
				ShardModulo:      35,
				ExcludedTagIDs:   []int{1},
				TagShards:        []TagShard{{TagID: 2, Shard: 11}},
				Version:          42,
				Created:          Timestamp{referenceTime},
				Updated:          Timestamp{referenceTime},
			},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMDeclarations.List returned %+v, want %+v", got, want)
	}
}

func TestMDMDeclarationsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/declarations/526efd25-c1f7-498c-82b5-94ff0b39ba8e/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mdGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMDeclarations.GetByID(ctx, "526efd25-c1f7-498c-82b5-94ff0b39ba8e")
	if err != nil {
		t.Errorf("MDMDeclarations.GetByID returned error: %v", err)
	}

	want := &MDMDeclaration{
		ID: "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
		Source: MDMDeclarationSource{
			Type:        "com.apple.configuration.passcode.settings",
			Identifier:  "com.example.zentral.pcs",
			ServerToken: "8cbb059c-326a-4ad8-8ffc-ea6c72e368a1",
			Payload:     map[string]interface{}{"MinimumLength": 10.0},
		},
		MDMArtifactVersion: MDMArtifactVersion{
			ArtifactID:       "b89d21e8-76de-4ae5-948d-5627474ab8be",
			IOS:              true,
			IOSMaxVersion:    "14",
			IOSMinVersion:    "13",
			IPadOS:           true,
			IPadOSMaxVersion: "16",
			IPadOSMinVersion: "15",
			MacOS:            true,
			MacOSMaxVersion:  "18",
			MacOSMinVersion:  "17",
			TVOS:             true,
			TVOSMaxVersion:   "20",
			TVOSMinVersion:   "19",
			DefaultShard:     17,
			ShardModulo:      35,
			ExcludedTagIDs:   []int{1},
			TagShards:        []TagShard{{TagID: 2, Shard: 11}},
			Version:          42,
			Created:          Timestamp{referenceTime},
			Updated:          Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMDeclarations.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMDMDeclarationsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &MDMDeclarationRequest{
		Source: MDMDeclarationSource{
			Type:        "com.apple.configuration.passcode.settings",
			Identifier:  "com.example.zentral.pcs",
			ServerToken: "8cbb059c-326a-4ad8-8ffc-ea6c72e368a1",
			Payload:     map[string]interface{}{"MinimumLength": 10.0},
		},
		MDMArtifactVersionRequest: MDMArtifactVersionRequest{
			ArtifactID:       "b89d21e8-76de-4ae5-948d-5627474ab8be",
			IOS:              true,
			IOSMaxVersion:    "14",
			IOSMinVersion:    "13",
			IPadOS:           true,
			IPadOSMaxVersion: "16",
			IPadOSMinVersion: "15",
			MacOS:            true,
			MacOSMaxVersion:  "18",
			MacOSMinVersion:  "17",
			TVOS:             true,
			TVOSMaxVersion:   "20",
			TVOSMinVersion:   "19",
			DefaultShard:     17,
			ShardModulo:      35,
			ExcludedTagIDs:   []int{1},
			TagShards:        []TagShard{{TagID: 2, Shard: 11}},
			Version:          42,
		},
	}

	mux.HandleFunc("/mdm/declarations/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMDeclarationRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, mdCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMDeclarations.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("MDMDeclarations.Create returned error: %v", err)
	}

	want := &MDMDeclaration{
		ID: "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
		Source: MDMDeclarationSource{
			Type:        "com.apple.configuration.passcode.settings",
			Identifier:  "com.example.zentral.pcs",
			ServerToken: "8cbb059c-326a-4ad8-8ffc-ea6c72e368a1",
			Payload:     map[string]interface{}{"MinimumLength": 10.0},
		},
		MDMArtifactVersion: MDMArtifactVersion{
			ArtifactID:       "b89d21e8-76de-4ae5-948d-5627474ab8be",
			IOS:              true,
			IOSMaxVersion:    "14",
			IOSMinVersion:    "13",
			IPadOS:           true,
			IPadOSMaxVersion: "16",
			IPadOSMinVersion: "15",
			MacOS:            true,
			MacOSMaxVersion:  "18",
			MacOSMinVersion:  "17",
			TVOS:             true,
			TVOSMaxVersion:   "20",
			TVOSMinVersion:   "19",
			DefaultShard:     17,
			ShardModulo:      35,
			ExcludedTagIDs:   []int{1},
			TagShards:        []TagShard{{TagID: 2, Shard: 11}},
			Version:          42,
			Created:          Timestamp{referenceTime},
			Updated:          Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMDeclarations.Create returned %+v, want %+v", got, want)
	}
}

func TestMDMDeclarationsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &MDMDeclarationRequest{
		Source: MDMDeclarationSource{
			Type:        "com.apple.configuration.passcode.settings",
			Identifier:  "com.example.zentral.pcs",
			ServerToken: "8cbb059c-326a-4ad8-8ffc-ea6c72e368a1",
			Payload:     map[string]interface{}{"MinimumLength": 10.0},
		},
		MDMArtifactVersionRequest: MDMArtifactVersionRequest{
			ArtifactID:       "b89d21e8-76de-4ae5-948d-5627474ab8be",
			IOS:              true,
			IOSMaxVersion:    "14",
			IOSMinVersion:    "13",
			IPadOS:           true,
			IPadOSMaxVersion: "16",
			IPadOSMinVersion: "15",
			MacOS:            true,
			MacOSMaxVersion:  "18",
			MacOSMinVersion:  "17",
			TVOS:             true,
			TVOSMaxVersion:   "20",
			TVOSMinVersion:   "19",
			DefaultShard:     17,
			ShardModulo:      35,
			ExcludedTagIDs:   []int{1},
			TagShards:        []TagShard{{TagID: 2, Shard: 11}},
			Version:          42,
		},
	}

	mux.HandleFunc("/mdm/declarations/526efd25-c1f7-498c-82b5-94ff0b39ba8e/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMDeclarationRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, mdUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMDeclarations.Update(ctx, "526efd25-c1f7-498c-82b5-94ff0b39ba8e", updateRequest)
	if err != nil {
		t.Errorf("MDMDeclarations.Update returned error: %v", err)
	}

	want := &MDMDeclaration{
		ID: "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
		Source: MDMDeclarationSource{
			Type:        "com.apple.configuration.passcode.settings",
			Identifier:  "com.example.zentral.pcs",
			ServerToken: "8cbb059c-326a-4ad8-8ffc-ea6c72e368a1",
			Payload:     map[string]interface{}{"MinimumLength": 10.0},
		},
		MDMArtifactVersion: MDMArtifactVersion{
			ArtifactID:       "b89d21e8-76de-4ae5-948d-5627474ab8be",
			IOS:              true,
			IOSMaxVersion:    "14",
			IOSMinVersion:    "13",
			IPadOS:           true,
			IPadOSMaxVersion: "16",
			IPadOSMinVersion: "15",
			MacOS:            true,
			MacOSMaxVersion:  "18",
			MacOSMinVersion:  "17",
			TVOS:             true,
			TVOSMaxVersion:   "20",
			TVOSMinVersion:   "19",
			DefaultShard:     17,
			ShardModulo:      35,
			ExcludedTagIDs:   []int{1},
			TagShards:        []TagShard{{TagID: 2, Shard: 11}},
			Version:          42,
			Created:          Timestamp{referenceTime},
			Updated:          Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMDeclarations.Update returned %+v, want %+v", got, want)
	}
}

func TestMDMDeclarationsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/declarations/526efd25-c1f7-498c-82b5-94ff0b39ba8e/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MDMDeclarations.Delete(ctx, "526efd25-c1f7-498c-82b5-94ff0b39ba8e")
	if err != nil {
		t.Errorf("MDMDeclarations.Delete returned error: %v", err)
	}
}
