package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const mrpcBasePath = "mdm/recovery_password_configs/"

// MDMRecoveryPasswordConfigsService is an interface for interfacing with the MDM recovery password configuration
// endpoints of the Zentral API
type MDMRecoveryPasswordConfigsService interface {
	List(context.Context, *ListOptions) ([]MDMRecoveryPasswordConfig, *Response, error)
	GetByID(context.Context, int) (*MDMRecoveryPasswordConfig, *Response, error)
	GetByName(context.Context, string) (*MDMRecoveryPasswordConfig, *Response, error)
	Create(context.Context, *MDMRecoveryPasswordConfigRequest) (*MDMRecoveryPasswordConfig, *Response, error)
	Update(context.Context, int, *MDMRecoveryPasswordConfigRequest) (*MDMRecoveryPasswordConfig, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// MDMRecoveryPasswordConfigsServiceOp handles communication with the MDM recovery password configurations related
// methods of the Zentral API.
type MDMRecoveryPasswordConfigsServiceOp struct {
	client *Client
}

var _ MDMRecoveryPasswordConfigsService = &MDMRecoveryPasswordConfigsServiceOp{}

// MDMRecoveryPasswordConfig represents a Zentral MDM recovery password configuration
type MDMRecoveryPasswordConfig struct {
	ID                     int       `json:"id"`
	Name                   string    `json:"name"`
	DynamicPassword        bool      `json:"dynamic_password"`
	StaticPassword         *string   `json:"static_password"`
	RotationIntervalDays   int       `json:"rotation_interval_days"`
	RotateFirmwarePassword bool      `json:"rotate_firmware_password"`
	Created                Timestamp `json:"created_at,omitempty"`
	Updated                Timestamp `json:"updated_at,omitempty"`
}

func (mrpc MDMRecoveryPasswordConfig) String() string {
	return Stringify(mrpc)
}

// MDMRecoveryPasswordConfigRequest represents a request to create or update a MDM recovery password configuration
type MDMRecoveryPasswordConfigRequest struct {
	Name                   string  `json:"name"`
	DynamicPassword        bool    `json:"dynamic_password"`
	StaticPassword         *string `json:"static_password"`
	RotationIntervalDays   int     `json:"rotation_interval_days"`
	RotateFirmwarePassword bool    `json:"rotate_firmware_password"`
}

type listMRPCOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the MDM recovery password configurations.
func (s *MDMRecoveryPasswordConfigsServiceOp) List(ctx context.Context, opt *ListOptions) ([]MDMRecoveryPasswordConfig, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a MDM recovery password configuration by id.
func (s *MDMRecoveryPasswordConfigsServiceOp) GetByID(ctx context.Context, mrpcID int) (*MDMRecoveryPasswordConfig, *Response, error) {
	if mrpcID < 1 {
		return nil, nil, NewArgError("mrpcID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mrpcBasePath, mrpcID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	mrpc := new(MDMRecoveryPasswordConfig)

	resp, err := s.client.Do(ctx, req, mrpc)
	if err != nil {
		return nil, resp, err
	}

	return mrpc, resp, err
}

// GetByName retrieves a MDM recovery password configuration by name.
func (s *MDMRecoveryPasswordConfigsServiceOp) GetByName(ctx context.Context, name string) (*MDMRecoveryPasswordConfig, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listMRPCOpt := &listMRPCOptions{Name: name}

	mrpcs, resp, err := s.list(ctx, nil, listMRPCOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(mrpcs) < 1 {
		return nil, resp, nil
	}

	return &mrpcs[0], resp, err
}

// Create a new MDM recovery password configuration.
func (s *MDMRecoveryPasswordConfigsServiceOp) Create(ctx context.Context, createRequest *MDMRecoveryPasswordConfigRequest) (*MDMRecoveryPasswordConfig, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, mrpcBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	mrpc := new(MDMRecoveryPasswordConfig)
	resp, err := s.client.Do(ctx, req, mrpc)
	if err != nil {
		return nil, resp, err
	}

	return mrpc, resp, err
}

// Update a MDM recovery password configuration.
func (s *MDMRecoveryPasswordConfigsServiceOp) Update(ctx context.Context, mrpcID int, updateRequest *MDMRecoveryPasswordConfigRequest) (*MDMRecoveryPasswordConfig, *Response, error) {
	if mrpcID < 1 {
		return nil, nil, NewArgError("mrpcID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", mrpcBasePath, mrpcID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	mrpc := new(MDMRecoveryPasswordConfig)
	resp, err := s.client.Do(ctx, req, mrpc)
	if err != nil {
		return nil, resp, err
	}

	return mrpc, resp, err
}

// Delete a MDM recovery password configuration.
func (s *MDMRecoveryPasswordConfigsServiceOp) Delete(ctx context.Context, mrpcID int) (*Response, error) {
	if mrpcID < 1 {
		return nil, NewArgError("mrpcID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mrpcBasePath, mrpcID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing MDM recovery password configurations
func (s *MDMRecoveryPasswordConfigsServiceOp) list(ctx context.Context, opt *ListOptions, mrpcOpt *listMRPCOptions) ([]MDMRecoveryPasswordConfig, *Response, error) {
	path := mrpcBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, mrpcOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var mrpcs []MDMRecoveryPasswordConfig
	resp, err := s.client.Do(ctx, req, &mrpcs)
	if err != nil {
		return nil, resp, err
	}

	return mrpcs, resp, err
}
