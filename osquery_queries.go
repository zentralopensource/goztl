package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const oqBasePath = "osquery/queries/"

// OsqueryQueriesService is an interface for interfacing with the Osquery queries
// endpoints of the Zentral API
type OsqueryQueriesService interface {
	List(context.Context, *ListOptions) ([]OsqueryQuery, *Response, error)
	GetByID(context.Context, int) (*OsqueryQuery, *Response, error)
	GetByName(context.Context, string) (*OsqueryQuery, *Response, error)
	Create(context.Context, *OsqueryQueryRequest) (*OsqueryQuery, *Response, error)
	Update(context.Context, int, *OsqueryQueryRequest) (*OsqueryQuery, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// OsqueryQueriesServiceOp handles communication with the Osquery queries related
// methods of the Zentral API.
type OsqueryQueriesServiceOp struct {
	client *Client
}

var _ OsqueryQueriesService = &OsqueryQueriesServiceOp{}

// OsqueryQuery represents a Zentral Osquery query
type OsqueryQuery struct {
	ID                     int       `json:"id,omitempty"`
	Name                   string    `json:"name"`
	SQL                    string    `json:"sql"`
	Platforms              []string  `json:"platforms"`
	MinOsqueryVersion      *string   `json:"minimum_osquery_version"`
	Description            string    `json:"description"`
	Value                  string    `json:"value"`
	Version                int       `json:"version"`
	ComplianceCheckEnabled bool      `json:"compliance_check_enabled"`
	Created                Timestamp `json:"created_at"`
	Updated                Timestamp `json:"updated_at"`
}

func (oq OsqueryQuery) String() string {
	return Stringify(oq)
}

// OsqueryQueryRequest represents a request to create or update a Osquery query
type OsqueryQueryRequest struct {
	Name                   string   `json:"name"`
	SQL                    string   `json:"sql"`
	Platforms              []string `json:"platforms"`
	MinOsqueryVersion      *string  `json:"minimum_osquery_version"`
	Description            string   `json:"description"`
	Value                  string   `json:"value"`
	ComplianceCheckEnabled bool     `json:"compliance_check_enabled"`
}

type listOQOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the Osquery queries.
func (s *OsqueryQueriesServiceOp) List(ctx context.Context, opt *ListOptions) ([]OsqueryQuery, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Osquery query by id.
func (s *OsqueryQueriesServiceOp) GetByID(ctx context.Context, oqID int) (*OsqueryQuery, *Response, error) {
	if oqID < 1 {
		return nil, nil, NewArgError("oqID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", oqBasePath, oqID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	oq := new(OsqueryQuery)

	resp, err := s.client.Do(ctx, req, oq)
	if err != nil {
		return nil, resp, err
	}

	return oq, resp, err
}

// GetByName retrieves a Osquery query by name.
func (s *OsqueryQueriesServiceOp) GetByName(ctx context.Context, name string) (*OsqueryQuery, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listOQOpt := &listOQOptions{Name: name}

	oqs, resp, err := s.list(ctx, nil, listOQOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(oqs) < 1 {
		return nil, resp, nil
	}

	return &oqs[0], resp, err
}

// Create a new Osquery query.
func (s *OsqueryQueriesServiceOp) Create(ctx context.Context, createRequest *OsqueryQueryRequest) (*OsqueryQuery, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, oqBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	oq := new(OsqueryQuery)
	resp, err := s.client.Do(ctx, req, oq)
	if err != nil {
		return nil, resp, err
	}

	return oq, resp, err
}

// Update a Osquery query.
func (s *OsqueryQueriesServiceOp) Update(ctx context.Context, oqID int, updateRequest *OsqueryQueryRequest) (*OsqueryQuery, *Response, error) {
	if oqID < 1 {
		return nil, nil, NewArgError("oqID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", oqBasePath, oqID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	oq := new(OsqueryQuery)
	resp, err := s.client.Do(ctx, req, oq)
	if err != nil {
		return nil, resp, err
	}

	return oq, resp, err
}

// Delete a Osquery query.
func (s *OsqueryQueriesServiceOp) Delete(ctx context.Context, oqID int) (*Response, error) {
	if oqID < 1 {
		return nil, NewArgError("oqID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", oqBasePath, oqID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Osquery queries.
func (s *OsqueryQueriesServiceOp) list(ctx context.Context, opt *ListOptions, oqOpt *listOQOptions) ([]OsqueryQuery, *Response, error) {
	path := oqBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, oqOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var oqs []OsqueryQuery
	resp, err := s.client.Do(ctx, req, &oqs)
	if err != nil {
		return nil, resp, err
	}

	return oqs, resp, err
}
