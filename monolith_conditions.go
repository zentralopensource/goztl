package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const mcoBasePath = "monolith/conditions/"

// MonolithConditionsService is an interface for interfacing with the Monolith conditions
// endpoints of the Zentral API
type MonolithConditionsService interface {
	List(context.Context, *ListOptions) ([]MonolithCondition, *Response, error)
	GetByID(context.Context, int) (*MonolithCondition, *Response, error)
	GetByName(context.Context, string) (*MonolithCondition, *Response, error)
	Create(context.Context, *MonolithConditionRequest) (*MonolithCondition, *Response, error)
	Update(context.Context, int, *MonolithConditionRequest) (*MonolithCondition, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// MonolithConditionsServiceOp handles comcunication with the Monolith conditions related
// methods of the Zentral API.
type MonolithConditionsServiceOp struct {
	client *Client
}

var _ MonolithConditionsService = &MonolithConditionsServiceOp{}

// MonolithCondition represents a Zentral MonolithCondition
type MonolithCondition struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Predicate string    `json:"predicate"`
	Created   Timestamp `json:"created_at"`
	Updated   Timestamp `json:"updated_at"`
}

func (se MonolithCondition) String() string {
	return Stringify(se)
}

// MonolithConditionRequest represents a request to create or update a Monolith condition
type MonolithConditionRequest struct {
	Name      string `json:"name"`
	Predicate string `json:"predicate"`
}

type listMCOOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the Monolith conditions.
func (s *MonolithConditionsServiceOp) List(ctx context.Context, opt *ListOptions) ([]MonolithCondition, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Monolith condition by id.
func (s *MonolithConditionsServiceOp) GetByID(ctx context.Context, mcID int) (*MonolithCondition, *Response, error) {
	if mcID < 1 {
		return nil, nil, NewArgError("mcID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mcoBasePath, mcID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	mc := new(MonolithCondition)

	resp, err := s.client.Do(ctx, req, mc)
	if err != nil {
		return nil, resp, err
	}

	return mc, resp, err
}

// GetByName retrieves a Monolith condition by name.
func (s *MonolithConditionsServiceOp) GetByName(ctx context.Context, name string) (*MonolithCondition, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listMCOOpt := &listMCOOptions{Name: name}

	mcs, resp, err := s.list(ctx, nil, listMCOOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(mcs) < 1 {
		return nil, resp, nil
	}

	return &mcs[0], resp, err
}

// Create a new Monolith condition.
func (s *MonolithConditionsServiceOp) Create(ctx context.Context, createRequest *MonolithConditionRequest) (*MonolithCondition, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, mcoBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	mc := new(MonolithCondition)
	resp, err := s.client.Do(ctx, req, mc)
	if err != nil {
		return nil, resp, err
	}

	return mc, resp, err
}

// Update a Monolith condition.
func (s *MonolithConditionsServiceOp) Update(ctx context.Context, mcID int, updateRequest *MonolithConditionRequest) (*MonolithCondition, *Response, error) {
	if mcID < 1 {
		return nil, nil, NewArgError("mcID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", mcoBasePath, mcID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	mc := new(MonolithCondition)
	resp, err := s.client.Do(ctx, req, mc)
	if err != nil {
		return nil, resp, err
	}

	return mc, resp, err
}

// Delete a Monolith condition.
func (s *MonolithConditionsServiceOp) Delete(ctx context.Context, mcID int) (*Response, error) {
	if mcID < 1 {
		return nil, NewArgError("mcID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mcoBasePath, mcID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Monolith conditions
func (s *MonolithConditionsServiceOp) list(ctx context.Context, opt *ListOptions, mcOpt *listMCOOptions) ([]MonolithCondition, *Response, error) {
	path := mcoBasePath
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

	var mcs []MonolithCondition
	resp, err := s.client.Do(ctx, req, &mcs)
	if err != nil {
		return nil, resp, err
	}

	return mcs, resp, err
}
