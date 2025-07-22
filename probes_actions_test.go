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

var probeActionListJSONResponse = `
[
  {
    "id": "eaabf092-caed-4b0e-a8d5-851205b2fa56",
    "name": "Default",
    "description": "Description",
    "backend": "HTTP_POST",
    "http_post_kwargs": {
      "url": "https://www.example.com/post"
    },
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
  }
]
`

var probeActionGetJSONResponse = `
{
  "id": "eaabf092-caed-4b0e-a8d5-851205b2fa56",
  "name": "Default",
  "description": "Description",
  "backend": "HTTP_POST",
  "http_post_kwargs": {
    "url": "https://www.example.com/post",
    "username": "yolo",
    "password": "fomo",
    "headers": [
      {"name": "Authorization",
       "value": "Bearer yolofomo"}
    ]
  },
  "created_at": "2022-07-22T01:02:03.444444",
  "updated_at": "2022-07-22T01:02:03.444444"
}
`

var probeActionCreateJSONResponse = `
{
  "id": "eaabf092-caed-4b0e-a8d5-851205b2fa56",
  "name": "Default",
  "description": "Description",
  "backend": "SLACK_INCOMING_WEBHOOK",
  "slack_incoming_webhook_kwargs": {
    "url": "https://www.example.com/post"
  },
  "created_at": "2022-07-22T01:02:03.444444",
  "updated_at": "2022-07-22T01:02:03.444444"
}
`

var probeActionUpdateJSONResponse = `
{
  "id": "eaabf092-caed-4b0e-a8d5-851205b2fa56",
  "name": "Default",
  "description": "",
  "backend": "HTTP_POST",
  "http_post_kwargs": {
    "url": "https://www.example.com/post"
  },
  "created_at": "2022-07-22T01:02:03.444444",
  "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestProbesActionsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/probes/actions/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, probeActionListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.ProbesActions.List(ctx, nil)
	if err != nil {
		t.Errorf("ProbesActions.List returned error: %v", err)
	}

	want := []ProbeAction{
		{
			ID:          "eaabf092-caed-4b0e-a8d5-851205b2fa56",
			Name:        "Default",
			Description: "Description",
			Backend:     "HTTP_POST",
			HTTPPost: &ProbeActionHTTPPost{
				URL: "https://www.example.com/post",
			},
			Created: Timestamp{referenceTime},
			Updated: Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("ProbesActions.List returned %+v, want %+v", got, want)
	}
}

func TestProbesActionsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/probes/actions/eaabf092-caed-4b0e-a8d5-851205b2fa56/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, probeActionGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.ProbesActions.GetByID(ctx, "eaabf092-caed-4b0e-a8d5-851205b2fa56")
	if err != nil {
		t.Errorf("ProbesActions.GetByID returned error: %v", err)
	}

	want := &ProbeAction{
		ID:          "eaabf092-caed-4b0e-a8d5-851205b2fa56",
		Name:        "Default",
		Description: "Description",
		Backend:     "HTTP_POST",
		HTTPPost: &ProbeActionHTTPPost{
			URL:      "https://www.example.com/post",
			Username: String("yolo"),
			Password: String("fomo"),
			Headers:  []HTTPHeader{{Name: "Authorization", Value: "Bearer yolofomo"}},
		},
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("ProbesActions.GetByID returned %+v, want %+v", got, want)
	}
}

func TestProbesActionsService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/probes/actions/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "Default")
		fmt.Fprint(w, probeActionListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.ProbesActions.GetByName(ctx, "Default")
	if err != nil {
		t.Errorf("ProbesActions.GetByName returned error: %v", err)
	}

	want := &ProbeAction{
		ID:          "eaabf092-caed-4b0e-a8d5-851205b2fa56",
		Name:        "Default",
		Description: "Description",
		Backend:     "HTTP_POST",
		HTTPPost: &ProbeActionHTTPPost{
			URL: "https://www.example.com/post",
		},
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("ProbesActions.GetByName returned %+v, want %+v", got, want)
	}
}

func TestProbesActionsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &ProbeActionRequest{
		Name:        "Default",
		Description: "Description",
		Backend:     "SLACK_INCOMING_WEBHOOK",
		SlackIncomingWebhook: &ProbeActionSlackIncomingWebhook{
			URL: "https://www.example.com/post",
		},
	}

	mux.HandleFunc("/probes/actions/", func(w http.ResponseWriter, r *http.Request) {
		v := new(ProbeActionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, probeActionCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.ProbesActions.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("ProbesActions.Create returned error: %v", err)
	}

	want := &ProbeAction{
		ID:          "eaabf092-caed-4b0e-a8d5-851205b2fa56",
		Name:        "Default",
		Description: "Description",
		Backend:     "SLACK_INCOMING_WEBHOOK",
		SlackIncomingWebhook: &ProbeActionSlackIncomingWebhook{
			URL: "https://www.example.com/post",
		},
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("ProbesActions.Create returned %+v, want %+v", got, want)
	}
}

func TestProbesActionsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &ProbeActionRequest{
		Name:    "Default",
		Backend: "HTTP_POST",
		HTTPPost: &ProbeActionHTTPPost{
			URL: "https://www.example.com/post",
		},
	}

	mux.HandleFunc("/probes/actions/eaabf092-caed-4b0e-a8d5-851205b2fa56/", func(w http.ResponseWriter, r *http.Request) {
		v := new(ProbeActionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, probeActionUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.ProbesActions.Update(ctx, "eaabf092-caed-4b0e-a8d5-851205b2fa56", updateRequest)
	if err != nil {
		t.Errorf("ProbesActions.Update returned error: %v", err)
	}

	want := &ProbeAction{
		ID:          "eaabf092-caed-4b0e-a8d5-851205b2fa56",
		Name:        "Default",
		Description: "",
		Backend:     "HTTP_POST",
		HTTPPost: &ProbeActionHTTPPost{
			URL: "https://www.example.com/post",
		},
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("ProbesActions.Update returned %+v, want %+v", got, want)
	}
}

func TestProbesActionsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/probes/actions/eaabf092-caed-4b0e-a8d5-851205b2fa56/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.ProbesActions.Delete(ctx, "eaabf092-caed-4b0e-a8d5-851205b2fa56")
	if err != nil {
		t.Errorf("ProbesActions.Delete returned error: %v", err)
	}
}
