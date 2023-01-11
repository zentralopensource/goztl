package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const seBasePath = "santa/enrollments/"

// SantaEnrollmentsService is an interface for interfacing with the Santa enrollments
// endpoints of the Zentral API
type SantaEnrollmentsService interface {
	List(context.Context, *ListOptions) ([]SantaEnrollment, *Response, error)
	GetByID(context.Context, int) (*SantaEnrollment, *Response, error)
	GetByConfigurationID(context.Context, int) (*SantaEnrollment, *Response, error)
	Create(context.Context, *SantaEnrollmentRequest) (*SantaEnrollment, *Response, error)
	Update(context.Context, int, *SantaEnrollmentRequest) (*SantaEnrollment, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// SantaEnrollmentsServiceOp handles communication with the Santa enrollments related
// methods of the Zentral API.
type SantaEnrollmentsServiceOp struct {
	client *Client
}

var _ SantaEnrollmentsService = &SantaEnrollmentsServiceOp{}

// SantaEnrollment represents a Zentral SantaEnrollment
type SantaEnrollment struct {
	ID                    int              `json:"id"`
	ConfigurationID       int              `json:"configuration"`
	EnrolledMachinesCount int              `json:"enrolled_machines_count"`
	Secret                EnrollmentSecret `json:"secret"`
	ConfigProfileURL      string           `json:"configuration_profile_download_url"`
	PlistURL              string           `json:"plist_download_url"`
	Version               int              `json:"version"`
	Created               Timestamp        `json:"created_at,omitempty"`
	Updated               Timestamp        `json:"updated_at,omitempty"`
}

func (se SantaEnrollment) String() string {
	return Stringify(se)
}

// SantaEnrollmentRequest represents a request to create or update a Santa enrollment
type SantaEnrollmentRequest struct {
	ConfigurationID int                     `json:"configuration"`
	Secret          EnrollmentSecretRequest `json:"secret"`
}

type listSEOptions struct {
	ConfigurationID int `url:"configuration_id,omitempty"`
}

// List lists all the Santa enrollments.
func (s *SantaEnrollmentsServiceOp) List(ctx context.Context, opt *ListOptions) ([]SantaEnrollment, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Santa enrollment by id.
func (s *SantaEnrollmentsServiceOp) GetByID(ctx context.Context, seID int) (*SantaEnrollment, *Response, error) {
	if seID < 1 {
		return nil, nil, NewArgError("seID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", seBasePath, seID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	se := new(SantaEnrollment)

	resp, err := s.client.Do(ctx, req, se)
	if err != nil {
		return nil, resp, err
	}

	return se, resp, err
}

// GetByConfigurationID retrieves a Santa enrollment by name.
func (s *SantaEnrollmentsServiceOp) GetByConfigurationID(ctx context.Context, configuration_id int) (*SantaEnrollment, *Response, error) {
	if configuration_id < 1 {
		return nil, nil, NewArgError("name", "cannot be negative")
	}

	listSEOpt := &listSEOptions{ConfigurationID: configuration_id}

	ses, resp, err := s.list(ctx, nil, listSEOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(ses) < 1 {
		return nil, resp, nil
	}

	return &ses[0], resp, err
}

// Create a new Santa enrollment.
func (s *SantaEnrollmentsServiceOp) Create(ctx context.Context, createRequest *SantaEnrollmentRequest) (*SantaEnrollment, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, seBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	se := new(SantaEnrollment)
	resp, err := s.client.Do(ctx, req, se)
	if err != nil {
		return nil, resp, err
	}

	return se, resp, err
}

// Update a Santa enrollment.
func (s *SantaEnrollmentsServiceOp) Update(ctx context.Context, seID int, updateRequest *SantaEnrollmentRequest) (*SantaEnrollment, *Response, error) {
	if seID < 1 {
		return nil, nil, NewArgError("seID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", seBasePath, seID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	se := new(SantaEnrollment)
	resp, err := s.client.Do(ctx, req, se)
	if err != nil {
		return nil, resp, err
	}

	return se, resp, err
}

// Delete a Santa enrollment.
func (s *SantaEnrollmentsServiceOp) Delete(ctx context.Context, seID int) (*Response, error) {
	if seID < 1 {
		return nil, NewArgError("seID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", seBasePath, seID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Santa enrollments
func (s *SantaEnrollmentsServiceOp) list(ctx context.Context, opt *ListOptions, seOpt *listSEOptions) ([]SantaEnrollment, *Response, error) {
	path := seBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, seOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var ses []SantaEnrollment
	resp, err := s.client.Do(ctx, req, &ses)
	if err != nil {
		return nil, resp, err
	}

	return ses, resp, err
}
