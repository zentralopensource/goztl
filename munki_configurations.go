package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const mucBasePath = "munki/configurations/"

// MunkiConfigurationsService is an interface for interfacing with the Munki configurations
// endpoints of the Zentral API
type MunkiConfigurationsService interface {
	List(context.Context, *ListOptions) ([]MunkiConfiguration, *Response, error)
	GetByID(context.Context, int) (*MunkiConfiguration, *Response, error)
	GetByName(context.Context, string) (*MunkiConfiguration, *Response, error)
	Create(context.Context, *MunkiConfigurationRequest) (*MunkiConfiguration, *Response, error)
	Update(context.Context, int, *MunkiConfigurationRequest) (*MunkiConfiguration, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// MunkiConfigurationsServiceOp handles communication with the Munki configurations related
// methods of the Zentral API.
type MunkiConfigurationsServiceOp struct {
	client *Client
}

var _ MunkiConfigurationsService = &MunkiConfigurationsServiceOp{}

// MunkiConfiguration represents a Zentral MunkiConfiguration
type MunkiConfiguration struct {
	ID                              int       `json:"id,omitempty"`
	Name                            string    `json:"name"`
	Description                     string    `json:"description"`
	InventoryAppsFullInfoShard      int       `json:"inventory_apps_full_info_shard"`
	PrincipalUserDetectionSources   []string  `json:"principal_user_detection_sources"`
	PrincipalUserDetectionDomains   []string  `json:"principal_user_detection_domains"`
	CollectedConditionKeys          []string  `json:"collected_condition_keys"`
	ManagedInstallsSyncIntervalDays int       `json:"managed_installs_sync_interval_days"`
	ScriptChecksRunIntervalSeconds  int       `json:"script_checks_run_interval_seconds"`
	AutoReinstallIncidents          bool      `json:"auto_reinstall_incidents"`
	AutoFailedInstallIncidents      bool      `json:"auto_failed_install_incidents"`
	Version                         int       `json:"version"`
	Created                         Timestamp `json:"created_at,omitempty"`
	Updated                         Timestamp `json:"updated_at,omitempty"`
}

func (mc MunkiConfiguration) String() string {
	return Stringify(mc)
}

// MunkiConfigurationRequest represents a request to create or update a Munki configuration
type MunkiConfigurationRequest struct {
	Name                            string   `json:"name"`
	Description                     string   `json:"description"`
	InventoryAppsFullInfoShard      int      `json:"inventory_apps_full_info_shard"`
	PrincipalUserDetectionSources   []string `json:"principal_user_detection_sources"`
	PrincipalUserDetectionDomains   []string `json:"principal_user_detection_domains"`
	CollectedConditionKeys          []string `json:"collected_condition_keys"`
	ManagedInstallsSyncIntervalDays int      `json:"managed_installs_sync_interval_days"`
	ScriptChecksRunIntervalSeconds  int      `json:"script_checks_run_interval_seconds"`
	AutoReinstallIncidents          bool     `json:"auto_reinstall_incidents"`
	AutoFailedInstallIncidents      bool     `json:"auto_failed_install_incidents"`
}

type listMUCOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the Munki configurations.
func (s *MunkiConfigurationsServiceOp) List(ctx context.Context, opt *ListOptions) ([]MunkiConfiguration, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Munki configuration by id.
func (s *MunkiConfigurationsServiceOp) GetByID(ctx context.Context, mcID int) (*MunkiConfiguration, *Response, error) {
	if mcID < 1 {
		return nil, nil, NewArgError("mcID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mucBasePath, mcID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	mc := new(MunkiConfiguration)

	resp, err := s.client.Do(ctx, req, mc)
	if err != nil {
		return nil, resp, err
	}

	return mc, resp, err
}

// GetByName retrieves a Munki configuration by name.
func (s *MunkiConfigurationsServiceOp) GetByName(ctx context.Context, name string) (*MunkiConfiguration, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listMUCOpt := &listMUCOptions{Name: name}

	mcs, resp, err := s.list(ctx, nil, listMUCOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(mcs) < 1 {
		return nil, resp, nil
	}

	return &mcs[0], resp, err
}

// Create a new Munki configuration.
func (s *MunkiConfigurationsServiceOp) Create(ctx context.Context, createRequest *MunkiConfigurationRequest) (*MunkiConfiguration, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, mucBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	mc := new(MunkiConfiguration)
	resp, err := s.client.Do(ctx, req, mc)
	if err != nil {
		return nil, resp, err
	}

	return mc, resp, err
}

// Update a Munki configuration.
func (s *MunkiConfigurationsServiceOp) Update(ctx context.Context, mcID int, updateRequest *MunkiConfigurationRequest) (*MunkiConfiguration, *Response, error) {
	if mcID < 1 {
		return nil, nil, NewArgError("mcID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", mucBasePath, mcID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	mc := new(MunkiConfiguration)
	resp, err := s.client.Do(ctx, req, mc)
	if err != nil {
		return nil, resp, err
	}

	return mc, resp, err
}

// Delete a Munki configuration.
func (s *MunkiConfigurationsServiceOp) Delete(ctx context.Context, mcID int) (*Response, error) {
	if mcID < 1 {
		return nil, NewArgError("mcID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mucBasePath, mcID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Munki configurations
func (s *MunkiConfigurationsServiceOp) list(ctx context.Context, opt *ListOptions, mcOpt *listMUCOptions) ([]MunkiConfiguration, *Response, error) {
	path := mucBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, mcOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var mcs []MunkiConfiguration
	resp, err := s.client.Do(ctx, req, &mcs)
	if err != nil {
		return nil, resp, err
	}

	return mcs, resp, err
}
