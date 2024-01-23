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

var mrListJSONResponse = `
[
  {
    "id": 4,
    "name": "Default",
    "backend": "VIRTUAL",
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
  }
]
`

var mrGetJSONResponse = `
{
  "id": 4,
  "meta_business_unit": 3,
  "name": "Default",
  "backend": "S3",
  "s3_kwargs": {
    "bucket": "bucket",
    "region_name": "us-east-1",
    "prefix": "prefix",
    "access_key_id": "11111111111111111111",
    "secret_access_key": "22222222222222222222",
    "assume_role_arn": "arn:aws:iam::123456789012:role/S3Access",
    "signature_version": "s3v4",
    "endpoint_url": "https://endpoint.example.com",
    "cloudfront_domain": "yolo.cloudfront.net",
    "cloudfront_key_id": "YOLO",
    "cloudfront_privkey_pem": "NOT A PEM KEY"
  },
  "created_at": "2022-07-22T01:02:03.444444",
  "updated_at": "2022-07-22T01:02:03.444444"
}
`

var mrCreateJSONResponse = `
{
  "id": 4,
  "name": "Default",
  "backend": "S3",
  "s3_kwargs": {
    "bucket": "bucket"
  },
  "created_at": "2022-07-22T01:02:03.444444",
  "updated_at": "2022-07-22T01:02:03.444444"
}
`

var mrUpdateJSONResponse = `
{
  "id": 4,
  "meta_business_unit": 3,
  "name": "Default",
  "backend": "S3",
  "s3_kwargs": {
    "bucket": "bucket",
    "region_name": "us-east-1",
    "prefix": "prefix",
    "access_key_id": "11111111111111111111",
    "secret_access_key": "22222222222222222222",
    "assume_role_arn": "arn:aws:iam::123456789012:role/S3Access",
    "signature_version": "s3v4",
    "endpoint_url": "https://endpoint.example.com",
    "cloudfront_domain": "yolo.cloudfront.net",
    "cloudfront_key_id": "YOLO",
    "cloudfront_privkey_pem": "NOT A PEM KEY"
  },
  "created_at": "2022-07-22T01:02:03.444444",
  "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestMonolithRepositoriesService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/repositories/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mrListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithRepositories.List(ctx, nil)
	if err != nil {
		t.Errorf("MonolithRepositories.List returned error: %v", err)
	}

	want := []MonolithRepository{
		{
			ID:      4,
			Name:    "Default",
			Backend: "VIRTUAL",
			Created: Timestamp{referenceTime},
			Updated: Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithRepositories.List returned %+v, want %+v", got, want)
	}
}

func TestMonolithRepositoriesService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/repositories/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, mrGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithRepositories.GetByID(ctx, 4)
	if err != nil {
		t.Errorf("MonolithRepositories.GetByID returned error: %v", err)
	}

	want := &MonolithRepository{
		ID:                 4,
		MetaBusinessUnitID: Int(3),
		Name:               "Default",
		Backend:            "S3",
		S3: &MonolithS3Backend{
			Bucket:               "bucket",
			RegionName:           "us-east-1",
			Prefix:               "prefix",
			AccessKeyID:          "11111111111111111111",
			SecretAccessKey:      "22222222222222222222",
			AssumeRoleARN:        "arn:aws:iam::123456789012:role/S3Access",
			SignatureVersion:     "s3v4",
			EndpointURL:          "https://endpoint.example.com",
			CloudfrontDomain:     "yolo.cloudfront.net",
			CloudfrontKeyID:      "YOLO",
			CloudfrontPrivkeyPEM: "NOT A PEM KEY",
		},
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithRepositories.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMonolithRepositoriesService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/repositories/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "Default")
		fmt.Fprint(w, mrListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithRepositories.GetByName(ctx, "Default")
	if err != nil {
		t.Errorf("MonolithRepositories.GetByName returned error: %v", err)
	}

	want := &MonolithRepository{
		ID:      4,
		Name:    "Default",
		Backend: "VIRTUAL",
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithRepositories.GetByName returned %+v, want %+v", got, want)
	}
}

func TestMonolithRepositoriesService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &MonolithRepositoryRequest{
		Name:    "Default",
		Backend: "S3",
		S3: &MonolithS3Backend{
			Bucket: "bucket",
		},
	}

	mux.HandleFunc("/monolith/repositories/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MonolithRepositoryRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, mrCreateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithRepositories.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("MonolithRepositories.Create returned error: %v", err)
	}

	want := &MonolithRepository{
		ID:      4,
		Name:    "Default",
		Backend: "S3",
		S3: &MonolithS3Backend{
			Bucket: "bucket",
		},
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithRepositories.Create returned %+v, want %+v", got, want)
	}
}

func TestMonolithRepositoriesService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &MonolithRepositoryRequest{
		Name:    "Default",
		Backend: "S3",
		S3: &MonolithS3Backend{
			Bucket:               "bucket",
			RegionName:           "us-east-1",
			Prefix:               "prefix",
			AccessKeyID:          "11111111111111111111",
			SecretAccessKey:      "22222222222222222222",
			AssumeRoleARN:        "arn:aws:iam::123456789012:role/S3Access",
			SignatureVersion:     "s3v4",
			EndpointURL:          "https://endpoint.example.com",
			CloudfrontDomain:     "yolo.cloudfront.net",
			CloudfrontKeyID:      "YOLO",
			CloudfrontPrivkeyPEM: "NOT A PEM KEY",
		},
	}

	mux.HandleFunc("/monolith/repositories/4/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MonolithRepositoryRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, mrUpdateJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.MonolithRepositories.Update(ctx, 4, updateRequest)
	if err != nil {
		t.Errorf("MonolithRepositories.Update returned error: %v", err)
	}

	want := &MonolithRepository{
		ID:                 4,
		MetaBusinessUnitID: Int(3),
		Name:               "Default",
		Backend:            "S3",
		S3: &MonolithS3Backend{
			Bucket:               "bucket",
			RegionName:           "us-east-1",
			Prefix:               "prefix",
			AccessKeyID:          "11111111111111111111",
			SecretAccessKey:      "22222222222222222222",
			AssumeRoleARN:        "arn:aws:iam::123456789012:role/S3Access",
			SignatureVersion:     "s3v4",
			EndpointURL:          "https://endpoint.example.com",
			CloudfrontDomain:     "yolo.cloudfront.net",
			CloudfrontKeyID:      "YOLO",
			CloudfrontPrivkeyPEM: "NOT A PEM KEY",
		},
		Created: Timestamp{referenceTime},
		Updated: Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MonolithRepositories.Update returned %+v, want %+v", got, want)
	}
}

func TestMonolithRepositoriesService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/monolith/repositories/4/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MonolithRepositories.Delete(ctx, 4)
	if err != nil {
		t.Errorf("MonolithRepositories.Delete returned error: %v", err)
	}
}
