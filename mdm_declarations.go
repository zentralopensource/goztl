package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const mdBasePath = "mdm/declarations/"

// MDMDeclarationsService is an interface for interfacing with the MDM declaration
// endpoints of the Zentral API
type MDMDeclarationsService interface {
	List(context.Context, *ListOptions) ([]MDMDeclaration, *Response, error)
	GetByID(context.Context, string) (*MDMDeclaration, *Response, error)
	Create(context.Context, *MDMDeclarationRequest) (*MDMDeclaration, *Response, error)
	Update(context.Context, string, *MDMDeclarationRequest) (*MDMDeclaration, *Response, error)
	Delete(context.Context, string) (*Response, error)
}

// MDMDeclarationsServiceOp handles communication with the MDM declarations related
// methods of the Zentral API.
type MDMDeclarationsServiceOp struct {
	client *Client
}

var _ MDMDeclarationsService = &MDMDeclarationsServiceOp{}

// MDMDeclarationSource represents the raw DDM declaration
// WARNING: the order of the attributes plays a role in the JSON serialization!!!
type MDMDeclarationSource struct {
	Identifier  string
	Payload     map[string]interface{}
	ServerToken string
	Type        string
}

// MDMDeclaration represents a Zentral MDM declaration
type MDMDeclaration struct {
	ID     string               `json:"id"`
	Source MDMDeclarationSource `json:"source"`
	MDMArtifactVersion
}

func (md MDMDeclaration) String() string {
	return Stringify(md)
}

// MDMDeclarationRequest represents a request to create or update a MDM declaration
type MDMDeclarationRequest struct {
	Source MDMDeclarationSource `json:"source"`
	MDMArtifactVersionRequest
}

type listMDOptions struct{}

// List lists all the MDM declarations.
func (s *MDMDeclarationsServiceOp) List(ctx context.Context, opt *ListOptions) ([]MDMDeclaration, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a MDM declaration by id.
func (s *MDMDeclarationsServiceOp) GetByID(ctx context.Context, mdID string) (*MDMDeclaration, *Response, error) {
	if len(mdID) < 1 {
		return nil, nil, NewArgError("mdID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", mdBasePath, mdID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	md := new(MDMDeclaration)

	resp, err := s.client.Do(ctx, req, md)
	if err != nil {
		return nil, resp, err
	}

	return md, resp, err
}

// Create a new MDM declaration.
func (s *MDMDeclarationsServiceOp) Create(ctx context.Context, createRequest *MDMDeclarationRequest) (*MDMDeclaration, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, mdBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	md := new(MDMDeclaration)
	resp, err := s.client.Do(ctx, req, md)
	if err != nil {
		return nil, resp, err
	}

	return md, resp, err
}

// Update a MDM declaration.
func (s *MDMDeclarationsServiceOp) Update(ctx context.Context, mdID string, updateRequest *MDMDeclarationRequest) (*MDMDeclaration, *Response, error) {
	if len(mdID) < 1 {
		return nil, nil, NewArgError("mdID", "cannot be blank")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%s/", mdBasePath, mdID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	md := new(MDMDeclaration)
	resp, err := s.client.Do(ctx, req, md)
	if err != nil {
		return nil, resp, err
	}

	return md, resp, err
}

// Delete a MDM declaration.
func (s *MDMDeclarationsServiceOp) Delete(ctx context.Context, mdID string) (*Response, error) {
	if len(mdID) < 1 {
		return nil, NewArgError("mdID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", mdBasePath, mdID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing MDM declarations
func (s *MDMDeclarationsServiceOp) list(ctx context.Context, opt *ListOptions, mdOpt *listMDOptions) ([]MDMDeclaration, *Response, error) {
	path := mdBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, mdOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var mds []MDMDeclaration
	resp, err := s.client.Do(ctx, req, &mds)
	if err != nil {
		return nil, resp, err
	}

	return mds, resp, err
}
