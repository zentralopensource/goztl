package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const probesBasePath = "probes/probes/"

// ProbesService is an interface for interfacing with the probes
// endpoints of the Zentral API
type ProbesService interface {
	List(context.Context, *ListOptions) ([]Probe, *Response, error)
	GetByID(context.Context, int) (*Probe, *Response, error)
	GetByName(context.Context, string) (*Probe, *Response, error)
	Create(context.Context, *ProbeRequest) (*Probe, *Response, error)
	Update(context.Context, int, *ProbeRequest) (*Probe, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// ProbesServiceOp handles communication with the probes related
// methods of the Zentral API.
type ProbesServiceOp struct {
	client *Client
}

var _ ProbesService = &ProbesServiceOp{}

// Probe represents a Zentral probe

type InventoryFilter struct {
	MetaBusinessUnitIDs []int    `json:"meta_business_unit_ids"`
	TagIDs              []int    `json:"tag_ids"`
	Platforms           []string `json:"platforms"`
	Types               []string `json:"types"`
}

type MetadataFilter struct {
	EventTypes       []string `json:"event_types"`
	EventTags        []string `json:"event_tags"`
	EventRoutingKeys []string `json:"event_routing_keys"`
}

type PayloadFilterItem struct {
	Attribute string   `json:"attribute"`
	Operator  string   `json:"operator"`
	Values    []string `json:"values"`
}

type Probe struct {
	ID               int                   `json:"id"`
	Name             string                `json:"name"`
	Slug             string                `json:"slug"`
	Description      string                `json:"description"`
	InventoryFilters []InventoryFilter     `json:"inventory_filters"`
	MetadataFilters  []MetadataFilter      `json:"metadata_filters"`
	PayloadFilters   [][]PayloadFilterItem `json:"payload_filters"`
	IncidentSeverity *int                  `json:"incident_severity"`
	ActionIDs        []string              `json:"actions"`
	Active           bool                  `json:"active"`
	Created          Timestamp             `json:"created_at"`
	Updated          Timestamp             `json:"updated_at"`
}

func (p Probe) String() string {
	return Stringify(p)
}

// ProbeRequest represents a request to create or update a probe
type ProbeRequest struct {
	Name             string                `json:"name"`
	Description      string                `json:"description"`
	InventoryFilters []InventoryFilter     `json:"inventory_filters"`
	MetadataFilters  []MetadataFilter      `json:"metadata_filters"`
	PayloadFilters   [][]PayloadFilterItem `json:"payload_filters"`
	IncidentSeverity *int                  `json:"incident_severity"`
	ActionIDs        []string              `json:"actions"`
	Active           bool                  `json:"active"`
}

type listProbeOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the probes
func (s *ProbesServiceOp) List(ctx context.Context, opt *ListOptions) ([]Probe, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a probe by id
func (s *ProbesServiceOp) GetByID(ctx context.Context, pID int) (*Probe, *Response, error) {
	if pID < 1 {
		return nil, nil, NewArgError("pID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", probesBasePath, pID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	p := new(Probe)

	resp, err := s.client.Do(ctx, req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, err
}

// GetByName retrieves a probe by name
func (s *ProbesServiceOp) GetByName(ctx context.Context, name string) (*Probe, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listProbeOpt := &listProbeOptions{Name: name}

	ps, resp, err := s.list(ctx, nil, listProbeOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(ps) < 1 {
		return nil, resp, nil
	}

	return &ps[0], resp, err
}

// Create a new probe
func (s *ProbesServiceOp) Create(ctx context.Context, createRequest *ProbeRequest) (*Probe, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, probesBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	p := new(Probe)
	resp, err := s.client.Do(ctx, req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, err
}

// Update a probe
func (s *ProbesServiceOp) Update(ctx context.Context, pID int, updateRequest *ProbeRequest) (*Probe, *Response, error) {
	if pID < 1 {
		return nil, nil, NewArgError("pID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", probesBasePath, pID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	p := new(Probe)
	resp, err := s.client.Do(ctx, req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, err
}

// Delete a probe
func (s *ProbesServiceOp) Delete(ctx context.Context, pID int) (*Response, error) {
	if pID < 1 {
		return nil, NewArgError("pID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", probesBasePath, pID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing probes
func (s *ProbesServiceOp) list(ctx context.Context, opt *ListOptions, pOpt *listProbeOptions) ([]Probe, *Response, error) {
	path := probesBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, pOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var ps []Probe
	resp, err := s.client.Do(ctx, req, &ps)
	if err != nil {
		return nil, resp, err
	}

	return ps, resp, err
}
