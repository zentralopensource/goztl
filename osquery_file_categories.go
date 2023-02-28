package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const ofcBasePath = "osquery/file_categories/"

// OsqueryFileCategoriesService is an interface for interfacing with the Osquery file categories
// endpoints of the Zentral API
type OsqueryFileCategoriesService interface {
	List(context.Context, *ListOptions) ([]OsqueryFileCategory, *Response, error)
	GetByID(context.Context, int) (*OsqueryFileCategory, *Response, error)
	GetByName(context.Context, string) (*OsqueryFileCategory, *Response, error)
	Create(context.Context, *OsqueryFileCategoryRequest) (*OsqueryFileCategory, *Response, error)
	Update(context.Context, int, *OsqueryFileCategoryRequest) (*OsqueryFileCategory, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// OsqueryFileCategoriesServiceOp handles communication with the Osquery file categories related
// methods of the Zentral API.
type OsqueryFileCategoriesServiceOp struct {
	client *Client
}

var _ OsqueryFileCategoriesService = &OsqueryFileCategoriesServiceOp{}

// OsqueryFileCategory represents a Zentral Osquery file category
type OsqueryFileCategory struct {
	ID               int       `json:"id,omitempty"`
	Name             string    `json:"name"`
	Slug             string    `json:"slug"`
	Description      string    `json:"description"`
	FilePaths        []string  `json:"file_paths"`
	ExcludePaths     []string  `json:"exclude_paths"`
	FilePathsQueries []string  `json:"file_paths_queries"`
	AccessMonitoring bool      `json:"access_monitoring"`
	Created          Timestamp `json:"created_at"`
	Updated          Timestamp `json:"updated_at"`
}

func (ofc OsqueryFileCategory) String() string {
	return Stringify(ofc)
}

// OsqueryFileCategoryRequest represents a request to create or update a Osquery file category
type OsqueryFileCategoryRequest struct {
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	FilePaths        []string `json:"file_paths"`
	ExcludePaths     []string `json:"exclude_paths"`
	FilePathsQueries []string `json:"file_paths_queries"`
	AccessMonitoring bool     `json:"access_monitoring"`
}

type listOFCOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the Osquery file categories.
func (s *OsqueryFileCategoriesServiceOp) List(ctx context.Context, opt *ListOptions) ([]OsqueryFileCategory, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Osquery file category by id.
func (s *OsqueryFileCategoriesServiceOp) GetByID(ctx context.Context, ofcID int) (*OsqueryFileCategory, *Response, error) {
	if ofcID < 1 {
		return nil, nil, NewArgError("ofcID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", ofcBasePath, ofcID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	ofc := new(OsqueryFileCategory)

	resp, err := s.client.Do(ctx, req, ofc)
	if err != nil {
		return nil, resp, err
	}

	return ofc, resp, err
}

// GetByName retrieves a Osquery file category by name.
func (s *OsqueryFileCategoriesServiceOp) GetByName(ctx context.Context, name string) (*OsqueryFileCategory, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listOFCOpt := &listOFCOptions{Name: name}

	ofcs, resp, err := s.list(ctx, nil, listOFCOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(ofcs) < 1 {
		return nil, resp, nil
	}

	return &ofcs[0], resp, err
}

// Create a new Osquery file category.
func (s *OsqueryFileCategoriesServiceOp) Create(ctx context.Context, createRequest *OsqueryFileCategoryRequest) (*OsqueryFileCategory, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, ofcBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	ofc := new(OsqueryFileCategory)
	resp, err := s.client.Do(ctx, req, ofc)
	if err != nil {
		return nil, resp, err
	}

	return ofc, resp, err
}

// Update a Osquery file category.
func (s *OsqueryFileCategoriesServiceOp) Update(ctx context.Context, ofcID int, updateRequest *OsqueryFileCategoryRequest) (*OsqueryFileCategory, *Response, error) {
	if ofcID < 1 {
		return nil, nil, NewArgError("ofcID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", ofcBasePath, ofcID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	ofc := new(OsqueryFileCategory)
	resp, err := s.client.Do(ctx, req, ofc)
	if err != nil {
		return nil, resp, err
	}

	return ofc, resp, err
}

// Delete a Osquery file category.
func (s *OsqueryFileCategoriesServiceOp) Delete(ctx context.Context, ofcID int) (*Response, error) {
	if ofcID < 1 {
		return nil, NewArgError("ofcID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", ofcBasePath, ofcID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Osquery file categories.
func (s *OsqueryFileCategoriesServiceOp) list(ctx context.Context, opt *ListOptions, ofcOpt *listOFCOptions) ([]OsqueryFileCategory, *Response, error) {
	path := ofcBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, ofcOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var ofcs []OsqueryFileCategory
	resp, err := s.client.Do(ctx, req, &ofcs)
	if err != nil {
		return nil, resp, err
	}

	return ofcs, resp, err
}
