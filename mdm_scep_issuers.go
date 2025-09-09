package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const mSCEPIssuerBasePath = "mdm/scep_issuers/"

// MDMSCEPIssuersService is an interface for interfacing with the MDM SCEP issuers
// endpoints of the Zentral API.
type MDMSCEPIssuersService interface {
	List(context.Context, *ListOptions) ([]MDMSCEPIssuer, *Response, error)
	GetByID(context.Context, string) (*MDMSCEPIssuer, *Response, error)
	GetByName(context.Context, string) (*MDMSCEPIssuer, *Response, error)
	Create(context.Context, *MDMSCEPIssuerRequest) (*MDMSCEPIssuer, *Response, error)
	Update(context.Context, string, *MDMSCEPIssuerRequest) (*MDMSCEPIssuer, *Response, error)
	Delete(context.Context, string) (*Response, error)
}

// MDMSCEPIssuersServiceOp handles communication with the MDM SCEP issuers related
// methods of the Zentral API.
type MDMSCEPIssuersServiceOp struct {
	client *Client
}

var _ MDMSCEPIssuersService = &MDMSCEPIssuersServiceOp{}

// MDMSCEPIssuer represents a Zentral MDM SCEP issuer.
type MDMSCEPIssuer struct {
	ID              string  `json:"id"`
	ProvisioningUID *string `json:"provisioning_uid"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`

	URL      string `json:"url"`
	KeySize  int    `json:"key_size"`
	KeyUsage int    `json:"key_usage"`

	Backend         *string          `json:"backend"`
	IDent           *IDent           `json:"ident_kwargs"`
	MicrosoftCA     *MicrosoftCA     `json:"microsoft_ca_kwargs"`
	OktaCA          *MicrosoftCA     `json:"okta_ca_kwargs"`
	StaticChallenge *StaticChallenge `json:"static_challenge_kwargs"`

	Version int       `json:"version"`
	Created Timestamp `json:"created_at"`
	Updated Timestamp `json:"updated_at"`
}

func (msi MDMSCEPIssuer) String() string {
	return Stringify(msi)
}

// MDMSCEPIssuerRequest represents a request to create or update a MDM SCEP issuer.
type MDMSCEPIssuerRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`

	URL      string `json:"url"`
	KeySize  int    `json:"key_size"`
	KeyUsage int    `json:"key_usage"`

	Backend         string           `json:"backend"`
	IDent           *IDent           `json:"ident_kwargs"`
	MicrosoftCA     *MicrosoftCA     `json:"microsoft_ca_kwargs"`
	OktaCA          *MicrosoftCA     `json:"okta_ca_kwargs"`
	StaticChallenge *StaticChallenge `json:"static_challenge_kwargs"`
}

type listMSIOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the MDM SCEP issuers.
func (s *MDMSCEPIssuersServiceOp) List(ctx context.Context, opt *ListOptions) ([]MDMSCEPIssuer, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a MDM SCEP issuer by id.
func (s *MDMSCEPIssuersServiceOp) GetByID(ctx context.Context, msiID string) (*MDMSCEPIssuer, *Response, error) {
	if len(msiID) < 1 {
		return nil, nil, NewArgError("msiID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", mSCEPIssuerBasePath, msiID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	msi := new(MDMSCEPIssuer)

	resp, err := s.client.Do(ctx, req, msi)
	if err != nil {
		return nil, resp, err
	}

	return msi, resp, err
}

// GetByName retrieves a MDM SCEP issuer by name.
func (s *MDMSCEPIssuersServiceOp) GetByName(ctx context.Context, name string) (*MDMSCEPIssuer, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listAIOpt := &listMSIOptions{Name: name}

	msis, resp, err := s.list(ctx, nil, listAIOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(msis) < 1 {
		return nil, resp, nil
	}

	return &msis[0], resp, err
}

// Create a new MDM SCEP issuer.
func (s *MDMSCEPIssuersServiceOp) Create(ctx context.Context, createRequest *MDMSCEPIssuerRequest) (*MDMSCEPIssuer, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, mSCEPIssuerBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	msi := new(MDMSCEPIssuer)
	resp, err := s.client.Do(ctx, req, msi)
	if err != nil {
		return nil, resp, err
	}

	return msi, resp, err
}

// Update a MDM SCEP issuer.
func (s *MDMSCEPIssuersServiceOp) Update(ctx context.Context, msiID string, updateRequest *MDMSCEPIssuerRequest) (*MDMSCEPIssuer, *Response, error) {
	if len(msiID) < 1 {
		return nil, nil, NewArgError("msiID", "cannot be blank")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%s/", mSCEPIssuerBasePath, msiID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	msi := new(MDMSCEPIssuer)
	resp, err := s.client.Do(ctx, req, msi)
	if err != nil {
		return nil, resp, err
	}

	return msi, resp, err
}

// Delete a MDM SCEP issuer.
func (s *MDMSCEPIssuersServiceOp) Delete(ctx context.Context, msiID string) (*Response, error) {
	if len(msiID) < 1 {
		return nil, NewArgError("msiID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", mSCEPIssuerBasePath, msiID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing MDM SCEP issuers.
func (s *MDMSCEPIssuersServiceOp) list(ctx context.Context, opt *ListOptions, msiOpt *listMSIOptions) ([]MDMSCEPIssuer, *Response, error) {
	path := mSCEPIssuerBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, msiOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var msis []MDMSCEPIssuer
	resp, err := s.client.Do(ctx, req, &msis)
	if err != nil {
		return nil, resp, err
	}

	return msis, resp, err
}
