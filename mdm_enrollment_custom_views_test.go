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

var enrollmentCustomViewCreateJsonResponse = `{
	"id": "23b69890-1037-44d8-bb93-f4f394eacf51", 
	"name": "jzMI1GT5urYZ", 
	"description": "P2vIRUmLk55GNJOiCKBN", 
	"html": "{{ serial_number }} {% if realm_user %}{{ realm_user.username }}{% else %}NO REALM USER{% endif %} CV", 
	"requires_authentication": false
}`
var enrollmentCustomViewGetJsonResponse = `{
	"id": "662f0711-819f-4589-9fec-cc5ce7daea8b", 
	"name": "xaAMXOaPBCrr", 
	"description": "pYetkUDaMWwy", 
	"html": "{{ serial_number }} {% if realm_user %}{{ realm_user.username }}{% else %}NO REALM USER{% endif %} CV", 
	"requires_authentication": false
}`
var enrollmentCustomViewListJsonResponse = `{
	"count": 1,
	"results": [
		{
			"id": "c4708e87-a6b0-43d4-9715-476fdf791209", 
			"name": "LpQJuTBrak73", 
			"description": "CXoZMEq6JjGn", 
			"html": "{{ serial_number }} {% if realm_user %}{{ realm_user.username }}{% else %}NO REALM USER{% endif %} CV", 
			"requires_authentication": false
		}
	]
}`
var enrollmentCustomViewListFirstPageJsonResponse = `{
	"count": 2,
	"next": "http://example.com/mdm/enrollment_custom_views/?page=2",
	"results": [
		{
			"id": "c4708e87-a6b0-43d4-9715-476fdf791209", 
			"name": "LpQJuTBrak73", 
			"description": "CXoZMEq6JjGn", 
			"html": "{{ serial_number }} {% if realm_user %}{{ realm_user.username }}{% else %}NO REALM USER{% endif %} CV", 
			"requires_authentication": false
		}
	]
}`
var enrollmentCustomViewListNextPageJsonResponse = `{
	"count": 2,
	"results": [
		{
			"id": "662f0711-819f-4589-9fec-cc5ce7daea8b", 
			"name": "xaAMXOaPBCrr", 
			"description": "pYetkUDaMWwy", 
			"html": "{{ serial_number }} {% if realm_user %}{{ realm_user.username }}{% else %}NO REALM USER{% endif %} CV", 
			"requires_authentication": false
		}
	]
}`
var enrollmentCustomViewUpdateJsonResponse = `{
	"id": "e2b112c8-7ba1-4ac9-abb7-da422596f494", 
	"name": "MFs7ziHMpYZb", 
	"description": "ZCRxD8Jz9lefXm8GRz61", 
	"html": "{{ serial_number }} {% if realm_user %}{{ realm_user.username }}{% else %}NO REALM USER{% endif %} CV", 
	"requires_authentication": false
}`

func TestMDMEnrollmentCustomViewsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/enrollment_custom_views/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")

		if r.URL.Query().Get("page") == "" {
			fmt.Fprint(w, enrollmentCustomViewListFirstPageJsonResponse)
			return
		}

		testQueryArg(t, r, "page", "2")
		fmt.Fprint(w, enrollmentCustomViewListNextPageJsonResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMEnrollmentCustomViews.List(ctx, nil)
	if err != nil {
		t.Errorf("MDMEnrollmentCustomViews.List returned error: %v", err)
	}

	want := []MDMEnrollmentCustomView{
		{
			ID:                     "c4708e87-a6b0-43d4-9715-476fdf791209",
			Name:                   "LpQJuTBrak73",
			Description:            "CXoZMEq6JjGn",
			HTML:                   "{{ serial_number }} {% if realm_user %}{{ realm_user.username }}{% else %}NO REALM USER{% endif %} CV",
			RequiresAuthentication: false,
		},
		{
			ID:                     "662f0711-819f-4589-9fec-cc5ce7daea8b",
			Name:                   "xaAMXOaPBCrr",
			Description:            "pYetkUDaMWwy",
			HTML:                   "{{ serial_number }} {% if realm_user %}{{ realm_user.username }}{% else %}NO REALM USER{% endif %} CV",
			RequiresAuthentication: false,
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMEnrollmentCustomViews.List returned %+v, want %+v", got, want)
	}
}

func TestMDMEnrollmentCustomViewsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/enrollment_custom_views/662f0711-819f-4589-9fec-cc5ce7daea8b/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, enrollmentCustomViewGetJsonResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMEnrollmentCustomViews.GetByID(ctx, "662f0711-819f-4589-9fec-cc5ce7daea8b")
	if err != nil {
		t.Errorf("MDMEnrollmentCustomViews.GetByID returned error: %v", err)
	}

	want := &MDMEnrollmentCustomView{
		ID:                     "662f0711-819f-4589-9fec-cc5ce7daea8b",
		Name:                   "xaAMXOaPBCrr",
		Description:            "pYetkUDaMWwy",
		HTML:                   "{{ serial_number }} {% if realm_user %}{{ realm_user.username }}{% else %}NO REALM USER{% endif %} CV",
		RequiresAuthentication: false,
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMEnrollmentCustomViews.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMDMEnrollmentCustomViewsService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/enrollment_custom_views/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "LpQJuTBrak73")
		fmt.Fprint(w, enrollmentCustomViewListJsonResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMEnrollmentCustomViews.GetByName(ctx, "LpQJuTBrak73")
	if err != nil {
		t.Errorf("MDMEnrollmentCustomViews.GetByName returned error: %v", err)
	}

	want := &MDMEnrollmentCustomView{
		ID:                     "c4708e87-a6b0-43d4-9715-476fdf791209",
		Name:                   "LpQJuTBrak73",
		Description:            "CXoZMEq6JjGn",
		HTML:                   "{{ serial_number }} {% if realm_user %}{{ realm_user.username }}{% else %}NO REALM USER{% endif %} CV",
		RequiresAuthentication: false,
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMEnrollmentCustomViews.GetByName returned %+v, want %+v", got, want)
	}

}

func TestMDMEnrollmentCustomViewsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &MDMEnrollmentCustomViewRequest{
		Name:        "jzMI1GT5urYZ",
		Description: "P2vIRUmLk55GNJOiCKBN",
		HTML:        "{{ serial_number }} {% if realm_user %}{{ realm_user.username }}{% else %}NO REALM USER{% endif %} CV",
	}

	mux.HandleFunc("/mdm/enrollment_custom_views/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMEnrollmentCustomViewRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, enrollmentCustomViewCreateJsonResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMEnrollmentCustomViews.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("MDMEnrollmentCustomViews.Create returned error: %v", err)
	}

	want := &MDMEnrollmentCustomView{
		ID:                     "23b69890-1037-44d8-bb93-f4f394eacf51",
		Name:                   "jzMI1GT5urYZ",
		Description:            "P2vIRUmLk55GNJOiCKBN",
		HTML:                   "{{ serial_number }} {% if realm_user %}{{ realm_user.username }}{% else %}NO REALM USER{% endif %} CV",
		RequiresAuthentication: false,
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMEnrollmentCustomViews.Create returned %+v, want %+v", got, want)
	}
}

func TestMDMEnrollmentCustomViewsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &MDMEnrollmentCustomViewRequest{
		Name:        "jzMI1GT5urYZ",
		Description: "P2vIRUmLk55GNJOiCKBN",
		HTML:        "{{ serial_number }} {% if realm_user %}{{ realm_user.username }}{% else %}NO REALM USER{% endif %} CV",
	}

	mux.HandleFunc("/mdm/enrollment_custom_views/e2b112c8-7ba1-4ac9-abb7-da422596f494/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMEnrollmentCustomViewRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, enrollmentCustomViewUpdateJsonResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMEnrollmentCustomViews.Update(ctx, "e2b112c8-7ba1-4ac9-abb7-da422596f494", updateRequest)
	if err != nil {
		t.Errorf("MDMEnrollmentCustomViews.Update returned error: %v", err)
	}

	want := &MDMEnrollmentCustomView{
		ID:                     "e2b112c8-7ba1-4ac9-abb7-da422596f494",
		Name:                   "MFs7ziHMpYZb",
		Description:            "ZCRxD8Jz9lefXm8GRz61",
		HTML:                   "{{ serial_number }} {% if realm_user %}{{ realm_user.username }}{% else %}NO REALM USER{% endif %} CV",
		RequiresAuthentication: false,
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMEnrollmentCustomViews.Update returned %+v, want %+v", got, want)
	}
}

func TestMDMEnrollmentCustomViewsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/enrollment_custom_views/e2b112c8-7ba1-4ac9-abb7-da422596f494/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MDMEnrollmentCustomViews.Delete(ctx, "e2b112c8-7ba1-4ac9-abb7-da422596f494")
	if err != nil {
		t.Errorf("MDMEnrollmentCustomViews.Delete returned error: %v", err)
	}
}
