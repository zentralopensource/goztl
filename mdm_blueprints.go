package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const mbBasePath = "mdm/blueprints/"

// MDMBlueprintsService is an interface for interfacing with the MDM blueprint
// endpoints of the Zentral API
type MDMBlueprintsService interface {
	List(context.Context, *ListOptions) ([]MDMBlueprint, *Response, error)
	GetByID(context.Context, int) (*MDMBlueprint, *Response, error)
	GetByName(context.Context, string) (*MDMBlueprint, *Response, error)
	Create(context.Context, *MDMBlueprintRequest) (*MDMBlueprint, *Response, error)
	Update(context.Context, int, *MDMBlueprintRequest) (*MDMBlueprint, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// MDMBlueprintsServiceOp handles communication with the MDM blueprints related
// methods of the Zentral API.
type MDMBlueprintsServiceOp struct {
	client *Client
}

var _ MDMBlueprintsService = &MDMBlueprintsServiceOp{}

// MDMBlueprint represents a Zentral MDM blueprint
type MDMBlueprint struct {
	ID                           int       `json:"id,omitempty"`
	Name                         string    `json:"name"`
	InventoryInterval            int       `json:"inventory_interval"`
	CollectApps                  int       `json:"collect_apps"`
	CollectCertificates          int       `json:"collect_certificates"`
	CollectProfiles              int       `json:"collect_profiles"`
	LegacyProfilesViaDDM         bool      `json:"legacy_profiles_via_ddm"`
	DefaultLocationID            *int      `json:"default_location"`
	FileVaultConfigID            *int      `json:"filevault_config"`
	RecoveryPasswordConfigID     *int      `json:"recovery_password_config"`
	SoftwareUpdateEnforcementIDs []int     `json:"software_update_enforcements"`
	Created                      Timestamp `json:"created_at,omitempty"`
	Updated                      Timestamp `json:"updated_at,omitempty"`
}

func (mb MDMBlueprint) String() string {
	return Stringify(mb)
}

// MDMBlueprintRequest represents a request to create or update a MDM blueprint
type MDMBlueprintRequest struct {
	Name                         string `json:"name"`
	InventoryInterval            int    `json:"inventory_interval"`
	CollectApps                  int    `json:"collect_apps"`
	CollectCertificates          int    `json:"collect_certificates"`
	CollectProfiles              int    `json:"collect_profiles"`
	LegacyProfilesViaDDM         bool   `json:"legacy_profiles_via_ddm"`
	DefaultLocationID            *int   `json:"default_location"`
	FileVaultConfigID            *int   `json:"filevault_config"`
	RecoveryPasswordConfigID     *int   `json:"recovery_password_config"`
	SoftwareUpdateEnforcementIDs []int  `json:"software_update_enforcements"`
}

type listMBOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the MDM blueprints.
func (s *MDMBlueprintsServiceOp) List(ctx context.Context, opt *ListOptions) ([]MDMBlueprint, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a MDM blueprint by id.
func (s *MDMBlueprintsServiceOp) GetByID(ctx context.Context, mbID int) (*MDMBlueprint, *Response, error) {
	if mbID < 1 {
		return nil, nil, NewArgError("mbID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mbBasePath, mbID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	sc := new(MDMBlueprint)

	resp, err := s.client.Do(ctx, req, sc)
	if err != nil {
		return nil, resp, err
	}

	return sc, resp, err
}

// GetByName retrieves a MDM blueprint by name.
func (s *MDMBlueprintsServiceOp) GetByName(ctx context.Context, name string) (*MDMBlueprint, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listSCOpt := &listMBOptions{Name: name}

	scs, resp, err := s.list(ctx, nil, listSCOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(scs) < 1 {
		return nil, resp, nil
	}

	return &scs[0], resp, err
}

// Create a new MDM blueprint.
func (s *MDMBlueprintsServiceOp) Create(ctx context.Context, createRequest *MDMBlueprintRequest) (*MDMBlueprint, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, mbBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	sc := new(MDMBlueprint)
	resp, err := s.client.Do(ctx, req, sc)
	if err != nil {
		return nil, resp, err
	}

	return sc, resp, err
}

// Update a MDM blueprint.
func (s *MDMBlueprintsServiceOp) Update(ctx context.Context, mbID int, updateRequest *MDMBlueprintRequest) (*MDMBlueprint, *Response, error) {
	if mbID < 1 {
		return nil, nil, NewArgError("mbID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", mbBasePath, mbID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	sc := new(MDMBlueprint)
	resp, err := s.client.Do(ctx, req, sc)
	if err != nil {
		return nil, resp, err
	}

	return sc, resp, err
}

// Delete a MDM blueprint.
func (s *MDMBlueprintsServiceOp) Delete(ctx context.Context, mbID int) (*Response, error) {
	if mbID < 1 {
		return nil, NewArgError("mbID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mbBasePath, mbID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing MDM blueprints
func (s *MDMBlueprintsServiceOp) list(ctx context.Context, opt *ListOptions, mbOpt *listMBOptions) ([]MDMBlueprint, *Response, error) {
	path := mbBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, mbOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var scs []MDMBlueprint
	resp, err := s.client.Do(ctx, req, &scs)
	if err != nil {
		return nil, resp, err
	}

	return scs, resp, err
}
