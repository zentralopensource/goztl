package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const srBasePath = "santa/rules/"

// SantaRulesService is an interface for interfacing with the Santa rules
// endpoints of the Zentral API
type SantaRulesService interface {
	List(context.Context, *ListOptions) ([]SantaRule, *Response, error)
	GetByID(context.Context, int) (*SantaRule, *Response, error)
	GetByConfigurationID(context.Context, int) ([]SantaRule, *Response, error)
	GetByTargetIdentifier(context.Context, string) ([]SantaRule, *Response, error)
	GetByTargetType(context.Context, string) ([]SantaRule, *Response, error)
	Create(context.Context, *SantaRuleRequest) (*SantaRule, *Response, error)
	Update(context.Context, int, *SantaRuleRequest) (*SantaRule, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// SantaRulesServiceOp handles communication with the Santa enrollments related
// methods of the Zentral API.
type SantaRulesServiceOp struct {
	client *Client
}

var _ SantaRulesService = &SantaRulesServiceOp{}

// SantaRule represents a Zentral SantaRule
type SantaRule struct {
	ID                    int       `json:"id"`
	ConfigurationID       int       `json:"configuration"`
	Policy                int       `json:"policy"`
	CELExpr               string    `json:"cel_expr"`
	TargetType            string    `json:"target_type"`
	TargetIdentifier      string    `json:"target_identifier"`
	Description           string    `json:"description"`
	CustomMessage         string    `json:"custom_msg"`
	RulesetID             *int      `json:"ruleset"`
	PrimaryUsers          []string  `json:"primary_users"`
	ExcludedPrimaryUsers  []string  `json:"excluded_primary_users"`
	SerialNumbers         []string  `json:"serial_numbers"`
	ExcludedSerialNumbers []string  `json:"excluded_serial_numbers"`
	TagIDs                []int     `json:"tags"`
	ExcludedTagIDs        []int     `json:"excluded_tags"`
	Version               int       `json:"version"`
	Created               Timestamp `json:"created_at"`
	Updated               Timestamp `json:"updated_at"`
}

func (sr SantaRule) String() string {
	return Stringify(sr)
}

// SantaRuleRequest represents a request to create or update a Santa rule
type SantaRuleRequest struct {
	ConfigurationID       int      `json:"configuration"`
	Policy                int      `json:"policy"`
	CELExpr               string   `json:"cel_expr"`
	TargetType            string   `json:"target_type"`
	TargetIdentifier      string   `json:"target_identifier"`
	Description           string   `json:"description"`
	CustomMessage         string   `json:"custom_msg"`
	PrimaryUsers          []string `json:"primary_users"`
	ExcludedPrimaryUsers  []string `json:"excluded_primary_users"`
	SerialNumbers         []string `json:"serial_numbers"`
	ExcludedSerialNumbers []string `json:"excluded_serial_numbers"`
	TagIDs                []int    `json:"tags"`
	ExcludedTagIDs        []int    `json:"excluded_tags"`
}

type listSROptions struct {
	ConfigurationID  int    `url:"configuration_id,omitempty"`
	TargetType       string `url:"target_type,omitempty"`
	TargetIdentifier string `url:"target_identifier,omitempty"`
}

// List lists all the Santa rules.
func (s *SantaRulesServiceOp) List(ctx context.Context, opt *ListOptions) ([]SantaRule, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Santa rule by id.
func (s *SantaRulesServiceOp) GetByID(ctx context.Context, srID int) (*SantaRule, *Response, error) {
	if srID < 1 {
		return nil, nil, NewArgError("srID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", srBasePath, srID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	sr := new(SantaRule)

	resp, err := s.client.Do(ctx, req, sr)
	if err != nil {
		return nil, resp, err
	}

	return sr, resp, err
}

// GetByConfigurationID retrieves the Santa rules for a given configuration.
func (s *SantaRulesServiceOp) GetByConfigurationID(ctx context.Context, configuration_id int) ([]SantaRule, *Response, error) {
	if configuration_id < 1 {
		return nil, nil, NewArgError("configuration_id", "cannot be negative")
	}

	listSROpt := &listSROptions{ConfigurationID: configuration_id}

	return s.list(ctx, nil, listSROpt)
}

// GetByTargetIdentifier retrieves the Santa rules for a given target identifier
func (s *SantaRulesServiceOp) GetByTargetIdentifier(ctx context.Context, target_identifier string) ([]SantaRule, *Response, error) {
	if len(target_identifier) == 0 {
		return nil, nil, NewArgError("target_identifier", "cannot be empty")
	}

	listSROpt := &listSROptions{TargetIdentifier: target_identifier}

	return s.list(ctx, nil, listSROpt)
}

// GetByTargetType retrieves the Santa rules for a given target type
func (s *SantaRulesServiceOp) GetByTargetType(ctx context.Context, target_type string) ([]SantaRule, *Response, error) {
	if len(target_type) == 0 {
		return nil, nil, NewArgError("target_type", "cannot be empty")
	}

	listSROpt := &listSROptions{TargetType: target_type}

	return s.list(ctx, nil, listSROpt)
}

// Create a new Santa rule
func (s *SantaRulesServiceOp) Create(ctx context.Context, createRequest *SantaRuleRequest) (*SantaRule, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, srBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	sr := new(SantaRule)
	resp, err := s.client.Do(ctx, req, sr)
	if err != nil {
		return nil, resp, err
	}

	return sr, resp, err
}

// Update a Santa rule
func (s *SantaRulesServiceOp) Update(ctx context.Context, srID int, updateRequest *SantaRuleRequest) (*SantaRule, *Response, error) {
	if srID < 1 {
		return nil, nil, NewArgError("srID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", srBasePath, srID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	sr := new(SantaRule)
	resp, err := s.client.Do(ctx, req, sr)
	if err != nil {
		return nil, resp, err
	}

	return sr, resp, err
}

// Delete a Santa rule
func (s *SantaRulesServiceOp) Delete(ctx context.Context, srID int) (*Response, error) {
	if srID < 1 {
		return nil, NewArgError("srID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", srBasePath, srID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Santa enrollments
func (s *SantaRulesServiceOp) list(ctx context.Context, opt *ListOptions, srOpt *listSROptions) ([]SantaRule, *Response, error) {
	path := srBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, srOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var srs []SantaRule
	resp, err := s.client.Do(ctx, req, &srs)
	if err != nil {
		return nil, resp, err
	}

	return srs, resp, err
}
