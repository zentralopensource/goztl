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

var storeListJSONResponse = `
[
  {
    "id": "eaabf092-caed-4b0e-a8d5-851205b2fa56",
    "provisioning_uid": null,
    "name": "Default",
    "description": "Description",
    "admin_console": true,
    "events_url_authorized_roles": [1, 2],
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444",
    "backend": "HTTP",
    "http_kwargs": {
      "endpoint_url": "https://www.example.com",
      "verify_tls": true,
      "username": null,
      "password": null,
      "concurrency": 1,
      "request_timeout": 120,
      "max_retries": 3
    }
  }
]
`

var storeGetJSONResponse = `
{
  "id": "eaabf092-caed-4b0e-a8d5-851205b2fa56",
  "provisioning_uid": null,
  "name": "Default",
  "description": "Description",
  "admin_console": false,
  "events_url_authorized_roles": [],
  "event_filters": {},
  "created_at": "2022-07-22T01:02:03.444444",
  "updated_at": "2022-07-22T01:02:03.444444",
  "backend": "HTTP",
  "http_kwargs": {
    "endpoint_url": "https://www.example.com",
    "verify_tls": true,
    "username": null,
    "password": null,
    "concurrency": 1,
    "request_timeout": 120,
    "max_retries": 3
  }
}
`

var storeCreateSplunkJSONResponse = `
{
  "id": "eaabf092-caed-4b0e-a8d5-851205b2fa56",
  "name": "Default",
  "description": "Description",
  "admin_console": false,
  "events_url_authorized_roles": [],
  "event_filters": {},
  "backend": "SPLUNK",
  "splunk_kwargs": {
    "hec_url": "https://www.example.com/hec",
    "hec_token": "yolo",
    "hec_request_timeout": 300,
    "serial_number_field": "machine_serial_number",
    "batch_size": 1,
    "search_request_timeout": 300,
    "verify_tls": true
  },
  "created_at": "2022-07-22T01:02:03.444444",
  "updated_at": "2022-07-22T01:02:03.444444"
}
`

var storeCreateKinesisJSONResponse = `
{
  "id": "eaabf092-caed-4b0e-a8d5-851205b2fa56",
  "name": "Default",
  "description": "Description",
  "admin_console": false,
  "events_url_authorized_roles": [],
  "event_filters": {},
  "backend": "KINESIS",
  "kinesis_kwargs": {
    "stream": "yolo-fomo",
    "region_name": "eu-central-1",
    "aws_access_key_id": "YOLO",
    "aws_secret_access_key": "FOMO",
    "batch_size": 17,
    "serialization_format": "firehose_v1"
  },
  "created_at": "2022-07-22T01:02:03.444444",
  "updated_at": "2022-07-22T01:02:03.444444"
}
`

var storeCreatePantherJSONResponse = `
{
  "id": "eaabf092-caed-4b0e-a8d5-851205b2fa56",
  "name": "Default",
  "description": "Description",
  "admin_console": false,
  "events_url_authorized_roles": [],
  "event_filters": {},
  "backend": "PANTHER",
  "panther_kwargs": {
    "endpoint_url": "https://panther.example.com/store/",
    "bearer_token": "YOLOFOMO",
    "batch_size": 17
  },
  "created_at": "2022-07-22T01:02:03.444444",
  "updated_at": "2022-07-22T01:02:03.444444"
}
`

