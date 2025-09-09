package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const moeBasePath = "mdm/ota_enrollments/"

// MDMOTAEnrollmentsService is an interface for interfacing with the MDM enrollments
// endpoints of the Zentral API
type MDMOTAEnrollmentsService interface {
	List(context.Context, *ListOptions) ([]MDMOTAEnrollment, *Response, error)
	GetByID(context.Context, int) (*MDMOTAEnrollment, *Response, error)
	GetByName(context.Context, string) (*MDMOTAEnrollment, *Response, error)
	Create(context.Context, *MDMOTAEnrollmentRequest) (*MDMOTAEnrollment, *Response, error)
	Update(context.Context, int, *MDMOTAEnrollmentRequest) (*MDMOTAEnrollment, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// MDMOTAEnrollmentsServiceOp handles communication with the MDM enrollments related
// moethods of the Zentral API.
type MDMOTAEnrollmentsServiceOp struct {
	client *Client
}

var _ MDMOTAEnrollmentsService = &MDMOTAEnrollmentsServiceOp{}

// MDMOTAEnrollment represents a Zentral MDM OTAEnrollment
type MDMOTAEnrollment struct {
	ID                int              `json:"id"`
	Name              string           `json:"name"`
	DisplayName       string           `json:"display_name"`
	BlueprintID       *int             `json:"blueprint"`
	PushCertificateID int              `json:"push_certificate"`
	RealmUUID         *string          `json:"realm"`
	ACMEIssuerUUID    *string          `json:"acme_issuer"`
	SCEPIssuerUUID    string           `json:"scep_issuer"`
	Secret            EnrollmentSecret `json:"enrollment_secret"`
	Created           Timestamp        `json:"created_at,omitempty"`
	Updated           Timestamp        `json:"updated_at,omitempty"`
}

func (oe MDMOTAEnrollment) String() string {
	return Stringify(oe)
}

// MDMOTAEnrollmentRequest represents a request to create or update a MDM OTA enrollment
type MDMOTAEnrollmentRequest struct {
	Name              string                  `json:"name"`
	DisplayName       *string                 `json:"display_name"`
	BlueprintID       *int                    `json:"blueprint"`
	PushCertificateID int                     `json:"push_certificate"`
	RealmUUID         *string                 `json:"realm"`
	ACMEIssuerUUID    *string                 `json:"acme_issuer"`
	SCEPIssuerUUID    string                  `json:"scep_issuer"`
	Secret            EnrollmentSecretRequest `json:"enrollment_secret"`
}

type listMOEOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the MDM OTA enrollments.
func (s *MDMOTAEnrollmentsServiceOp) List(ctx context.Context, opt *ListOptions) ([]MDMOTAEnrollment, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a MDM OTA enrollment by id.
func (s *MDMOTAEnrollmentsServiceOp) GetByID(ctx context.Context, moeID int) (*MDMOTAEnrollment, *Response, error) {
	if moeID < 1 {
		return nil, nil, NewArgError("moeID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", moeBasePath, moeID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	moe := new(MDMOTAEnrollment)

	resp, err := s.client.Do(ctx, req, moe)
	if err != nil {
		return nil, resp, err
	}

	return moe, resp, err
}

// GetByName retrieves a MDM OTA enrollment by name.
func (s *MDMOTAEnrollmentsServiceOp) GetByName(ctx context.Context, name string) (*MDMOTAEnrollment, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listMOEOpt := &listMOEOptions{Name: name}

	moes, resp, err := s.list(ctx, nil, listMOEOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(moes) < 1 {
		return nil, resp, nil
	}

	return &moes[0], resp, err
}

// Create a new MDM OTA enrollment.
func (s *MDMOTAEnrollmentsServiceOp) Create(ctx context.Context, createRequest *MDMOTAEnrollmentRequest) (*MDMOTAEnrollment, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, moeBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	moe := new(MDMOTAEnrollment)
	resp, err := s.client.Do(ctx, req, moe)
	if err != nil {
		return nil, resp, err
	}

	return moe, resp, err
}

// Update a MDM OTA enrollment.
func (s *MDMOTAEnrollmentsServiceOp) Update(ctx context.Context, moeID int, updateRequest *MDMOTAEnrollmentRequest) (*MDMOTAEnrollment, *Response, error) {
	if moeID < 1 {
		return nil, nil, NewArgError("moeID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", moeBasePath, moeID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	moe := new(MDMOTAEnrollment)
	resp, err := s.client.Do(ctx, req, moe)
	if err != nil {
		return nil, resp, err
	}

	return moe, resp, err
}

// Delete a MDM OTA enrollment.
func (s *MDMOTAEnrollmentsServiceOp) Delete(ctx context.Context, moeID int) (*Response, error) {
	if moeID < 1 {
		return nil, NewArgError("moeID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", moeBasePath, moeID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper moethod for listing MDM OTA enrollments
func (s *MDMOTAEnrollmentsServiceOp) list(ctx context.Context, opt *ListOptions, moeOpt *listMOEOptions) ([]MDMOTAEnrollment, *Response, error) {
	path := moeBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, moeOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var moes []MDMOTAEnrollment
	resp, err := s.client.Do(ctx, req, &moes)
	if err != nil {
		return nil, resp, err
	}

	return moes, resp, err
}
