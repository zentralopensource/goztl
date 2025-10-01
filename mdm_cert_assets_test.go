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

var mcaListJSONResponse = `
[
    {
        "id": "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
	"artifact": "b89d21e8-76de-4ae5-948d-5627474ab8be",
	"accessible": "Default",
	"acme_issuer": "d5656a00-ca75-4970-80b6-991d58dd9528",
	"scep_issuer": "cc939995-4e22-4ee4-b8f2-4a17838320af",
	"subject": [{"type": "CN", "value": "yolo"}],
	"subject_alt_name": {
	  "dNSName": "yolo.example.com",
	  "ntPrincipalName": "yolo@example.com",
	  "rfc822Name": "yolo@example.com",
	  "uniformResourceIdentifier": null
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

var mcaGetJSONResponse = `
{
    "id": "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
    "artifact": "b89d21e8-76de-4ae5-948d-5627474ab8be",
    "accessible": "Default",
    "acme_issuer": "d5656a00-ca75-4970-80b6-991d58dd9528",
    "scep_issuer": "cc939995-4e22-4ee4-b8f2-4a17838320af",
    "subject": [{"type": "CN", "value": "yolo"}],
    "subject_alt_name": {
      "dNSName": "dns",
      "ntPrincipalName": "nt",
      "rfc822Name": "email@example.com",
      "uniformResourceIdentifier": "https://uri"
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

var mcaCreateJSONResponse = `
{
    "id": "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
    "artifact": "b89d21e8-76de-4ae5-948d-5627474ab8be",
    "accessible": "AfterFirstUnlock",
    "acme_issuer": "d5656a00-ca75-4970-80b6-991d58dd9528",
    "scep_issuer": null,
    "subject": [{"type": "CN", "value": "yolo"}],
    "subject_alt_name": {
      "dNSName": "yolo.example.com",
      "ntPrincipalName": "yolo@example.com",
      "rfc822Name": "yolo@example.com",
      "uniformResourceIdentifier": null
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

var mcaUpdateJSONResponse = `
{
    "id": "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
    "artifact": "b89d21e8-76de-4ae5-948d-5627474ab8be",
    "accessible": "AfterFirstUnlock",
    "acme_issuer": "d5656a00-ca75-4970-80b6-991d58dd9528",
    "scep_issuer": null,
    "subject": [],
    "subject_alt_name": {
      "dNSName": null,
      "ntPrincipalName": null,
      "rfc822Name": "yolo@example.com",
      "uniformResourceIdentifier": null
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

func TestMDMCertAssetsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/cert_assets/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mcaListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMCertAssets.List(ctx, nil)
	if err != nil {
		t.Errorf("MDMCertAssets.List returned error: %v", err)
	}

	want := []MDMCertAsset{
		{
			ID:             "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
			Accessible:     "Default",
			ACMEIssuerUUID: String("d5656a00-ca75-4970-80b6-991d58dd9528"),
			SCEPIssuerUUID: String("cc939995-4e22-4ee4-b8f2-4a17838320af"),
			Subject:        []MDMCertAssetRDN{{Type: "CN", Value: "yolo"}},
			SubjectAltName: MDMCertAssetSubjectAltName{
				DNSName:         String("yolo.example.com"),
				NTPrincipalName: String("yolo@example.com"),
				RFC822Name:      String("yolo@example.com"),
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
		t.Errorf("MDMCertAssets.List returned %+v, want %+v", got, want)
	}
}

func TestMDMCertAssetsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/cert_assets/526efd25-c1f7-498c-82b5-94ff0b39ba8e/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mcaGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMCertAssets.GetByID(ctx, "526efd25-c1f7-498c-82b5-94ff0b39ba8e")
	if err != nil {
		t.Errorf("MDMCertAssets.GetByID returned error: %v", err)
	}

	want := &MDMCertAsset{
		ID:             "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
		Accessible:     "Default",
		ACMEIssuerUUID: String("d5656a00-ca75-4970-80b6-991d58dd9528"),
		SCEPIssuerUUID: String("cc939995-4e22-4ee4-b8f2-4a17838320af"),
		Subject:        []MDMCertAssetRDN{{Type: "CN", Value: "yolo"}},
		SubjectAltName: MDMCertAssetSubjectAltName{
			DNSName:         String("dns"),
			NTPrincipalName: String("nt"),
			RFC822Name:      String("email@example.com"),
			URI:             String("https://uri"),
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
		t.Errorf("MDMCertAssets.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMDMCertAssetsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &MDMCertAssetRequest{
		Accessible:     "AfterFirstUnlock",
		ACMEIssuerUUID: String("d5656a00-ca75-4970-80b6-991d58dd9528"),
		Subject:        []MDMCertAssetRDN{{Type: "CN", Value: "yolo"}},
		SubjectAltName: MDMCertAssetSubjectAltName{
			DNSName:         String("yolo.example.com"),
			NTPrincipalName: String("yolo@example.com"),
			RFC822Name:      String("yolo@example.com"),
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

	mux.HandleFunc("/mdm/cert_assets/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMCertAssetRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, mcaCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMCertAssets.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("MDMCertAssets.Create returned error: %v", err)
	}

	want := &MDMCertAsset{
		ID:             "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
		Accessible:     "AfterFirstUnlock",
		ACMEIssuerUUID: String("d5656a00-ca75-4970-80b6-991d58dd9528"),
		Subject:        []MDMCertAssetRDN{{Type: "CN", Value: "yolo"}},
		SubjectAltName: MDMCertAssetSubjectAltName{
			DNSName:         String("yolo.example.com"),
			NTPrincipalName: String("yolo@example.com"),
			RFC822Name:      String("yolo@example.com"),
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
		t.Errorf("MDMCertAssets.Create returned %+v, want %+v", got, want)
	}
}

func TestMDMCertAssetsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &MDMCertAssetRequest{
		Accessible:     "AfterFirstUnlock",
		ACMEIssuerUUID: String("d5656a00-ca75-4970-80b6-991d58dd9528"),
		Subject:        []MDMCertAssetRDN{},
		SubjectAltName: MDMCertAssetSubjectAltName{
			RFC822Name: String("yolo@example.com"),
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

	mux.HandleFunc("/mdm/cert_assets/526efd25-c1f7-498c-82b5-94ff0b39ba8e/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMCertAssetRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, mcaUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMCertAssets.Update(ctx, "526efd25-c1f7-498c-82b5-94ff0b39ba8e", updateRequest)
	if err != nil {
		t.Errorf("MDMCertAssets.Update returned error: %v", err)
	}

	want := &MDMCertAsset{
		ID:             "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
		Accessible:     "AfterFirstUnlock",
		ACMEIssuerUUID: String("d5656a00-ca75-4970-80b6-991d58dd9528"),
		Subject:        []MDMCertAssetRDN{},
		SubjectAltName: MDMCertAssetSubjectAltName{
			RFC822Name: String("yolo@example.com"),
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
		t.Errorf("MDMCertAssets.Update returned %+v, want %+v", got, want)
	}
}

func TestMDMCertAssetsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/cert_assets/526efd25-c1f7-498c-82b5-94ff0b39ba8e/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MDMCertAssets.Delete(ctx, "526efd25-c1f7-498c-82b5-94ff0b39ba8e")
	if err != nil {
		t.Errorf("MDMCertAssets.Delete returned error: %v", err)
	}
}
