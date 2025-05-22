package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const mdaBasePath = "mdm/data_assets/"

// MDMDataAssetsService is an interface for interfacing with the MDM data asset
// endpoints of the Zentral API
type MDMDataAssetsService interface {
	List(context.Context, *ListOptions) ([]MDMDataAsset, *Response, error)
	GetByID(context.Context, string) (*MDMDataAsset, *Response, error)
	Create(context.Context, *MDMDataAssetRequest) (*MDMDataAsset, *Response, error)
	Update(context.Context, string, *MDMDataAssetRequest) (*MDMDataAsset, *Response, error)
	Delete(context.Context, string) (*Response, error)
}

// MDMDataAssetsServiceOp handles communication with the MDM data assets related
// methods of the Zentral API.
type MDMDataAssetsServiceOp struct {
	client *Client
}

var _ MDMDataAssetsService = &MDMDataAssetsServiceOp{}

// MDMDataAsset represents a Zentral MDM data asset
type MDMDataAsset struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	FileURI    string `json:"file_uri"`
	FileSHA256 string `json:"file_sha256"`
	FileSize   int64  `json:"file_size"`
	Filename   string `json:"filename"`
	MDMArtifactVersion
}

func (mda MDMDataAsset) String() string {
	return Stringify(mda)
}

// MDMDataAssetRequest represents a request to create or update a MDM data asset
type MDMDataAssetRequest struct {
	Type       string `json:"type"`
	FileURI    string `json:"file_uri"`
	FileSHA256 string `json:"file_sha256"`
	MDMArtifactVersionRequest
}

type listMDAOptions struct{}

// List lists all the MDM data assets.
func (s *MDMDataAssetsServiceOp) List(ctx context.Context, opt *ListOptions) ([]MDMDataAsset, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a MDM data asset by id.
func (s *MDMDataAssetsServiceOp) GetByID(ctx context.Context, mdaID string) (*MDMDataAsset, *Response, error) {
	if len(mdaID) < 1 {
		return nil, nil, NewArgError("mdaID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", mdaBasePath, mdaID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	mda := new(MDMDataAsset)

	resp, err := s.client.Do(ctx, req, mda)
	if err != nil {
		return nil, resp, err
	}

	return mda, resp, err
}

// Create a new MDM data asset.
func (s *MDMDataAssetsServiceOp) Create(ctx context.Context, createRequest *MDMDataAssetRequest) (*MDMDataAsset, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, mdaBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	mda := new(MDMDataAsset)
	resp, err := s.client.Do(ctx, req, mda)
	if err != nil {
		return nil, resp, err
	}

	return mda, resp, err
}

// Update a MDM data asset.
func (s *MDMDataAssetsServiceOp) Update(ctx context.Context, mdaID string, updateRequest *MDMDataAssetRequest) (*MDMDataAsset, *Response, error) {
	if len(mdaID) < 1 {
		return nil, nil, NewArgError("mdaID", "cannot be blank")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%s/", mdaBasePath, mdaID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	mda := new(MDMDataAsset)
	resp, err := s.client.Do(ctx, req, mda)
	if err != nil {
		return nil, resp, err
	}

	return mda, resp, err
}

// Delete a MDM data asset.
func (s *MDMDataAssetsServiceOp) Delete(ctx context.Context, mdaID string) (*Response, error) {
	if len(mdaID) < 1 {
		return nil, NewArgError("mdaID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", mdaBasePath, mdaID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing MDM data assets
func (s *MDMDataAssetsServiceOp) list(ctx context.Context, opt *ListOptions, mdaOpt *listMDAOptions) ([]MDMDataAsset, *Response, error) {
	path := mdaBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, mdaOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var mdas []MDMDataAsset
	resp, err := s.client.Do(ctx, req, &mdas)
	if err != nil {
		return nil, resp, err
	}

	return mdas, resp, err
}
