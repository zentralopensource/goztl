package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const opBasePath = "osquery/packs/"

// OsqueryPacksService is an interface for interfacing with the Osquery packs
// endpoints of the Zentral API
type OsqueryPacksService interface {
	List(context.Context, *ListOptions) ([]OsqueryPack, *Response, error)
	GetByID(context.Context, int) (*OsqueryPack, *Response, error)
	GetByName(context.Context, string) (*OsqueryPack, *Response, error)
	Create(context.Context, *OsqueryPackRequest) (*OsqueryPack, *Response, error)
	Update(context.Context, int, *OsqueryPackRequest) (*OsqueryPack, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// OsqueryPacksServiceOp handles communication with the Osquery packs related
// methods of the Zentral API.
type OsqueryPacksServiceOp struct {
	client *Client
}

var _ OsqueryPacksService = &OsqueryPacksServiceOp{}

// OsqueryPack represents a Zentral Osquery pack
type OsqueryPack struct {
	ID               int       `json:"id,omitempty"`
	Name             string    `json:"name"`
	Slug             string    `json:"slug"`
	Description      string    `json:"description"`
	DiscoveryQueries []string  `json:"discovery_queries"`
	Shard            *int      `json:"shard,omitempty"`
	EventRoutingKey  string    `json:"event_routing_key"`
	Created          Timestamp `json:"created_at"`
	Updated          Timestamp `json:"updated_at"`
}

func (op OsqueryPack) String() string {
	return Stringify(op)
}

// OsqueryPackRequest represents a request to create or update a Osquery pack
type OsqueryPackRequest struct {
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	DiscoveryQueries []string `json:"discovery_queries"`
	Shard            *int     `json:"shard,omitempty"`
	EventRoutingKey  string   `json:"event_routing_key"`
}

type listOPOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the Osquery packs.
func (s *OsqueryPacksServiceOp) List(ctx context.Context, opt *ListOptions) ([]OsqueryPack, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Osquery pack by id.
func (s *OsqueryPacksServiceOp) GetByID(ctx context.Context, opID int) (*OsqueryPack, *Response, error) {
	if opID < 1 {
		return nil, nil, NewArgError("opID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", opBasePath, opID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	op := new(OsqueryPack)

	resp, err := s.client.Do(ctx, req, op)
	if err != nil {
		return nil, resp, err
	}

	return op, resp, err
}

// GetByName retrieves a Osquery pack by name.
func (s *OsqueryPacksServiceOp) GetByName(ctx context.Context, name string) (*OsqueryPack, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listOPOpt := &listOPOptions{Name: name}

	ops, resp, err := s.list(ctx, nil, listOPOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(ops) < 1 {
		return nil, resp, nil
	}

	return &ops[0], resp, err
}

// Create a new Osquery pack.
func (s *OsqueryPacksServiceOp) Create(ctx context.Context, createRequest *OsqueryPackRequest) (*OsqueryPack, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, opBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	op := new(OsqueryPack)
	resp, err := s.client.Do(ctx, req, op)
	if err != nil {
		return nil, resp, err
	}

	return op, resp, err
}

// Update a Osquery pack.
func (s *OsqueryPacksServiceOp) Update(ctx context.Context, opID int, updateRequest *OsqueryPackRequest) (*OsqueryPack, *Response, error) {
	if opID < 1 {
		return nil, nil, NewArgError("opID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", opBasePath, opID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	op := new(OsqueryPack)
	resp, err := s.client.Do(ctx, req, op)
	if err != nil {
		return nil, resp, err
	}

	return op, resp, err
}

// Delete a Osquery pack.
func (s *OsqueryPacksServiceOp) Delete(ctx context.Context, opID int) (*Response, error) {
	if opID < 1 {
		return nil, NewArgError("opID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", opBasePath, opID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Osquery packs.
func (s *OsqueryPacksServiceOp) list(ctx context.Context, opt *ListOptions, opOpt *listOPOptions) ([]OsqueryPack, *Response, error) {
	path := opBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, opOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var ops []OsqueryPack
	resp, err := s.client.Do(ctx, req, &ops)
	if err != nil {
		return nil, resp, err
	}

	return ops, resp, err
}
