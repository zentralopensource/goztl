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

var depCustomViewCreateJsonResponse = `{
	"id": "63bae1f7-b296-418c-af84-53aed2a9eeab", 
	"dep_enrollment": 30465, 
	"custom_view": "f95d6606-de61-42c9-bef0-fa5756d097cc", 
	"weight": 10
}`
var depCustomViewGetJsonResponse = `{
	"id": "71a9c3e4-e8db-4992-b8b5-85c098d2738d", 
	"dep_enrollment": 30475, 
	"custom_view": "1bffe93a-37c6-452a-a556-19767f3b06a7", 
	"weight": 10
}`
var depCustomViewListFirstPageJsonResponse = `{
	"count": 2,
	"next": "http://example.com/mdm/dep_enrollment_custom_views/?page=2",
	"results": [
		{
			"id": "437d47bf-1e05-4f16-8091-b3b865fa69a7", 
			"dep_enrollment": 30483, 
			"custom_view": "3ae17ae9-f638-48ac-bede-0df35653c3e3",
			"weight": 10
		}
	]
}`
var depCustomViewListNextPageJsonResponse = `{
	"count": 2,
	"results": [
		{
			"id": "71a9c3e4-e8db-4992-b8b5-85c098d2738d", 
			"dep_enrollment": 30475, 
			"custom_view": "1bffe93a-37c6-452a-a556-19767f3b06a7", 
			"weight": 10
		}
	]
}`
var depCustomViewUpdateJsonResponse = `{
	"id": "6ece1982-5c9f-48bb-8ea7-a467cfeeca2b", 
	"dep_enrollment": 30492, 
	"custom_view": "cef15372-6a06-4127-9c8b-93cab341e109", 
	"weight": 10
}`

func TestMDMDEPEnrollmentCustomViewsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/dep_enrollment_custom_views/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")

		if r.URL.Query().Get("page") == "" {
			fmt.Fprint(w, depCustomViewListFirstPageJsonResponse)
			return
		}

		testQueryArg(t, r, "page", "2")
		fmt.Fprint(w, depCustomViewListNextPageJsonResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMDEPEnrollmentCustomViews.List(ctx, nil)
	if err != nil {
		t.Errorf("MDMDEPEnrollmentCustomViews.List returned error: %v", err)
	}

	want := []MDMDEPEnrollmentCustomView{
		{
			ID:              "437d47bf-1e05-4f16-8091-b3b865fa69a7",
			DEPEnrollmentID: 30483,
			CustomViewID:    "3ae17ae9-f638-48ac-bede-0df35653c3e3",
			Weight:          10,
		},
		{
			ID:              "71a9c3e4-e8db-4992-b8b5-85c098d2738d",
			DEPEnrollmentID: 30475,
			CustomViewID:    "1bffe93a-37c6-452a-a556-19767f3b06a7",
			Weight:          10,
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMDEPEnrollmentCustomViews.List returned %+v, want %+v", got, want)
	}
}

func TestMDMDEPEnrollmentCustomViewsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/dep_enrollment_custom_views/662f0711-819f-4589-9fec-cc5ce7daea8b/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, depCustomViewGetJsonResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMDEPEnrollmentCustomViews.GetByID(ctx, "662f0711-819f-4589-9fec-cc5ce7daea8b")
	if err != nil {
		t.Errorf("MDMDEPEnrollmentCustomViews.GetByID returned error: %v", err)
	}

	want := &MDMDEPEnrollmentCustomView{
		ID:              "71a9c3e4-e8db-4992-b8b5-85c098d2738d",
		DEPEnrollmentID: 30475,
		CustomViewID:    "1bffe93a-37c6-452a-a556-19767f3b06a7",
		Weight:          10,
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMDEPEnrollmentCustomViews.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMDMDEPEnrollmentCustomViewsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &MDMDEPEnrollmentCustomViewRequest{
		DEPEnrollmentID: 30465,
		CustomViewID:    "f95d6606-de61-42c9-bef0-fa5756d097cc",
		Weight:          10,
	}

	mux.HandleFunc("/mdm/dep_enrollment_custom_views/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMDEPEnrollmentCustomViewRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, depCustomViewCreateJsonResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMDEPEnrollmentCustomViews.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("MDMDEPEnrollmentCustomViews.Create returned error: %v", err)
	}

	want := &MDMDEPEnrollmentCustomView{
		ID:              "63bae1f7-b296-418c-af84-53aed2a9eeab",
		DEPEnrollmentID: 30465,
		CustomViewID:    "f95d6606-de61-42c9-bef0-fa5756d097cc",
		Weight:          10,
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMDEPEnrollmentCustomViews.Create returned %+v, want %+v", got, want)
	}
}

func TestMDMDEPEnrollmentCustomViewsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &MDMDEPEnrollmentCustomViewRequest{
		DEPEnrollmentID: 30492,
		CustomViewID:    "cef15372-6a06-4127-9c8b-93cab341e109",
		Weight:          10,
	}

	mux.HandleFunc("/mdm/dep_enrollment_custom_views/6ece1982-5c9f-48bb-8ea7-a467cfeeca2b/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMDEPEnrollmentCustomViewRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, depCustomViewUpdateJsonResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMDEPEnrollmentCustomViews.Update(ctx, "6ece1982-5c9f-48bb-8ea7-a467cfeeca2b", updateRequest)
	if err != nil {
		t.Errorf("MDMDEPEnrollmentCustomViews.Update returned error: %v", err)
	}

	want := &MDMDEPEnrollmentCustomView{
		ID:              "6ece1982-5c9f-48bb-8ea7-a467cfeeca2b",
		DEPEnrollmentID: 30492,
		CustomViewID:    "cef15372-6a06-4127-9c8b-93cab341e109",
		Weight:          10,
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMDEPEnrollmentCustomViews.Update returned %+v, want %+v", got, want)
	}
}

func TestMDMDEPEnrollmentCustomViewsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/dep_enrollment_custom_views/e2b112c8-7ba1-4ac9-abb7-da422596f494/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MDMDEPEnrollmentCustomViews.Delete(ctx, "e2b112c8-7ba1-4ac9-abb7-da422596f494")
	if err != nil {
		t.Errorf("MDMDEPEnrollmentCustomViews.Delete returned error: %v", err)
	}
}
