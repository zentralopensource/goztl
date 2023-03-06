package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const mcBasePath = "monolith/catalogs/"

// MonolithCatalogsService is an interface for interfacing with the Monolith catalogs
// endpoints of the Zentral API
type MonolithCatalogsService interface {
	List(context.Context, *ListOptions) ([]MonolithCatalog, *Response, error)
	GetByID(context.Context, int) (*MonolithCatalog, *Response, error)
	GetByName(context.Context, string) (*MonolithCatalog, *Response, error)
	Create(context.Context, *MonolithCatalogRequest) (*MonolithCatalog, *Response, error)
	Update(context.Context, int, *MonolithCatalogRequest) (*MonolithCatalog, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// MonolithCatalogsServiceOp handles comcunication with the Monolith catalogs related
// methods of the Zentral API.
type MonolithCatalogsServiceOp struct {
	client *Client
}

var _ MonolithCatalogsService = &MonolithCatalogsServiceOp{}

// MonolithCatalog represents a Zentral MonolithCatalog
type MonolithCatalog struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Priority   int        `json:"priority"`
	Created    Timestamp  `json:"created_at"`
	Updated    Timestamp  `json:"updated_at"`
	ArchivedAt *Timestamp `json:"archived_at"`
}

func (se MonolithCatalog) String() string {
	return Stringify(se)
}

// MonolithCatalogRequest represents a request to create or update a Monolith catalog
type MonolithCatalogRequest struct {
	Name     string `json:"name"`
	Priority int    `json:"priority"`
}

type listMCOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the Monolith catalogs.
func (s *MonolithCatalogsServiceOp) List(ctx context.Context, opt *ListOptions) ([]MonolithCatalog, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Monolith catalog by id.
func (s *MonolithCatalogsServiceOp) GetByID(ctx context.Context, mcID int) (*MonolithCatalog, *Response, error) {
	if mcID < 1 {
		return nil, nil, NewArgError("mcID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mcBasePath, mcID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	mc := new(MonolithCatalog)

	resp, err := s.client.Do(ctx, req, mc)
	if err != nil {
		return nil, resp, err
	}

	return mc, resp, err
}

// GetByName retrieves a Monolith catalog by name.
func (s *MonolithCatalogsServiceOp) GetByName(ctx context.Context, name string) (*MonolithCatalog, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listMCOpt := &listMCOptions{Name: name}

	mcs, resp, err := s.list(ctx, nil, listMCOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(mcs) < 1 {
		return nil, resp, nil
	}

	return &mcs[0], resp, err
}

// Create a new Monolith catalog.
func (s *MonolithCatalogsServiceOp) Create(ctx context.Context, createRequest *MonolithCatalogRequest) (*MonolithCatalog, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, mcBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	mc := new(MonolithCatalog)
	resp, err := s.client.Do(ctx, req, mc)
	if err != nil {
		return nil, resp, err
	}

	return mc, resp, err
}

// Update a Monolith catalog.
func (s *MonolithCatalogsServiceOp) Update(ctx context.Context, mcID int, updateRequest *MonolithCatalogRequest) (*MonolithCatalog, *Response, error) {
	if mcID < 1 {
		return nil, nil, NewArgError("mcID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", mcBasePath, mcID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	mc := new(MonolithCatalog)
	resp, err := s.client.Do(ctx, req, mc)
	if err != nil {
		return nil, resp, err
	}

	return mc, resp, err
}

// Delete a Monolith catalog.
func (s *MonolithCatalogsServiceOp) Delete(ctx context.Context, mcID int) (*Response, error) {
	if mcID < 1 {
		return nil, NewArgError("mcID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mcBasePath, mcID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Monolith catalogs
func (s *MonolithCatalogsServiceOp) list(ctx context.Context, opt *ListOptions, mcOpt *listMCOptions) ([]MonolithCatalog, *Response, error) {
	path := mcBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, mcOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var mcs []MonolithCatalog
	resp, err := s.client.Do(ctx, req, &mcs)
	if err != nil {
		return nil, resp, err
	}

	return mcs, resp, err
}
