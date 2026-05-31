package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const mpkgBasePath = "mdm/packages/"

// MDMPackagesService is an interface for interfacing with the MDM package
// endpoints of the Zentral API
type MDMPackagesService interface {
	List(context.Context, *ListOptions) ([]MDMPackage, *Response, error)
	GetByID(context.Context, string) (*MDMPackage, *Response, error)
	GetByName(context.Context, string) ([]MDMPackage, *Response, error)
	Create(context.Context, *MDMPackageCreateRequest) (*MDMPackage, *Response, error)
	Update(context.Context, string, *MDMPackageUpdateRequest) (*MDMPackage, *Response, error)
	Delete(context.Context, string) (*Response, error)
}

// MDMPackagesServiceOp handles communication with the MDM packages related
// methods of the Zentral API.
type MDMPackagesServiceOp struct {
	client *Client
}

var _ MDMPackagesService = &MDMPackagesServiceOp{}

// MDMPackage represents a Zentral MDM package
type MDMPackage struct {
	ID             string                   `json:"id"`
	Name           string                   `json:"name"`
	Description    string                   `json:"description"`
	Type           string                   `json:"type"`
	SourceURI      string                   `json:"source_uri"`
	SHA256         string                   `json:"sha256"`
	Size           int64                    `json:"size"`
	Filename       string                   `json:"filename"`
	ProductID      string                   `json:"product_id"`
	ProductVersion string                   `json:"product_version"`
	Bundles        []map[string]interface{} `json:"bundles"`
	Manifest       map[string]interface{}   `json:"manifest"`
	Created        Timestamp                `json:"created_at"`
	Updated        Timestamp                `json:"updated_at"`
}

func (mp MDMPackage) String() string {
	return Stringify(mp)
}

// MDMPackageCreateRequest represents a request to create a MDM package.
type MDMPackageCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	SourceURI   string `json:"source_uri"`
	SHA256      string `json:"sha256"`
}

// MDMPackageUpdateRequest represents a request to update a MDM package. Only
// name and description are mutable post-create — the underlying file, and
// therefore source_uri and sha256, are fixed at creation time.
type MDMPackageUpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type listMPKGOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the MDM packages.
func (s *MDMPackagesServiceOp) List(ctx context.Context, opt *ListOptions) ([]MDMPackage, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a MDM package by id.
func (s *MDMPackagesServiceOp) GetByID(ctx context.Context, mpID string) (*MDMPackage, *Response, error) {
	if len(mpID) < 1 {
		return nil, nil, NewArgError("mpID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", mpkgBasePath, mpID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	mp := new(MDMPackage)

	resp, err := s.client.Do(ctx, req, mp)
	if err != nil {
		return nil, resp, err
	}

	return mp, resp, err
}

// GetByName retrieves the MDM packages matching name. Package names are not
// unique server-side, so this returns a slice.
func (s *MDMPackagesServiceOp) GetByName(ctx context.Context, name string) ([]MDMPackage, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	return s.list(ctx, nil, &listMPKGOptions{Name: name})
}

// Create a new MDM package.
func (s *MDMPackagesServiceOp) Create(ctx context.Context, createRequest *MDMPackageCreateRequest) (*MDMPackage, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, mpkgBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	mp := new(MDMPackage)
	resp, err := s.client.Do(ctx, req, mp)
	if err != nil {
		return nil, resp, err
	}

	return mp, resp, err
}

// Update a MDM package.
func (s *MDMPackagesServiceOp) Update(ctx context.Context, mpID string, updateRequest *MDMPackageUpdateRequest) (*MDMPackage, *Response, error) {
	if len(mpID) < 1 {
		return nil, nil, NewArgError("mpID", "cannot be blank")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%s/", mpkgBasePath, mpID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	mp := new(MDMPackage)
	resp, err := s.client.Do(ctx, req, mp)
	if err != nil {
		return nil, resp, err
	}

	return mp, resp, err
}

// Delete a MDM package.
func (s *MDMPackagesServiceOp) Delete(ctx context.Context, mpID string) (*Response, error) {
	if len(mpID) < 1 {
		return nil, NewArgError("mpID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", mpkgBasePath, mpID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing MDM packages
func (s *MDMPackagesServiceOp) list(ctx context.Context, opt *ListOptions, mpOpt *listMPKGOptions) ([]MDMPackage, *Response, error) {
	path := mpkgBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, mpOpt)
	if err != nil {
		return nil, nil, err
	}
	return resolveAllPages[MDMPackage](ctx, s.client, path)
}
