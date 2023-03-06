package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const mmBasePath = "monolith/manifests/"

// MonolithManifestsService is an interface for interfacing with the Monolith manifests
// endpoints of the Zentral API
type MonolithManifestsService interface {
	List(context.Context, *ListOptions) ([]MonolithManifest, *Response, error)
	GetByID(context.Context, int) (*MonolithManifest, *Response, error)
	GetByName(context.Context, string) (*MonolithManifest, *Response, error)
	Create(context.Context, *MonolithManifestRequest) (*MonolithManifest, *Response, error)
	Update(context.Context, int, *MonolithManifestRequest) (*MonolithManifest, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// MonolithManifestsServiceOp handles communication with the Monolith manifests related
// methods of the Zentral API.
type MonolithManifestsServiceOp struct {
	client *Client
}

var _ MonolithManifestsService = &MonolithManifestsServiceOp{}

// MonolithManifest represents a Zentral MonolithManifest
type MonolithManifest struct {
	ID                 int       `json:"id"`
	Name               string    `json:"name"`
	MetaBusinessUnitID int       `json:"meta_business_unit"`
	Version            int       `json:"version"`
	Created            Timestamp `json:"created_at,omitempty"`
	Updated            Timestamp `json:"updated_at,omitempty"`
}

func (se MonolithManifest) String() string {
	return Stringify(se)
}

// MonolithManifestRequest represents a request to create or update a Monolith manifest
type MonolithManifestRequest struct {
	Name               string `json:"name"`
	MetaBusinessUnitID int    `json:"meta_business_unit"`
}

type listMMOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the Monolith manifests.
func (s *MonolithManifestsServiceOp) List(ctx context.Context, opt *ListOptions) ([]MonolithManifest, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Monolith manifest by id.
func (s *MonolithManifestsServiceOp) GetByID(ctx context.Context, mmID int) (*MonolithManifest, *Response, error) {
	if mmID < 1 {
		return nil, nil, NewArgError("mmID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mmBasePath, mmID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	mm := new(MonolithManifest)

	resp, err := s.client.Do(ctx, req, mm)
	if err != nil {
		return nil, resp, err
	}

	return mm, resp, err
}

// GetByName retrieves a Monolith manifest by name.
func (s *MonolithManifestsServiceOp) GetByName(ctx context.Context, name string) (*MonolithManifest, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listMMOpt := &listMMOptions{Name: name}

	mms, resp, err := s.list(ctx, nil, listMMOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(mms) < 1 {
		return nil, resp, nil
	}

	return &mms[0], resp, err
}

// Create a new Monolith manifest.
func (s *MonolithManifestsServiceOp) Create(ctx context.Context, createRequest *MonolithManifestRequest) (*MonolithManifest, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, mmBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	mm := new(MonolithManifest)
	resp, err := s.client.Do(ctx, req, mm)
	if err != nil {
		return nil, resp, err
	}

	return mm, resp, err
}

// Update a Monolith manifest.
func (s *MonolithManifestsServiceOp) Update(ctx context.Context, mmID int, updateRequest *MonolithManifestRequest) (*MonolithManifest, *Response, error) {
	if mmID < 1 {
		return nil, nil, NewArgError("mmID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", mmBasePath, mmID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	mm := new(MonolithManifest)
	resp, err := s.client.Do(ctx, req, mm)
	if err != nil {
		return nil, resp, err
	}

	return mm, resp, err
}

// Delete a Monolith manifest.
func (s *MonolithManifestsServiceOp) Delete(ctx context.Context, mmID int) (*Response, error) {
	if mmID < 1 {
		return nil, NewArgError("mmID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mmBasePath, mmID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Monolith manifests
func (s *MonolithManifestsServiceOp) list(ctx context.Context, opt *ListOptions, mmOpt *listMMOptions) ([]MonolithManifest, *Response, error) {
	path := mmBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, mmOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var mms []MonolithManifest
	resp, err := s.client.Do(ctx, req, &mms)
	if err != nil {
		return nil, resp, err
	}

	return mms, resp, err
}
