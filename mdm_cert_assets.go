package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const mcaBasePath = "mdm/cert_assets/"

// MDMCertAssetsService is an interface for interfacing with the MDM cert asset
// endpoints of the Zentral API
type MDMCertAssetsService interface {
	List(context.Context, *ListOptions) ([]MDMCertAsset, *Response, error)
	GetByID(context.Context, string) (*MDMCertAsset, *Response, error)
	Create(context.Context, *MDMCertAssetRequest) (*MDMCertAsset, *Response, error)
	Update(context.Context, string, *MDMCertAssetRequest) (*MDMCertAsset, *Response, error)
	Delete(context.Context, string) (*Response, error)
}

// MDMCertAssetsServiceOp handles communication with the MDM cert assets related
// methods of the Zentral API.
type MDMCertAssetsServiceOp struct {
	client *Client
}

var _ MDMCertAssetsService = &MDMCertAssetsServiceOp{}

// MDMCertAsset represents a Zentral MDM cert asset

type MDMCertAssetRDN struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type MDMCertAssetSubjectAltName struct {
	DNSName         *string `json:"dNSName"`
	NTPrincipalName *string `json:"ntPrincipalName"`
	RFC822Name      *string `json:"rfc822Name"`
	URI             *string `json:"uniformResourceIdentifier"`
}

type MDMCertAsset struct {
	ID             string                     `json:"id"`
	ACMEIssuerUUID *string                    `json:"acme_issuer"`
	SCEPIssuerUUID *string                    `json:"scep_issuer"`
	Accessible     string                     `json:"accessible"`
	Subject        []MDMCertAssetRDN          `json:"subject"`
	SubjectAltName MDMCertAssetSubjectAltName `json:"subject_alt_name"`
	MDMArtifactVersion
}

func (mca MDMCertAsset) String() string {
	return Stringify(mca)
}

// MDMCertAssetRequest represents a request to create or update a MDM cert asset
type MDMCertAssetRequest struct {
	ACMEIssuerUUID *string                    `json:"acme_issuer"`
	SCEPIssuerUUID *string                    `json:"scep_issuer"`
	Accessible     string                     `json:"accessible"`
	Subject        []MDMCertAssetRDN          `json:"subject"`
	SubjectAltName MDMCertAssetSubjectAltName `json:"subject_alt_name"`
	MDMArtifactVersionRequest
}

type listMCAOptions struct{}

// List lists all the MDM cert assets.
func (s *MDMCertAssetsServiceOp) List(ctx context.Context, opt *ListOptions) ([]MDMCertAsset, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a MDM cert asset by id.
func (s *MDMCertAssetsServiceOp) GetByID(ctx context.Context, mcaID string) (*MDMCertAsset, *Response, error) {
	if len(mcaID) < 1 {
		return nil, nil, NewArgError("mcaID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", mcaBasePath, mcaID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	mca := new(MDMCertAsset)

	resp, err := s.client.Do(ctx, req, mca)
	if err != nil {
		return nil, resp, err
	}

	return mca, resp, err
}

// Create a new MDM cert asset.
func (s *MDMCertAssetsServiceOp) Create(ctx context.Context, createRequest *MDMCertAssetRequest) (*MDMCertAsset, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, mcaBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	mca := new(MDMCertAsset)
	resp, err := s.client.Do(ctx, req, mca)
	if err != nil {
		return nil, resp, err
	}

	return mca, resp, err
}

// Update a MDM cert asset.
func (s *MDMCertAssetsServiceOp) Update(ctx context.Context, mcaID string, updateRequest *MDMCertAssetRequest) (*MDMCertAsset, *Response, error) {
	if len(mcaID) < 1 {
		return nil, nil, NewArgError("mcaID", "cannot be blank")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%s/", mcaBasePath, mcaID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	mca := new(MDMCertAsset)
	resp, err := s.client.Do(ctx, req, mca)
	if err != nil {
		return nil, resp, err
	}

	return mca, resp, err
}

// Delete a MDM cert asset.
func (s *MDMCertAssetsServiceOp) Delete(ctx context.Context, mcaID string) (*Response, error) {
	if len(mcaID) < 1 {
		return nil, NewArgError("mcaID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", mcaBasePath, mcaID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing MDM cert assets
func (s *MDMCertAssetsServiceOp) list(ctx context.Context, opt *ListOptions, mcaOpt *listMCAOptions) ([]MDMCertAsset, *Response, error) {
	path := mcaBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, mcaOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var mcas []MDMCertAsset
	resp, err := s.client.Do(ctx, req, &mcas)
	if err != nil {
		return nil, resp, err
	}

	return mcas, resp, err
}
