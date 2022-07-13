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
}

// MetaBusinessUnitsServiceOp handles communication with the meta business units related
// methods of the Zentral API.
type MetaBusinessUnitsServiceOp struct {
	client *Client
}

var _ MetaBusinessUnitsService = &MetaBusinessUnitsServiceOp{}

// MetaBusinessUnit represents a Zentral MetaBusinessUnit
type MetaBusinessUnit struct {
	ID      int       `json:"id,float64,omitempty"`
	Name    string    `json:"name,omitempty"`
	Created Timestamp `json:"created_at,omitempty"`
	Updated Timestamp `json:"updated_at,omitempty"`
}

func (mbu MetaBusinessUnit) String() string {
	return Stringify(mbu)
}

type listMBUOptions struct {
	Name string `url:"name,omitempty`
}

// List lists all the meta business units.
func (s *MetaBusinessUnitsServiceOp) List(ctx context.Context, opt *ListOptions) ([]MetaBusinessUnit, *Response, error) {
	return s.list(ctx, opt, nil)
}

// ListByName lists all the meta business units filtered by name returning only exact matches.
// It is case-insensitive
func (s *MetaBusinessUnitsServiceOp) ListByName(ctx context.Context, name string, opt *ListOptions) ([]MetaBusinessUnit, *Response, error) {
	listMBUOpt := &listMBUOptions{Name: name}
	return s.list(ctx, opt, listMBUOpt)
}

// GetByID retrieves a meta business unit by id.
func (s *MetaBusinessUnitsServiceOp) GetByID(ctx context.Context, mbuID int) (*MetaBusinessUnit, *Response, error) {
	if mbuID < 1 {
		return nil, nil, NewArgError("mbuID", "cannot be less than 1")
	}

	return s.get(ctx, interface{}(mbuID))
}

// Helper method for getting an individual meta business unit
func (s *MetaBusinessUnitsServiceOp) get(ctx context.Context, ID interface{}) (*MetaBusinessUnit, *Response, error) {
	path := fmt.Sprintf("%s%v", mbuBasePath, ID)

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
