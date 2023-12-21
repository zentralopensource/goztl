package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const mscBasePath = "munki/script_checks/"

// MunkiScriptChecksService is an interface for interfacing with the Munki script checks
// endpoints of the Zentral API.
type MunkiScriptChecksService interface {
	List(context.Context, *ListOptions) ([]MunkiScriptCheck, *Response, error)
	GetByID(context.Context, int) (*MunkiScriptCheck, *Response, error)
	GetByName(context.Context, string) (*MunkiScriptCheck, *Response, error)
	Create(context.Context, *MunkiScriptCheckRequest) (*MunkiScriptCheck, *Response, error)
	Update(context.Context, int, *MunkiScriptCheckRequest) (*MunkiScriptCheck, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// MunkiScriptChecksServiceOp handles communication with the Munki script checks related
// methods of the Zentral API.
type MunkiScriptChecksServiceOp struct {
	client *Client
}

var _ MunkiScriptChecksService = &MunkiScriptChecksServiceOp{}

// MunkiScriptCheck represents a Zentral Munki script check.
type MunkiScriptCheck struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Type           string    `json:"type"`
	Source         string    `json:"source"`
	ExpectedResult string    `json:"expected_result"`
	ArchAMD64      bool      `json:"arch_amd64"`
	ArchARM64      bool      `json:"arch_arm64"`
	MinOSVersion   string    `json:"min_os_version"`
	MaxOSVersion   string    `json:"max_os_version"`
	TagIDs         []int     `json:"tags"`
	ExcludedTagIDs []int     `json:"excluded_tags"`
	Version        int       `json:"version"`
	Created        Timestamp `json:"created_at"`
	Updated        Timestamp `json:"updated_at"`
}

// MunkiScriptCheckRequest represents a request to create or update a Munki script check.
type MunkiScriptCheckRequest struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	Type           string `json:"type"`
	Source         string `json:"source"`
	ExpectedResult string `json:"expected_result"`
	ArchAMD64      bool   `json:"arch_amd64"`
	ArchARM64      bool   `json:"arch_arm64"`
	MinOSVersion   string `json:"min_os_version"`
	MaxOSVersion   string `json:"max_os_version"`
	TagIDs         []int  `json:"tags"`
	ExcludedTagIDs []int  `json:"excluded_tags"`
}

func (msc MunkiScriptCheck) String() string {
	return Stringify(msc)
}

type listMunkiScriptCheckOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the Munki script checks.
func (s *MunkiScriptChecksServiceOp) List(ctx context.Context, opt *ListOptions) ([]MunkiScriptCheck, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Munki script check by id.
func (s *MunkiScriptChecksServiceOp) GetByID(ctx context.Context, mscID int) (*MunkiScriptCheck, *Response, error) {
	if mscID < 1 {
		return nil, nil, NewArgError("mscID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mscBasePath, mscID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	msc := new(MunkiScriptCheck)

	resp, err := s.client.Do(ctx, req, msc)
	if err != nil {
		return nil, resp, err
	}

	return msc, resp, err
}

// GetByName retrieves a Munki script check by name.
func (s *MunkiScriptChecksServiceOp) GetByName(ctx context.Context, name string) (*MunkiScriptCheck, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listMunkiScriptCheckOpt := &listMunkiScriptCheckOptions{Name: name}

	mscs, resp, err := s.list(ctx, nil, listMunkiScriptCheckOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(mscs) < 1 {
		return nil, resp, nil
	}

	return &mscs[0], resp, err
}

// Create a new Munki script check.
func (s *MunkiScriptChecksServiceOp) Create(ctx context.Context, createRequest *MunkiScriptCheckRequest) (*MunkiScriptCheck, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, mscBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	msc := new(MunkiScriptCheck)
	resp, err := s.client.Do(ctx, req, msc)
	if err != nil {
		return nil, resp, err
	}

	return msc, resp, err
}

// Update a Munki script check.
func (s *MunkiScriptChecksServiceOp) Update(ctx context.Context, mscID int, updateRequest *MunkiScriptCheckRequest) (*MunkiScriptCheck, *Response, error) {
	if mscID < 1 {
		return nil, nil, NewArgError("mscID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", mscBasePath, mscID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	msc := new(MunkiScriptCheck)
	resp, err := s.client.Do(ctx, req, msc)
	if err != nil {
		return nil, resp, err
	}

	return msc, resp, err
}

// Delete a Munki script check.
func (s *MunkiScriptChecksServiceOp) Delete(ctx context.Context, mscID int) (*Response, error) {
	if mscID < 1 {
		return nil, NewArgError("mscID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mscBasePath, mscID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Munki script checks.
func (s *MunkiScriptChecksServiceOp) list(ctx context.Context, opt *ListOptions, mscOpt *listMunkiScriptCheckOptions) ([]MunkiScriptCheck, *Response, error) {
	path := mscBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, mscOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var mscs []MunkiScriptCheck
	resp, err := s.client.Do(ctx, req, &mscs)
	if err != nil {
		return nil, resp, err
	}

	return mscs, resp, err
}
