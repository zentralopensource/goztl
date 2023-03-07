package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const mmsmBasePath = "monolith/manifest_sub_manifests/"

// MonolithManifestSubManifestsService is an interface for interfacing with the Monolith manifest sub manifests
// endpoints of the Zentral API
type MonolithManifestSubManifestsService interface {
	List(context.Context, *ListOptions) ([]MonolithManifestSubManifest, *Response, error)
	GetByID(context.Context, int) (*MonolithManifestSubManifest, *Response, error)
	GetByManifestID(context.Context, int) ([]MonolithManifestSubManifest, *Response, error)
	GetBySubManifestID(context.Context, int) ([]MonolithManifestSubManifest, *Response, error)
	Create(context.Context, *MonolithManifestSubManifestRequest) (*MonolithManifestSubManifest, *Response, error)
	Update(context.Context, int, *MonolithManifestSubManifestRequest) (*MonolithManifestSubManifest, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// MonolithManifestSubManifestsServiceOp handles comsmunication with the Monolith manifest sub manifests related
// methods of the Zentral API.
type MonolithManifestSubManifestsServiceOp struct {
	client *Client
}

var _ MonolithManifestSubManifestsService = &MonolithManifestSubManifestsServiceOp{}

// MonolithManifestSubManifest represents a Zentral manifest sub manifest.
type MonolithManifestSubManifest struct {
	ID            int   `json:"id"`
	ManifestID    int   `json:"manifest"`
	SubManifestID int   `json:"sub_manifest"`
	TagIDs        []int `json:"tags"`
}

func (se MonolithManifestSubManifest) String() string {
	return Stringify(se)
}

// MonolithManifestSubManifestRequest represents a request to create or update a Monolith manifest sub manifest.
type MonolithManifestSubManifestRequest struct {
	ManifestID    int   `json:"manifest"`
	SubManifestID int   `json:"sub_manifest"`
	TagIDs        []int `json:"tags"`
}

type listMMSMOptions struct {
	SubManifestID int `url:"sub_manifest_id,omitempty"`
	ManifestID    int `url:"manifest_id,omitempty"`
}

// List lists all the Monolith manifest sub manifests.
func (s *MonolithManifestSubManifestsServiceOp) List(ctx context.Context, opt *ListOptions) ([]MonolithManifestSubManifest, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Monolith manifest sub manifest by id.
func (s *MonolithManifestSubManifestsServiceOp) GetByID(ctx context.Context, msmID int) (*MonolithManifestSubManifest, *Response, error) {
	if msmID < 1 {
		return nil, nil, NewArgError("msmID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mmsmBasePath, msmID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	msm := new(MonolithManifestSubManifest)

	resp, err := s.client.Do(ctx, req, msm)
	if err != nil {
		return nil, resp, err
	}

	return msm, resp, err
}

// GetBySubManifestID retrieves Monolith manifest sub manifests by sub manifest ID.
func (s *MonolithManifestSubManifestsServiceOp) GetBySubManifestID(ctx context.Context, msmID int) ([]MonolithManifestSubManifest, *Response, error) {
	if msmID < 1 {
		return nil, nil, NewArgError("msmID", "cannot be < 1")
	}

	listMMSMOpt := &listMMSMOptions{SubManifestID: msmID}

	return s.list(ctx, nil, listMMSMOpt)
}

// GetByManifestID retrieves Monolith manifest sub manifests by manifest ID.
func (s *MonolithManifestSubManifestsServiceOp) GetByManifestID(ctx context.Context, mID int) ([]MonolithManifestSubManifest, *Response, error) {
	if mID < 1 {
		return nil, nil, NewArgError("mID", "cannot be < 1")
	}

	listMMSMOpt := &listMMSMOptions{ManifestID: mID}

	return s.list(ctx, nil, listMMSMOpt)
}

// Create a new Monolith manifest sub manifest.
func (s *MonolithManifestSubManifestsServiceOp) Create(ctx context.Context, createRequest *MonolithManifestSubManifestRequest) (*MonolithManifestSubManifest, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, mmsmBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	msm := new(MonolithManifestSubManifest)
	resp, err := s.client.Do(ctx, req, msm)
	if err != nil {
		return nil, resp, err
	}

	return msm, resp, err
}

// Update a Monolith manifest sub manifest.
func (s *MonolithManifestSubManifestsServiceOp) Update(ctx context.Context, msmID int, updateRequest *MonolithManifestSubManifestRequest) (*MonolithManifestSubManifest, *Response, error) {
	if msmID < 1 {
		return nil, nil, NewArgError("msmID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", mmsmBasePath, msmID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	msm := new(MonolithManifestSubManifest)
	resp, err := s.client.Do(ctx, req, msm)
	if err != nil {
		return nil, resp, err
	}

	return msm, resp, err
}

// Delete a Monolith manifest sub manifest.
func (s *MonolithManifestSubManifestsServiceOp) Delete(ctx context.Context, msmID int) (*Response, error) {
	if msmID < 1 {
		return nil, NewArgError("msmID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mmsmBasePath, msmID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Monolith manifest sub manifests.
func (s *MonolithManifestSubManifestsServiceOp) list(ctx context.Context, opt *ListOptions, msmOpt *listMMSMOptions) ([]MonolithManifestSubManifest, *Response, error) {
	path := mmsmBasePath
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

	var msms []MonolithManifestSubManifest
	resp, err := s.client.Do(ctx, req, &msms)
	if err != nil {
		return nil, resp, err
	}

	return msms, resp, err
}
