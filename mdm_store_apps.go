package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const msaBasePath = "mdm/store_apps/"

// MDMStoreAppsService is an interface for interfacing with the MDM store app
// endpoints of the Zentral API
type MDMStoreAppsService interface {
	List(context.Context, *ListOptions) ([]MDMStoreApp, *Response, error)
	GetByID(context.Context, string) (*MDMStoreApp, *Response, error)
	Create(context.Context, *MDMStoreAppRequest) (*MDMStoreApp, *Response, error)
	Update(context.Context, string, *MDMStoreAppRequest) (*MDMStoreApp, *Response, error)
	Delete(context.Context, string) (*Response, error)
}

// MDMStoreAppsServiceOp handles communication with the MDM store apps related
// methods of the Zentral API.
type MDMStoreAppsServiceOp struct {
	client *Client
}

var _ MDMStoreAppsService = &MDMStoreAppsServiceOp{}

// MDMStoreApp represents a Zentral MDM store app
type MDMStoreApp struct {
	ID                                     string   `json:"id"`
	LocationAssetID                        int      `json:"location_asset"`
	AssociatedDomains                      []string `json:"associated_domains"`
	AssociatedDomainsEnableDirectDownloads bool     `json:"associated_domains_enable_direct_downloads"`
	Configuration                          *string  `json:"configuration"`
	ContentFilterUUID                      *string  `json:"content_filter_uuid"`
	DNSProxyUUID                           *string  `json:"dns_proxy_uuid"`
	VPNUUID                                *string  `json:"vpn_uuid"`
	PreventBackup                          bool     `json:"prevent_backup"`
	Removable                              bool     `json:"removable"`
	RemoveOnUnenroll                       bool     `json:"remove_on_unenroll"`
	MDMArtifactVersion
}

func (msa MDMStoreApp) String() string {
	return Stringify(msa)
}

// MDMStoreAppRequest represents a request to create or update a MDM store app
type MDMStoreAppRequest struct {
	LocationAssetID                        int      `json:"location_asset"`
	AssociatedDomains                      []string `json:"associated_domains"`
	AssociatedDomainsEnableDirectDownloads bool     `json:"associated_domains_enable_direct_downloads"`
	Configuration                          *string  `json:"configuration"`
	ContentFilterUUID                      *string  `json:"content_filter_uuid"`
	DNSProxyUUID                           *string  `json:"dns_proxy_uuid"`
	VPNUUID                                *string  `json:"vpn_uuid"`
	PreventBackup                          bool     `json:"prevent_backup"`
	Removable                              bool     `json:"removable"`
	RemoveOnUnenroll                       bool     `json:"remove_on_unenroll"`
	MDMArtifactVersionRequest
}

type listMSAOptions struct{}

// List lists all the MDM store apps.
func (s *MDMStoreAppsServiceOp) List(ctx context.Context, opt *ListOptions) ([]MDMStoreApp, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a MDM store app by id.
func (s *MDMStoreAppsServiceOp) GetByID(ctx context.Context, msaID string) (*MDMStoreApp, *Response, error) {
	if len(msaID) < 1 {
		return nil, nil, NewArgError("msaID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", msaBasePath, msaID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	msa := new(MDMStoreApp)

	resp, err := s.client.Do(ctx, req, msa)
	if err != nil {
		return nil, resp, err
	}

	return msa, resp, err
}

// Create a new MDM store app.
func (s *MDMStoreAppsServiceOp) Create(ctx context.Context, createRequest *MDMStoreAppRequest) (*MDMStoreApp, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, msaBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	msa := new(MDMStoreApp)
	resp, err := s.client.Do(ctx, req, msa)
	if err != nil {
		return nil, resp, err
	}

	return msa, resp, err
}

// Update a MDM store app.
func (s *MDMStoreAppsServiceOp) Update(ctx context.Context, msaID string, updateRequest *MDMStoreAppRequest) (*MDMStoreApp, *Response, error) {
	if len(msaID) < 1 {
		return nil, nil, NewArgError("msaID", "cannot be blank")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%s/", msaBasePath, msaID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	msa := new(MDMStoreApp)
	resp, err := s.client.Do(ctx, req, msa)
	if err != nil {
		return nil, resp, err
	}

	return msa, resp, err
}

// Delete a MDM store app.
func (s *MDMStoreAppsServiceOp) Delete(ctx context.Context, msaID string) (*Response, error) {
	if len(msaID) < 1 {
		return nil, NewArgError("msaID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", msaBasePath, msaID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing MDM store apps
func (s *MDMStoreAppsServiceOp) list(ctx context.Context, opt *ListOptions, msaOpt *listMSAOptions) ([]MDMStoreApp, *Response, error) {
	path := msaBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, msaOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var msas []MDMStoreApp
	resp, err := s.client.Do(ctx, req, &msas)
	if err != nil {
		return nil, resp, err
	}

	return msas, resp, err
}
