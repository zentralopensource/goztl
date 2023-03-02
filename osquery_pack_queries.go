package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const opqBasePath = "osquery/pack_queries/"

// OsqueryPackQueriesService is an interface for interfacing with the Osquery pack queries
// endpoints of the Zentral API
type OsqueryPackQueriesService interface {
	List(context.Context, *ListOptions) ([]OsqueryPackQuery, *Response, error)
	GetByID(context.Context, int) (*OsqueryPackQuery, *Response, error)
	GetByPackID(context.Context, int) ([]OsqueryPackQuery, *Response, error)
	Create(context.Context, *OsqueryPackQueryRequest) (*OsqueryPackQuery, *Response, error)
	Update(context.Context, int, *OsqueryPackQueryRequest) (*OsqueryPackQuery, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// OsqueryPackQueriesServiceOp handles communication with the Osquery pack queries related
// methods of the Zentral API.
type OsqueryPackQueriesServiceOp struct {
	client *Client
}

var _ OsqueryPackQueriesService = &OsqueryPackQueriesServiceOp{}

// OsqueryPackQuery represents a Zentral Osquery pack query
type OsqueryPackQuery struct {
	ID                int       `json:"id"`
	PackID            int       `json:"pack"`
	QueryID           int       `json:"query"`
	Slug              string    `json:"slug"`
	Interval          int       `json:"interval"`
	LogRemovedActions bool      `json:"log_removed_actions"`
	SnapshotMode      bool      `json:"snapshot_mode"`
	Shard             *int      `json:"shard"`
	CanBeDenyListed   bool      `json:"can_be_denylisted"`
	Created           Timestamp `json:"created_at"`
	Updated           Timestamp `json:"updated_at"`
}

func (opq OsqueryPackQuery) String() string {
	return Stringify(opq)
}

// OsqueryPackQueryRequest represents a request to create or update a Osquery pack query
type OsqueryPackQueryRequest struct {
	PackID            int  `json:"pack"`
	QueryID           int  `json:"query"`
	Interval          int  `json:"interval"`
	LogRemovedActions bool `json:"log_removed_actions"`
	SnapshotMode      bool `json:"snapshot_mode"`
	Shard             *int `json:"shard"`
	CanBeDenyListed   bool `json:"can_be_denylisted"`
}

type listOPQOptions struct {
	PackID int `url:"pack_id,omitempty"`
}

// List lists all the Osquery pack queries.
func (s *OsqueryPackQueriesServiceOp) List(ctx context.Context, opt *ListOptions) ([]OsqueryPackQuery, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Osquery pack query by id.
func (s *OsqueryPackQueriesServiceOp) GetByID(ctx context.Context, opqID int) (*OsqueryPackQuery, *Response, error) {
	if opqID < 1 {
		return nil, nil, NewArgError("opqID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", opqBasePath, opqID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	opq := new(OsqueryPackQuery)

	resp, err := s.client.Do(ctx, req, opq)
	if err != nil {
		return nil, resp, err
	}

	return opq, resp, err
}

// GetByPackID retrieves Osquery pack queries by pack ID.
func (s *OsqueryPackQueriesServiceOp) GetByPackID(ctx context.Context, packID int) ([]OsqueryPackQuery, *Response, error) {
	if packID < 1 {
		return nil, nil, NewArgError("packID", "cannot be less than 1")
	}

	listOPQOpt := &listOPQOptions{PackID: packID}

	return s.list(ctx, nil, listOPQOpt)
}

// Create a new Osquery pack query.
func (s *OsqueryPackQueriesServiceOp) Create(ctx context.Context, createRequest *OsqueryPackQueryRequest) (*OsqueryPackQuery, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, opqBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	opq := new(OsqueryPackQuery)
	resp, err := s.client.Do(ctx, req, opq)
	if err != nil {
		return nil, resp, err
	}

	return opq, resp, err
}

// Update a Osquery pack query.
func (s *OsqueryPackQueriesServiceOp) Update(ctx context.Context, opqID int, updateRequest *OsqueryPackQueryRequest) (*OsqueryPackQuery, *Response, error) {
	if opqID < 1 {
		return nil, nil, NewArgError("opqID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", opqBasePath, opqID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	opq := new(OsqueryPackQuery)
	resp, err := s.client.Do(ctx, req, opq)
	if err != nil {
		return nil, resp, err
	}

	return opq, resp, err
}

// Delete a Osquery pack query.
func (s *OsqueryPackQueriesServiceOp) Delete(ctx context.Context, opqID int) (*Response, error) {
	if opqID < 1 {
		return nil, NewArgError("opqID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", opqBasePath, opqID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Osquery pack queries.
func (s *OsqueryPackQueriesServiceOp) list(ctx context.Context, opt *ListOptions, opqOpt *listOPQOptions) ([]OsqueryPackQuery, *Response, error) {
	path := opqBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, opqOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var opqs []OsqueryPackQuery
	resp, err := s.client.Do(ctx, req, &opqs)
	if err != nil {
		return nil, resp, err
	}

	return opqs, resp, err
}
