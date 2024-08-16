package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const rBasePath = "realms/realms/"

// RealmsRealmsService is an interface for interfacing with the realms
// endpoints of the Zentral API
type RealmsRealmsService interface {
	List(context.Context, *ListOptions) ([]RealmsRealm, *Response, error)
	GetByUUID(context.Context, string) (*RealmsRealm, *Response, error)
	GetByName(context.Context, string) (*RealmsRealm, *Response, error)
}

// RealmsRealmsServiceOp handles communication with the realms related
// methods of the Zentral API.
type RealmsRealmsServiceOp struct {
	client *Client
}

var _ RealmsRealmsService = &RealmsRealmsServiceOp{}

// LDAPConfig represents a Zentral Realm LDAP config
type LDAPConfig struct {
	Host         string `json:"host"`
	BindDN       string `json:"bind_dn"`
	BindPassword string `json:"bind_password"`
	UsersBaseDN  string `json:"users_base_dn"`
}

// OpenIDCConfig represents a Zentral Realm OpenIDC config
type OpenIDCConfig struct {
	DiscoveryURL string   `json:"discovery_url"`
	ClientID     string   `json:"client_id"`
	ClientSecret *string  `json:"client_secret"`
	ExtraScopes  []string `json:"extra_scopes"`
}

// SAMLConfig represents a Zentral Realm SAML config
type SAMLConfig struct {
	DefaultRelayState string `json:"default_relay_state"`
	IDPMetadata       string `json:"idp_metadata"`
}

// RealmsRealm represents a Zentral realm
type RealmsRealm struct {
	UUID               string         `json:"uuid"`
	Name               string         `json:"name"`
	Backend            string         `json:"backend"`
	LDAPConfig         *LDAPConfig    `json:"ldap_config"`
	OpenIDCConfig      *OpenIDCConfig `json:"openidc_config"`
	SAMLConfig         *SAMLConfig    `json:"saml_config"`
	EnabledForLogin    bool           `json:"enabled_for_login"`
	LoginSessionExpiry int            `json:"login_session_expiry"`
	UsernameClaim      string         `json:"username_claim"`
	EmailClaim         string         `json:"email_claim"`
	FirstNameClaim     string         `json:"first_name_claim"`
	LastNameClaim      string         `json:"last_name_claim"`
	FullNameClaim      string         `json:"full_name_claim"`
	CustomAttr1Claim   string         `json:"custom_attr_1_claim"`
	CustomAttr2Claim   string         `json:"custom_attr_2_claim"`
	SCIMEnabled        bool           `json:"scim_enabled"`
	Created            Timestamp      `json:"created_at"`
	Updated            Timestamp      `json:"updated_at"`
}

func (r RealmsRealm) String() string {
	return Stringify(r)
}

type listROptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the Realms realms.
func (s *RealmsRealmsServiceOp) List(ctx context.Context, opt *ListOptions) ([]RealmsRealm, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Realms realm by id.
func (s *RealmsRealmsServiceOp) GetByUUID(ctx context.Context, rUUID string) (*RealmsRealm, *Response, error) {
	if len(rUUID) < 1 {
		return nil, nil, NewArgError("rUUID", "cannot be empty")
	}

	path := fmt.Sprintf("%s%s/", rBasePath, rUUID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	r := new(RealmsRealm)

	resp, err := s.client.Do(ctx, req, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, err
}

// GetByName retrieves a Realms realm by name.
func (s *RealmsRealmsServiceOp) GetByName(ctx context.Context, name string) (*RealmsRealm, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listROpt := &listROptions{Name: name}

	rs, resp, err := s.list(ctx, nil, listROpt)
	if err != nil {
		return nil, resp, err
	}
	if len(rs) < 1 {
		return nil, resp, nil
	}

	return &rs[0], resp, err
}

// Helper method for listing Realms realms
func (s *RealmsRealmsServiceOp) list(ctx context.Context, opt *ListOptions, rOpt *listROptions) ([]RealmsRealm, *Response, error) {
	path := rBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, rOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var rs []RealmsRealm
	resp, err := s.client.Do(ctx, req, &rs)
	if err != nil {
		return nil, resp, err
	}

	return rs, resp, err
}
