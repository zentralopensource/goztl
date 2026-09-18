package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const ssprBasePath = "santa/scoped_path_regexes/"

// SantaScopedPathRegexesService is an interface for interfacing with the Santa scoped path regex
// endpoints of the Zentral API
type SantaScopedPathRegexesService interface {
	GetByID(context.Context, int) (*SantaScopedPathRegex, *Response, error)
	GetByConfigurationID(context.Context, int) ([]SantaScopedPathRegex, *Response, error)
	Create(context.Context, *SantaScopedPathRegexRequest) (*SantaScopedPathRegex, *Response, error)
	Update(context.Context, int, *SantaScopedPathRegexRequest) (*SantaScopedPathRegex, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// SantaScopedPathRegexesServiceOp handles communication with the Santa scoped path regexes related
// methods of the Zentral API.
type SantaScopedPathRegexesServiceOp struct {
	client *Client
}

var _ SantaScopedPathRegexesService = &SantaScopedPathRegexesServiceOp{}

// SantaScopedPathRegex represents a Zentral Santa scoped path regex
type SantaScopedPathRegex struct {
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
	Policy                string    `json:"policy"`
	Regex                 string    `json:"regex"`
	Created               Timestamp `json:"created_at"`
	Updated               Timestamp `json:"updated_at"`
}

func (sspr SantaScopedPathRegex) String() string {
	return Stringify(sspr)
}

// SantaScopedPathRegexRequest represents a request to create or update a Santa scoped path regex
type SantaScopedPathRegexRequest struct {
	ConfigurationID int    `json:"configuration"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	// nothing here is omitted. An update has to carry the scope attributes, because the
	// server reads each one with its counterpart: a stored value read back would answer for
	// a caller that never sent it. An empty dimension goes out as an empty list, because the
	// server refuses the null of a nil one
	SerialNumbers         []string `json:"serial_numbers"`
	ExcludedSerialNumbers []string `json:"excluded_serial_numbers"`
	PrimaryUsers          []string `json:"primary_users"`
	ExcludedPrimaryUsers  []string `json:"excluded_primary_users"`
	TagIDs                []int    `json:"tags"`
	ExcludedTagIDs        []int    `json:"excluded_tags"`
	Policy                string   `json:"policy"`
	Regex                 string   `json:"regex"`
}

type listSSPROptions struct {
	ConfigurationID int `url:"configuration_id,omitempty"`
}

// GetByID retrieves a Santa scoped path regex by id.
func (s *SantaScopedPathRegexesServiceOp) GetByID(ctx context.Context, ssprID int) (*SantaScopedPathRegex, *Response, error) {
	if ssprID < 1 {
		return nil, nil, NewArgError("ssprID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", ssprBasePath, ssprID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	sspr := new(SantaScopedPathRegex)

	resp, err := s.client.Do(ctx, req, sspr)
	if err != nil {
		return nil, resp, err
	}

	return sspr, resp, err
}

// GetByConfigurationID retrieves the Santa scoped path regexes of a given configuration.
func (s *SantaScopedPathRegexesServiceOp) GetByConfigurationID(ctx context.Context, configurationID int) ([]SantaScopedPathRegex, *Response, error) {
	if configurationID < 1 {
		return nil, nil, NewArgError("configurationID", "cannot be less than 1")
	}

	listSSPROpt := &listSSPROptions{ConfigurationID: configurationID}

	return s.list(ctx, listSSPROpt)
}

// Create a new Santa scoped path regex.
func (s *SantaScopedPathRegexesServiceOp) Create(ctx context.Context, createRequest *SantaScopedPathRegexRequest) (*SantaScopedPathRegex, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, ssprBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	sspr := new(SantaScopedPathRegex)
	resp, err := s.client.Do(ctx, req, sspr)
	if err != nil {
		return nil, resp, err
	}

	return sspr, resp, err
}

// Update a Santa scoped path regex.
func (s *SantaScopedPathRegexesServiceOp) Update(ctx context.Context, ssprID int, updateRequest *SantaScopedPathRegexRequest) (*SantaScopedPathRegex, *Response, error) {
	if ssprID < 1 {
		return nil, nil, NewArgError("ssprID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", ssprBasePath, ssprID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	sspr := new(SantaScopedPathRegex)
	resp, err := s.client.Do(ctx, req, sspr)
	if err != nil {
		return nil, resp, err
	}

	return sspr, resp, err
}

// Delete a Santa scoped path regex.
func (s *SantaScopedPathRegexesServiceOp) Delete(ctx context.Context, ssprID int) (*Response, error) {
	if ssprID < 1 {
		return nil, NewArgError("ssprID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", ssprBasePath, ssprID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Santa scoped path regexes
func (s *SantaScopedPathRegexesServiceOp) list(ctx context.Context, ssprOpt *listSSPROptions) ([]SantaScopedPathRegex, *Response, error) {
	path, err := addOptions(ssprBasePath, ssprOpt)
	if err != nil {
		return nil, nil, err
	}
	return resolveAllPages[SantaScopedPathRegex](ctx, s.client, path)
}
