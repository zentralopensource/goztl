package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const sscmBasePath = "santa/scoped_client_modes/"

// SantaScopedClientModesService is an interface for interfacing with the Santa scoped client mode
// endpoints of the Zentral API
type SantaScopedClientModesService interface {
	GetByID(context.Context, int) (*SantaScopedClientMode, *Response, error)
	GetByConfigurationID(context.Context, int) ([]SantaScopedClientMode, *Response, error)
	Create(context.Context, *SantaScopedClientModeRequest) (*SantaScopedClientMode, *Response, error)
	Update(context.Context, int, *SantaScopedClientModeRequest) (*SantaScopedClientMode, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// SantaScopedClientModesServiceOp handles communication with the Santa scoped client modes related
// methods of the Zentral API.
type SantaScopedClientModesServiceOp struct {
	client *Client
}

var _ SantaScopedClientModesService = &SantaScopedClientModesServiceOp{}

// SantaScopedClientMode represents a Zentral Santa scoped client mode
type SantaScopedClientMode struct {
	ID                    int       `json:"id"`
	ConfigurationID       int       `json:"configuration"`
	Name                  string    `json:"name"`
	Description           string    `json:"description"`
	SerialNumbers         []string  `json:"serial_numbers"`
	ExcludedSerialNumbers []string  `json:"excluded_serial_numbers"`
	PrimaryUsers          []string  `json:"primary_users"`
	ExcludedPrimaryUsers  []string  `json:"excluded_primary_users"`
	TagIDs                []int     `json:"tags"`
	ExcludedTagIDs        []int     `json:"excluded_tags"`
	ClientMode            int       `json:"client_mode"`
	EventDetailSource     string    `json:"event_detail_source"`
	EventDetailURL        string    `json:"event_detail_url"`
	EventDetailText       string    `json:"event_detail_text"`
	Created               Timestamp `json:"created_at"`
	Updated               Timestamp `json:"updated_at"`
}

func (sscm SantaScopedClientMode) String() string {
	return Stringify(sscm)
}

// SantaScopedClientModeRequest represents a request to create or update a Santa scoped client mode
type SantaScopedClientModeRequest struct {
	ConfigurationID int    `json:"configuration"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	// nothing here is omitted. An update has to carry the scope attributes, and the event
	// detail source and URL below, because the server reads each one with another one: a
	// stored value read back would answer for a caller that never sent it. An empty
	// dimension goes out as an empty list, because the server refuses the null of a nil one
	SerialNumbers         []string `json:"serial_numbers"`
	ExcludedSerialNumbers []string `json:"excluded_serial_numbers"`
	PrimaryUsers          []string `json:"primary_users"`
	ExcludedPrimaryUsers  []string `json:"excluded_primary_users"`
	TagIDs                []int    `json:"tags"`
	ExcludedTagIDs        []int    `json:"excluded_tags"`
	ClientMode            int      `json:"client_mode"`
	EventDetailSource     string   `json:"event_detail_source"`
	EventDetailURL        string   `json:"event_detail_url"`
	EventDetailText       string   `json:"event_detail_text"`
}

type listSSCMOptions struct {
	ConfigurationID int `url:"configuration_id,omitempty"`
}

// GetByID retrieves a Santa scoped client mode by id.
func (s *SantaScopedClientModesServiceOp) GetByID(ctx context.Context, sscmID int) (*SantaScopedClientMode, *Response, error) {
	if sscmID < 1 {
		return nil, nil, NewArgError("sscmID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", sscmBasePath, sscmID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	sscm := new(SantaScopedClientMode)

	resp, err := s.client.Do(ctx, req, sscm)
	if err != nil {
		return nil, resp, err
	}

	return sscm, resp, err
}

// GetByConfigurationID retrieves the Santa scoped client modes of a given configuration.
func (s *SantaScopedClientModesServiceOp) GetByConfigurationID(ctx context.Context, configurationID int) ([]SantaScopedClientMode, *Response, error) {
	if configurationID < 1 {
		return nil, nil, NewArgError("configurationID", "cannot be less than 1")
	}

	listSSCMOpt := &listSSCMOptions{ConfigurationID: configurationID}

	return s.list(ctx, listSSCMOpt)
}

// Create a new Santa scoped client mode.
func (s *SantaScopedClientModesServiceOp) Create(ctx context.Context, createRequest *SantaScopedClientModeRequest) (*SantaScopedClientMode, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, sscmBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	sscm := new(SantaScopedClientMode)
	resp, err := s.client.Do(ctx, req, sscm)
	if err != nil {
		return nil, resp, err
	}

	return sscm, resp, err
}

// Update a Santa scoped client mode.
func (s *SantaScopedClientModesServiceOp) Update(ctx context.Context, sscmID int, updateRequest *SantaScopedClientModeRequest) (*SantaScopedClientMode, *Response, error) {
	if sscmID < 1 {
		return nil, nil, NewArgError("sscmID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", sscmBasePath, sscmID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	sscm := new(SantaScopedClientMode)
	resp, err := s.client.Do(ctx, req, sscm)
	if err != nil {
		return nil, resp, err
	}

	return sscm, resp, err
}

// Delete a Santa scoped client mode.
func (s *SantaScopedClientModesServiceOp) Delete(ctx context.Context, sscmID int) (*Response, error) {
	if sscmID < 1 {
		return nil, NewArgError("sscmID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", sscmBasePath, sscmID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Santa scoped client modes
func (s *SantaScopedClientModesServiceOp) list(ctx context.Context, sscmOpt *listSSCMOptions) ([]SantaScopedClientMode, *Response, error) {
	path, err := addOptions(sscmBasePath, sscmOpt)
	if err != nil {
		return nil, nil, err
	}
	return resolveAllPages[SantaScopedClientMode](ctx, s.client, path)
}
