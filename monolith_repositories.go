package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const mrBasePath = "monolith/repositories/"

// MonolithRepositoriesService is an interface for interfacing with the Monolith manifests
// endpoints of the Zentral API
type MonolithRepositoriesService interface {
	List(context.Context, *ListOptions) ([]MonolithRepository, *Response, error)
	GetByID(context.Context, int) (*MonolithRepository, *Response, error)
	GetByName(context.Context, string) (*MonolithRepository, *Response, error)
	Create(context.Context, *MonolithRepositoryRequest) (*MonolithRepository, *Response, error)
	Update(context.Context, int, *MonolithRepositoryRequest) (*MonolithRepository, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// MonolithRepositoriesServiceOp handles comrunication with the Monolith manifests related
// methods of the Zentral API.
type MonolithRepositoriesServiceOp struct {
	client *Client
}

var _ MonolithRepositoriesService = &MonolithRepositoriesServiceOp{}

// MonolithRepository represents a Zentral MonolithRepository

type MonolithS3Backend struct {
	Bucket               string `json:"bucket"`
	RegionName           string `json:"region_name"`
	Prefix               string `json:"prefix"`
	AccessKeyID          string `json:"access_key_id"`
	SecretAccessKey      string `json:"secret_access_key"`
	AssumeRoleARN        string `json:"assume_role_arn"`
	SignatureVersion     string `json:"signature_version"`
	EndpointURL          string `json:"endpoint_url"`
	CloudfrontDomain     string `json:"cloudfront_domain"`
	CloudfrontKeyID      string `json:"cloudfront_key_id"`
	CloudfrontPrivkeyPEM string `json:"cloudfront_privkey_pem"`
}

type MonolithAzureBackend struct {
	StorageAccount string `json:"storage_account"`
	Container      string `json:"container"`
	Prefix         string `json:"prefix"`
	ClientID       string `json:"client_id"`
	TenantID       string `json:"tenant_id"`
	ClientSecret   string `json:"client_secret"`
}

type MonolithRepository struct {
	ID                 int                   `json:"id"`
	Name               string                `json:"name"`
	MetaBusinessUnitID *int                  `json:"meta_business_unit"`
	Backend            string                `json:"backend"`
	Azure              *MonolithAzureBackend `json:"azure_kwargs"`
	S3                 *MonolithS3Backend    `json:"s3_kwargs"`
	Created            Timestamp             `json:"created_at"`
	Updated            Timestamp             `json:"updated_at"`
}

func (mr MonolithRepository) String() string {
	return Stringify(mr)
}

// MonolithRepositoryRequest represents a request to create or update a Monolith manifest
type MonolithRepositoryRequest struct {
	Name               string                `json:"name"`
	MetaBusinessUnitID *int                  `json:"meta_business_unit"`
	Backend            string                `json:"backend"`
	Azure              *MonolithAzureBackend `json:"azure_kwargs,omitempty"`
	S3                 *MonolithS3Backend    `json:"s3_kwargs,omitempty"`
}

type listMROptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the Monolith manifests.
func (s *MonolithRepositoriesServiceOp) List(ctx context.Context, opt *ListOptions) ([]MonolithRepository, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Monolith manifest by id.
func (s *MonolithRepositoriesServiceOp) GetByID(ctx context.Context, mrID int) (*MonolithRepository, *Response, error) {
	if mrID < 1 {
		return nil, nil, NewArgError("mrID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mrBasePath, mrID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	mr := new(MonolithRepository)

	resp, err := s.client.Do(ctx, req, mr)
	if err != nil {
		return nil, resp, err
	}

	return mr, resp, err
}

// GetByName retrieves a Monolith manifest by name.
func (s *MonolithRepositoriesServiceOp) GetByName(ctx context.Context, name string) (*MonolithRepository, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listMROpt := &listMROptions{Name: name}

	mrs, resp, err := s.list(ctx, nil, listMROpt)
	if err != nil {
		return nil, resp, err
	}
	if len(mrs) < 1 {
		return nil, resp, nil
	}

	return &mrs[0], resp, err
}

// Create a new Monolith manifest.
func (s *MonolithRepositoriesServiceOp) Create(ctx context.Context, createRequest *MonolithRepositoryRequest) (*MonolithRepository, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, mrBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	mr := new(MonolithRepository)
	resp, err := s.client.Do(ctx, req, mr)
	if err != nil {
		return nil, resp, err
	}

	return mr, resp, err
}

// Update a Monolith manifest.
func (s *MonolithRepositoriesServiceOp) Update(ctx context.Context, mrID int, updateRequest *MonolithRepositoryRequest) (*MonolithRepository, *Response, error) {
	if mrID < 1 {
		return nil, nil, NewArgError("mrID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", mrBasePath, mrID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	mr := new(MonolithRepository)
	resp, err := s.client.Do(ctx, req, mr)
	if err != nil {
		return nil, resp, err
	}

	return mr, resp, err
}

// Delete a Monolith manifest.
func (s *MonolithRepositoriesServiceOp) Delete(ctx context.Context, mrID int) (*Response, error) {
	if mrID < 1 {
		return nil, NewArgError("mrID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mrBasePath, mrID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Monolith manifests
func (s *MonolithRepositoriesServiceOp) list(ctx context.Context, opt *ListOptions, mrOpt *listMROptions) ([]MonolithRepository, *Response, error) {
	path := mrBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, mrOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var mrs []MonolithRepository
	resp, err := s.client.Do(ctx, req, &mrs)
	if err != nil {
		return nil, resp, err
	}

	return mrs, resp, err
}
