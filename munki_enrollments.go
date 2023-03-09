package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const mueBasePath = "munki/enrollments/"

// MunkiEnrollmentsService is an interface for interfacing with the Munki enrollments
// endpoints of the Zentral API
type MunkiEnrollmentsService interface {
	List(context.Context, *ListOptions) ([]MunkiEnrollment, *Response, error)
	GetByID(context.Context, int) (*MunkiEnrollment, *Response, error)
	GetByConfigurationID(context.Context, int) ([]MunkiEnrollment, *Response, error)
	Create(context.Context, *MunkiEnrollmentRequest) (*MunkiEnrollment, *Response, error)
	Update(context.Context, int, *MunkiEnrollmentRequest) (*MunkiEnrollment, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// MunkiEnrollmentsServiceOp handles communication with the Munki enrollments related
// methods of the Zentral API.
type MunkiEnrollmentsServiceOp struct {
	client *Client
}

var _ MunkiEnrollmentsService = &MunkiEnrollmentsServiceOp{}

// MunkiEnrollment represents a Zentral MunkiEnrollment
type MunkiEnrollment struct {
	ID                    int              `json:"id"`
	ConfigurationID       int              `json:"configuration"`
	EnrolledMachinesCount int              `json:"enrolled_machines_count"`
	Secret                EnrollmentSecret `json:"secret"`
	PackageURL            string           `json:"package_download_url"`
	Version               int              `json:"version"`
	Created               Timestamp        `json:"created_at,omitempty"`
	Updated               Timestamp        `json:"updated_at,omitempty"`
}

func (se MunkiEnrollment) String() string {
	return Stringify(se)
}

// MunkiEnrollmentRequest represents a request to create or update a Munki enrollment
type MunkiEnrollmentRequest struct {
	ConfigurationID int                     `json:"configuration"`
	Secret          EnrollmentSecretRequest `json:"secret"`
}

type listMUEOptions struct {
	ConfigurationID int `url:"configuration_id,omitempty"`
}

// List lists all the Munki enrollments.
func (s *MunkiEnrollmentsServiceOp) List(ctx context.Context, opt *ListOptions) ([]MunkiEnrollment, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Munki enrollment by id.
func (s *MunkiEnrollmentsServiceOp) GetByID(ctx context.Context, meID int) (*MunkiEnrollment, *Response, error) {
	if meID < 1 {
		return nil, nil, NewArgError("meID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mueBasePath, meID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	se := new(MunkiEnrollment)

	resp, err := s.client.Do(ctx, req, se)
	if err != nil {
		return nil, resp, err
	}

	return se, resp, err
}

// GetByConfigurationID retrieves the Munki enrollments for a given configuration.
func (s *MunkiEnrollmentsServiceOp) GetByConfigurationID(ctx context.Context, configuration_id int) ([]MunkiEnrollment, *Response, error) {
	if configuration_id < 1 {
		return nil, nil, NewArgError("configuration_id", "cannot be negative")
	}

	listSEOpt := &listMUEOptions{ConfigurationID: configuration_id}

	return s.list(ctx, nil, listSEOpt)
}

// Create a new Munki enrollment.
func (s *MunkiEnrollmentsServiceOp) Create(ctx context.Context, createRequest *MunkiEnrollmentRequest) (*MunkiEnrollment, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, mueBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	se := new(MunkiEnrollment)
	resp, err := s.client.Do(ctx, req, se)
	if err != nil {
		return nil, resp, err
	}

	return se, resp, err
}

// Update a Munki enrollment.
func (s *MunkiEnrollmentsServiceOp) Update(ctx context.Context, meID int, updateRequest *MunkiEnrollmentRequest) (*MunkiEnrollment, *Response, error) {
	if meID < 1 {
		return nil, nil, NewArgError("meID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", mueBasePath, meID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	se := new(MunkiEnrollment)
	resp, err := s.client.Do(ctx, req, se)
	if err != nil {
		return nil, resp, err
	}

	return se, resp, err
}

// Delete a Munki enrollment.
func (s *MunkiEnrollmentsServiceOp) Delete(ctx context.Context, meID int) (*Response, error) {
	if meID < 1 {
		return nil, NewArgError("meID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mueBasePath, meID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Munki enrollments
func (s *MunkiEnrollmentsServiceOp) list(ctx context.Context, opt *ListOptions, meOpt *listMUEOptions) ([]MunkiEnrollment, *Response, error) {
	path := mueBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, meOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var ses []MunkiEnrollment
	resp, err := s.client.Do(ctx, req, &ses)
	if err != nil {
		return nil, resp, err
	}

	return ses, resp, err
}
