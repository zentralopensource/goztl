package goztl

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMetaBusinessUnitsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/inventory/meta_business_units/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, fmt.Sprintf(`[{"id":1,"name":"yolo","api_enrollment_enabled":false,"created_at":%[1]s,"updated_at":%[1]s}]`, referenceTimeStr))
	})

	ctx := context.Background()
	got, _, err := client.MetaBusinessUnits.List(ctx, nil)
	if err != nil {
		t.Errorf("MetaBusinessUnits.List returned error: %v", err)
	}

	want := []MetaBusinessUnit{{ID: 1, Name: "yolo", APIEnrollmentEnabled: false, Created: Timestamp{referenceTime}, Updated: Timestamp{referenceTime}}}
	if !cmp.Equal(got, want) {
		t.Errorf("MetaBusinessUnits.List returned %+v, want %+v", got, want)
	}
}

func TestMetaBusinessUnitsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/inventory/meta_business_units/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, fmt.Sprintf(`{"id":1,"name":"yolo","api_enrollment_enabled":false,"created_at":%[1]s,"updated_at":%[1]s}`, referenceTimeStr))
	})

	ctx := context.Background()
	got, _, err := client.MetaBusinessUnits.GetByID(ctx, 1)
	if err != nil {
		t.Errorf("MetaBusinessUnits.GetByID returned error: %v", err)
	}

	want := &MetaBusinessUnit{ID: 1, Name: "yolo", APIEnrollmentEnabled: false, Created: Timestamp{referenceTime}, Updated: Timestamp{referenceTime}}
	if !cmp.Equal(got, want) {
		t.Errorf("MetaBusinessUnits.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMetaBusinessUnitsService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/inventory/meta_business_units/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "yolo")
		fmt.Fprint(w, fmt.Sprintf(`[{"id":1,"name":"yolo","api_enrollment_enabled":false,"created_at":%[1]s,"updated_at":%[1]s}]`, referenceTimeStr))
	})

	ctx := context.Background()
	got, _, err := client.MetaBusinessUnits.GetByName(ctx, "yolo")
	if err != nil {
		t.Errorf("MetaBusinessUnits.GetByName returned error: %v", err)
	}

	want := &MetaBusinessUnit{ID: 1, Name: "yolo", APIEnrollmentEnabled: false, Created: Timestamp{referenceTime}, Updated: Timestamp{referenceTime}}
	if !cmp.Equal(got, want) {
		t.Errorf("MetaBusinessUnits.GetByName returned %+v, want %+v", got, want)
	}
}

func TestMetaBusinessUnitsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/inventory/meta_business_units/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		testBody(t, r, `{"name":"yolo","api_enrollment_enabled":false}`+"\n")
		fmt.Fprint(w, fmt.Sprintf(`{"id":1,"name":"yolo","api_enrollment_enabled":false,"created_at":%[1]s,"updated_at":%[1]s}`, referenceTimeStr))
	})

	ctx := context.Background()
	got, _, err := client.MetaBusinessUnits.Create(ctx, &MetaBusinessUnitCreateRequest{Name: "yolo", APIEnrollmentEnabled: false})
	if err != nil {
		t.Errorf("MetaBusinessUnits.Create returned error: %v", err)
	}

	want := &MetaBusinessUnit{ID: 1, Name: "yolo", APIEnrollmentEnabled: false, Created: Timestamp{referenceTime}, Updated: Timestamp{referenceTime}}
	if !cmp.Equal(got, want) {
		t.Errorf("MetaBusinessUnits.Create returned %+v, want %+v", got, want)
	}
}

func TestMetaBusinessUnitsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/inventory/meta_business_units/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		testBody(t, r, `{"name":"yolo1","api_enrollment_enabled":false}`+"\n")
		fmt.Fprint(w, fmt.Sprintf(`{"id":1,"name":"yolo1","api_enrollment_enabled":false,"created_at":%[1]s,"updated_at":%[1]s}`, referenceTimeStr))
	})

	ctx := context.Background()
	got, _, err := client.MetaBusinessUnits.Update(ctx, 1, &MetaBusinessUnitUpdateRequest{Name: "yolo1", APIEnrollmentEnabled: false})
	if err != nil {
		t.Errorf("MetaBusinessUnits.Update returned error: %v", err)
	}

	want := &MetaBusinessUnit{ID: 1, Name: "yolo1", APIEnrollmentEnabled: false, Created: Timestamp{referenceTime}, Updated: Timestamp{referenceTime}}
	if !cmp.Equal(got, want) {
		t.Errorf("MetaBusinessUnits.Update returned %+v, want %+v", got, want)
	}
}

func TestMetaBusinessUnitsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/inventory/meta_business_units/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MetaBusinessUnits.Delete(ctx, 1)
	if err != nil {
		t.Errorf("MetaBusinessUnits.Delete returned error: %v", err)
	}
}
