package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const mbuBasePath = "inventory/meta_business_units/"

// MetaBusinessUnitsService is an interface for interfacing with the meta business units
// endpoints of the Zentral API
type MetaBusinessUnitsService interface {
	List(context.Context, *ListOptions) ([]MetaBusinessUnit, *Response, error)
	GetByID(context.Context, int) (*MetaBusinessUnit, *Response, error)
	GetByName(context.Context, string) (*MetaBusinessUnit, *Response, error)
	Create(context.Context, *MetaBusinessUnitCreateRequest) (*MetaBusinessUnit, *Response, error)
	Update(context.Context, int, *MetaBusinessUnitUpdateRequest) (*MetaBusinessUnit, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// MetaBusinessUnitsServiceOp handles communication with the meta business units related
// methods of the Zentral API.
type MetaBusinessUnitsServiceOp struct {
	client *Client
}

var _ MetaBusinessUnitsService = &MetaBusinessUnitsServiceOp{}

// MetaBusinessUnit represents a Zentral MetaBusinessUnit
type MetaBusinessUnit struct {
	ID                   int       `json:"id,omitempty"`
	Name                 string    `json:"name,omitempty"`
	APIEnrollmentEnabled bool      `json:"api_enrollment_enabled"`
	Created              Timestamp `json:"created_at,omitempty"`
	Updated              Timestamp `json:"updated_at,omitempty"`
}

// MetaBusinessUnitCreateRequest represents a request to create a meta business unit.
type MetaBusinessUnitCreateRequest struct {
	Name string `json:"name"`

	// Boolean to enable API enrollments.
	APIEnrollmentEnabled bool `json:"api_enrollment_enabled"`
}

// MetaBusinessUnitUpdateRequest represents a request to update a meta business unit.
type MetaBusinessUnitUpdateRequest struct {
	Name string `json:"name"`

	// Boolean to enable API enrollments.
	// If set, it cannot be unset.
	APIEnrollmentEnabled bool `json:"api_enrollment_enabled"`
}

func (mbu MetaBusinessUnit) String() string {
	return Stringify(mbu)
}

type listMBUOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the meta business units.
func (s *MetaBusinessUnitsServiceOp) List(ctx context.Context, opt *ListOptions) ([]MetaBusinessUnit, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a meta business unit by id.
func (s *MetaBusinessUnitsServiceOp) GetByID(ctx context.Context, mbuID int) (*MetaBusinessUnit, *Response, error) {
	if mbuID < 1 {
		return nil, nil, NewArgError("mbuID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mbuBasePath, mbuID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	mbu := new(MetaBusinessUnit)

	resp, err := s.client.Do(ctx, req, mbu)
	if err != nil {
		return nil, resp, err
	}

	return mbu, resp, err
}

// GetByName retrieves a meta business unit by name.
func (s *MetaBusinessUnitsServiceOp) GetByName(ctx context.Context, name string) (*MetaBusinessUnit, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listMBUOpt := &listMBUOptions{Name: name}

	mbus, resp, err := s.list(ctx, nil, listMBUOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(mbus) < 1 {
		return nil, resp, nil
	}

	return &mbus[0], resp, err
}

// Create a new meta business unit.
func (s *MetaBusinessUnitsServiceOp) Create(ctx context.Context, createRequest *MetaBusinessUnitCreateRequest) (*MetaBusinessUnit, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, mbuBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	mbu := new(MetaBusinessUnit)
	resp, err := s.client.Do(ctx, req, mbu)
	if err != nil {
		return nil, resp, err
	}

	return mbu, resp, err
}

// Update a meta business unit.
func (s *MetaBusinessUnitsServiceOp) Update(ctx context.Context, mbuID int, updateRequest *MetaBusinessUnitUpdateRequest) (*MetaBusinessUnit, *Response, error) {
	if mbuID < 1 {
		return nil, nil, NewArgError("mbuID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", mbuBasePath, mbuID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	mbu := new(MetaBusinessUnit)
	resp, err := s.client.Do(ctx, req, mbu)
	if err != nil {
		return nil, resp, err
	}

	return mbu, resp, err
}

// Delete a meta business unit.
func (s *MetaBusinessUnitsServiceOp) Delete(ctx context.Context, mbuID int) (*Response, error) {
	if mbuID < 1 {
		return nil, NewArgError("mbuID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mbuBasePath, mbuID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing meta business units
func (s *MetaBusinessUnitsServiceOp) list(ctx context.Context, opt *ListOptions, mbuOpt *listMBUOptions) ([]MetaBusinessUnit, *Response, error) {
	path := mbuBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, mbuOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var mbus []MetaBusinessUnit
	resp, err := s.client.Do(ctx, req, &mbus)
	if err != nil {
		return nil, resp, err
	}

	return mbus, resp, err
}
