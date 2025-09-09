package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const mACMEIssuerBasePath = "mdm/acme_issuers/"

// MDMACMEIssuersService is an interface for interfacing with the MDM ACME issuers
// endpoints of the Zentral API.
type MDMACMEIssuersService interface {
	List(context.Context, *ListOptions) ([]MDMACMEIssuer, *Response, error)
	GetByID(context.Context, string) (*MDMACMEIssuer, *Response, error)
	GetByName(context.Context, string) (*MDMACMEIssuer, *Response, error)
	Create(context.Context, *MDMACMEIssuerRequest) (*MDMACMEIssuer, *Response, error)
	Update(context.Context, string, *MDMACMEIssuerRequest) (*MDMACMEIssuer, *Response, error)
	Delete(context.Context, string) (*Response, error)
}

// MDMACMEIssuersServiceOp handles communication with the MDM ACME issuers related
// methods of the Zentral API.
type MDMACMEIssuersServiceOp struct {
	client *Client
}

var _ MDMACMEIssuersService = &MDMACMEIssuersServiceOp{}

// MDMACMEIssuer represents a Zentral MDM ACME issuer.
type MDMACMEIssuer struct {
	ID              string  `json:"id"`
	ProvisioningUID *string `json:"provisioning_uid"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`

	DirectoryURL     string   `json:"directory_url"`
	KeySize          int      `json:"key_size"`
	KeyType          string   `json:"key_type"`
	UsageFlags       int      `json:"usage_flags"`
	ExtendedKeyUsage []string `json:"extended_key_usage"`
	HardwareBound    bool     `json:"hardware_bound"`
	Attest           bool     `json:"attest"`

	Backend         *string          `json:"backend"`
	IDent           *IDent           `json:"ident_kwargs"`
	MicrosoftCA     *MicrosoftCA     `json:"microsoft_ca_kwargs"`
	OktaCA          *MicrosoftCA     `json:"okta_ca_kwargs"`
	StaticChallenge *StaticChallenge `json:"static_challenge_kwargs"`

	Version int       `json:"version"`
	Created Timestamp `json:"created_at"`
	Updated Timestamp `json:"updated_at"`
}

func (mai MDMACMEIssuer) String() string {
	return Stringify(mai)
}

// MDMACMEIssuerRequest represents a request to create or update a MDM ACME issuer.
type MDMACMEIssuerRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`

	DirectoryURL     string   `json:"directory_url"`
	KeySize          int      `json:"key_size"`
	KeyType          string   `json:"key_type"`
	UsageFlags       int      `json:"usage_flags"`
	ExtendedKeyUsage []string `json:"extended_key_usage"`
	HardwareBound    bool     `json:"hardware_bound"`
	Attest           bool     `json:"attest"`

	Backend         string           `json:"backend"`
	IDent           *IDent           `json:"ident_kwargs"`
	MicrosoftCA     *MicrosoftCA     `json:"microsoft_ca_kwargs"`
	OktaCA          *MicrosoftCA     `json:"okta_ca_kwargs"`
	StaticChallenge *StaticChallenge `json:"static_challenge_kwargs"`
}

type listMAIOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the MDM ACME issuers.
func (s *MDMACMEIssuersServiceOp) List(ctx context.Context, opt *ListOptions) ([]MDMACMEIssuer, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a MDM ACME issuer by id.
func (s *MDMACMEIssuersServiceOp) GetByID(ctx context.Context, maiID string) (*MDMACMEIssuer, *Response, error) {
	if len(maiID) < 1 {
		return nil, nil, NewArgError("maiID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", mACMEIssuerBasePath, maiID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	mai := new(MDMACMEIssuer)

	resp, err := s.client.Do(ctx, req, mai)
	if err != nil {
		return nil, resp, err
	}

	return mai, resp, err
}

// GetByName retrieves a MDM ACME issuer by name.
func (s *MDMACMEIssuersServiceOp) GetByName(ctx context.Context, name string) (*MDMACMEIssuer, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listAIOpt := &listMAIOptions{Name: name}

	mais, resp, err := s.list(ctx, nil, listAIOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(mais) < 1 {
		return nil, resp, nil
	}

	return &mais[0], resp, err
}

// Create a new MDM ACME issuer.
func (s *MDMACMEIssuersServiceOp) Create(ctx context.Context, createRequest *MDMACMEIssuerRequest) (*MDMACMEIssuer, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, mACMEIssuerBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	mai := new(MDMACMEIssuer)
	resp, err := s.client.Do(ctx, req, mai)
	if err != nil {
		return nil, resp, err
	}

	return mai, resp, err
}

// Update a MDM ACME issuer.
func (s *MDMACMEIssuersServiceOp) Update(ctx context.Context, maiID string, updateRequest *MDMACMEIssuerRequest) (*MDMACMEIssuer, *Response, error) {
	if len(maiID) < 1 {
		return nil, nil, NewArgError("maiID", "cannot be blank")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%s/", mACMEIssuerBasePath, maiID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	mai := new(MDMACMEIssuer)
	resp, err := s.client.Do(ctx, req, mai)
	if err != nil {
		return nil, resp, err
	}

	return mai, resp, err
}

// Delete a MDM ACME issuer.
func (s *MDMACMEIssuersServiceOp) Delete(ctx context.Context, maiID string) (*Response, error) {
	if len(maiID) < 1 {
		return nil, NewArgError("maiID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", mACMEIssuerBasePath, maiID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing MDM ACME issuers.
func (s *MDMACMEIssuersServiceOp) list(ctx context.Context, opt *ListOptions, maiOpt *listMAIOptions) ([]MDMACMEIssuer, *Response, error) {
	path := mACMEIssuerBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, maiOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var mais []MDMACMEIssuer
	resp, err := s.client.Do(ctx, req, &mais)
	if err != nil {
		return nil, resp, err
	}

	return mais, resp, err
}
