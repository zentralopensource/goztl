package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const oeBasePath = "osquery/enrollments/"

// OsqueryEnrollmentsService is an interface for interfacing with the Osquery enrollments
// endpoints of the Zentral API
type OsqueryEnrollmentsService interface {
	List(context.Context, *ListOptions) ([]OsqueryEnrollment, *Response, error)
	GetByID(context.Context, int) (*OsqueryEnrollment, *Response, error)
	GetByConfigurationID(context.Context, int) ([]OsqueryEnrollment, *Response, error)
	Create(context.Context, *OsqueryEnrollmentRequest) (*OsqueryEnrollment, *Response, error)
	Update(context.Context, int, *OsqueryEnrollmentRequest) (*OsqueryEnrollment, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// OsqueryEnrollmentsServiceOp handles communication with the Osquery enrollments related
// methods of the Zentral API.
type OsqueryEnrollmentsServiceOp struct {
	client *Client
}

var _ OsqueryEnrollmentsService = &OsqueryEnrollmentsServiceOp{}

// OsqueryEnrollment represents a Zentral OsqueryEnrollment
type OsqueryEnrollment struct {
	ID                    int              `json:"id"`
	ConfigurationID       int              `json:"configuration"`
	OsqueryRelease        string           `json:"osquery_release"`
	EnrolledMachinesCount int              `json:"enrolled_machines_count"`
	Secret                EnrollmentSecret `json:"secret"`
	PackageURL            string           `json:"package_download_url"`
	ScriptURL             string           `json:"script_download_url"`
	PowershellScriptURL   string           `json:"powershell_script_download_url"`
	Version               int              `json:"version"`
	Created               Timestamp        `json:"created_at,omitempty"`
	Updated               Timestamp        `json:"updated_at,omitempty"`
}

func (se OsqueryEnrollment) String() string {
	return Stringify(se)
}

// OsqueryEnrollmentRequest represents a request to create or update a Osquery enrollment
type OsqueryEnrollmentRequest struct {
	ConfigurationID int                     `json:"configuration"`
	OsqueryRelease  string                  `json:"osquery_release"`
	Secret          EnrollmentSecretRequest `json:"secret"`
}

type listOEOptions struct {
	ConfigurationID int `url:"configuration_id,omitempty"`
}

// List lists all the Osquery enrollments.
func (s *OsqueryEnrollmentsServiceOp) List(ctx context.Context, opt *ListOptions) ([]OsqueryEnrollment, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Osquery enrollment by id.
func (s *OsqueryEnrollmentsServiceOp) GetByID(ctx context.Context, oeID int) (*OsqueryEnrollment, *Response, error) {
	if oeID < 1 {
		return nil, nil, NewArgError("oeID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", oeBasePath, oeID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	oe := new(OsqueryEnrollment)

	resp, err := s.client.Do(ctx, req, oe)
	if err != nil {
		return nil, resp, err
	}

	return oe, resp, err
}

// GetByConfigurationID retrieves the Osquery enrollments for a given configuration.
func (s *OsqueryEnrollmentsServiceOp) GetByConfigurationID(ctx context.Context, configuration_id int) ([]OsqueryEnrollment, *Response, error) {
	if configuration_id < 1 {
		return nil, nil, NewArgError("configuration_id", "cannot be negative")
	}

	listSEOpt := &listOEOptions{ConfigurationID: configuration_id}

	return s.list(ctx, nil, listSEOpt)
}

// Create a new Osquery enrollment.
func (s *OsqueryEnrollmentsServiceOp) Create(ctx context.Context, createRequest *OsqueryEnrollmentRequest) (*OsqueryEnrollment, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, oeBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	oe := new(OsqueryEnrollment)
	resp, err := s.client.Do(ctx, req, oe)
	if err != nil {
		return nil, resp, err
	}

	return oe, resp, err
}

// Update a Osquery enrollment.
func (s *OsqueryEnrollmentsServiceOp) Update(ctx context.Context, oeID int, updateRequest *OsqueryEnrollmentRequest) (*OsqueryEnrollment, *Response, error) {
	if oeID < 1 {
		return nil, nil, NewArgError("oeID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", oeBasePath, oeID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	oe := new(OsqueryEnrollment)
	resp, err := s.client.Do(ctx, req, oe)
	if err != nil {
		return nil, resp, err
	}

	return oe, resp, err
}

// Delete a Osquery enrollment.
func (s *OsqueryEnrollmentsServiceOp) Delete(ctx context.Context, oeID int) (*Response, error) {
	if oeID < 1 {
		return nil, NewArgError("oeID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", oeBasePath, oeID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Osquery enrollments
func (s *OsqueryEnrollmentsServiceOp) list(ctx context.Context, opt *ListOptions, oeOpt *listOEOptions) ([]OsqueryEnrollment, *Response, error) {
	path := oeBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, oeOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var oes []OsqueryEnrollment
	resp, err := s.client.Do(ctx, req, &oes)
	if err != nil {
		return nil, resp, err
	}

	return oes, resp, err
}
