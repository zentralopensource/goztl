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

var mbaListJSONResponse = `
[
    {
        "id": 4,
	"blueprint": 5,
	"artifact": "b89d21e8-76de-4ae5-948d-5627474ab8be",
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
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var mbaGetJSONResponse = `
{
    "id": 4,
    "blueprint": 5,
    "artifact": "b89d21e8-76de-4ae5-948d-5627474ab8be",
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
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var mbaCreateJSONResponse = `
{
    "id": 4,
    "blueprint": 5,
    "artifact": "b89d21e8-76de-4ae5-948d-5627474ab8be",
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
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var mbaUpdateJSONResponse = `
{
    "id": 4,
    "blueprint": 5,
    "artifact": "b89d21e8-76de-4ae5-948d-5627474ab8be",
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
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestMDMBlueprintArtifactsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/blueprint_artifacts/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mbaListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMBlueprintArtifacts.List(ctx, nil)
	if err != nil {
		t.Errorf("MDMBlueprintArtifacts.List returned error: %v", err)
	}

	want := []MDMBlueprintArtifact{
		{
			ID:               4,
			BlueprintID:      5,
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
			Created:          Timestamp{referenceTime},
			Updated:          Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMBlueprintArtifacts.List returned %+v, want %+v", got, want)
	}
}

func TestMDMBlueprintArtifactsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/blueprint_artifacts/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mbaGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMBlueprintArtifacts.GetByID(ctx, 4)
	if err != nil {
		t.Errorf("MDMBlueprintArtifacts.GetByID returned error: %v", err)
	}

	want := &MDMBlueprintArtifact{
		ID:               4,
		BlueprintID:      5,
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
		Created:          Timestamp{referenceTime},
		Updated:          Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMBlueprintArtifacts.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMDMBlueprintArtifactsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &MDMBlueprintArtifactRequest{
		BlueprintID:      5,
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
	}

	mux.HandleFunc("/mdm/blueprint_artifacts/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMBlueprintArtifactRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, mbaCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMBlueprintArtifacts.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("MDMBlueprintArtifacts.Create returned error: %v", err)
	}

	want := &MDMBlueprintArtifact{
		ID:               4,
		BlueprintID:      5,
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
		Created:          Timestamp{referenceTime},
		Updated:          Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMBlueprintArtifacts.Create returned %+v, want %+v", got, want)
	}
}

func TestMDMBlueprintArtifactsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &MDMBlueprintArtifactRequest{
		BlueprintID:      5,
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
	}

	mux.HandleFunc("/mdm/blueprint_artifacts/4/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMBlueprintArtifactRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, mbaUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMBlueprintArtifacts.Update(ctx, 4, updateRequest)
	if err != nil {
		t.Errorf("MDMBlueprintArtifacts.Update returned error: %v", err)
	}

	want := &MDMBlueprintArtifact{
		ID:               4,
		BlueprintID:      5,
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
		Created:          Timestamp{referenceTime},
		Updated:          Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMBlueprintArtifacts.Update returned %+v, want %+v", got, want)
	}
}

func TestMDMBlueprintArtifactsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/blueprint_artifacts/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MDMBlueprintArtifacts.Delete(ctx, 4)
	if err != nil {
		t.Errorf("MDMBlueprintArtifacts.Delete returned error: %v", err)
	}
}
