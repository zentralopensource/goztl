package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const ocBasePath = "osquery/configurations/"

// OsqueryConfigurationsService is an interface for interfacing with the Osquery configurations
// endpoints of the Zentral API
type OsqueryConfigurationsService interface {
	List(context.Context, *ListOptions) ([]OsqueryConfiguration, *Response, error)
	GetByID(context.Context, int) (*OsqueryConfiguration, *Response, error)
	GetByName(context.Context, string) (*OsqueryConfiguration, *Response, error)
	Create(context.Context, *OsqueryConfigurationRequest) (*OsqueryConfiguration, *Response, error)
	Update(context.Context, int, *OsqueryConfigurationRequest) (*OsqueryConfiguration, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// OsqueryConfigurationsServiceOp handles communication with the Osquery configurations related
// methods of the Zentral API.
type OsqueryConfigurationsServiceOp struct {
	client *Client
}

var _ OsqueryConfigurationsService = &OsqueryConfigurationsServiceOp{}

// OsqueryConfiguration represents a Zentral Osquery configuration
type OsqueryConfiguration struct {
	ID                int                    `json:"id,omitempty"`
	Name              string                 `json:"name"`
	Description       string                 `json:"description"`
	Inventory         bool                   `json:"inventory"`
	InventoryApps     bool                   `json:"inventory_apps"`
	InventoryEC2      bool                   `json:"inventory_ec2"`
	InventoryInterval int                    `json:"inventory_interval"`
	Options           map[string]interface{} `json:"options"`
	Created           Timestamp              `json:"created_at,omitempty"`
	Updated           Timestamp              `json:"updated_at,omitempty"`
}

func (oc OsqueryConfiguration) String() string {
	return Stringify(oc)
}

// OsqueryConfigurationRequest represents a request to create or update a Osquery configuration
type OsqueryConfigurationRequest struct {
	Name              string                 `json:"name"`
	Description       string                 `json:"description"`
	Inventory         bool                   `json:"inventory"`
	InventoryApps     bool                   `json:"inventory_apps"`
	InventoryEC2      bool                   `json:"inventory_ec2"`
	InventoryInterval int                    `json:"inventory_interval"`
	Options           map[string]interface{} `json:"options"`
}

type listOCOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the Osquery configurations.
func (s *OsqueryConfigurationsServiceOp) List(ctx context.Context, opt *ListOptions) ([]OsqueryConfiguration, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Osquery configuration by id.
func (s *OsqueryConfigurationsServiceOp) GetByID(ctx context.Context, ocID int) (*OsqueryConfiguration, *Response, error) {
	if ocID < 1 {
		return nil, nil, NewArgError("ocID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", ocBasePath, ocID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	oc := new(OsqueryConfiguration)

	resp, err := s.client.Do(ctx, req, oc)
	if err != nil {
		return nil, resp, err
	}

	return oc, resp, err
}

// GetByName retrieves a Osquery configuration by name.
func (s *OsqueryConfigurationsServiceOp) GetByName(ctx context.Context, name string) (*OsqueryConfiguration, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listSCOpt := &listOCOptions{Name: name}

	ocs, resp, err := s.list(ctx, nil, listSCOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(ocs) < 1 {
		return nil, resp, nil
	}

	return &ocs[0], resp, err
}

// Create a new Osquery configuration.
func (s *OsqueryConfigurationsServiceOp) Create(ctx context.Context, createRequest *OsqueryConfigurationRequest) (*OsqueryConfiguration, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, ocBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	oc := new(OsqueryConfiguration)
	resp, err := s.client.Do(ctx, req, oc)
	if err != nil {
		return nil, resp, err
	}

	return oc, resp, err
}

// Update a Osquery configuration.
func (s *OsqueryConfigurationsServiceOp) Update(ctx context.Context, ocID int, updateRequest *OsqueryConfigurationRequest) (*OsqueryConfiguration, *Response, error) {
	if ocID < 1 {
		return nil, nil, NewArgError("ocID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", ocBasePath, ocID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	oc := new(OsqueryConfiguration)
	resp, err := s.client.Do(ctx, req, oc)
	if err != nil {
		return nil, resp, err
	}

	return oc, resp, err
}

// Delete a Osquery configuration.
func (s *OsqueryConfigurationsServiceOp) Delete(ctx context.Context, ocID int) (*Response, error) {
	if ocID < 1 {
		return nil, NewArgError("ocID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", ocBasePath, ocID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Osquery configurations
func (s *OsqueryConfigurationsServiceOp) list(ctx context.Context, opt *ListOptions, ocOpt *listOCOptions) ([]OsqueryConfiguration, *Response, error) {
	path := ocBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, ocOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var ocs []OsqueryConfiguration
	resp, err := s.client.Do(ctx, req, &ocs)
	if err != nil {
		return nil, resp, err
	}

	return ocs, resp, err
}
