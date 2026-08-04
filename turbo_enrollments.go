package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const tenrBasePath = "turbo/enrollments/"

// TurboEnrollmentsService is an interface for interfacing with the Turbo enrollments
// endpoints of the Zentral API
type TurboEnrollmentsService interface {
	List(context.Context, *ListOptions) ([]TurboEnrollment, *Response, error)
	GetByID(context.Context, int) (*TurboEnrollment, *Response, error)
	GetByConfigurationID(context.Context, string) ([]TurboEnrollment, *Response, error)
	Create(context.Context, *TurboEnrollmentRequest) (*TurboEnrollment, *Response, error)
	Update(context.Context, int, *TurboEnrollmentRequest) (*TurboEnrollment, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// TurboEnrollmentsServiceOp handles communication with the Turbo enrollments related
// methods of the Zentral API.
type TurboEnrollmentsServiceOp struct {
	client *Client
}

var _ TurboEnrollmentsService = &TurboEnrollmentsServiceOp{}

// TurboEnrollment represents a Zentral TurboEnrollment
type TurboEnrollment struct {
	ID                    int              `json:"id"`
	ConfigurationID       string           `json:"configuration"`
	EnrolledMachinesCount int              `json:"enrolled_machines_count"`
	Secret                EnrollmentSecret `json:"secret"`
	ConfigProfileURL      string           `json:"configuration_profile_download_url"`
	PlistURL              string           `json:"plist_download_url"`
	Version               int              `json:"version"`
	Created               Timestamp        `json:"created_at,omitempty"`
	Updated               Timestamp        `json:"updated_at,omitempty"`
}

func (te TurboEnrollment) String() string {
	return Stringify(te)
}

// TurboEnrollmentRequest represents a request to create or update a Turbo enrollment
type TurboEnrollmentRequest struct {
	ConfigurationID string                  `json:"configuration"`
	Secret          EnrollmentSecretRequest `json:"secret"`
}

type listTEnrOptions struct {
	ConfigurationID string `url:"configuration,omitempty"`
}

// List lists all the Turbo enrollments.
func (s *TurboEnrollmentsServiceOp) List(ctx context.Context, opt *ListOptions) ([]TurboEnrollment, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Turbo enrollment by id.
func (s *TurboEnrollmentsServiceOp) GetByID(ctx context.Context, teID int) (*TurboEnrollment, *Response, error) {
	if teID < 1 {
		return nil, nil, NewArgError("teID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", tenrBasePath, teID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	te := new(TurboEnrollment)

	resp, err := s.client.Do(ctx, req, te)
	if err != nil {
		return nil, resp, err
	}

	return te, resp, err
}

// GetByConfigurationID retrieves the Turbo enrollments for a given configuration.
func (s *TurboEnrollmentsServiceOp) GetByConfigurationID(ctx context.Context, configurationID string) ([]TurboEnrollment, *Response, error) {
	if len(configurationID) < 1 {
		return nil, nil, NewArgError("configurationID", "cannot be blank")
	}

	listTEnrOpt := &listTEnrOptions{ConfigurationID: configurationID}

	return s.list(ctx, nil, listTEnrOpt)
}

// Create a new Turbo enrollment.
func (s *TurboEnrollmentsServiceOp) Create(ctx context.Context, createRequest *TurboEnrollmentRequest) (*TurboEnrollment, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, tenrBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	te := new(TurboEnrollment)
	resp, err := s.client.Do(ctx, req, te)
	if err != nil {
		return nil, resp, err
	}

	return te, resp, err
}

// Update a Turbo enrollment.
func (s *TurboEnrollmentsServiceOp) Update(ctx context.Context, teID int, updateRequest *TurboEnrollmentRequest) (*TurboEnrollment, *Response, error) {
	if teID < 1 {
		return nil, nil, NewArgError("teID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", tenrBasePath, teID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	te := new(TurboEnrollment)
	resp, err := s.client.Do(ctx, req, te)
	if err != nil {
		return nil, resp, err
	}

	return te, resp, err
}

// Delete a Turbo enrollment.
func (s *TurboEnrollmentsServiceOp) Delete(ctx context.Context, teID int) (*Response, error) {
	if teID < 1 {
		return nil, NewArgError("teID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", tenrBasePath, teID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Turbo enrollments
func (s *TurboEnrollmentsServiceOp) list(ctx context.Context, opt *ListOptions, teOpt *listTEnrOptions) ([]TurboEnrollment, *Response, error) {
	path := tenrBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, teOpt)
	if err != nil {
		return nil, nil, err
	}

	return resolveAllPages[TurboEnrollment](ctx, s.client, path)
}
