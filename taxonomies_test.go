package goztl

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTaxonomiesService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/inventory/taxonomies/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, fmt.Sprintf(`[{"id":1,"meta_business_unit":1,"name":"yolo","created_at":%[1]s,"updated_at":%[1]s}]`, referenceTimeStr))
	})

	ctx := context.Background()
	got, _, err := client.Taxonomies.List(ctx, nil)
	if err != nil {
		t.Errorf("Taxonomies.List returned error: %v", err)
	}

	want := []Taxonomy{{ID: 1, MetaBusinessUnitID: Int(1), Name: "yolo", Created: Timestamp{referenceTime}, Updated: Timestamp{referenceTime}}}
	if !cmp.Equal(got, want) {
		t.Errorf("Taxonomies.List returned %+v, want %+v", got, want)
	}
}

func TestTaxonomiesService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/inventory/taxonomies/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, fmt.Sprintf(`{"id":1,"meta_business_unit":null,"name":"yolo","created_at":%[1]s,"updated_at":%[1]s}`, referenceTimeStr))
	})

	ctx := context.Background()
	got, _, err := client.Taxonomies.GetByID(ctx, 1)
	if err != nil {
		t.Errorf("Taxonomies.GetByID returned error: %v", err)
	}

	want := &Taxonomy{ID: 1, MetaBusinessUnitID: nil, Name: "yolo", Created: Timestamp{referenceTime}, Updated: Timestamp{referenceTime}}
	if !cmp.Equal(got, want) {
		t.Errorf("Taxonomies.GetByID returned %+v, want %+v", got, want)
	}
}

func TestTaxonomiesService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/inventory/taxonomies/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "yolo")
		fmt.Fprint(w, fmt.Sprintf(`[{"id":1,"meta_business_unit":1,"name":"yolo","created_at":%[1]s,"updated_at":%[1]s}]`, referenceTimeStr))
	})

	ctx := context.Background()
	got, _, err := client.Taxonomies.GetByName(ctx, "yolo")
	if err != nil {
		t.Errorf("Taxonomies.GetByName returned error: %v", err)
	}

	want := &Taxonomy{ID: 1, MetaBusinessUnitID: Int(1), Name: "yolo", Created: Timestamp{referenceTime}, Updated: Timestamp{referenceTime}}
	if !cmp.Equal(got, want) {
		t.Errorf("Taxonomies.GetByName returned %+v, want %+v", got, want)
	}
}

func TestTaxonomiesService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/inventory/taxonomies/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		testBody(t, r, `{"name":"yolo","meta_business_unit":1}`+"\n")
		fmt.Fprint(w, fmt.Sprintf(`{"id":1,"meta_business_unit":1,"name":"yolo","created_at":%[1]s,"updated_at":%[1]s}`, referenceTimeStr))
	})

	ctx := context.Background()
	got, _, err := client.Taxonomies.Create(ctx, &TaxonomyCreateRequest{Name: "yolo", MetaBusinessUnitID: Int(1)})
	if err != nil {
		t.Errorf("Taxonomies.Create returned error: %v", err)
	}

	want := &Taxonomy{ID: 1, MetaBusinessUnitID: Int(1), Name: "yolo", Created: Timestamp{referenceTime}, Updated: Timestamp{referenceTime}}
	if !cmp.Equal(got, want) {
		t.Errorf("Taxonomies.Create returned %+v, want %+v", got, want)
	}
}

func TestTaxonomiesService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/inventory/taxonomies/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		testBody(t, r, `{"name":"yolo1","meta_business_unit":null}`+"\n")
		fmt.Fprint(w, fmt.Sprintf(`{"id":1,"meta_business_unit":null,"name":"yolo1","created_at":%[1]s,"updated_at":%[1]s}`, referenceTimeStr))
	})

	ctx := context.Background()
	got, _, err := client.Taxonomies.Update(ctx, 1, &TaxonomyUpdateRequest{Name: "yolo1", MetaBusinessUnitID: nil})
	if err != nil {
		t.Errorf("Taxonomies.Update returned error: %v", err)
	}

	want := &Taxonomy{ID: 1, MetaBusinessUnitID: nil, Name: "yolo1", Created: Timestamp{referenceTime}, Updated: Timestamp{referenceTime}}
	if !cmp.Equal(got, want) {
		t.Errorf("Taxonomies.Update returned %+v, want %+v", got, want)
	}
}

func TestTaxonomiesService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/inventory/taxonomies/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.Taxonomies.Delete(ctx, 1)
	if err != nil {
		t.Errorf("Taxonomies.Delete returned error: %v", err)
	}
}
