package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const gwsGroupTagMappingsBasePath = "google_workspace/group_tag_mappings/"

// GWSGroupTagMappingsService is an interface for interfacing with the Google Workspace
// group tag mapping endpoints of the Zentral API.
type GWSGroupTagMappingsService interface {
	List(context.Context, *ListOptions) ([]GWSGroupTagMapping, *Response, error)
	GetByID(context.Context, string) (*GWSGroupTagMapping, *Response, error)
	GetByConnectionID(context.Context, string) ([]GWSGroupTagMapping, *Response, error)
	GetByGroupEmail(context.Context, string) ([]GWSGroupTagMapping, *Response, error)
	Create(context.Context, *GWSGroupTagMappingRequest) (*GWSGroupTagMapping, *Response, error)
	Update(context.Context, string, *GWSGroupTagMappingRequest) (*GWSGroupTagMapping, *Response, error)
	Delete(context.Context, string) (*Response, error)
}

// GWSConnectionsServiceOp handles communication with the Google Workspace
// group tag mapping endpoints of the Zentral API.
type GWSGroupTagMappingsServiceOp struct {
	client *Client
}

var _ GWSGroupTagMappingsService = &GWSGroupTagMappingsServiceOp{}

// GWSGroupTagMapping represents a Zentral Google Workspace group tag mapping
type GWSGroupTagMapping struct {
	ID           string `json:"id"`
	GroupEmail   string `json:"group_email"`
	ConnectionID string `json:"connection"`
	TagIDs       []int  `json:"tags"`

	Created Timestamp `json:"created_at"`
	Updated Timestamp `json:"updated_at"`
}

// GWSGroupTagMappingRequest represents a request to create or update a
// Google Workspace group tag mapping.
type GWSGroupTagMappingRequest struct {
	GroupEmail   string `json:"group_email"`
	ConnectionID string `json:"connection"`
	TagIDs       []int  `json:"tags"`
}

func (mapping GWSGroupTagMapping) String() string {
	return Stringify(mapping)
}

type listGWSGroupTagMappingsOptions struct {
	GroupEmail string `url:"group_email"`
	Connection string `url:"connection_id"`
}

// List lists all Google Workspace group tag mappings.
func (s *GWSGroupTagMappingsServiceOp) List(ctx context.Context, opt *ListOptions) ([]GWSGroupTagMapping, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Google Workspace group tag mapping by id.
func (s *GWSGroupTagMappingsServiceOp) GetByID(ctx context.Context, gwsGroupTagMappingID string) (*GWSGroupTagMapping, *Response, error) {
	if len(gwsGroupTagMappingID) < 1 {
		return nil, nil, NewArgError("gwsGroupTagMappingID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", gwsGroupTagMappingsBasePath, gwsGroupTagMappingID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	mapping := new(GWSGroupTagMapping)

	resp, err := s.client.Do(ctx, req, mapping)
	if err != nil {
		return nil, resp, err
	}

	return mapping, resp, err
}

// GetByConnectionID retrieves Google Workspace group tag mappings by connection.
func (s *GWSGroupTagMappingsServiceOp) GetByConnectionID(ctx context.Context, connectionID string) ([]GWSGroupTagMapping, *Response, error) {
	if len(connectionID) < 1 {
		return nil, nil, NewArgError("connectionID", "cannot be blank")
	}

	listOpt := &listGWSGroupTagMappingsOptions{Connection: connectionID}

	mappings, resp, err := s.list(ctx, nil, listOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(mappings) < 1 {
		return nil, resp, nil
	}

	return mappings, resp, err
}

// GetByGroupEmail retrieves Google Workspace group tag mappings by group email.
func (s *GWSGroupTagMappingsServiceOp) GetByGroupEmail(ctx context.Context, groupEmail string) ([]GWSGroupTagMapping, *Response, error) {
	if len(groupEmail) < 1 {
		return nil, nil, NewArgError("groupEmail", "cannot be blank")
	}

	listOpt := &listGWSGroupTagMappingsOptions{GroupEmail: groupEmail}

	mappings, resp, err := s.list(ctx, nil, listOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(mappings) < 1 {
		return nil, resp, nil
	}

	return mappings, resp, err
}

// Create a new Google Workspace group tag mapping.
func (s *GWSGroupTagMappingsServiceOp) Create(ctx context.Context, createRequest *GWSGroupTagMappingRequest) (*GWSGroupTagMapping, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, gwsGroupTagMappingsBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	gwsGroupTagMapping := new(GWSGroupTagMapping)
	resp, err := s.client.Do(ctx, req, gwsGroupTagMapping)
	if err != nil {
		return nil, resp, err
	}

	return gwsGroupTagMapping, resp, err
}

// Update a Google Workspace group tag mapping.
func (s *GWSGroupTagMappingsServiceOp) Update(ctx context.Context, gwsGroupTagMappingID string, updateRequest *GWSGroupTagMappingRequest) (*GWSGroupTagMapping, *Response, error) {
	if len(gwsGroupTagMappingID) < 1 {
		return nil, nil, NewArgError("gwsGroupTagMappingID", "cannot be blank")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%s/", gwsGroupTagMappingsBasePath, gwsGroupTagMappingID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	gwsGroupTagMapping := new(GWSGroupTagMapping)
	resp, err := s.client.Do(ctx, req, gwsGroupTagMapping)
	if err != nil {
		return nil, resp, err
	}

	return gwsGroupTagMapping, resp, err
}

// Delete a Google Workspace group tag mapping.
func (s *GWSGroupTagMappingsServiceOp) Delete(ctx context.Context, gwsGroupTagMappingID string) (*Response, error) {
	if len(gwsGroupTagMappingID) < 1 {
		return nil, NewArgError("gwsGroupTagMappingID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", gwsGroupTagMappingsBasePath, gwsGroupTagMappingID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Google Workspace group tag mappings.
func (s *GWSGroupTagMappingsServiceOp) list(ctx context.Context, opt *ListOptions, gwsGroupTagMappingOpt *listGWSGroupTagMappingsOptions) ([]GWSGroupTagMapping, *Response, error) {
	path := gwsGroupTagMappingsBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, gwsGroupTagMappingOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var gwsGroupTagMappings []GWSGroupTagMapping
	resp, err := s.client.Do(ctx, req, &gwsGroupTagMappings)
	if err != nil {
		return nil, resp, err
	}

	return gwsGroupTagMappings, resp, err
}
