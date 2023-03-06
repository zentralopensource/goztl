package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const msmBasePath = "monolith/sub_manifests/"

// MonolithSubManifestsService is an interface for interfacing with the Monolith sub manifests
// endpoints of the Zentral API
type MonolithSubManifestsService interface {
	List(context.Context, *ListOptions) ([]MonolithSubManifest, *Response, error)
	GetByID(context.Context, int) (*MonolithSubManifest, *Response, error)
	GetByName(context.Context, string) (*MonolithSubManifest, *Response, error)
	Create(context.Context, *MonolithSubManifestRequest) (*MonolithSubManifest, *Response, error)
	Update(context.Context, int, *MonolithSubManifestRequest) (*MonolithSubManifest, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// MonolithSubManifestsServiceOp handles comsmunication with the Monolith sub manifests related
// methods of the Zentral API.
type MonolithSubManifestsServiceOp struct {
	client *Client
}

var _ MonolithSubManifestsService = &MonolithSubManifestsServiceOp{}

// MonolithSubManifest represents a Zentral MonolithSubManifest
type MonolithSubManifest struct {
	ID                 int       `json:"id"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	MetaBusinessUnitID *int      `json:"meta_business_unit"`
	Created            Timestamp `json:"created_at,omitempty"`
	Updated            Timestamp `json:"updated_at,omitempty"`
}

func (se MonolithSubManifest) String() string {
	return Stringify(se)
}

// MonolithSubManifestRequest represents a request to create or update a Monolith sub manifest
type MonolithSubManifestRequest struct {
	Name               string `json:"name"`
	Description        string `json:"description"`
	MetaBusinessUnitID *int   `json:"meta_business_unit"`
}

type listMSMOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the Monolith sub manifests.
func (s *MonolithSubManifestsServiceOp) List(ctx context.Context, opt *ListOptions) ([]MonolithSubManifest, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Monolith sub manifest by id.
func (s *MonolithSubManifestsServiceOp) GetByID(ctx context.Context, msmID int) (*MonolithSubManifest, *Response, error) {
	if msmID < 1 {
		return nil, nil, NewArgError("msmID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", msmBasePath, msmID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	msm := new(MonolithSubManifest)

	resp, err := s.client.Do(ctx, req, msm)
	if err != nil {
		return nil, resp, err
	}

	return msm, resp, err
}

// GetByName retrieves a Monolith sub manifest by name.
func (s *MonolithSubManifestsServiceOp) GetByName(ctx context.Context, name string) (*MonolithSubManifest, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listMSMOpt := &listMSMOptions{Name: name}

	msms, resp, err := s.list(ctx, nil, listMSMOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(msms) < 1 {
		return nil, resp, nil
	}

	return &msms[0], resp, err
}

// Create a new Monolith sub manifest.
func (s *MonolithSubManifestsServiceOp) Create(ctx context.Context, createRequest *MonolithSubManifestRequest) (*MonolithSubManifest, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, msmBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	msm := new(MonolithSubManifest)
	resp, err := s.client.Do(ctx, req, msm)
	if err != nil {
		return nil, resp, err
	}

	return msm, resp, err
}

// Update a Monolith sub manifest.
func (s *MonolithSubManifestsServiceOp) Update(ctx context.Context, msmID int, updateRequest *MonolithSubManifestRequest) (*MonolithSubManifest, *Response, error) {
	if msmID < 1 {
		return nil, nil, NewArgError("msmID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", msmBasePath, msmID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	msm := new(MonolithSubManifest)
	resp, err := s.client.Do(ctx, req, msm)
	if err != nil {
		return nil, resp, err
	}

	return msm, resp, err
}

// Delete a Monolith sub manifest.
func (s *MonolithSubManifestsServiceOp) Delete(ctx context.Context, msmID int) (*Response, error) {
	if msmID < 1 {
		return nil, NewArgError("msmID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", msmBasePath, msmID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Monolith sub manifests
func (s *MonolithSubManifestsServiceOp) list(ctx context.Context, opt *ListOptions, msmOpt *listMSMOptions) ([]MonolithSubManifest, *Response, error) {
	path := msmBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, msmOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var msms []MonolithSubManifest
	resp, err := s.client.Do(ctx, req, &msms)
	if err != nil {
		return nil, resp, err
	}

	return msms, resp, err
}
