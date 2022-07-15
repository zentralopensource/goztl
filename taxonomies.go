package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const TaxonomyBasePath = "inventory/taxonomies/"

// TaxonomiesService is an interface for interfacing with the Taxonomies
// endpoints of the Zentral API
type TaxonomiesService interface {
	List(context.Context, *ListOptions) ([]Taxonomy, *Response, error)
	GetByID(context.Context, int) (*Taxonomy, *Response, error)
	GetByName(context.Context, string) (*Taxonomy, *Response, error)
	Create(context.Context, *TaxonomyCreateRequest) (*Taxonomy, *Response, error)
	Update(context.Context, int, *TaxonomyUpdateRequest) (*Taxonomy, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// TaxonomiesServiceOp handles communication with the Taxonomies related
// methods of the Zentral API.
type TaxonomiesServiceOp struct {
	client *Client
}

var _ TaxonomiesService = &TaxonomiesServiceOp{}

// Taxonomy represents a Zentral Taxonomy
type Taxonomy struct {
	ID                 int       `json:"id,omitempty"`
	MetaBusinessUnitID int       `json:"meta_business_unit,omitempty"`
	Name               string    `json:"name,omitempty"`
	Created            Timestamp `json:"created_at,omitempty"`
	Updated            Timestamp `json:"updated_at,omitempty"`
}

// TaxonomyCreateRequest represents a request to create a Taxonomy.
type TaxonomyCreateRequest struct {
	Name               string `json:"name"`
	MetaBusinessUnitID int    `json:"meta_business_unit,omitempty"`
}

// TaxonomyUpdateRequest represents a request to create a Taxonomy.
type TaxonomyUpdateRequest struct {
	Name               string `json:"name"`
	MetaBusinessUnitID int    `json:"meta_business_unit,omitempty"`
}

func (Taxonomy Taxonomy) String() string {
	return Stringify(Taxonomy)
}

type listTaxonomyOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the Taxonomies.
func (s *TaxonomiesServiceOp) List(ctx context.Context, opt *ListOptions) ([]Taxonomy, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Taxonomy by id.
func (s *TaxonomiesServiceOp) GetByID(ctx context.Context, TaxonomyID int) (*Taxonomy, *Response, error) {
	if TaxonomyID < 1 {
		return nil, nil, NewArgError("TaxonomyID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", TaxonomyBasePath, TaxonomyID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	Taxonomy := new(Taxonomy)

	resp, err := s.client.Do(ctx, req, Taxonomy)
	if err != nil {
		return nil, resp, err
	}

	return Taxonomy, resp, err
}

// GetByName retrieves a Taxonomy by name.
func (s *TaxonomiesServiceOp) GetByName(ctx context.Context, name string) (*Taxonomy, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listTaxonomyOpt := &listTaxonomyOptions{Name: name}

	Taxonomies, resp, err := s.list(ctx, nil, listTaxonomyOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(Taxonomies) < 1 {
		return nil, resp, nil
	}

	return &Taxonomies[0], resp, err
}

// Create a new Taxonomy.
func (s *TaxonomiesServiceOp) Create(ctx context.Context, createRequest *TaxonomyCreateRequest) (*Taxonomy, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, TaxonomyBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	Taxonomy := new(Taxonomy)
	resp, err := s.client.Do(ctx, req, Taxonomy)
	if err != nil {
		return nil, resp, err
	}

	return Taxonomy, resp, err
}

// Update a Taxonomy.
func (s *TaxonomiesServiceOp) Update(ctx context.Context, TaxonomyID int, updateRequest *TaxonomyUpdateRequest) (*Taxonomy, *Response, error) {
	if TaxonomyID < 1 {
		return nil, nil, NewArgError("TaxonomyID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", TaxonomyBasePath, TaxonomyID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	Taxonomy := new(Taxonomy)
	resp, err := s.client.Do(ctx, req, Taxonomy)
	if err != nil {
		return nil, resp, err
	}

	return Taxonomy, resp, err
}

// Delete a Taxonomy.
func (s *TaxonomiesServiceOp) Delete(ctx context.Context, TaxonomyID int) (*Response, error) {
	if TaxonomyID < 1 {
		return nil, NewArgError("TaxonomyID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", TaxonomyBasePath, TaxonomyID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Taxonomies
func (s *TaxonomiesServiceOp) list(ctx context.Context, opt *ListOptions, TaxonomyOpt *listTaxonomyOptions) ([]Taxonomy, *Response, error) {
	path := TaxonomyBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, TaxonomyOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var Taxonomies []Taxonomy
	resp, err := s.client.Do(ctx, req, &Taxonomies)
	if err != nil {
		return nil, resp, err
	}

	return Taxonomies, resp, err
}
