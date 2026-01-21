package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const enrollmentCustomViewBasePath = "mdm/enrollment_custom_views/"

// MDMEnrollmentCustomViewsService is an interface for interfacing with the MDM enrollment custom views
// endpoints of the Zentral API
type MDMEnrollmentCustomViewsService interface {
	List(context.Context, *ListOptions) ([]MDMEnrollmentCustomView, *Response, error)
	GetByID(context.Context, string) (*MDMEnrollmentCustomView, *Response, error)
	GetByName(context.Context, string) (*MDMEnrollmentCustomView, *Response, error)
	Create(context.Context, *MDMEnrollmentCustomViewRequest) (*MDMEnrollmentCustomView, *Response, error)
	Update(context.Context, string, *MDMEnrollmentCustomViewRequest) (*MDMEnrollmentCustomView, *Response, error)
	Delete(context.Context, string) (*Response, error)
}

// MDMEnrollmentCustomViewsServiceOp handles communication with the MDM enrollments related
// moethods of the Zentral API.
type MDMEnrollmentCustomViewsServiceOp struct {
	client *Client
}

var _ MDMEnrollmentCustomViewsService = &MDMEnrollmentCustomViewsServiceOp{}

// MDMEnrollmentCustomView represents a Zentral MDM enrollment custom view
type MDMEnrollmentCustomView struct {
	ID                     string    `json:"id"`
	Name                   string    `json:"name"`
	Description            string    `json:"description"`
	HTML                   string    `json:"html"`
	RequiresAuthentication bool      `json:"requires_authentication"`
	Created                Timestamp `json:"created_at,omitempty"`
	Updated                Timestamp `json:"updated_at,omitempty"`
}

func (customView MDMEnrollmentCustomView) String() string {
	return Stringify(customView)
}

// MDMEnrollmentCustomViewRequest represents a request to create or update a MDM enrollment custom view
type MDMEnrollmentCustomViewRequest struct {
	Name                   string `json:"name"`
	Description            string `json:"description"`
	HTML                   string `json:"html"`
	RequiresAuthentication bool   `json:"requires_authentication"`
}

type listMDMEnrollmentCustomViewOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the MDM enrollment custom view
func (service *MDMEnrollmentCustomViewsServiceOp) List(ctx context.Context, opt *ListOptions) ([]MDMEnrollmentCustomView, *Response, error) {
	return service.list(ctx, opt, nil)
}

// GetByID retrieves a  MDM enrollment custom view by id.
func (service *MDMEnrollmentCustomViewsServiceOp) GetByID(ctx context.Context, customViewID string) (*MDMEnrollmentCustomView, *Response, error) {
	if len(customViewID) < 1 {
		return nil, nil, NewArgError("customViewID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", enrollmentCustomViewBasePath, customViewID)

	req, err := service.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	customView := new(MDMEnrollmentCustomView)

	resp, err := service.client.Do(ctx, req, customView)
	if err != nil {
		return nil, resp, err
	}

	return customView, resp, err
}

// GetByName retrieves a MDM enrollment custom view by name.
func (service *MDMEnrollmentCustomViewsServiceOp) GetByName(ctx context.Context, name string) (*MDMEnrollmentCustomView, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listOpt := &listMDMEnrollmentCustomViewOptions{Name: name}

	customViews, resp, err := service.list(ctx, nil, listOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(customViews) < 1 {
		return nil, resp, nil
	}

	return &customViews[0], resp, err
}

// Create a new MDM enrollment custom view.
func (s *MDMEnrollmentCustomViewsServiceOp) Create(ctx context.Context, createRequest *MDMEnrollmentCustomViewRequest) (*MDMEnrollmentCustomView, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, enrollmentCustomViewBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	customView := new(MDMEnrollmentCustomView)
	resp, err := s.client.Do(ctx, req, customView)
	if err != nil {
		return nil, resp, err
	}

	return customView, resp, err
}

// Update a MDM enrollment custom view.
func (s *MDMEnrollmentCustomViewsServiceOp) Update(ctx context.Context, customViewID string, updateRequest *MDMEnrollmentCustomViewRequest) (*MDMEnrollmentCustomView, *Response, error) {
	if len(customViewID) < 1 {
		return nil, nil, NewArgError("customViewID", "cannot be blank")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%s/", enrollmentCustomViewBasePath, customViewID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	customView := new(MDMEnrollmentCustomView)
	resp, err := s.client.Do(ctx, req, customView)
	if err != nil {
		return nil, resp, err
	}

	return customView, resp, err
}

// Delete a MDM enrollment custom view.
func (s *MDMEnrollmentCustomViewsServiceOp) Delete(ctx context.Context, customViewID string) (*Response, error) {
	if len(customViewID) < 1 {
		return nil, NewArgError("customViewID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", enrollmentCustomViewBasePath, customViewID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper moethod for listing MDM enrollment custom view
func (service *MDMEnrollmentCustomViewsServiceOp) list(ctx context.Context, opt *ListOptions, listOpt *listMDMEnrollmentCustomViewOptions) ([]MDMEnrollmentCustomView, *Response, error) {
	path := enrollmentCustomViewBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, listOpt)
	if err != nil {
		return nil, nil, err
	}

	return resolveAllPages[MDMEnrollmentCustomView](ctx, service.client, path)
}