var storeUpdateJSONResponse = `
{
  "id": "eaabf092-caed-4b0e-a8d5-851205b2fa56",
  "name": "Default",
  "description": "Description",
  "admin_console": true,
  "events_url_authorized_roles": [1, 2],
  "event_filters": {
    "included_event_filters": [{
      "routing_key": ["RK1"]
    }]
  },
  "backend": "SPLUNK",
  "splunk_kwargs": {
    "hec_url": "https://www.example.com/hec",
    "hec_token": "hec_token",
    "hec_extra_headers": [
      {
        "name": "X-HEC-Yolo",
        "value": "Fomo"
      }
    ],
    "hec_request_timeout": 123,
    "hec_index": "HECIndex",
    "hec_source": "HECSource",
    "computer_name_as_host_sources": [
      "munki",
      "osquery"
    ],
    "custom_host_field": "my_host",
    "serial_number_field": "serial_number",
    "batch_size": 50,
    "search_app_url": "https://www.example.com/search_app",
    "search_url": "https://www.example.com/search",
    "search_token": "search_token",
    "search_extra_headers": [
      {
        "name": "X-Search-Yolo",
        "value": "Fomo"
      }
    ],
    "search_index": "SearchIndex",
    "search_source": "SearchSource",
    "search_request_timeout": 456,
    "verify_tls": false
  },
  "created_at": "2022-07-22T01:02:03.444444",
  "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestStoresService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/stores/stores/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, storeListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.Stores.List(ctx, nil)
	if err != nil {
		t.Errorf("Stores.List returned error: %v", err)
	}

	want := []Store{
		{
			ID:                         "eaabf092-caed-4b0e-a8d5-851205b2fa56",
			Name:                       "Default",
			Description:                "Description",
			AdminConsole:               true,
			EventsURLAuthorizedRoleIDs: []int{1, 2},
			Backend:                    "HTTP",
			HTTP: &StoreHTTP{
				EndpointURL:    "https://www.example.com",
				VerifyTLS:      true,
				Concurrency:    1,
				RequestTimeout: 120,
				MaxRetries:     3,
			},
			Created: Timestamp{referenceTime},
			Updated: Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("Stores.List returned %+v, want %+v", got, want)
	}
}

func TestStoresService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/stores/stores/eaabf092-caed-4b0e-a8d5-851205b2fa56/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, storeGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.Stores.GetByID(ctx, "eaabf092-caed-4b0e-a8d5-851205b2fa56")
	if err != nil {
		t.Errorf("Stores.GetByID returned error: %v", err)
	}

	want := &Store{
		ID:                         "eaabf092-caed-4b0e-a8d5-851205b2fa56",
		Name:                       "Default",
		Description:                "Description",
		AdminConsole:               false,
		EventsURLAuthorizedRoleIDs: []int{},
		EventFilters:               &EventFilterSet{},
		Backend:                    "HTTP",
		HTTP: &StoreHTTP{
			EndpointURL:    "https://www.example.com",
			VerifyTLS:      true,
			Concurrency:    1,
			RequestTimeout: 120,
			MaxRetries:     3,
		},
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("Stores.GetByID returned %+v, want %+v", got, want)
	}
}

func TestStoresService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/stores/stores/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "Default")
		fmt.Fprint(w, storeListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.Stores.GetByName(ctx, "Default")
	if err != nil {
		t.Errorf("Stores.GetByName returned error: %v", err)
	}

	want := &Store{
		ID:                         "eaabf092-caed-4b0e-a8d5-851205b2fa56",
		Name:                       "Default",
		Description:                "Description",
		AdminConsole:               true,
		EventsURLAuthorizedRoleIDs: []int{1, 2},
		Backend:                    "HTTP",
		HTTP: &StoreHTTP{
			EndpointURL:    "https://www.example.com",
			VerifyTLS:      true,
			Concurrency:    1,
			RequestTimeout: 120,
			MaxRetries:     3,
		},
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("Stores.GetByName returned %+v, want %+v", got, want)
	}
}

func TestStoresService_CreateSplunk(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &StoreRequest{
		Name:        "Default",
		Description: "Description",
		Backend:     "SPLUNK",
		Splunk: &StoreSplunk{
			HECURL:               "https://www.example.com/hec",
			HECToken:             "yolo",
			HECRequestTimeout:    300,
			SerialNumberField:    "machine_serial_number",
			BatchSize:            1,
			SearchRequestTimeout: 300,
			VerifyTLS:            true,
		},
	}

	mux.HandleFunc("/stores/stores/", func(w http.ResponseWriter, r *http.Request) {
		v := new(StoreRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, storeCreateSplunkJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.Stores.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("Stores.Create returned error: %v", err)
	}

	want := &Store{
		ID:                         "eaabf092-caed-4b0e-a8d5-851205b2fa56",
		Name:                       "Default",
		Description:                "Description",
		AdminConsole:               false,
		EventsURLAuthorizedRoleIDs: []int{},
		EventFilters:               &EventFilterSet{},
		Backend:                    "SPLUNK",
		Splunk: &StoreSplunk{
			HECURL:               "https://www.example.com/hec",
			HECToken:             "yolo",
			HECRequestTimeout:    300,
			SerialNumberField:    "machine_serial_number",
			BatchSize:            1,
			SearchRequestTimeout: 300,
			VerifyTLS:            true,
		},
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("Stores.Create returned %+v, want %+v", got, want)
	}
}

func TestStoresService_CreateKinesis(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &StoreRequest{
		Name:        "Default",
		Description: "Description",
		Backend:     "KINESIS",
		Kinesis: &StoreKinesis{
			Stream:              "yolo-fomo",
			RegionName:          "eu-central-1",
			AWSAccessKeyID:      String("YOLO"),
			AWSSecretAccessKey:  String("FOMO"),
			BatchSize:           17,
			SerializationFormat: "firehose_v1",
		},
	}

	mux.HandleFunc("/stores/stores/", func(w http.ResponseWriter, r *http.Request) {
		v := new(StoreRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, storeCreateKinesisJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.Stores.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("Stores.Create Kinesis returned error: %v", err)
	}

	want := &Store{
		ID:                         "eaabf092-caed-4b0e-a8d5-851205b2fa56",
		Name:                       "Default",
		Description:                "Description",
		AdminConsole:               false,
		EventsURLAuthorizedRoleIDs: []int{},
		EventFilters:               &EventFilterSet{},
		Backend:                    "KINESIS",
		Kinesis: &StoreKinesis{
			Stream:              "yolo-fomo",
			RegionName:          "eu-central-1",
			AWSAccessKeyID:      String("YOLO"),
			AWSSecretAccessKey:  String("FOMO"),
			BatchSize:           17,
			SerializationFormat: "firehose_v1",
		},
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("Stores.Create Kinesis returned %+v, want %+v", got, want)
	}
}

func TestStoresService_CreatePanther(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &StoreRequest{
		Name:        "Default",
		Description: "Description",
		Backend:     "PANTHER",
		Panther: &StorePanther{
			EndpointURL: "https://panther.example.com/store/",
			BearerToken: "YOLOFOMO",
			BatchSize:   17,
		},
	}

	mux.HandleFunc("/stores/stores/", func(w http.ResponseWriter, r *http.Request) {
		v := new(StoreRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, storeCreatePantherJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.Stores.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("Stores.Create Panther returned error: %v", err)
	}

	want := &Store{
		ID:                         "eaabf092-caed-4b0e-a8d5-851205b2fa56",
		Name:                       "Default",
		Description:                "Description",
		AdminConsole:               false,
		EventsURLAuthorizedRoleIDs: []int{},
		EventFilters:               &EventFilterSet{},
		Backend:                    "PANTHER",
		Panther: &StorePanther{
			EndpointURL: "https://panther.example.com/store/",
			BearerToken: "YOLOFOMO",
			BatchSize:   17,
		},
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("Stores.Create Panther returned %+v, want %+v", got, want)
	}
}

func TestStoresService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &StoreRequest{
		Name:                       "Default",
		Description:                "Description",
		AdminConsole:               true,
		EventsURLAuthorizedRoleIDs: []int{1, 2},
		EventFilters: &EventFilterSet{
			IncludedEventFilters: []EventFilter{{
				RoutingKey: []string{"RK1"},
			}},
		},
		Backend: "SPLUNK",
		Splunk: &StoreSplunk{
			HECURL:   "https://www.example.com/hec",
			HECToken: "hec_token",
			HECExtraHeaders: []HTTPHeader{{
				Name:  "X-HEC-Yolo",
				Value: "Fomo",
			}},
			HECRequestTimeout:         123,
			HECIndex:                  String("HECIndex"),
			HECSource:                 String("HECSource"),
			ComputerNameAsHostSources: []string{"munki", "osquery"},
			CustomHostField:           String("my_host"),
			SerialNumberField:         "serial_number",
			BatchSize:                 50,
			SearchAppURL:              String("https://www.example.com/search_app"),
			SearchURL:                 String("https://www.example.com/search"),
			SearchToken:               String("search_token"),
			SearchExtraHeaders: []HTTPHeader{{
				Name:  "X-Search-Yolo",
				Value: "Fomo",
			}},
			SearchIndex:          String("SearchIndex"),
			SearchSource:         String("SearchSource"),
			SearchRequestTimeout: 456,
			VerifyTLS:            true,
		},
	}

	mux.HandleFunc("/stores/stores/eaabf092-caed-4b0e-a8d5-851205b2fa56/", func(w http.ResponseWriter, r *http.Request) {
		v := new(StoreRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, storeUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.Stores.Update(ctx, "eaabf092-caed-4b0e-a8d5-851205b2fa56", updateRequest)
	if err != nil {
		t.Errorf("Stores.Update returned error: %v", err)
	}

	want := &Store{
		ID:                         "eaabf092-caed-4b0e-a8d5-851205b2fa56",
		Name:                       "Default",
		Description:                "Description",
		AdminConsole:               true,
		EventsURLAuthorizedRoleIDs: []int{1, 2},
		EventFilters: &EventFilterSet{
			IncludedEventFilters: []EventFilter{{
				RoutingKey: []string{"RK1"},
			}},
		},
		Backend: "SPLUNK",
		Splunk: &StoreSplunk{
			HECURL:   "https://www.example.com/hec",
			HECToken: "hec_token",
			HECExtraHeaders: []HTTPHeader{{
				Name:  "X-HEC-Yolo",
				Value: "Fomo",
			}},
			HECRequestTimeout:         123,
			HECIndex:                  String("HECIndex"),
			HECSource:                 String("HECSource"),
			ComputerNameAsHostSources: []string{"munki", "osquery"},
			CustomHostField:           String("my_host"),
			SerialNumberField:         "serial_number",
			BatchSize:                 50,
			SearchAppURL:              String("https://www.example.com/search_app"),
			SearchURL:                 String("https://www.example.com/search"),
			SearchToken:               String("search_token"),
			SearchExtraHeaders: []HTTPHeader{{
				Name:  "X-Search-Yolo",
				Value: "Fomo",
			}},
			SearchIndex:          String("SearchIndex"),
			SearchSource:         String("SearchSource"),
			SearchRequestTimeout: 456,
			VerifyTLS:            false,
		},
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("Stores.Update returned %+v, want %+v", got, want)
	}
}

func TestStoresService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/stores/stores/eaabf092-caed-4b0e-a8d5-851205b2fa56/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.Stores.Delete(ctx, "eaabf092-caed-4b0e-a8d5-851205b2fa56")
	if err != nil {
		t.Errorf("Stores.Delete returned error: %v", err)
	}
}
