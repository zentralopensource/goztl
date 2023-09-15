package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const meaBasePath = "mdm/enterprise_apps/"

// MDMEnterpriseAppsService is an interface for interfacing with the MDM enterprise app
// endpoints of the Zentral API
type MDMEnterpriseAppsService interface {
	List(context.Context, *ListOptions) ([]MDMEnterpriseApp, *Response, error)
	GetByID(context.Context, string) (*MDMEnterpriseApp, *Response, error)
	Create(context.Context, *MDMEnterpriseAppRequest) (*MDMEnterpriseApp, *Response, error)
	Update(context.Context, string, *MDMEnterpriseAppRequest) (*MDMEnterpriseApp, *Response, error)
	Delete(context.Context, string) (*Response, error)
}

// MDMEnterpriseAppsServiceOp handles communication with the MDM enterprise apps related
// methods of the Zentral API.
type MDMEnterpriseAppsServiceOp struct {
	client *Client
}

var _ MDMEnterpriseAppsService = &MDMEnterpriseAppsServiceOp{}

// MDMEnterpriseApp represents a Zentral MDM enterprise app
type MDMEnterpriseApp struct {
	ID               string  `json:"id"`
	Filename         string  `json:"filename"`
	ProductID        string  `json:"product_id"`
	ProductVersion   string  `json:"product_version"`
	IOSApp           bool    `json:"ios_app"`
	Configuration    *string `json:"configuration"`
	InstallAsManaged bool    `json:"install_as_managed"`
	RemoveOnUnenroll bool    `json:"remove_on_unenroll"`
	MDMArtifactVersion
}

func (mea MDMEnterpriseApp) String() string {
	return Stringify(mea)
}

// MDMEnterpriseAppRequest represents a request to create or update a MDM enterprise app
type MDMEnterpriseAppRequest struct {
	SourceURI        string  `json:"source_uri"`
	SourceSHA256     string  `json:"source_sha256"`
	IOSApp           bool    `json:"ios_app"`
	Configuration    *string `json:"configuration"`
	InstallAsManaged bool    `json:"installed_as_managed"`
	RemoveOnUnenroll bool    `json:"remove_on_unenroll"`
	MDMArtifactVersionRequest
}

type listMEAOptions struct{}

// List lists all the MDM enterprise apps.
func (s *MDMEnterpriseAppsServiceOp) List(ctx context.Context, opt *ListOptions) ([]MDMEnterpriseApp, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a MDM enterprise app by id.
func (s *MDMEnterpriseAppsServiceOp) GetByID(ctx context.Context, meaID string) (*MDMEnterpriseApp, *Response, error) {
	if len(meaID) < 1 {
		return nil, nil, NewArgError("meaID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", meaBasePath, meaID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	mea := new(MDMEnterpriseApp)

	resp, err := s.client.Do(ctx, req, mea)
	if err != nil {
		return nil, resp, err
	}

	return mea, resp, err
}

// Create a new MDM enterprise app.
func (s *MDMEnterpriseAppsServiceOp) Create(ctx context.Context, createRequest *MDMEnterpriseAppRequest) (*MDMEnterpriseApp, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, meaBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	mea := new(MDMEnterpriseApp)
	resp, err := s.client.Do(ctx, req, mea)
	if err != nil {
		return nil, resp, err
	}

	return mea, resp, err
}

// Update a MDM enterprise app.
func (s *MDMEnterpriseAppsServiceOp) Update(ctx context.Context, meaID string, updateRequest *MDMEnterpriseAppRequest) (*MDMEnterpriseApp, *Response, error) {
	if len(meaID) < 1 {
		return nil, nil, NewArgError("meaID", "cannot be blank")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%s/", meaBasePath, meaID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	mea := new(MDMEnterpriseApp)
	resp, err := s.client.Do(ctx, req, mea)
	if err != nil {
		return nil, resp, err
	}

	return mea, resp, err
}

// Delete a MDM enterprise app.
func (s *MDMEnterpriseAppsServiceOp) Delete(ctx context.Context, meaID string) (*Response, error) {
	if len(meaID) < 1 {
		return nil, NewArgError("meaID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", meaBasePath, meaID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing MDM enterprise apps
func (s *MDMEnterpriseAppsServiceOp) list(ctx context.Context, opt *ListOptions, meaOpt *listMEAOptions) ([]MDMEnterpriseApp, *Response, error) {
	path := meaBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, meaOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var meas []MDMEnterpriseApp
	resp, err := s.client.Do(ctx, req, &meas)
	if err != nil {
		return nil, resp, err
	}

	return meas, resp, err
}
