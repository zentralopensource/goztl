package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const mscepBasePath = "mdm/scep_configs/"

// MDMSCEPConfigsService is an interface for interfacing with the MDM SCEP config
// endpoints of the Zentral API
type MDMSCEPConfigsService interface {
	List(context.Context, *ListOptions) ([]MDMSCEPConfig, *Response, error)
	GetByID(context.Context, int) (*MDMSCEPConfig, *Response, error)
	GetByName(context.Context, string) (*MDMSCEPConfig, *Response, error)
}

// MDMSCEPConfigsServiceOp handles communication with the MDM SCEP configs related
// methods of the Zentral API.
type MDMSCEPConfigsServiceOp struct {
	client *Client
}

var _ MDMSCEPConfigsService = &MDMSCEPConfigsServiceOp{}

// MDMSCEPConfig represents a Zentral MDM SCEP config

type MicrosoftCASCEPChallenge struct {
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type StaticSCEPChallenge struct {
	Challenge string `json:"challenge"`
}

type MDMSCEPConfig struct {
	ID                 int     `json:"id"`
	ProvisioningUID    *string `json:"provisioning_uid"`
	Name               string  `json:"name"`
	URL                string  `json:"url"`
	KeyUsage           int     `json:"key_usage"`
	KeyIsExtractable   bool    `json:"key_is_extractable"`
	KeySize            int     `json:"keysize"`
	AllowAllAppsAccess bool    `json:"allow_all_apps_access"`

	// Challenge info only present if not provisioned
	ChallengeType        *string                   `json:"challenge_type"`
	MicrosoftCAChallenge *MicrosoftCASCEPChallenge `json:"microsoft_ca_challenge_kwargs"`
	OktaCAChallenge      *MicrosoftCASCEPChallenge `json:"okta_ca_challenge_kwargs"`
	StaticChallenge      *StaticSCEPChallenge      `json:"static_challenge_kwargs"`

	Created Timestamp `json:"created_at"`
	Updated Timestamp `json:"updated_at"`
}

func (msc MDMSCEPConfig) String() string {
	return Stringify(msc)
}

type listMSCOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the MDM SCEP configs.
func (s *MDMSCEPConfigsServiceOp) List(ctx context.Context, opt *ListOptions) ([]MDMSCEPConfig, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a MDM SCEP config by id.
func (s *MDMSCEPConfigsServiceOp) GetByID(ctx context.Context, mscID int) (*MDMSCEPConfig, *Response, error) {
	if mscID < 1 {
		return nil, nil, NewArgError("mscID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mscepBasePath, mscID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	sc := new(MDMSCEPConfig)

	resp, err := s.client.Do(ctx, req, sc)
	if err != nil {
		return nil, resp, err
	}

	return sc, resp, err
}

// GetByName retrieves a MDM SCEP config by name.
func (s *MDMSCEPConfigsServiceOp) GetByName(ctx context.Context, name string) (*MDMSCEPConfig, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listSCOpt := &listMSCOptions{Name: name}

	scs, resp, err := s.list(ctx, nil, listSCOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(scs) < 1 {
		return nil, resp, nil
	}

	return &scs[0], resp, err
}

// Helper method for listing MDM SCEP configs
func (s *MDMSCEPConfigsServiceOp) list(ctx context.Context, opt *ListOptions, mscOpt *listMSCOptions) ([]MDMSCEPConfig, *Response, error) {
	path := mscepBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, mscOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var scs []MDMSCEPConfig
	resp, err := s.client.Do(ctx, req, &scs)
	if err != nil {
		return nil, resp, err
	}

	return scs, resp, err
}
