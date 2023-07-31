package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const mfcBasePath = "mdm/filevault_configs/"

// MDMFileVaultConfigsService is an interface for interfacing with the MDM FileVault configuration
// endpoints of the Zentral API
type MDMFileVaultConfigsService interface {
	List(context.Context, *ListOptions) ([]MDMFileVaultConfig, *Response, error)
	GetByID(context.Context, int) (*MDMFileVaultConfig, *Response, error)
	GetByName(context.Context, string) (*MDMFileVaultConfig, *Response, error)
	Create(context.Context, *MDMFileVaultConfigRequest) (*MDMFileVaultConfig, *Response, error)
	Update(context.Context, int, *MDMFileVaultConfigRequest) (*MDMFileVaultConfig, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// MDMFileVaultConfigsServiceOp handles communication with the MDM FileVault configurations related
// methods of the Zentral API.
type MDMFileVaultConfigsServiceOp struct {
	client *Client
}

var _ MDMFileVaultConfigsService = &MDMFileVaultConfigsServiceOp{}

// MDMFileVaultConfig represents a Zentral MDM FileVault configuration
type MDMFileVaultConfig struct {
	ID                        int       `json:"id,omitempty"`
	Name                      string    `json:"name"`
	EscrowLocationDisplayName string    `json:"escrow_location_display_name"`
	AtLoginOnly               bool      `json:"at_login_only"`
	BypassAttempts            int       `json:"bypass_attempts"`
	ShowRecoveryKey           bool      `json:"show_recovery_key"`
	DestroyKeyOnStandby       bool      `json:"destroy_key_on_standby"`
	PRKRotationIntervalDays   int       `json:"prk_rotation_interval_days"`
	Created                   Timestamp `json:"created_at,omitempty"`
	Updated                   Timestamp `json:"updated_at,omitempty"`
}

func (mfc MDMFileVaultConfig) String() string {
	return Stringify(mfc)
}

// MDMFileVaultConfigRequest represents a request to create or update a MDM FileVault configuration
type MDMFileVaultConfigRequest struct {
	Name                      string `json:"name"`
	EscrowLocationDisplayName string `json:"escrow_location_display_name"`
	AtLoginOnly               bool   `json:"at_login_only"`
	BypassAttempts            int    `json:"bypass_attempts"`
	ShowRecoveryKey           bool   `json:"show_recovery_key"`
	DestroyKeyOnStandby       bool   `json:"destroy_key_on_standby"`
	PRKRotationIntervalDays   int    `json:"prk_rotation_interval_days"`
}

type listMFCOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the MDM FileVault configurations.
func (s *MDMFileVaultConfigsServiceOp) List(ctx context.Context, opt *ListOptions) ([]MDMFileVaultConfig, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a MDM FileVault configuration by id.
func (s *MDMFileVaultConfigsServiceOp) GetByID(ctx context.Context, mfcID int) (*MDMFileVaultConfig, *Response, error) {
	if mfcID < 1 {
		return nil, nil, NewArgError("mfcID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mfcBasePath, mfcID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	mfc := new(MDMFileVaultConfig)

	resp, err := s.client.Do(ctx, req, mfc)
	if err != nil {
		return nil, resp, err
	}

	return mfc, resp, err
}

// GetByName retrieves a MDM FileVault configuration by name.
func (s *MDMFileVaultConfigsServiceOp) GetByName(ctx context.Context, name string) (*MDMFileVaultConfig, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listMFCOpt := &listMFCOptions{Name: name}

	mfcs, resp, err := s.list(ctx, nil, listMFCOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(mfcs) < 1 {
		return nil, resp, nil
	}

	return &mfcs[0], resp, err
}

// Create a new MDM FileVault configuration.
func (s *MDMFileVaultConfigsServiceOp) Create(ctx context.Context, createRequest *MDMFileVaultConfigRequest) (*MDMFileVaultConfig, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, mfcBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	mfc := new(MDMFileVaultConfig)
	resp, err := s.client.Do(ctx, req, mfc)
	if err != nil {
		return nil, resp, err
	}

	return mfc, resp, err
}

// Update a MDM FileVault configuration.
func (s *MDMFileVaultConfigsServiceOp) Update(ctx context.Context, mfcID int, updateRequest *MDMFileVaultConfigRequest) (*MDMFileVaultConfig, *Response, error) {
	if mfcID < 1 {
		return nil, nil, NewArgError("mfcID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", mfcBasePath, mfcID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	mfc := new(MDMFileVaultConfig)
	resp, err := s.client.Do(ctx, req, mfc)
	if err != nil {
		return nil, resp, err
	}

	return mfc, resp, err
}

// Delete a MDM FileVault configuration.
func (s *MDMFileVaultConfigsServiceOp) Delete(ctx context.Context, mfcID int) (*Response, error) {
	if mfcID < 1 {
		return nil, NewArgError("mfcID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mfcBasePath, mfcID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing MDM FileVault configurations
func (s *MDMFileVaultConfigsServiceOp) list(ctx context.Context, opt *ListOptions, mfcOpt *listMFCOptions) ([]MDMFileVaultConfig, *Response, error) {
	path := mfcBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, mfcOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var mfcs []MDMFileVaultConfig
	resp, err := s.client.Do(ctx, req, &mfcs)
	if err != nil {
		return nil, resp, err
	}

	return mfcs, resp, err
}
