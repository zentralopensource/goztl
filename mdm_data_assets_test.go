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

var mdaListJSONResponse = `
[
    {
        "id": "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
	"type": "ZIP",
	"file_uri": "s3://bucket/test123.pkg",
	"file_sha256": "0000000000000000000000000000000000000000000000000000000000000000",
	"file_size": 12345678,
	"filename": "test123.pkg",
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
	"version": 42,
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var mdaGetJSONResponse = `
{
    "id": "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
    "type": "ZIP",
    "file_uri": "s3://bucket/test123.pkg",
    "file_sha256": "0000000000000000000000000000000000000000000000000000000000000000",
    "file_size": 12345678,
    "filename": "test123.pkg",
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
    "version": 42,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var mdaCreateJSONResponse = `
{
    "id": "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
    "type": "ZIP",
    "file_uri": "s3://bucket/test123.pkg",
    "file_sha256": "0000000000000000000000000000000000000000000000000000000000000000",
    "file_size": 12345678,
    "filename": "test123.pkg",
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
    "version": 42,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var mdaUpdateJSONResponse = `
{
    "id": "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
    "type": "ZIP",
    "file_uri": "s3://bucket/test123.zip",
    "file_sha256": "0000000000000000000000000000000000000000000000000000000000000000",
    "file_size": 12345678,
    "filename": "test123.zip",
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
    "version": 42,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestMDMDataAssetsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/data_assets/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mdaListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMDataAssets.List(ctx, nil)
	if err != nil {
		t.Errorf("MDMDataAssets.List returned error: %v", err)
	}

	want := []MDMDataAsset{
		{
			ID:         "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
			Type:       "ZIP",
			FileURI:    "s3://bucket/test123.pkg",
			FileSHA256: "0000000000000000000000000000000000000000000000000000000000000000",
			FileSize:   12345678,
			Filename:   "test123.pkg",
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
		t.Errorf("MDMDataAssets.List returned %+v, want %+v", got, want)
	}
}

func TestMDMDataAssetsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/data_assets/526efd25-c1f7-498c-82b5-94ff0b39ba8e/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mdaGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMDataAssets.GetByID(ctx, "526efd25-c1f7-498c-82b5-94ff0b39ba8e")
	if err != nil {
		t.Errorf("MDMDataAssets.GetByID returned error: %v", err)
	}

	want := &MDMDataAsset{
		ID:         "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
		Type:       "ZIP",
		FileURI:    "s3://bucket/test123.pkg",
		FileSHA256: "0000000000000000000000000000000000000000000000000000000000000000",
		FileSize:   12345678,
		Filename:   "test123.pkg",
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
		t.Errorf("MDMDataAssets.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMDMDataAssetsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &MDMDataAssetRequest{
		Type:       "ZIP",
		FileURI:    "s3://bucket/test123.pkg",
		FileSHA256: "0000000000000000000000000000000000000000000000000000000000000000",
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

	mux.HandleFunc("/mdm/data_assets/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMDataAssetRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, mdaCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMDataAssets.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("MDMDataAssets.Create returned error: %v", err)
	}

	want := &MDMDataAsset{
		ID:         "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
		Type:       "ZIP",
		FileURI:    "s3://bucket/test123.pkg",
		FileSHA256: "0000000000000000000000000000000000000000000000000000000000000000",
		FileSize:   12345678,
		Filename:   "test123.pkg",
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
		t.Errorf("MDMDataAssets.Create returned %+v, want %+v", got, want)
	}
}

func TestMDMDataAssetsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &MDMDataAssetRequest{
		Type:       "ZIP",
		FileURI:    "s3://bucket/test123.zip",
		FileSHA256: "0000000000000000000000000000000000000000000000000000000000000000",
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

	mux.HandleFunc("/mdm/data_assets/526efd25-c1f7-498c-82b5-94ff0b39ba8e/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMDataAssetRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, mdaUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMDataAssets.Update(ctx, "526efd25-c1f7-498c-82b5-94ff0b39ba8e", updateRequest)
	if err != nil {
		t.Errorf("MDMDataAssets.Update returned error: %v", err)
	}

	want := &MDMDataAsset{
		ID:         "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
		Type:       "ZIP",
		FileURI:    "s3://bucket/test123.zip",
		FileSHA256: "0000000000000000000000000000000000000000000000000000000000000000",
		FileSize:   12345678,
		Filename:   "test123.zip",
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
		t.Errorf("MDMDataAssets.Update returned %+v, want %+v", got, want)
	}
}

func TestMDMDataAssetsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/data_assets/526efd25-c1f7-498c-82b5-94ff0b39ba8e/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MDMDataAssets.Delete(ctx, "526efd25-c1f7-498c-82b5-94ff0b39ba8e")
	if err != nil {
		t.Errorf("MDMDataAssets.Delete returned error: %v", err)
	}
}
