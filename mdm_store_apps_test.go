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

var msaListJSONResponse = `
[
    {
        "id": "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
	"location_asset": 3,
	"associated_domains": [],
	"associated_domains_enable_direct_downloads": false,
	"configuration": null,
	"content_filter_uuid": null,
	"dns_proxy_uuid": null,
	"vpn_uuid": null,
	"prevent_backup": false,
	"removable": false,
	"remove_on_unenroll": false,
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

var msaGetJSONResponse = `
{
    "id": "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
    "location_asset": 3,
    "associated_domains": ["www.example.com"],
    "associated_domains_enable_direct_downloads": true,
    "configuration": null,
    "content_filter_uuid": "123",
    "dns_proxy_uuid": "456",
    "vpn_uuid": "789",
    "prevent_backup": true,
    "removable": true,
    "remove_on_unenroll": true,
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

var msaCreateJSONResponse = `
{
    "id": "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
    "location_asset": 3,
    "associated_domains": [],
    "associated_domains_enable_direct_downloads": false,
    "configuration": null,
    "content_filter_uuid": null,
    "dns_proxy_uuid": null,
    "vpn_uuid": null,
    "prevent_backup": false,
    "removable": false,
    "remove_on_unenroll": false,
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

var msaUpdateJSONResponse = `
{
    "id": "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
    "location_asset": 3,
    "associated_domains": ["www.example.com"],
    "associated_domains_enable_direct_downloads": true,
    "configuration": "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<!DOCTYPE plist PUBLIC \"-//Apple//DTD PLIST 1.0//EN\" \"http://www.apple.com/DTDs/PropertyList-1.0.dtd\">\n<plist version=\"1.0\">\n<dict>\n\t<key>un</key>\n\t<integer>1</integer>\n</dict>\n</plist>\n",
    "content_filter_uuid": "123",
    "dns_proxy_uuid": "456",
    "vpn_uuid": "789",
    "prevent_backup": true,
    "removable": true,
    "remove_on_unenroll": false,
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

func TestMDMStoreAppsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/store_apps/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, msaListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMStoreApps.List(ctx, nil)
	if err != nil {
		t.Errorf("MDMStoreApps.List returned error: %v", err)
	}

	want := []MDMStoreApp{
		{
			ID:                                     "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
			LocationAssetID:                        3,
			AssociatedDomains:                      []string{},
			AssociatedDomainsEnableDirectDownloads: false,
			PreventBackup:                          false,
			Removable:                              false,
			RemoveOnUnenroll:                       false,
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
		t.Errorf("MDMStoreApps.List returned %+v, want %+v", got, want)
	}
}

func TestMDMStoreAppsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/store_apps/526efd25-c1f7-498c-82b5-94ff0b39ba8e/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, msaGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMStoreApps.GetByID(ctx, "526efd25-c1f7-498c-82b5-94ff0b39ba8e")
	if err != nil {
		t.Errorf("MDMStoreApps.GetByID returned error: %v", err)
	}

	want := &MDMStoreApp{
		ID:                                     "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
		LocationAssetID:                        3,
		AssociatedDomains:                      []string{"www.example.com"},
		AssociatedDomainsEnableDirectDownloads: true,
		ContentFilterUUID:                      String("123"),
		DNSProxyUUID:                           String("456"),
		VPNUUID:                                String("789"),
		PreventBackup:                          true,
		Removable:                              true,
		RemoveOnUnenroll:                       true,
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
		t.Errorf("MDMStoreApps.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMDMStoreAppsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &MDMStoreAppRequest{
		LocationAssetID:                        3,
		AssociatedDomains:                      []string{},
		AssociatedDomainsEnableDirectDownloads: false,
		PreventBackup:                          false,
		Removable:                              false,
		RemoveOnUnenroll:                       false,
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

	mux.HandleFunc("/mdm/store_apps/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMStoreAppRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, msaCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMStoreApps.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("MDMStoreApps.Create returned error: %v", err)
	}

	want := &MDMStoreApp{
		ID:                                     "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
		LocationAssetID:                        3,
		AssociatedDomains:                      []string{},
		AssociatedDomainsEnableDirectDownloads: false,
		PreventBackup:                          false,
		Removable:                              false,
		RemoveOnUnenroll:                       false,
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
		t.Errorf("MDMStoreApps.Create returned %+v, want %+v", got, want)
	}
}

func TestMDMStoreAppsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &MDMStoreAppRequest{
		LocationAssetID:                        3,
		AssociatedDomains:                      []string{"www.example.com"},
		AssociatedDomainsEnableDirectDownloads: true,
		Configuration: String(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
        <key>un</key>
        <integer>1</integer>
</dict>
</plist>
`),
		ContentFilterUUID: String("123"),
		DNSProxyUUID:      String("456"),
		VPNUUID:           String("789"),
		PreventBackup:     true,
		Removable:         true,
		RemoveOnUnenroll:  false,
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

	mux.HandleFunc("/mdm/store_apps/526efd25-c1f7-498c-82b5-94ff0b39ba8e/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMStoreAppRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, msaUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMStoreApps.Update(ctx, "526efd25-c1f7-498c-82b5-94ff0b39ba8e", updateRequest)
	if err != nil {
		t.Errorf("MDMStoreApps.Update returned error: %v", err)
	}

	want := &MDMStoreApp{
		ID:                                     "526efd25-c1f7-498c-82b5-94ff0b39ba8e",
		LocationAssetID:                        3,
		AssociatedDomains:                      []string{"www.example.com"},
		AssociatedDomainsEnableDirectDownloads: true,
		Configuration: String(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>un</key>
	<integer>1</integer>
</dict>
</plist>
`),
		ContentFilterUUID: String("123"),
		DNSProxyUUID:      String("456"),
		VPNUUID:           String("789"),
		PreventBackup:     true,
		Removable:         true,
		RemoveOnUnenroll:  false,
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
		t.Errorf("MDMStoreApps.Update returned %+v, want %+v", got, want)
	}
}

func TestMDMStoreAppsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/store_apps/526efd25-c1f7-498c-82b5-94ff0b39ba8e/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MDMStoreApps.Delete(ctx, "526efd25-c1f7-498c-82b5-94ff0b39ba8e")
	if err != nil {
		t.Errorf("MDMStoreApps.Delete returned error: %v", err)
	}
}
