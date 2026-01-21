package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const depEnrollmentBasePath = "mdm/dep_enrollments/"

// MDMDEPEnrollmentsService is an interface for interfacing with the MDM dep enrollments
// endpoints of the Zentral API
type MDMDEPEnrollmentsService interface {
	List(context.Context, *ListOptions) ([]MDMDEPEnrollment, *Response, error)
	GetByID(context.Context, int) (*MDMDEPEnrollment, *Response, error)
	GetByName(context.Context, string) (*MDMDEPEnrollment, *Response, error)
	Create(context.Context, *MDMDEPEnrollmentCreationRequest) (*MDMDEPEnrollment, *Response, error)
	Update(context.Context, int, *MDMDEPEnrollmentUpdateRequest) (*MDMDEPEnrollment, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// MDMDEPEnrollmentsServiceOp handles communication with the MDM enrollments related
// moethods of the Zentral API.
type MDMDEPEnrollmentsServiceOp struct {
	client *Client
}

var _ MDMDEPEnrollmentsService = &MDMDEPEnrollmentsServiceOp{}

// MDMDEPEnrollment represents a Zentral MDM DEPEnrollment
type MDMDEPEnrollment struct {
	ID                         int              `json:"id"`
	Name                       string           `json:"name"`
	DisplayName                string           `json:"display_name"`
	Secret                     EnrollmentSecret `json:"enrollment_secret"`
	UseRealmUser               bool             `json:"use_realm_user"`
	UsernamePattern            string           `json:"username_pattern"`
	RealmUserIsAdmin           bool             `json:"realm_user_is_admin"`
	AdminFullName              *string          `json:"admin_full_name"`
	AdminShortName             *string          `json:"admin_short_name"`
	HiddenAdmin                bool             `json:"hidden_admin"`
	AdminPasswordComplexity    int              `json:"admin_password_complexity"`
	AdminPasswordRotationDelay int              `json:"admin_password_rotation_delay"`
	AllowPairing               bool             `json:"allow_pairing"`
	AutoAdvanceSetup           bool             `json:"auto_advance_setup"`
	AwaitDeviceConfigured      bool             `json:"await_device_configured"`
	Department                 string           `json:"department"`
	IsMandatory                bool             `json:"is_mandatory"`
	IsMDMRemovable             bool             `json:"is_mdm_removable"`
	IsMultiUser                bool             `json:"is_multi_user"`
	IsSupervised               bool             `json:"is_supervised"`
	Language                   string           `json:"language"`
	OrgMagic                   string           `json:"org_magic"`
	Region                     string           `json:"region"`
	SkipSetupItems             []string         `json:"skip_setup_items"`
	SupportEmailAddress        string           `json:"support_email_address"`
	SupportPhoneNumber         string           `json:"support_phone_number"`
	IncludeTLSCertificates     bool             `json:"include_tls_certificates"`
	IOSMaxVersion              string           `json:"ios_max_version"`
	IOSMinVersion              string           `json:"ios_min_version"`
	MacOSMaxVersion            string           `json:"macos_max_version"`
	MacOSMinVersion            string           `json:"macos_min_version"`
	PushCertificateID          int              `json:"push_certificate"`
	ACMEIssuerUUID             *string          `json:"acme_issuer"`
	SCEPIssuerUUID             string           `json:"scep_issuer"`
	BlueprintID                *int             `json:"blueprint"`
	RealmUUID                  *string          `json:"realm"`
	VirtualServerID            int              `json:"virtual_server"`
	Created                    Timestamp        `json:"created_at,omitempty"`
	Updated                    Timestamp        `json:"updated_at,omitempty"`
}

func (enrollment MDMDEPEnrollment) String() string {
	return Stringify(enrollment)
}

// MDMDEPEnrollmentCreationRequest represents a request to create MDM DEPEnrollment
type MDMDEPEnrollmentCreationRequest struct {
	Name                       string                  `json:"name"`
	DisplayName                string                  `json:"display_name"`
	Secret                     EnrollmentSecretRequest `json:"enrollment_secret"`
	UseRealmUser               bool                    `json:"use_realm_user"`
	UsernamePattern            string                  `json:"username_pattern"`
	RealmUserIsAdmin           bool                    `json:"realm_user_is_admin"`
	AdminFullName              *string                 `json:"admin_full_name"`
	AdminShortName             *string                 `json:"admin_short_name"`
	HiddenAdmin                bool                    `json:"hidden_admin"`
	AdminPasswordComplexity    int                     `json:"admin_password_complexity"`
	AdminPasswordRotationDelay int                     `json:"admin_password_rotation_delay"`
	AllowPairing               bool                    `json:"allow_pairing"`
	AutoAdvanceSetup           bool                    `json:"auto_advance_setup"`
	AwaitDeviceConfigured      bool                    `json:"await_device_configured"`
	Department                 string                  `json:"department"`
	IsMandatory                bool                    `json:"is_mandatory"`
	IsMDMRemovable             bool                    `json:"is_mdm_removable"`
	IsMultiUser                bool                    `json:"is_multi_user"`
	IsSupervised               bool                    `json:"is_supervised"`
	Language                   string                  `json:"language"`
	OrgMagic                   string                  `json:"org_magic"`
	Region                     string                  `json:"region"`
	SkipSetupItems             []string                `json:"skip_setup_items"`
	SupportEmailAddress        string                  `json:"support_email_address"`
	SupportPhoneNumber         string                  `json:"support_phone_number"`
	IncludeTLSCertificates     bool                    `json:"include_tls_certificates"`
	IOSMaxVersion              string                  `json:"ios_max_version"`
	IOSMinVersion              string                  `json:"ios_min_version"`
	MacOSMaxVersion            string                  `json:"macos_max_version"`
	MacOSMinVersion            string                  `json:"macos_min_version"`
	PushCertificateID          int                     `json:"push_certificate"`
	ACMEIssuerUUID             *string                 `json:"acme_issuer"`
	SCEPIssuerUUID             string                  `json:"scep_issuer"`
	BlueprintID                *int                    `json:"blueprint"`
	RealmUUID                  *string                 `json:"realm"`
	VirtualServerID            int                     `json:"virtual_server"`
}

// MDMDEPEnrollmentUpdateRequest represents a request to update a MDM DEPEnrollment
type MDMDEPEnrollmentUpdateRequest struct {
	Name                       string                  `json:"name"`
	DisplayName                string                  `json:"display_name"`
	Secret                     EnrollmentSecretRequest `json:"enrollment_secret"`
	UseRealmUser               bool                    `json:"use_realm_user"`
	UsernamePattern            string                  `json:"username_pattern"`
	RealmUserIsAdmin           bool                    `json:"realm_user_is_admin"`
	AdminFullName              *string                 `json:"admin_full_name"`
	AdminShortName             *string                 `json:"admin_short_name"`
	HiddenAdmin                bool                    `json:"hidden_admin"`
	AdminPasswordComplexity    int                     `json:"admin_password_complexity"`
	AdminPasswordRotationDelay int                     `json:"admin_password_rotation_delay"`
	AllowPairing               bool                    `json:"allow_pairing"`
	AutoAdvanceSetup           bool                    `json:"auto_advance_setup"`
	AwaitDeviceConfigured      bool                    `json:"await_device_configured"`
	Department                 string                  `json:"department"`
	IsMandatory                bool                    `json:"is_mandatory"`
	IsMDMRemovable             bool                    `json:"is_mdm_removable"`
	IsMultiUser                bool                    `json:"is_multi_user"`
	IsSupervised               bool                    `json:"is_supervised"`
	Language                   string                  `json:"language"`
	OrgMagic                   string                  `json:"org_magic"`
	Region                     string                  `json:"region"`
	SkipSetupItems             []string                `json:"skip_setup_items"`
	SupportEmailAddress        string                  `json:"support_email_address"`
	SupportPhoneNumber         string                  `json:"support_phone_number"`
	IncludeTLSCertificates     bool                    `json:"include_tls_certificates"`
	IOSMaxVersion              string                  `json:"ios_max_version"`
	IOSMinVersion              string                  `json:"ios_min_version"`
	MacOSMaxVersion            string                  `json:"macos_max_version"`
	MacOSMinVersion            string                  `json:"macos_min_version"`
	PushCertificateID          int                     `json:"push_certificate"`
	ACMEIssuerUUID             *string                 `json:"acme_issuer"`
	SCEPIssuerUUID             string                  `json:"scep_issuer"`
	BlueprintID                *int                    `json:"blueprint"`
	RealmUUID                  *string                 `json:"realm"`
}

type listMDMDEPEnrollmentOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the MDM DEP enrollments.
func (service *MDMDEPEnrollmentsServiceOp) List(ctx context.Context, opt *ListOptions) ([]MDMDEPEnrollment, *Response, error) {
	return service.list(ctx, opt, nil)
}

// GetByID retrieves a MDM DEP enrollment by id.
func (service *MDMDEPEnrollmentsServiceOp) GetByID(ctx context.Context, enrollmentID int) (*MDMDEPEnrollment, *Response, error) {
	if enrollmentID < 1 {
		return nil, nil, NewArgError("enrollmentID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", depEnrollmentBasePath, enrollmentID)

	req, err := service.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	enrollment := new(MDMDEPEnrollment)

	resp, err := service.client.Do(ctx, req, enrollment)
	if err != nil {
		return nil, resp, err
	}

	return enrollment, resp, err
}

// GetByName retrieves a MDM DEP enrollment by name.
func (service *MDMDEPEnrollmentsServiceOp) GetByName(ctx context.Context, name string) (*MDMDEPEnrollment, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listOpt := &listMDMDEPEnrollmentOptions{Name: name}

	enrollments, resp, err := service.list(ctx, nil, listOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(enrollments) < 1 {
		return nil, resp, nil
	}

	return &enrollments[0], resp, err
}

// Create a new MDM DEP enrollment.
func (s *MDMDEPEnrollmentsServiceOp) Create(ctx context.Context, createRequest *MDMDEPEnrollmentCreationRequest) (*MDMDEPEnrollment, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, depEnrollmentBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	enrollment := new(MDMDEPEnrollment)
	resp, err := s.client.Do(ctx, req, enrollment)
	if err != nil {
		return nil, resp, err
	}

	return enrollment, resp, err
}

// Update a MDM DEP enrollment.
func (s *MDMDEPEnrollmentsServiceOp) Update(ctx context.Context, enrollmentID int, updateRequest *MDMDEPEnrollmentUpdateRequest) (*MDMDEPEnrollment, *Response, error) {
	if enrollmentID < 1 {
		return nil, nil, NewArgError("enrollmentID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", depEnrollmentBasePath, enrollmentID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	enrollment := new(MDMDEPEnrollment)
	resp, err := s.client.Do(ctx, req, enrollment)
	if err != nil {
		return nil, resp, err
	}

	return enrollment, resp, err
}

// Delete a MDM DEP enrollment.
func (s *MDMDEPEnrollmentsServiceOp) Delete(ctx context.Context, enrollmentID int) (*Response, error) {
	if enrollmentID < 1 {
		return nil, NewArgError("enrollmentID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", depEnrollmentBasePath, enrollmentID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper moethod for listing MDM DEP enrollments
func (service *MDMDEPEnrollmentsServiceOp) list(ctx context.Context, opt *ListOptions, listOpt *listMDMDEPEnrollmentOptions) ([]MDMDEPEnrollment, *Response, error) {
	path := depEnrollmentBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, listOpt)
	if err != nil {
		return nil, nil, err
	}
	return resolveAllPages[MDMDEPEnrollment](ctx, service.client, path)
}
