package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const depVirtualServersBasePath = "mdm/dep/virtual_servers/"

// MDMDEPVirtualServersService is an interface for interfacing with the MDM DEP virtual servers
// endpoints of the Zentral API
type MDMDEPVirtualServersService interface {
	List(context.Context, *ListOptions) ([]MDMDEPVirtualServer, *Response, error)
	GetByID(context.Context, int) (*MDMDEPVirtualServer, *Response, error)
	GetByName(context.Context, string) ([]MDMDEPVirtualServer, *Response, error)
}

// MDMDEPVirtualServersServiceOp handles communication with the MDM DEP virtual servers
// methods of the Zentral API.
type MDMDEPVirtualServersServiceOp struct {
	client *Client
}

var _ MDMDEPVirtualServersService = &MDMDEPVirtualServersServiceOp{}

// MDMDEPVirtualServer represents a Zentral MDM DEP virtual server
type MDMDEPVirtualServer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	UUID string `json:"uuid"`

	Created Timestamp `json:"created_at"`
	Updated Timestamp `json:"updated_at"`
}

func (connection MDMDEPVirtualServer) String() string {
	return Stringify(connection)
}

type listMDMDEPVirtualServerOptions struct {
	Name string `url:"name"`
}

// List lists all the Zentral MDM DEP virtual server.
func (service *MDMDEPVirtualServersServiceOp) List(ctx context.Context, opt *ListOptions) ([]MDMDEPVirtualServer, *Response, error) {
	return service.list(ctx, opt, nil)
}

// GetByID retrieves a Zentral MDM DEP virtual server by id.
func (service *MDMDEPVirtualServersServiceOp) GetByID(ctx context.Context, virtualServerID int) (*MDMDEPVirtualServer, *Response, error) {
	if virtualServerID < 1 {
		return nil, nil, NewArgError("virtualServerID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", depVirtualServersBasePath, virtualServerID)

	req, err := service.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	virtualServer := new(MDMDEPVirtualServer)

	resp, err := service.client.Do(ctx, req, virtualServer)
	if err != nil {
		return nil, resp, err
	}

	return virtualServer, resp, err
}

// GetByName retrieves a MDM DEP virtual server by name.
func (service *MDMDEPVirtualServersServiceOp) GetByName(ctx context.Context, name string) ([]MDMDEPVirtualServer, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listOpt := &listMDMDEPVirtualServerOptions{Name: name}

	virtualServers, resp, err := service.list(ctx, nil, listOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(virtualServers) < 1 {
		return nil, resp, nil
	}

	return virtualServers, resp, err
}

// Helper method for listing MDM DEP virtual servers.
func (service *MDMDEPVirtualServersServiceOp) list(ctx context.Context, opt *ListOptions, listOpt *listMDMDEPVirtualServerOptions) ([]MDMDEPVirtualServer, *Response, error) {
	path := depVirtualServersBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, listOpt)
	if err != nil {
		return nil, nil, err
	}

	return resolveAllPages[MDMDEPVirtualServer](ctx, service.client, path)
}
