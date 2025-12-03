package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const gwsConnctionsBasePath = "google_workspace/connections/"

// GWSCpnnectionsService is an interface for interfacing with the Google Workspace connections
// endpoints of the Zentral API.
type GWSConnectionsService interface {
	List(context.Context, *ListOptions) ([]GWSConnection, *Response, error)
	GetByID(context.Context, string) (*GWSConnection, *Response, error)
	GetByName(context.Context, string) (*GWSConnection, *Response, error)
}

// GWSConnectionsServiceOp handles communication with the Google Workspace connections related
// methods of the Zentral API.
type GWSConnectionsServiceOp struct {
	client *Client
}

var _ GWSConnectionsService = &GWSConnectionsServiceOp{}

// GWSConnection represents a Zentral Google Workspace connection
type GWSConnection struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	Created Timestamp `json:"created_at"`
	Updated Timestamp `json:"updated_at"`
}

func (connection GWSConnection) String() string {
	return Stringify(connection)
}

type listGWSConnectionsOptions struct {
	Name string `url:"name"`
}

// List lists all the Zentral Google Workspace connections.
func (s *GWSConnectionsServiceOp) List(ctx context.Context, opt *ListOptions) ([]GWSConnection, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Zentral Google Workspace connection by id.
func (s *GWSConnectionsServiceOp) GetByID(ctx context.Context, gwsConnectionID string) (*GWSConnection, *Response, error) {
	if len(gwsConnectionID) < 1 {
		return nil, nil, NewArgError("gwsConnectionID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", gwsConnctionsBasePath, gwsConnectionID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	gwsConnection := new(GWSConnection)

	resp, err := s.client.Do(ctx, req, gwsConnection)
	if err != nil {
		return nil, resp, err
	}

	return gwsConnection, resp, err
}

// GetByName retrieves a Google Workspace connection by name.
func (s *GWSConnectionsServiceOp) GetByName(ctx context.Context, name string) (*GWSConnection, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listOpt := &listGWSConnectionsOptions{Name: name}

	gwsConnections, resp, err := s.list(ctx, nil, listOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(gwsConnections) < 1 {
		return nil, resp, nil
	}

	return &gwsConnections[0], resp, err
}

// Helper method for listing Google Workspace connections.
func (s *GWSConnectionsServiceOp) list(ctx context.Context, opt *ListOptions, maiOpt *listGWSConnectionsOptions) ([]GWSConnection, *Response, error) {
	path := gwsConnctionsBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, maiOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var gwsConnections []GWSConnection
	resp, err := s.client.Do(ctx, req, &gwsConnections)
	if err != nil {
		return nil, resp, err
	}

	return gwsConnections, resp, err
}
