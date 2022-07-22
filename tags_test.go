package goztl

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTagsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/inventory/tags/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, `[{"id":1,"taxonomy":1,"meta_business_unit":1,"name":"yolo","slug":"yolo","color":"0079bf"}]`)
	})

	ctx := context.Background()
	got, _, err := client.Tags.List(ctx, nil)
	if err != nil {
		t.Errorf("Tags.List returned error: %v", err)
	}

	want := []Tag{{ID: 1, TaxonomyID: Int(1), MetaBusinessUnitID: Int(1), Name: "yolo", Slug: "yolo", Color: "0079bf"}}
	if !cmp.Equal(got, want) {
		t.Errorf("Tags.List returned %+v, want %+v", got, want)
	}
}

func TestTagsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/inventory/tags/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, `{"id":1,"taxonomy":null,"meta_business_unit":null,"name":"yolo","slug":"yolo","color":"0079bf"}`)
	})

	ctx := context.Background()
	got, _, err := client.Tags.GetByID(ctx, 1)
	if err != nil {
		t.Errorf("Tags.GetByID returned error: %v", err)
	}

	want := &Tag{ID: 1, TaxonomyID: nil, MetaBusinessUnitID: nil, Name: "yolo", Slug: "yolo", Color: "0079bf"}
	if !cmp.Equal(got, want) {
		t.Errorf("Tags.GetByID returned %+v, want %+v", got, want)
	}
}

func TestTagsService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/inventory/tags/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "yolo")
		fmt.Fprint(w, `[{"id":1,"taxonomy":1,"meta_business_unit":1,"name":"yolo","slug":"yolo","color":"00ff00"}]`)
	})

	ctx := context.Background()
	got, _, err := client.Tags.GetByName(ctx, "yolo")
	if err != nil {
		t.Errorf("Tags.GetByName returned error: %v", err)
	}

	want := &Tag{ID: 1, TaxonomyID: Int(1), MetaBusinessUnitID: Int(1), Name: "yolo", Slug: "yolo", Color: "00ff00"}
	if !cmp.Equal(got, want) {
		t.Errorf("Tags.GetByName returned %+v, want %+v", got, want)
	}
}

func TestTagsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/inventory/tags/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		testBody(t, r, `{"name":"yolo","taxonomy":1,"meta_business_unit":null,"color":"ff0000"}`+"\n")
		fmt.Fprint(w, `{"id":1,"taxonomy":1,"meta_business_unit":null,"name":"yolo","slug":"yolo","color":"ff0000"}`)
	})

	ctx := context.Background()
	got, _, err := client.Tags.Create(ctx, &TagCreateRequest{Name: "yolo", TaxonomyID: Int(1), MetaBusinessUnitID: nil, Color: "ff0000"})
	if err != nil {
		t.Errorf("Tags.Create returned error: %v", err)
	}

	want := &Tag{ID: 1, TaxonomyID: Int(1), MetaBusinessUnitID: nil, Name: "yolo", Slug: "yolo", Color: "ff0000"}
	if !cmp.Equal(got, want) {
		t.Errorf("Tags.Create returned %+v, want %+v", got, want)
	}
}

func TestTagsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/inventory/tags/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		testBody(t, r, `{"name":"yolo1","taxonomy":null,"meta_business_unit":1,"color":"0000ff"}`+"\n")
		fmt.Fprint(w, `{"id":1,"taxonomy":null,"meta_business_unit":1,"name":"yolo1","slug":"yolo1","color":"0000ff"}`)
	})

	ctx := context.Background()
	got, _, err := client.Tags.Update(ctx, 1, &TagUpdateRequest{Name: "yolo1", TaxonomyID: nil, MetaBusinessUnitID: Int(1), Color: "0000ff"})
	if err != nil {
		t.Errorf("Tags.Update returned error: %v", err)
	}

	want := &Tag{ID: 1, TaxonomyID: nil, MetaBusinessUnitID: Int(1), Name: "yolo1", Slug: "yolo1", Color: "0000ff"}
	if !cmp.Equal(got, want) {
		t.Errorf("Tags.Update returned %+v, want %+v", got, want)
	}
}

func TestTagsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/inventory/tags/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.Tags.Delete(ctx, 1)
	if err != nil {
		t.Errorf("Tags.Delete returned error: %v", err)
	}
}
