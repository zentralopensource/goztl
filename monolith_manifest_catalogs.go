package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const mmcBasePath = "monolith/manifest_catalogs/"

// MonolithManifestCatalogsService is an interface for interfacing with the Monolith manifest catalogs
// endpoints of the Zentral API
type MonolithManifestCatalogsService interface {
	List(context.Context, *ListOptions) ([]MonolithManifestCatalog, *Response, error)
	GetByID(context.Context, int) (*MonolithManifestCatalog, *Response, error)
	GetByCatalogID(context.Context, int) ([]MonolithManifestCatalog, *Response, error)
	GetByManifestID(context.Context, int) ([]MonolithManifestCatalog, *Response, error)
	Create(context.Context, *MonolithManifestCatalogRequest) (*MonolithManifestCatalog, *Response, error)
	Update(context.Context, int, *MonolithManifestCatalogRequest) (*MonolithManifestCatalog, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// MonolithManifestCatalogsServiceOp handles commcunication with the Monolith manifest catalogs related
// methods of the Zentral API.
type MonolithManifestCatalogsServiceOp struct {
	client *Client
}

var _ MonolithManifestCatalogsService = &MonolithManifestCatalogsServiceOp{}

// MonolithManifestCatalog represents a Zentral manifest catalog.
type MonolithManifestCatalog struct {
	ID         int   `json:"id"`
	ManifestID int   `json:"manifest"`
	CatalogID  int   `json:"catalog"`
	TagIDs     []int `json:"tags"`
}

func (se MonolithManifestCatalog) String() string {
	return Stringify(se)
}

// MonolithManifestCatalogRequest represents a request to create or update a Monolith manifest catalog.
type MonolithManifestCatalogRequest struct {
	ManifestID int   `json:"manifest"`
	CatalogID  int   `json:"catalog"`
	TagIDs     []int `json:"tags"`
}

type listMMCOptions struct {
	CatalogID  int `url:"catalog_id,omitempty"`
	ManifestID int `url:"manifest_id,omitempty"`
}

// List lists all the Monolith manifest catalogs.
func (s *MonolithManifestCatalogsServiceOp) List(ctx context.Context, opt *ListOptions) ([]MonolithManifestCatalog, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Monolith manifest catalog by id.
func (s *MonolithManifestCatalogsServiceOp) GetByID(ctx context.Context, mmcID int) (*MonolithManifestCatalog, *Response, error) {
	if mmcID < 1 {
		return nil, nil, NewArgError("mmcID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mmcBasePath, mmcID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	mmc := new(MonolithManifestCatalog)

	resp, err := s.client.Do(ctx, req, mmc)
	if err != nil {
		return nil, resp, err
	}

	return mmc, resp, err
}

// GetByCatalogID retrieves Monolith manifest catalogs by catalog ID.
func (s *MonolithManifestCatalogsServiceOp) GetByCatalogID(ctx context.Context, mcID int) ([]MonolithManifestCatalog, *Response, error) {
	if mcID < 1 {
		return nil, nil, NewArgError("mcID", "cannot be < 1")
	}

	listMMCOpt := &listMMCOptions{CatalogID: mcID}

	return s.list(ctx, nil, listMMCOpt)
}

// GetByManifestID retrieves Monolith manifest catalogs by manifest ID.
func (s *MonolithManifestCatalogsServiceOp) GetByManifestID(ctx context.Context, mmID int) ([]MonolithManifestCatalog, *Response, error) {
	if mmID < 1 {
		return nil, nil, NewArgError("mmID", "cannot be < 1")
	}

	listMMCOpt := &listMMCOptions{ManifestID: mmID}

	return s.list(ctx, nil, listMMCOpt)
}

// Create a new Monolith manifest catalog.
func (s *MonolithManifestCatalogsServiceOp) Create(ctx context.Context, createRequest *MonolithManifestCatalogRequest) (*MonolithManifestCatalog, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, mmcBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	mmc := new(MonolithManifestCatalog)
	resp, err := s.client.Do(ctx, req, mmc)
	if err != nil {
		return nil, resp, err
	}

	return mmc, resp, err
}

// Update a Monolith manifest catalog.
func (s *MonolithManifestCatalogsServiceOp) Update(ctx context.Context, mmcID int, updateRequest *MonolithManifestCatalogRequest) (*MonolithManifestCatalog, *Response, error) {
	if mmcID < 1 {
		return nil, nil, NewArgError("mmcID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", mmcBasePath, mmcID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	mmc := new(MonolithManifestCatalog)
	resp, err := s.client.Do(ctx, req, mmc)
	if err != nil {
		return nil, resp, err
	}

	return mmc, resp, err
}

// Delete a Monolith manifest catalog.
func (s *MonolithManifestCatalogsServiceOp) Delete(ctx context.Context, mmcID int) (*Response, error) {
	if mmcID < 1 {
		return nil, NewArgError("mmcID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mmcBasePath, mmcID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Monolith manifest catalogs.
func (s *MonolithManifestCatalogsServiceOp) list(ctx context.Context, opt *ListOptions, mmcOpt *listMMCOptions) ([]MonolithManifestCatalog, *Response, error) {
	path := mmcBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, mmcOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var mmcs []MonolithManifestCatalog
	resp, err := s.client.Do(ctx, req, &mmcs)
	if err != nil {
		return nil, resp, err
	}

	return mmcs, resp, err
}
