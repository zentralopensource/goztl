package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const meBasePath = "monolith/enrollments/"

// MonolithEnrollmentsService is an interface for interfacing with the Monolith enrollments
// endpoints of the Zentral API
type MonolithEnrollmentsService interface {
	List(context.Context, *ListOptions) ([]MonolithEnrollment, *Response, error)
	GetByID(context.Context, int) (*MonolithEnrollment, *Response, error)
	GetByManifestID(context.Context, int) ([]MonolithEnrollment, *Response, error)
	Create(context.Context, *MonolithEnrollmentRequest) (*MonolithEnrollment, *Response, error)
	Update(context.Context, int, *MonolithEnrollmentRequest) (*MonolithEnrollment, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// MonolithEnrollmentsServiceOp handles communication with the Monolith enrollments related
// methods of the Zentral API.
type MonolithEnrollmentsServiceOp struct {
	client *Client
}

var _ MonolithEnrollmentsService = &MonolithEnrollmentsServiceOp{}

// MonolithEnrollment represents a Zentral MonolithEnrollment
type MonolithEnrollment struct {
	ID                    int              `json:"id"`
	ManifestID            int              `json:"manifest"`
	EnrolledMachinesCount int              `json:"enrolled_machines_count"`
	Secret                EnrollmentSecret `json:"secret"`
	Version               int              `json:"version"`
	ConfigProfileURL      string           `json:"configuration_profile_download_url"`
	PlistURL              string           `json:"plist_download_url"`
	Created               Timestamp        `json:"created_at,omitempty"`
	Updated               Timestamp        `json:"updated_at,omitempty"`
}

func (se MonolithEnrollment) String() string {
	return Stringify(se)
}

// MonolithEnrollmentRequest represents a request to create or update a Monolith enrollment
type MonolithEnrollmentRequest struct {
	ManifestID int                     `json:"manifest"`
	Secret     EnrollmentSecretRequest `json:"secret"`
}

type listMEOptions struct {
	ManifestID int `url:"manifest_id,omitempty"`
}

// List lists all the Monolith enrollments.
func (s *MonolithEnrollmentsServiceOp) List(ctx context.Context, opt *ListOptions) ([]MonolithEnrollment, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Monolith enrollment by id.
func (s *MonolithEnrollmentsServiceOp) GetByID(ctx context.Context, meID int) (*MonolithEnrollment, *Response, error) {
	if meID < 1 {
		return nil, nil, NewArgError("meID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", meBasePath, meID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	me := new(MonolithEnrollment)

	resp, err := s.client.Do(ctx, req, me)
	if err != nil {
		return nil, resp, err
	}

	return me, resp, err
}

// GetByManifestID retrieves the Monolith enrollments for a given manifest.
func (s *MonolithEnrollmentsServiceOp) GetByManifestID(ctx context.Context, mID int) ([]MonolithEnrollment, *Response, error) {
	if mID < 1 {
		return nil, nil, NewArgError("mID", "cannot be negative")
	}

	listMEOpt := &listMEOptions{ManifestID: mID}

	return s.list(ctx, nil, listMEOpt)
}

// Create a new Monolith enrollment.
func (s *MonolithEnrollmentsServiceOp) Create(ctx context.Context, createRequest *MonolithEnrollmentRequest) (*MonolithEnrollment, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, meBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	me := new(MonolithEnrollment)
	resp, err := s.client.Do(ctx, req, me)
	if err != nil {
		return nil, resp, err
	}

	return me, resp, err
}

// Update a Monolith enrollment.
func (s *MonolithEnrollmentsServiceOp) Update(ctx context.Context, meID int, updateRequest *MonolithEnrollmentRequest) (*MonolithEnrollment, *Response, error) {
	if meID < 1 {
		return nil, nil, NewArgError("meID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", meBasePath, meID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	me := new(MonolithEnrollment)
	resp, err := s.client.Do(ctx, req, me)
	if err != nil {
		return nil, resp, err
	}

	return me, resp, err
}

// Delete a Monolith enrollment.
func (s *MonolithEnrollmentsServiceOp) Delete(ctx context.Context, meID int) (*Response, error) {
	if meID < 1 {
		return nil, NewArgError("meID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", meBasePath, meID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Monolith enrollments
func (s *MonolithEnrollmentsServiceOp) list(ctx context.Context, opt *ListOptions, meOpt *listMEOptions) ([]MonolithEnrollment, *Response, error) {
	path := meBasePath
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

	var mes []MonolithEnrollment
	resp, err := s.client.Do(ctx, req, &mes)
	if err != nil {
		return nil, resp, err
	}

	return mes, resp, err
}
