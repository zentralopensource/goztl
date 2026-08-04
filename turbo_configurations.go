package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const tconfBasePath = "turbo/configurations/"

// TurboConfigurationsService is an interface for interfacing with the Turbo configurations
// endpoints of the Zentral API
type TurboConfigurationsService interface {
	List(context.Context, *ListOptions) ([]TurboConfiguration, *Response, error)
	GetByID(context.Context, string) (*TurboConfiguration, *Response, error)
	GetByName(context.Context, string) (*TurboConfiguration, *Response, error)
	Create(context.Context, *TurboConfigurationRequest) (*TurboConfiguration, *Response, error)
	Update(context.Context, string, *TurboConfigurationRequest) (*TurboConfiguration, *Response, error)
	Delete(context.Context, string) (*Response, error)
}

// TurboConfigurationsServiceOp handles communication with the Turbo configurations related
// methods of the Zentral API.
type TurboConfigurationsServiceOp struct {
	client *Client
}

var _ TurboConfigurationsService = &TurboConfigurationsServiceOp{}

// TurboConfiguration represents a Zentral Turbo configuration
type TurboConfiguration struct {
	ID                    string    `json:"id,omitempty"`
	Name                  string    `json:"name"`
	Description           string    `json:"description"`
	CollectInventory      bool      `json:"collect_inventory"`
	InventoryInterval     int       `json:"inventory_interval"`
	DefaultCheckInterval  int       `json:"default_check_interval"`
	ConfigRefreshInterval int       `json:"config_refresh_interval"`
	ResultsBatchSize      int       `json:"results_batch_size"`
	Created               Timestamp `json:"created_at,omitempty"`
	Updated               Timestamp `json:"updated_at,omitempty"`
}

func (tc TurboConfiguration) String() string {
	return Stringify(tc)
}

// TurboConfigurationRequest represents a request to create or update a Turbo configuration
type TurboConfigurationRequest struct {
	Name                  string `json:"name"`
	Description           string `json:"description"`
	CollectInventory      bool   `json:"collect_inventory"`
	InventoryInterval     int    `json:"inventory_interval"`
	DefaultCheckInterval  int    `json:"default_check_interval"`
	ConfigRefreshInterval int    `json:"config_refresh_interval"`
	ResultsBatchSize      int    `json:"results_batch_size"`
}

type listTConfOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the Turbo configurations.
func (s *TurboConfigurationsServiceOp) List(ctx context.Context, opt *ListOptions) ([]TurboConfiguration, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Turbo configuration by id.
func (s *TurboConfigurationsServiceOp) GetByID(ctx context.Context, tcID string) (*TurboConfiguration, *Response, error) {
	if len(tcID) < 1 {
		return nil, nil, NewArgError("tcID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", tconfBasePath, tcID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	tc := new(TurboConfiguration)

	resp, err := s.client.Do(ctx, req, tc)
	if err != nil {
		return nil, resp, err
	}

	return tc, resp, err
}

// GetByName retrieves a Turbo configuration by name.
func (s *TurboConfigurationsServiceOp) GetByName(ctx context.Context, name string) (*TurboConfiguration, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listTConfOpt := &listTConfOptions{Name: name}

	tcs, resp, err := s.list(ctx, nil, listTConfOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(tcs) < 1 {
		return nil, resp, nil
	}

	return &tcs[0], resp, err
}

// Create a new Turbo configuration.
func (s *TurboConfigurationsServiceOp) Create(ctx context.Context, createRequest *TurboConfigurationRequest) (*TurboConfiguration, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, tconfBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	tc := new(TurboConfiguration)
	resp, err := s.client.Do(ctx, req, tc)
	if err != nil {
		return nil, resp, err
	}

	return tc, resp, err
}

// Update a Turbo configuration.
func (s *TurboConfigurationsServiceOp) Update(ctx context.Context, tcID string, updateRequest *TurboConfigurationRequest) (*TurboConfiguration, *Response, error) {
	if len(tcID) < 1 {
		return nil, nil, NewArgError("tcID", "cannot be blank")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%s/", tconfBasePath, tcID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	tc := new(TurboConfiguration)
	resp, err := s.client.Do(ctx, req, tc)
	if err != nil {
		return nil, resp, err
	}

	return tc, resp, err
}

// Delete a Turbo configuration.
func (s *TurboConfigurationsServiceOp) Delete(ctx context.Context, tcID string) (*Response, error) {
	if len(tcID) < 1 {
		return nil, NewArgError("tcID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", tconfBasePath, tcID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Turbo configurations
func (s *TurboConfigurationsServiceOp) list(ctx context.Context, opt *ListOptions, tcOpt *listTConfOptions) ([]TurboConfiguration, *Response, error) {
	path := tconfBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, tcOpt)
	if err != nil {
		return nil, nil, err
	}

	return resolveAllPages[TurboConfiguration](ctx, s.client, path)
}
