package goztl

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestJMESPathChecksService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/inventory/jmespath_checks/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, fmt.Sprintf(`[{"id":1,"name":"yolo","description":"desc",
                                             "source_name":"source","platforms":["MACOS"],"tags":[18,29],
					     "jmespath_expression":"ok","version":3,
		                             "created_at":%[1]s,"updated_at":%[1]s}]`, referenceTimeStr))
	})

	ctx := context.Background()
	got, _, err := client.JMESPathChecks.List(ctx, nil)
	if err != nil {
		t.Errorf("JMESPathChecks.List returned error: %v", err)
	}

	want := []JMESPathCheck{
		{ID: 1,
			Name:               "yolo",
			Description:        "desc",
			SourceName:         "source",
			Platforms:          []string{"MACOS"},
			TagIDs:             []int{18, 29},
			JMESPathExpression: "ok",
			Version:            3,
			Created:            Timestamp{referenceTime},
			Updated:            Timestamp{referenceTime}},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("JMESPathChecks.List returned %+v, want %+v", got, want)
	}
}

func TestJMESPathChecksService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/inventory/jmespath_checks/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, fmt.Sprintf(`{"id":17,"name":"yolo","description":"desc",
                                            "source_name":"source","platforms":["MACOS"],
					    "jmespath_expression":"ok","version":3,
		                            "created_at":%[1]s,"updated_at":%[1]s}`, referenceTimeStr))
	})

	ctx := context.Background()
	got, _, err := client.JMESPathChecks.GetByID(ctx, 1)
	if err != nil {
		t.Errorf("JMESPathChecks.GetByID returned error: %v", err)
	}

	want := &JMESPathCheck{
		ID:                 17,
		Name:               "yolo",
		Description:        "desc",
		SourceName:         "source",
		Platforms:          []string{"MACOS"},
		JMESPathExpression: "ok",
		Version:            3,
		Created:            Timestamp{referenceTime},
		Updated:            Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("JMESPathChecks.GetByID returned %+v, want %+v", got, want)
	}
}

func TestJMESPathChecksService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/inventory/jmespath_checks/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "yolo")
		fmt.Fprint(w, fmt.Sprintf(`[{"id":18,"name":"yolo","description":"desc",
                                             "source_name":"source",
					     "jmespath_expression":"ok","version":3,
		                             "created_at":%[1]s,"updated_at":%[1]s}]`, referenceTimeStr))
	})

	ctx := context.Background()
	got, _, err := client.JMESPathChecks.GetByName(ctx, "yolo")
	if err != nil {
		t.Errorf("JMESPathChecks.GetByName returned error: %v", err)
	}

	want := &JMESPathCheck{
		ID:                 18,
		Name:               "yolo",
		Description:        "desc",
		SourceName:         "source",
		JMESPathExpression: "ok",
		Version:            3,
		Created:            Timestamp{referenceTime},
		Updated:            Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("JMESPathChecks.GetByName returned %+v, want %+v", got, want)
	}
}

func TestJMESPathChecksService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/inventory/jmespath_checks/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		testBody(t, r, `{"name":"yolo","description":"desc","source_name":"source","platforms":["MACOS"],"tags":[18,29],"jmespath_expression":"ok"}`+"\n")
		fmt.Fprint(w, fmt.Sprintf(`{"id":19,"name":"yolo","description":"desc",
                                            "source_name":"source","platforms":["MACOS"],"tags":[18,29],
					    "jmespath_expression":"ok","version":3,
		                            "created_at":%[1]s,"updated_at":%[1]s}`, referenceTimeStr))
	})

	ctx := context.Background()
	got, _, err := client.JMESPathChecks.Create(
		ctx,
		&JMESPathCheckCreateRequest{
			Name:               "yolo",
			Description:        "desc",
			SourceName:         "source",
			Platforms:          []string{"MACOS"},
			TagIDs:             []int{18, 29},
			JMESPathExpression: "ok",
		},
	)
	if err != nil {
		t.Errorf("JMESPathChecks.Create returned error: %v", err)
	}

	want := &JMESPathCheck{
		ID:                 19,
		Name:               "yolo",
		Description:        "desc",
		SourceName:         "source",
		Platforms:          []string{"MACOS"},
		TagIDs:             []int{18, 29},
		JMESPathExpression: "ok",
		Version:            3,
		Created:            Timestamp{referenceTime},
		Updated:            Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("JMESPathChecks.Create returned %+v, want %+v", got, want)
	}
}

func TestJMESPathChecksService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/inventory/jmespath_checks/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		testBody(t, r, `{"name":"yolo1","description":"","source_name":"source","platforms":["MACOS"],"tags":[],"jmespath_expression":"ok"}`+"\n")
		fmt.Fprint(w, fmt.Sprintf(`{"id":1,"name":"yolo1","description":"",
                                            "source_name":"source","platforms":["MACOS"],"tags":[],
					    "jmespath_expression":"ok","version":3,
		                            "created_at":%[1]s,"updated_at":%[1]s}`, referenceTimeStr))
	})

	ctx := context.Background()
	got, _, err := client.JMESPathChecks.Update(
		ctx, 1,
		&JMESPathCheckUpdateRequest{
			Name:               "yolo1",
			SourceName:         "source",
			Platforms:          []string{"MACOS"},
			TagIDs:             make([]int, 0),
			JMESPathExpression: "ok",
		},
	)
	if err != nil {
		t.Errorf("JMESPathChecks.Update returned error: %v", err)
	}

	want := &JMESPathCheck{
		ID:                 1,
		Name:               "yolo1",
		Description:        "",
		SourceName:         "source",
		Platforms:          []string{"MACOS"},
		TagIDs:             make([]int, 0),
		JMESPathExpression: "ok",
		Version:            3,
		Created:            Timestamp{referenceTime},
		Updated:            Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("JMESPathChecks.Update returned %+v, want %+v", got, want)
	}
}

func TestJMESPathChecksService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/inventory/jmespath_checks/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.JMESPathChecks.Delete(ctx, 1)
	if err != nil {
		t.Errorf("JMESPathChecks.Delete returned error: %v", err)
	}
}
