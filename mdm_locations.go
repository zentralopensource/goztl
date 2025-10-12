package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const mlBasePath = "mdm/locations/"

// MDMLocationsService is an interface for interfacing with the MDM location
// endpoints of the Zentral API
type MDMLocationsService interface {
	List(context.Context, *ListOptions) ([]MDMLocation, *Response, error)
	GetByID(context.Context, int) (*MDMLocation, *Response, error)
	GetByMDMInfoID(context.Context, string) (*MDMLocation, *Response, error)
	GetByName(context.Context, string) (*MDMLocation, *Response, error)
}

// MDMLocationsServiceOp handles communication with the MDM locations related
// methods of the Zentral API.
type MDMLocationsServiceOp struct {
	client *Client
}

var _ MDMLocationsService = &MDMLocationsServiceOp{}

// MDMLocation represents a Zentral MDM location
type MDMLocation struct {
	ID                        int       `json:"id"`
	OrganizationName          string    `json:"organization_name"`
	Name                      string    `json:"name"`
	CountryCode               string    `json:"country_code"`
	LibraryUID                string    `json:"library_uid"`
	MDMInfoID                 string    `json:"mdm_info_id"`
	Platform                  string    `json:"platform"`
	WebsiteURL                string    `json:"website_url"`
	ServerTokenExpirationDate Timestamp `json:"server_token_expiration_date"`
	Created                   Timestamp `json:"created_at"`
	Updated                   Timestamp `json:"updated_at"`
}

func (ml MDMLocation) String() string {
	return Stringify(ml)
}

type listMLOptions struct {
	Name      string `url:"name,omitempty"`
	MDMInfoID string `url:"mdm_info_id,omitempty"`
}

// List lists all the MDM locations.
func (s *MDMLocationsServiceOp) List(ctx context.Context, opt *ListOptions) ([]MDMLocation, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a MDM location by id.
func (s *MDMLocationsServiceOp) GetByID(ctx context.Context, mlID int) (*MDMLocation, *Response, error) {
	if mlID < 1 {
		return nil, nil, NewArgError("mlID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mlBasePath, mlID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	l := new(MDMLocation)

	resp, err := s.client.Do(ctx, req, l)
	if err != nil {
		return nil, resp, err
	}

	return l, resp, err
}

// GetByMDMInfoID retrieves a MDM location by its MDM info ID.
func (s *MDMLocationsServiceOp) GetByMDMInfoID(ctx context.Context, mii string) (*MDMLocation, *Response, error) {
	if len(mii) < 1 {
		return nil, nil, NewArgError("mii", "cannot be blank")
	}

	listMLOpt := &listMLOptions{MDMInfoID: mii}

	ls, resp, err := s.list(ctx, nil, listMLOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(ls) < 1 {
		return nil, resp, nil
	}

	return &ls[0], resp, err
}

// GetByName retrieves a MDM location by name.
func (s *MDMLocationsServiceOp) GetByName(ctx context.Context, name string) (*MDMLocation, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listMLOpt := &listMLOptions{Name: name}

	ls, resp, err := s.list(ctx, nil, listMLOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(ls) < 1 {
		return nil, resp, nil
	}

	return &ls[0], resp, err
}

// Helper method for listing MDM locations
func (s *MDMLocationsServiceOp) list(ctx context.Context, opt *ListOptions, mlOpt *listMLOptions) ([]MDMLocation, *Response, error) {
	path := mlBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, mlOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var ls []MDMLocation
	resp, err := s.client.Do(ctx, req, &ls)
	if err != nil {
		return nil, resp, err
	}

	return ls, resp, err
}
