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

var msmpiListJSONResponse = `
[
    {
        "id": 4,
	"sub_manifest": 5,
	"key": "managed_installs",
	"pkg_info_name": "Nudge",
	"featured_item": false,
	"condition": null,
	"shard_modulo": 6,
	"default_shard": 1,
	"excluded_tags": [7],
	"tag_shards": [
	    {"tag": 8, "shard": 5}
	],
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var msmpiGetJSONResponse = `
{
    "id": 4,
    "sub_manifest": 5,
    "key": "managed_installs",
    "pkg_info_name": "Nudge",
    "featured_item": false,
    "condition": 9,
    "shard_modulo": 6,
    "default_shard": 1,
    "excluded_tags": [7],
    "tag_shards": [
	{"tag": 8, "shard": 5}
    ],
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var msmpiCreateJSONResponse = `
{
    "id": 4,
    "sub_manifest": 5,
    "key": "managed_installs",
    "pkg_info_name": "Nudge",
    "featured_item": false,
    "condition": null,
    "shard_modulo": 6,
    "default_shard": 1,
    "excluded_tags": [7],
    "tag_shards": [
	{"tag": 8, "shard": 5}
    ],
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

var msmpiUpdateJSONResponse = `
{
    "id": 4,
    "sub_manifest": 5,
    "key": "managed_installs",
    "pkg_info_name": "Nudge",
    "featured_item": false,
    "condition": 9,
    "shard_modulo": 6,
    "default_shard": 1,
    "excluded_tags": [7],
    "tag_shards": [
	{"tag": 8, "shard": 5}
    ],
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestMonolithSubManifestPkgInfosService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/sub_manifest_pkg_infos/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, msmpiListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithSubManifestPkgInfos.List(ctx, nil)
	if err != nil {
		t.Errorf("MonolithSubManifestPkgInfos.List returned error: %v", err)
	}

	want := []MonolithSubManifestPkgInfo{
		{
			ID:             4,
			SubManifestID:  5,
			Key:            "managed_installs",
			PkgInfoName:    "Nudge",
			FeaturedItem:   false,
			ShardModulo:    6,
			DefaultShard:   1,
			ExcludedTagIDs: []int{7},
			TagShards:      []TagShard{{TagID: 8, Shard: 5}},
			Created:        Timestamp{referenceTime},
			Updated:        Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithSubManifestPkgInfos.List returned %+v, want %+v", got, want)
	}
}

func TestMonolithSubManifestPkgInfosService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/sub_manifest_pkg_infos/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, msmpiGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithSubManifestPkgInfos.GetByID(ctx, 4)
	if err != nil {
		t.Errorf("MonolithSubManifestPkgInfos.GetByID returned error: %v", err)
	}

	want := &MonolithSubManifestPkgInfo{
		ID:             4,
		SubManifestID:  5,
		Key:            "managed_installs",
		PkgInfoName:    "Nudge",
		FeaturedItem:   false,
		ConditionID:    Int(9),
		ShardModulo:    6,
		DefaultShard:   1,
		ExcludedTagIDs: []int{7},
		TagShards:      []TagShard{{TagID: 8, Shard: 5}},
		Created:        Timestamp{referenceTime},
		Updated:        Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithSubManifestPkgInfos.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMonolithSubManifestPkgInfosService_GetBySubManifestID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/sub_manifest_pkg_infos/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "sub_manifest_id", "5")
		fmt.Fprint(w, msmpiListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithSubManifestPkgInfos.GetBySubManifestID(ctx, 5)
	if err != nil {
		t.Errorf("MonolithSubManifestPkgInfos.GetBySubManifestID returned error: %v", err)
	}

	want := []MonolithSubManifestPkgInfo{
		{
			ID:             4,
			SubManifestID:  5,
			Key:            "managed_installs",
			PkgInfoName:    "Nudge",
			FeaturedItem:   false,
			ShardModulo:    6,
			DefaultShard:   1,
			ExcludedTagIDs: []int{7},
			TagShards:      []TagShard{{TagID: 8, Shard: 5}},
			Created:        Timestamp{referenceTime},
			Updated:        Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithSubManifestPkgInfos.GetBySubManifestID returned %+v, want %+v", got, want)
	}
}

func TestMonolithSubManifestPkgInfosService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &MonolithSubManifestPkgInfoRequest{
		SubManifestID:  5,
		Key:            "managed_installs",
		PkgInfoName:    "Nudge",
		FeaturedItem:   false,
		ShardModulo:    6,
		DefaultShard:   1,
		ExcludedTagIDs: []int{7},
		TagShards:      []TagShard{{TagID: 8, Shard: 5}},
	}

	mux.HandleFunc("/monolith/sub_manifest_pkg_infos/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MonolithSubManifestPkgInfoRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, msmpiCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithSubManifestPkgInfos.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("MonolithSubManifestPkgInfos.Create returned error: %v", err)
	}

	want := &MonolithSubManifestPkgInfo{
		ID:             4,
		SubManifestID:  5,
		Key:            "managed_installs",
		PkgInfoName:    "Nudge",
		FeaturedItem:   false,
		ShardModulo:    6,
		DefaultShard:   1,
		ExcludedTagIDs: []int{7},
		TagShards:      []TagShard{{TagID: 8, Shard: 5}},
		Created:        Timestamp{referenceTime},
		Updated:        Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithSubManifestPkgInfos.Create returned %+v, want %+v", got, want)
	}
}

func TestMonolithSubManifestPkgInfosService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &MonolithSubManifestPkgInfoRequest{
		SubManifestID:  5,
		Key:            "managed_installs",
		PkgInfoName:    "Nudge",
		FeaturedItem:   false,
		ConditionID:    Int(9),
		ShardModulo:    6,
		DefaultShard:   1,
		ExcludedTagIDs: []int{7},
		TagShards:      []TagShard{{TagID: 8, Shard: 5}},
	}

	mux.HandleFunc("/monolith/sub_manifest_pkg_infos/4/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MonolithSubManifestPkgInfoRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, msmpiUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithSubManifestPkgInfos.Update(ctx, 4, updateRequest)
	if err != nil {
		t.Errorf("MonolithSubManifestPkgInfos.Update returned error: %v", err)
	}

	want := &MonolithSubManifestPkgInfo{
		ID:             4,
		SubManifestID:  5,
		Key:            "managed_installs",
		PkgInfoName:    "Nudge",
		FeaturedItem:   false,
		ConditionID:    Int(9),
		ShardModulo:    6,
		DefaultShard:   1,
		ExcludedTagIDs: []int{7},
		TagShards:      []TagShard{{TagID: 8, Shard: 5}},
		Created:        Timestamp{referenceTime},
		Updated:        Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithSubManifestPkgInfos.Update returned %+v, want %+v", got, want)
	}
}

func TestMonolithSubManifestPkgInfosService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/sub_manifest_pkg_infos/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MonolithSubManifestPkgInfos.Delete(ctx, 4)
	if err != nil {
		t.Errorf("MonolithSubManifestPkgInfos.Delete returned error: %v", err)
	}
}
