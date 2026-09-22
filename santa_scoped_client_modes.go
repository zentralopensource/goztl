package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const sscmBasePath = "santa/scoped_client_modes/"

// SantaScopedClientModesService is an interface for interfacing with the Santa scoped client modes
// endpoints of the Zentral API
type SantaScopedClientModesService interface {
	List(context.Context, int, *ListOptions) ([]SantaScopedClientMode, *Response, error)
	GetByID(context.Context, int) (*SantaScopedClientMode, *Response, error)
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

// SantaScopedClientMode represents a Zentral SantaScopedClientMode
type SantaScopedClientMode struct {
	ID                    int       `json:"id"`
	ConfigurationID       int       `json:"configuration"`
	Name                  string    `json:"name"`
	Description           string    `json:"description"`
	ClientMode            int       `json:"client_mode"`
	EventDetailSource     string    `json:"event_detail_source"`
	EventDetailURL        string    `json:"event_detail_url"`
	EventDetailText       string    `json:"event_detail_text"`
	PrimaryUsers          []string  `json:"primary_users"`
	ExcludedPrimaryUsers  []string  `json:"excluded_primary_users"`
	SerialNumbers         []string  `json:"serial_numbers"`
	ExcludedSerialNumbers []string  `json:"excluded_serial_numbers"`
	TagIDs                []int     `json:"tags"`
	ExcludedTagIDs        []int     `json:"excluded_tags"`
	Created               Timestamp `json:"created_at"`
	Updated               Timestamp `json:"updated_at"`
}

func (sscm SantaScopedClientMode) String() string {
	return Stringify(sscm)
}

// SantaScopedClientModeRequest represents a request to create or update a Santa scoped client mode.
// An update is a full update: the scope attributes and the two event detail attributes below are
// read against one another by the server, so they are always written, and a nil slice is sent as
// null and refused. The configuration of an entry cannot change.
type SantaScopedClientModeRequest struct {
	ConfigurationID       int      `json:"configuration"`
	Name                  string   `json:"name"`
	Description           string   `json:"description"`
	ClientMode            int      `json:"client_mode"`
	EventDetailSource     string   `json:"event_detail_source"`
	EventDetailURL        string   `json:"event_detail_url"`
	EventDetailText       string   `json:"event_detail_text"`
	PrimaryUsers          []string `json:"primary_users"`
	ExcludedPrimaryUsers  []string `json:"excluded_primary_users"`
	SerialNumbers         []string `json:"serial_numbers"`
	ExcludedSerialNumbers []string `json:"excluded_serial_numbers"`
	TagIDs                []int    `json:"tags"`
	ExcludedTagIDs        []int    `json:"excluded_tags"`
}

type listSSCMOptions struct {
	ConfigurationID int `url:"configuration_id,omitempty"`
}

// List lists the Santa scoped client modes of a configuration. The endpoint answers the entries of
// one configuration, so it has no listing of its own without one.
func (s *SantaScopedClientModesServiceOp) List(ctx context.Context, cfgID int, opt *ListOptions) ([]SantaScopedClientMode, *Response, error) {
	if cfgID < 1 {
		return nil, nil, NewArgError("cfgID", "cannot be less than 1")
	}

	return s.list(ctx, opt, &listSSCMOptions{ConfigurationID: cfgID})
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

// Create a new Santa scoped client mode
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

// Update a Santa scoped client mode
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

// Delete a Santa scoped client mode
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
func (s *SantaScopedClientModesServiceOp) list(ctx context.Context, opt *ListOptions, sscmOpt *listSSCMOptions) ([]SantaScopedClientMode, *Response, error) {
	path := sscmBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, sscmOpt)
	if err != nil {
		return nil, nil, err
	}
	return resolveAllPages[SantaScopedClientMode](ctx, s.client, path)
}
