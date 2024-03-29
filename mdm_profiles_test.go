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

var mpListJSONResponse = `
[
    {
        "id": "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
	"artifact": "b89d21e8-76de-4ae5-948d-5627474ab8be",
        "source": "bm90IGEgcGxpc3Q=",
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

var mpGetJSONResponse = `
{
    "id": "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
    "artifact": "b89d21e8-76de-4ae5-948d-5627474ab8be",
    "source": "bm90IGEgcGxpc3Q=",
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

var mpCreateJSONResponse = `
{
    "id": "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
    "artifact": "b89d21e8-76de-4ae5-948d-5627474ab8be",
    "source": "bm90IGEgcGxpc3Q=",
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

var mpUpdateJSONResponse = `
{
    "id": "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
    "artifact": "b89d21e8-76de-4ae5-948d-5627474ab8be",
    "source": "bm90IGEgcGxpc3Q=",
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

func TestMDMProfilesService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/profiles/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mpListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMProfiles.List(ctx, nil)
	if err != nil {
		t.Errorf("MDMProfiles.List returned error: %v", err)
	}

	want := []MDMProfile{
		{
			ID:     "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
			Source: "bm90IGEgcGxpc3Q=",
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
		t.Errorf("MDMProfiles.List returned %+v, want %+v", got, want)
	}
}

func TestMDMProfilesService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/profiles/526efd25-c1f7-498c-82b5-94ff0b39ba8e/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mpGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMProfiles.GetByID(ctx, "526efd25-c1f7-498c-82b5-94ff0b39ba8e")
	if err != nil {
		t.Errorf("MDMProfiles.GetByID returned error: %v", err)
	}

	want := &MDMProfile{
		ID:     "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
		Source: "bm90IGEgcGxpc3Q=",
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
		t.Errorf("MDMProfiles.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMDMProfilesService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &MDMProfileRequest{
		Source: "bm90IGEgcGxpc3Q=",
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

	mux.HandleFunc("/mdm/profiles/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMProfileRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, mpCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMProfiles.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("MDMProfiles.Create returned error: %v", err)
	}

	want := &MDMProfile{
		ID:     "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
		Source: "bm90IGEgcGxpc3Q=",
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
		t.Errorf("MDMProfiles.Create returned %+v, want %+v", got, want)
	}
}

func TestMDMProfilesService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &MDMProfileRequest{
		Source: "bm90IGEgcGxpc3Q=",
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

	mux.HandleFunc("/mdm/profiles/526efd25-c1f7-498c-82b5-94ff0b39ba8e/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMProfileRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, mpUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMProfiles.Update(ctx, "526efd25-c1f7-498c-82b5-94ff0b39ba8e", updateRequest)
	if err != nil {
		t.Errorf("MDMProfiles.Update returned error: %v", err)
	}

	want := &MDMProfile{
		ID:     "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
		Source: "bm90IGEgcGxpc3Q=",
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
		t.Errorf("MDMProfiles.Update returned %+v, want %+v", got, want)
	}
}

func TestMDMProfilesService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/profiles/526efd25-c1f7-498c-82b5-94ff0b39ba8e/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MDMProfiles.Delete(ctx, "526efd25-c1f7-498c-82b5-94ff0b39ba8e")
	if err != nil {
		t.Errorf("MDMProfiles.Delete returned error: %v", err)
	}
}
