package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const scBasePath = "santa/configurations/"

// SantaConfigurationsService is an interface for interfacing with the Santa configurations
// endpoints of the Zentral API
type SantaConfigurationsService interface {
	List(context.Context, *ListOptions) ([]SantaConfiguration, *Response, error)
	GetByID(context.Context, int) (*SantaConfiguration, *Response, error)
	GetByName(context.Context, string) (*SantaConfiguration, *Response, error)
	Create(context.Context, *SantaConfigurationRequest) (*SantaConfiguration, *Response, error)
	Update(context.Context, int, *SantaConfigurationRequest) (*SantaConfiguration, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// SantaConfigurationsServiceOp handles communication with the Santa configurations related
// methods of the Zentral API.
type SantaConfigurationsServiceOp struct {
	client *Client
}

var _ SantaConfigurationsService = &SantaConfigurationsServiceOp{}

// SantaConfiguration represents a Zentral SantaConfiguration
type SantaConfiguration struct {
	ID                        int       `json:"id,omitempty"`
	Name                      string    `json:"name"`
	ClientMode                int       `json:"client_mode"`
	ClientCertificateAuth     bool      `json:"client_certificate_auth"`
	BatchSize                 int       `json:"batch_size"`
	FullSyncInterval          int       `json:"full_sync_interval"`
	EnableBundles             bool      `json:"enable_bundles"`
	EnableTransitiveRules     bool      `json:"enable_transitive_rules"`
	AllowedPathRegex          string    `json:"allowed_path_regex"`
	BlockedPathRegex          string    `json:"blocked_path_regex"`
	BlockUSBMount             bool      `json:"block_usb_mount"`
	RemountUSBMode            []string  `json:"remount_usb_mode"`
	AllowUnknownShard         int       `json:"allow_unknown_shard"`
	EnableAllEventUploadShard int       `json:"enable_all_event_upload_shard"`
	SyncIncidentSeverity      int       `json:"sync_incident_severity"`
	Created                   Timestamp `json:"created_at,omitempty"`
	Updated                   Timestamp `json:"updated_at,omitempty"`
}

func (sc SantaConfiguration) String() string {
	return Stringify(sc)
}

// SantaConfigurationRequest represents a request to create or update a Santa configuration
type SantaConfigurationRequest struct {
	Name                      string   `json:"name"`
	ClientMode                int      `json:"client_mode"`
	ClientCertificateAuth     bool     `json:"client_certificate_auth"`
	BatchSize                 int      `json:"batch_size"`
	FullSyncInterval          int      `json:"full_sync_interval"`
	EnableBundles             bool     `json:"enable_bundles"`
	EnableTransitiveRules     bool     `json:"enable_transitive_rules"`
	AllowedPathRegex          string   `json:"allowed_path_regex"`
	BlockedPathRegex          string   `json:"blocked_path_regex"`
	BlockUSBMount             bool     `json:"block_usb_mount"`
	RemountUSBMode            []string `json:"remount_usb_mode"`
	AllowUnknownShard         int      `json:"allow_unknown_shard"`
	EnableAllEventUploadShard int      `json:"enable_all_even_upload_shard"`
	SyncIncidentSeverity      int      `json:"sync_incident_severity"`
}

type listSCOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the Santa configurations.
func (s *SantaConfigurationsServiceOp) List(ctx context.Context, opt *ListOptions) ([]SantaConfiguration, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Santa configuration by id.
func (s *SantaConfigurationsServiceOp) GetByID(ctx context.Context, scID int) (*SantaConfiguration, *Response, error) {
	if scID < 1 {
		return nil, nil, NewArgError("scID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", scBasePath, scID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	sc := new(SantaConfiguration)

	resp, err := s.client.Do(ctx, req, sc)
	if err != nil {
		return nil, resp, err
	}

	return sc, resp, err
}

// GetByName retrieves a Santa configuration by name.
func (s *SantaConfigurationsServiceOp) GetByName(ctx context.Context, name string) (*SantaConfiguration, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listSCOpt := &listSCOptions{Name: name}

	scs, resp, err := s.list(ctx, nil, listSCOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(scs) < 1 {
		return nil, resp, nil
	}

	return &scs[0], resp, err
}

// Create a new Santa configuration.
func (s *SantaConfigurationsServiceOp) Create(ctx context.Context, createRequest *SantaConfigurationRequest) (*SantaConfiguration, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, scBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	sc := new(SantaConfiguration)
	resp, err := s.client.Do(ctx, req, sc)
	if err != nil {
		return nil, resp, err
	}

	return sc, resp, err
}

// Update a Santa configuration.
func (s *SantaConfigurationsServiceOp) Update(ctx context.Context, scID int, updateRequest *SantaConfigurationRequest) (*SantaConfiguration, *Response, error) {
	if scID < 1 {
		return nil, nil, NewArgError("scID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", scBasePath, scID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	sc := new(SantaConfiguration)
	resp, err := s.client.Do(ctx, req, sc)
	if err != nil {
		return nil, resp, err
	}

	return sc, resp, err
}

// Delete a Santa configuration.
func (s *SantaConfigurationsServiceOp) Delete(ctx context.Context, scID int) (*Response, error) {
	if scID < 1 {
		return nil, NewArgError("scID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", scBasePath, scID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Santa configurations
func (s *SantaConfigurationsServiceOp) list(ctx context.Context, opt *ListOptions, scOpt *listSCOptions) ([]SantaConfiguration, *Response, error) {
	path := scBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, scOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var scs []SantaConfiguration
	resp, err := s.client.Do(ctx, req, &scs)
	if err != nil {
		return nil, resp, err
	}

	return scs, resp, err
}
