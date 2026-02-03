package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const mppBasePath = "mdm/provisioning_profiles/"

// MDMProvisioningProfilesService is an interface for interfacing with the MDM provisioning profile
// endpoints of the Zentral API
type MDMProvisioningProfilesService interface {
	List(context.Context, *ListOptions) ([]MDMProvisioningProfile, *Response, error)
	GetByID(context.Context, string) (*MDMProvisioningProfile, *Response, error)
	Create(context.Context, *MDMProvisioningProfileRequest) (*MDMProvisioningProfile, *Response, error)
	Update(context.Context, string, *MDMProvisioningProfileRequest) (*MDMProvisioningProfile, *Response, error)
	Delete(context.Context, string) (*Response, error)
}

// MDMProvisioningProfilesServiceOp handles communication with the MDM provisioning profiles related
// methods of the Zentral API.
type MDMProvisioningProfilesServiceOp struct {
	client *Client
}

var _ MDMProvisioningProfilesService = &MDMProvisioningProfilesServiceOp{}

// MDMProvisioningProfile represents a Zentral MDM provisioning profile
type MDMProvisioningProfile struct {
	ID     string `json:"id"`
	Source string `json:"source"`
	MDMArtifactVersion
}

func (mpp MDMProvisioningProfile) String() string {
	return Stringify(mpp)
}

// MDMProvisioningProfileRequest represents a request to create or update a MDM provisioning profile
type MDMProvisioningProfileRequest struct {
	Source string `json:"source"`
	MDMArtifactVersionRequest
}

type listMPPOptions struct{}

// List lists all the MDM provisioning profiles.
func (s *MDMProvisioningProfilesServiceOp) List(ctx context.Context, opt *ListOptions) ([]MDMProvisioningProfile, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a MDM provisioning profile by id.
func (s *MDMProvisioningProfilesServiceOp) GetByID(ctx context.Context, mppID string) (*MDMProvisioningProfile, *Response, error) {
	if len(mppID) < 1 {
		return nil, nil, NewArgError("mppID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", mppBasePath, mppID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	mpp := new(MDMProvisioningProfile)

	resp, err := s.client.Do(ctx, req, mpp)
	if err != nil {
		return nil, resp, err
	}

	return mpp, resp, err
}

// Create a new MDM provisioning profile.
func (s *MDMProvisioningProfilesServiceOp) Create(ctx context.Context, createRequest *MDMProvisioningProfileRequest) (*MDMProvisioningProfile, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, mppBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	mpp := new(MDMProvisioningProfile)
	resp, err := s.client.Do(ctx, req, mpp)
	if err != nil {
		return nil, resp, err
	}

	return mpp, resp, err
}

// Update a MDM provisioning profile.
func (s *MDMProvisioningProfilesServiceOp) Update(ctx context.Context, mppID string, updateRequest *MDMProvisioningProfileRequest) (*MDMProvisioningProfile, *Response, error) {
	if len(mppID) < 1 {
		return nil, nil, NewArgError("mppID", "cannot be blank")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%s/", mppBasePath, mppID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	mpp := new(MDMProvisioningProfile)
	resp, err := s.client.Do(ctx, req, mpp)
	if err != nil {
		return nil, resp, err
	}

	return mpp, resp, err
}

// Delete a MDM provisioning profile.
func (s *MDMProvisioningProfilesServiceOp) Delete(ctx context.Context, mppID string) (*Response, error) {
	if len(mppID) < 1 {
		return nil, NewArgError("mppID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", mppBasePath, mppID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing MDM provisioning profiles
func (s *MDMProvisioningProfilesServiceOp) list(ctx context.Context, opt *ListOptions, mppOpt *listMPPOptions) ([]MDMProvisioningProfile, *Response, error) {
	path := mppBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, mppOpt)
	if err != nil {
		return nil, nil, err
	}

	return resolveAllPages[MDMProvisioningProfile](ctx, s.client, path)
}
