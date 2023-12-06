package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const msueBasePath = "mdm/software_update_enforcements/"

// MDMSoftwareUpdateEnforcementsService is an interface for interfacing with the MDM software update enforcement
// endpoints of the Zentral API
type MDMSoftwareUpdateEnforcementsService interface {
	List(context.Context, *ListOptions) ([]MDMSoftwareUpdateEnforcement, *Response, error)
	GetByID(context.Context, int) (*MDMSoftwareUpdateEnforcement, *Response, error)
	GetByName(context.Context, string) (*MDMSoftwareUpdateEnforcement, *Response, error)
	Create(context.Context, *MDMSoftwareUpdateEnforcementRequest) (*MDMSoftwareUpdateEnforcement, *Response, error)
	Update(context.Context, int, *MDMSoftwareUpdateEnforcementRequest) (*MDMSoftwareUpdateEnforcement, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// MDMSoftwareUpdateEnforcementsServiceOp handles communication with the MDM software update enforcements related
// methods of the Zentral API.
type MDMSoftwareUpdateEnforcementsServiceOp struct {
	client *Client
}

var _ MDMSoftwareUpdateEnforcementsService = &MDMSoftwareUpdateEnforcementsServiceOp{}

// MDMSoftwareUpdateEnforcement represents a Zentral MDM software update enforcement
type MDMSoftwareUpdateEnforcement struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	DetailsURL    string    `json:"details_url"`
	Platforms     []string  `json:"platforms"`
	TagIDs        []int     `json:"tags"`
	OSVersion     string    `json:"os_version"`
	BuildVersion  string    `json:"build_version"`
	LocalDateTime *string   `json:"local_datetime"`
	MaxOSVersion  string    `json:"max_os_version"`
	DelayDays     *int      `json:"delay_days"`
	LocalTime     *string   `json:"local_time"`
	Created       Timestamp `json:"created_at,omitempty"`
	Updated       Timestamp `json:"updated_at,omitempty"`
}

func (msue MDMSoftwareUpdateEnforcement) String() string {
	return Stringify(msue)
}

// MDMSoftwareUpdateEnforcementRequest represents a request to create or update a MDM software update enforcement
type MDMSoftwareUpdateEnforcementRequest struct {
	Name          string   `json:"name"`
	DetailsURL    string   `json:"details_url"`
	Platforms     []string `json:"platforms"`
	TagIDs        []int    `json:"tags"`
	OSVersion     string   `json:"os_version"`
	BuildVersion  string   `json:"build_version"`
	LocalDateTime *string  `json:"local_datetime"`
	MaxOSVersion  string   `json:"max_os_version"`
	DelayDays     *int     `json:"delay_days"`
	LocalTime     *string  `json:"local_time"`
}

type listMSUEOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the MDM software update enforcements.
func (s *MDMSoftwareUpdateEnforcementsServiceOp) List(ctx context.Context, opt *ListOptions) ([]MDMSoftwareUpdateEnforcement, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a MDM software update enforcement by id.
func (s *MDMSoftwareUpdateEnforcementsServiceOp) GetByID(ctx context.Context, msueID int) (*MDMSoftwareUpdateEnforcement, *Response, error) {
	if msueID < 1 {
		return nil, nil, NewArgError("msueID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", msueBasePath, msueID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	msue := new(MDMSoftwareUpdateEnforcement)

	resp, err := s.client.Do(ctx, req, msue)
	if err != nil {
		return nil, resp, err
	}

	return msue, resp, err
}

// GetByName retrieves a MDM software update enforcement by name.
func (s *MDMSoftwareUpdateEnforcementsServiceOp) GetByName(ctx context.Context, name string) (*MDMSoftwareUpdateEnforcement, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listMSUEOpt := &listMSUEOptions{Name: name}

	msues, resp, err := s.list(ctx, nil, listMSUEOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(msues) < 1 {
		return nil, resp, nil
	}

	return &msues[0], resp, err
}

// Create a new MDM software update enforcement.
func (s *MDMSoftwareUpdateEnforcementsServiceOp) Create(ctx context.Context, createRequest *MDMSoftwareUpdateEnforcementRequest) (*MDMSoftwareUpdateEnforcement, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, msueBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	msue := new(MDMSoftwareUpdateEnforcement)
	resp, err := s.client.Do(ctx, req, msue)
	if err != nil {
		return nil, resp, err
	}

	return msue, resp, err
}

// Update a MDM software update enforcement.
func (s *MDMSoftwareUpdateEnforcementsServiceOp) Update(ctx context.Context, msueID int, updateRequest *MDMSoftwareUpdateEnforcementRequest) (*MDMSoftwareUpdateEnforcement, *Response, error) {
	if msueID < 1 {
		return nil, nil, NewArgError("msueID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", msueBasePath, msueID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	msue := new(MDMSoftwareUpdateEnforcement)
	resp, err := s.client.Do(ctx, req, msue)
	if err != nil {
		return nil, resp, err
	}

	return msue, resp, err
}

// Delete a MDM software update enforcement.
func (s *MDMSoftwareUpdateEnforcementsServiceOp) Delete(ctx context.Context, msueID int) (*Response, error) {
	if msueID < 1 {
		return nil, NewArgError("msueID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", msueBasePath, msueID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing MDM software update enforcements
func (s *MDMSoftwareUpdateEnforcementsServiceOp) list(ctx context.Context, opt *ListOptions, msueOpt *listMSUEOptions) ([]MDMSoftwareUpdateEnforcement, *Response, error) {
	path := msueBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, msueOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var msues []MDMSoftwareUpdateEnforcement
	resp, err := s.client.Do(ctx, req, &msues)
	if err != nil {
		return nil, resp, err
	}

	return msues, resp, err
}
