package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const depEnrollmentCustomViewBasePath = "mdm/dep_enrollment_custom_views/"

// MDMDEPEnrollmentCustomViewsService is an interface for interfacing with the MDM dep enrollment custom views
// endpoints of the Zentral API
type MDMDEPEnrollmentCustomViewsService interface {
	List(context.Context, *ListOptions) ([]MDMDEPEnrollmentCustomView, *Response, error)
	GetByID(context.Context, string) (*MDMDEPEnrollmentCustomView, *Response, error)
	Create(context.Context, *MDMDEPEnrollmentCustomViewRequest) (*MDMDEPEnrollmentCustomView, *Response, error)
	Update(context.Context, string, *MDMDEPEnrollmentCustomViewRequest) (*MDMDEPEnrollmentCustomView, *Response, error)
	Delete(context.Context, string) (*Response, error)
}

// MDMEnrollmentCustomViewsServiceOp handles communication with the MDM enrollments related
// moethods of the Zentral API.
type MDMDEPEnrollmentCustomViewsServiceOp struct {
	client *Client
}

var _ MDMDEPEnrollmentCustomViewsService = &MDMDEPEnrollmentCustomViewsServiceOp{}

// MDMDEPEnrollmentCustomView represents a Zentral MDM DEP enrollment custom view
type MDMDEPEnrollmentCustomView struct {
	ID              string    `json:"id"`
	DEPEnrollmentID int       `json:"dep_enrollment"`
	CustomViewID    string    `json:"custom_view"`
	Weight          int       `json:"weight"`
	Created         Timestamp `json:"created_at,omitempty"`
	Updated         Timestamp `json:"updated_at,omitempty"`
}

func (depCustomView MDMDEPEnrollmentCustomView) String() string {
	return Stringify(depCustomView)
}

// MDMDEPEnrollmentCustomViewRequest represents a request to create or update a MDM DEP enrollment custom view
type MDMDEPEnrollmentCustomViewRequest struct {
	DEPEnrollmentID int    `json:"dep_enrollment"`
	CustomViewID    string `json:"custom_view"`
	Weight          int    `json:"weight"`
}

type listMDMDEPEnrollmentCustomViewOptions struct {
	Name string `url:"omitempty"`
}

// List lists all the MDM enrollment custom view
func (service *MDMDEPEnrollmentCustomViewsServiceOp) List(ctx context.Context, opt *ListOptions) ([]MDMDEPEnrollmentCustomView, *Response, error) {
	return service.list(ctx, opt, nil)
}

// GetByID retrieves a MDM DEP enrollment custom view by id.
func (service *MDMDEPEnrollmentCustomViewsServiceOp) GetByID(ctx context.Context, depEnrollmentID string) (*MDMDEPEnrollmentCustomView, *Response, error) {
	if len(depEnrollmentID) < 1 {
		return nil, nil, NewArgError("depEnrollmentID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", depEnrollmentCustomViewBasePath, depEnrollmentID)

	req, err := service.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	depCustomView := new(MDMDEPEnrollmentCustomView)

	resp, err := service.client.Do(ctx, req, depCustomView)
	if err != nil {
		return nil, resp, err
	}

	return depCustomView, resp, err
}

// Create a new MDM DEP enrollment custom view.
func (s *MDMDEPEnrollmentCustomViewsServiceOp) Create(ctx context.Context, createRequest *MDMDEPEnrollmentCustomViewRequest) (*MDMDEPEnrollmentCustomView, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, depEnrollmentCustomViewBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	depCustomView := new(MDMDEPEnrollmentCustomView)
	resp, err := s.client.Do(ctx, req, depCustomView)
	if err != nil {
		return nil, resp, err
	}

	return depCustomView, resp, err
}

// Update a MDM DEP enrollment custom view.
func (s *MDMDEPEnrollmentCustomViewsServiceOp) Update(ctx context.Context, depCustomViewID string, updateRequest *MDMDEPEnrollmentCustomViewRequest) (*MDMDEPEnrollmentCustomView, *Response, error) {
	if len(depCustomViewID) < 1 {
		return nil, nil, NewArgError("depCustomViewID", "cannot be blank")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%s/", depEnrollmentCustomViewBasePath, depCustomViewID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	depCustomView := new(MDMDEPEnrollmentCustomView)
	resp, err := s.client.Do(ctx, req, depCustomView)
	if err != nil {
		return nil, resp, err
	}

	return depCustomView, resp, err
}

// Delete a MDM DEP enrollment custom view.
func (s *MDMDEPEnrollmentCustomViewsServiceOp) Delete(ctx context.Context, depCustomViewID string) (*Response, error) {
	if len(depCustomViewID) < 1 {
		return nil, NewArgError("depCustomViewID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", depEnrollmentCustomViewBasePath, depCustomViewID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper moethod for listing MDM DEP enrollment custom view
func (service *MDMDEPEnrollmentCustomViewsServiceOp) list(ctx context.Context, opt *ListOptions, listOpt *listMDMDEPEnrollmentCustomViewOptions) ([]MDMDEPEnrollmentCustomView, *Response, error) {
	path := depEnrollmentCustomViewBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, listOpt)
	if err != nil {
		return nil, nil, err
	}

	return resolveAllPages[MDMDEPEnrollmentCustomView](ctx, service.client, path)
}
