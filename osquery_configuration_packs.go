package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const ocpBasePath = "osquery/configuration_packs/"

// OsqueryConfigurationPacksService is an interface for interfacing with the Osquery configuration packs
// endpoints of the Zentral API
type OsqueryConfigurationPacksService interface {
	List(context.Context, *ListOptions) ([]OsqueryConfigurationPack, *Response, error)
	GetByID(context.Context, int) (*OsqueryConfigurationPack, *Response, error)
	GetByConfigurationID(context.Context, int) ([]OsqueryConfigurationPack, *Response, error)
	GetByPackID(context.Context, int) ([]OsqueryConfigurationPack, *Response, error)
	Create(context.Context, *OsqueryConfigurationPackRequest) (*OsqueryConfigurationPack, *Response, error)
	Update(context.Context, int, *OsqueryConfigurationPackRequest) (*OsqueryConfigurationPack, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// OsqueryConfigurationPacksServiceOp handles communication with the Osquery configuration packs related
// methods of the Zentral API.
type OsqueryConfigurationPacksServiceOp struct {
	client *Client
}

var _ OsqueryConfigurationPacksService = &OsqueryConfigurationPacksServiceOp{}

// OsqueryConfigurationPack represents a Zentral Osquery configuration pack
type OsqueryConfigurationPack struct {
	ID              int   `json:"id"`
	ConfigurationID int   `json:"configuration"`
	PackID          int   `json:"pack"`
	TagIDs          []int `json:"tags"`
}

func (ocp OsqueryConfigurationPack) String() string {
	return Stringify(ocp)
}

// OsqueryConfigurationPackRequest represents a request to create or update a Osquery configuration pack
type OsqueryConfigurationPackRequest struct {
	ConfigurationID int   `json:"configuration"`
	PackID          int   `json:"pack"`
	TagIDs          []int `json:"tags"`
}

type listOCPOptions struct {
	ConfigurationID int `url:"configuration_id,omitempty"`
	PackID          int `url:"pack_id,omitempty"`
}

// List lists all the Osquery configuration packs.
func (s *OsqueryConfigurationPacksServiceOp) List(ctx context.Context, opt *ListOptions) ([]OsqueryConfigurationPack, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Osquery configuration pack by id.
func (s *OsqueryConfigurationPacksServiceOp) GetByID(ctx context.Context, ocpID int) (*OsqueryConfigurationPack, *Response, error) {
	if ocpID < 1 {
		return nil, nil, NewArgError("ocpID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", ocpBasePath, ocpID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	ocp := new(OsqueryConfigurationPack)

	resp, err := s.client.Do(ctx, req, ocp)
	if err != nil {
		return nil, resp, err
	}

	return ocp, resp, err
}

// GetByConfigurationID retrieves Osquery configuration packs by configuration ID.
func (s *OsqueryConfigurationPacksServiceOp) GetByConfigurationID(ctx context.Context, cfgID int) ([]OsqueryConfigurationPack, *Response, error) {
	if cfgID < 1 {
		return nil, nil, NewArgError("cfgID", "cannot be less than 1")
	}

	listOCPOpt := &listOCPOptions{ConfigurationID: cfgID}

	return s.list(ctx, nil, listOCPOpt)
}

// GetByPackID retrieves Osquery configuration packs by pack ID.
func (s *OsqueryConfigurationPacksServiceOp) GetByPackID(ctx context.Context, packID int) ([]OsqueryConfigurationPack, *Response, error) {
	if packID < 1 {
		return nil, nil, NewArgError("packID", "cannot be less than 1")
	}

	listOCPOpt := &listOCPOptions{PackID: packID}

	return s.list(ctx, nil, listOCPOpt)
}

// Create a new Osquery configuration pack.
func (s *OsqueryConfigurationPacksServiceOp) Create(ctx context.Context, createRequest *OsqueryConfigurationPackRequest) (*OsqueryConfigurationPack, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, ocpBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	ocp := new(OsqueryConfigurationPack)
	resp, err := s.client.Do(ctx, req, ocp)
	if err != nil {
		return nil, resp, err
	}

	return ocp, resp, err
}

// Update a Osquery configuration pack.
func (s *OsqueryConfigurationPacksServiceOp) Update(ctx context.Context, ocpID int, updateRequest *OsqueryConfigurationPackRequest) (*OsqueryConfigurationPack, *Response, error) {
	if ocpID < 1 {
		return nil, nil, NewArgError("ocpID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", ocpBasePath, ocpID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	ocp := new(OsqueryConfigurationPack)
	resp, err := s.client.Do(ctx, req, ocp)
	if err != nil {
		return nil, resp, err
	}

	return ocp, resp, err
}

// Delete a Osquery configuration pack.
func (s *OsqueryConfigurationPacksServiceOp) Delete(ctx context.Context, ocpID int) (*Response, error) {
	if ocpID < 1 {
		return nil, NewArgError("ocpID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", ocpBasePath, ocpID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Osquery configuration packs.
func (s *OsqueryConfigurationPacksServiceOp) list(ctx context.Context, opt *ListOptions, ocpOpt *listOCPOptions) ([]OsqueryConfigurationPack, *Response, error) {
	path := ocpBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, ocpOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var ocps []OsqueryConfigurationPack
	resp, err := s.client.Do(ctx, req, &ocps)
	if err != nil {
		return nil, resp, err
	}

	return ocps, resp, err
}
