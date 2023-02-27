package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const oaBasePath = "osquery/atcs/"

// OsqueryATCService is an interface for interfacing with the Osquery automatic table construction
// endpoints of the Zentral API
type OsqueryATCService interface {
	List(context.Context, *ListOptions) ([]OsqueryATC, *Response, error)
	GetByID(context.Context, int) (*OsqueryATC, *Response, error)
	GetByName(context.Context, string) (*OsqueryATC, *Response, error)
	Create(context.Context, *OsqueryATCRequest) (*OsqueryATC, *Response, error)
	Update(context.Context, int, *OsqueryATCRequest) (*OsqueryATC, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// OsqueryATCServiceOp handles communication with the Osquery automatic table construction related
// methods of the Zentral API.
type OsqueryATCServiceOp struct {
	client *Client
}

var _ OsqueryATCService = &OsqueryATCServiceOp{}

// OsqueryATC represents a Zentral Osquery ATC
type OsqueryATC struct {
	ID          int       `json:"id,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	TableName   string    `json:"table_name"`
	Query       string    `json:"query"`
	Path        string    `json:"path"`
	Columns     []string  `json:"columns"`
	Platforms   []string  `json:"platforms"`
	Created     Timestamp `json:"created_at"`
	Updated     Timestamp `json:"updated_at"`
}

func (oa OsqueryATC) String() string {
	return Stringify(oa)
}

// OsqueryATCRequest represents a request to create or update a Osquery ATC
type OsqueryATCRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	TableName   string   `json:"table_name"`
	Query       string   `json:"query"`
	Path        string   `json:"path"`
	Columns     []string `json:"columns"`
	Platforms   []string `json:"platforms"`
}

type listOAOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the Osquery ATC.
func (s *OsqueryATCServiceOp) List(ctx context.Context, opt *ListOptions) ([]OsqueryATC, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Osquery ATC by id.
func (s *OsqueryATCServiceOp) GetByID(ctx context.Context, oaID int) (*OsqueryATC, *Response, error) {
	if oaID < 1 {
		return nil, nil, NewArgError("oaID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", oaBasePath, oaID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	oa := new(OsqueryATC)

	resp, err := s.client.Do(ctx, req, oa)
	if err != nil {
		return nil, resp, err
	}

	return oa, resp, err
}

// GetByName retrieves a Osquery ATC by name.
func (s *OsqueryATCServiceOp) GetByName(ctx context.Context, name string) (*OsqueryATC, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listSCOpt := &listOAOptions{Name: name}

	oas, resp, err := s.list(ctx, nil, listSCOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(oas) < 1 {
		return nil, resp, nil
	}

	return &oas[0], resp, err
}

// Create a new Osquery ATC.
func (s *OsqueryATCServiceOp) Create(ctx context.Context, createRequest *OsqueryATCRequest) (*OsqueryATC, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, oaBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	oa := new(OsqueryATC)
	resp, err := s.client.Do(ctx, req, oa)
	if err != nil {
		return nil, resp, err
	}

	return oa, resp, err
}

// Update a Osquery ATC.
func (s *OsqueryATCServiceOp) Update(ctx context.Context, oaID int, updateRequest *OsqueryATCRequest) (*OsqueryATC, *Response, error) {
	if oaID < 1 {
		return nil, nil, NewArgError("oaID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", oaBasePath, oaID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	oa := new(OsqueryATC)
	resp, err := s.client.Do(ctx, req, oa)
	if err != nil {
		return nil, resp, err
	}

	return oa, resp, err
}

// Delete a Osquery ATC.
func (s *OsqueryATCServiceOp) Delete(ctx context.Context, oaID int) (*Response, error) {
	if oaID < 1 {
		return nil, NewArgError("oaID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", oaBasePath, oaID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Osquery ATC.
func (s *OsqueryATCServiceOp) list(ctx context.Context, opt *ListOptions, oaOpt *listOAOptions) ([]OsqueryATC, *Response, error) {
	path := oaBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, oaOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var oas []OsqueryATC
	resp, err := s.client.Do(ctx, req, &oas)
	if err != nil {
		return nil, resp, err
	}

	return oas, resp, err
}
