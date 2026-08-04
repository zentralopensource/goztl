package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const tscrBasePath = "turbo/scripts/"

// TurboScriptsService is an interface for interfacing with the Turbo scripts
// endpoints of the Zentral API
type TurboScriptsService interface {
	List(context.Context, *ListOptions) ([]TurboScript, *Response, error)
	GetByID(context.Context, string) (*TurboScript, *Response, error)
	GetByName(context.Context, string) (*TurboScript, *Response, error)
	Create(context.Context, *TurboScriptRequest) (*TurboScript, *Response, error)
	Update(context.Context, string, *TurboScriptRequest) (*TurboScript, *Response, error)
	Delete(context.Context, string) (*Response, error)
}

// TurboScriptsServiceOp handles communication with the Turbo scripts related
// methods of the Zentral API.
type TurboScriptsServiceOp struct {
	client *Client
}

var _ TurboScriptsService = &TurboScriptsServiceOp{}

// TurboScript represents a Zentral Turbo script
type TurboScript struct {
	ID                     string    `json:"id,omitempty"`
	Name                   string    `json:"name"`
	Description            string    `json:"description"`
	Source                 string    `json:"source"`
	TagID                  *int      `json:"tag"`
	ArchAMD64              bool      `json:"arch_amd64"`
	ArchARM64              bool      `json:"arch_arm64"`
	MinOSVersion           string    `json:"min_os_version"`
	MaxOSVersion           string    `json:"max_os_version"`
	Version                int       `json:"version"`
	JobID                  string    `json:"job_id"`
	ComplianceCheckEnabled bool      `json:"compliance_check_enabled"`
	ComplianceCheckID      *int      `json:"compliance_check_id"`
	Created                Timestamp `json:"created_at,omitempty"`
	Updated                Timestamp `json:"updated_at,omitempty"`
}

func (ts TurboScript) String() string {
	return Stringify(ts)
}

// TurboScriptRequest represents a request to create or update a Turbo script
type TurboScriptRequest struct {
	Name                   string `json:"name"`
	Description            string `json:"description"`
	Source                 string `json:"source"`
	TagID                  *int   `json:"tag"`
	ArchAMD64              bool   `json:"arch_amd64"`
	ArchARM64              bool   `json:"arch_arm64"`
	MinOSVersion           string `json:"min_os_version"`
	MaxOSVersion           string `json:"max_os_version"`
	ComplianceCheckEnabled bool   `json:"compliance_check_enabled"`
}

type listTScrOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the Turbo scripts.
func (s *TurboScriptsServiceOp) List(ctx context.Context, opt *ListOptions) ([]TurboScript, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Turbo script by id.
func (s *TurboScriptsServiceOp) GetByID(ctx context.Context, tsID string) (*TurboScript, *Response, error) {
	if len(tsID) < 1 {
		return nil, nil, NewArgError("tsID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", tscrBasePath, tsID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	ts := new(TurboScript)

	resp, err := s.client.Do(ctx, req, ts)
	if err != nil {
		return nil, resp, err
	}

	return ts, resp, err
}

// GetByName retrieves a Turbo script by name.
func (s *TurboScriptsServiceOp) GetByName(ctx context.Context, name string) (*TurboScript, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listTScrOpt := &listTScrOptions{Name: name}

	tss, resp, err := s.list(ctx, nil, listTScrOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(tss) < 1 {
		return nil, resp, nil
	}

	return &tss[0], resp, err
}

// Create a new Turbo script.
func (s *TurboScriptsServiceOp) Create(ctx context.Context, createRequest *TurboScriptRequest) (*TurboScript, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, tscrBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	ts := new(TurboScript)
	resp, err := s.client.Do(ctx, req, ts)
	if err != nil {
		return nil, resp, err
	}

	return ts, resp, err
}

// Update a Turbo script.
func (s *TurboScriptsServiceOp) Update(ctx context.Context, tsID string, updateRequest *TurboScriptRequest) (*TurboScript, *Response, error) {
	if len(tsID) < 1 {
		return nil, nil, NewArgError("tsID", "cannot be blank")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%s/", tscrBasePath, tsID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	ts := new(TurboScript)
	resp, err := s.client.Do(ctx, req, ts)
	if err != nil {
		return nil, resp, err
	}

	return ts, resp, err
}

// Delete a Turbo script.
func (s *TurboScriptsServiceOp) Delete(ctx context.Context, tsID string) (*Response, error) {
	if len(tsID) < 1 {
		return nil, NewArgError("tsID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", tscrBasePath, tsID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Turbo scripts
func (s *TurboScriptsServiceOp) list(ctx context.Context, opt *ListOptions, tsOpt *listTScrOptions) ([]TurboScript, *Response, error) {
	path := tscrBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, tsOpt)
	if err != nil {
		return nil, nil, err
	}

	return resolveAllPages[TurboScript](ctx, s.client, path)
}
