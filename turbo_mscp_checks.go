package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const tmscBasePath = "turbo/mscp_checks/"

// TurboMSCPChecksService is an interface for interfacing with the Turbo mSCP checks
// endpoints of the Zentral API
type TurboMSCPChecksService interface {
	List(context.Context, *ListOptions) ([]TurboMSCPCheck, *Response, error)
	GetByID(context.Context, string) (*TurboMSCPCheck, *Response, error)
	GetByRuleID(context.Context, string) ([]TurboMSCPCheck, *Response, error)
	Create(context.Context, *TurboMSCPCheckRequest) (*TurboMSCPCheck, *Response, error)
	Update(context.Context, string, *TurboMSCPCheckRequest) (*TurboMSCPCheck, *Response, error)
	Delete(context.Context, string) (*Response, error)
}

// TurboMSCPChecksServiceOp handles communication with the Turbo mSCP checks related
// methods of the Zentral API.
type TurboMSCPChecksServiceOp struct {
	client *Client
}

var _ TurboMSCPChecksService = &TurboMSCPChecksServiceOp{}

// TurboMSCPCheck represents a Zentral Turbo mSCP check
type TurboMSCPCheck struct {
	ID                string    `json:"id,omitempty"`
	RuleID            string    `json:"rule_id"`
	Baseline          string    `json:"baseline"`
	ODVInt            *int      `json:"odv_int"`
	ODVString         *string   `json:"odv_string"`
	ODVBool           *bool     `json:"odv_bool"`
	Version           int       `json:"version"`
	JobID             string    `json:"job_id"`
	ComplianceCheckID int       `json:"compliance_check_id"`
	Created           Timestamp `json:"created_at,omitempty"`
	Updated           Timestamp `json:"updated_at,omitempty"`
}

func (tmc TurboMSCPCheck) String() string {
	return Stringify(tmc)
}

// TurboMSCPCheckRequest represents a request to create or update a Turbo mSCP check
type TurboMSCPCheckRequest struct {
	RuleID    string  `json:"rule_id"`
	Baseline  string  `json:"baseline"`
	ODVInt    *int    `json:"odv_int"`
	ODVString *string `json:"odv_string"`
	ODVBool   *bool   `json:"odv_bool"`
}

type listTMSCPOptions struct {
	RuleID string `url:"rule_id,omitempty"`
}

// List lists all the Turbo mSCP checks.
func (s *TurboMSCPChecksServiceOp) List(ctx context.Context, opt *ListOptions) ([]TurboMSCPCheck, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Turbo mSCP check by id.
func (s *TurboMSCPChecksServiceOp) GetByID(ctx context.Context, tmcID string) (*TurboMSCPCheck, *Response, error) {
	if len(tmcID) < 1 {
		return nil, nil, NewArgError("tmcID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", tmscBasePath, tmcID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	tmc := new(TurboMSCPCheck)

	resp, err := s.client.Do(ctx, req, tmc)
	if err != nil {
		return nil, resp, err
	}

	return tmc, resp, err
}

// GetByRuleID retrieves the Turbo mSCP checks for a given rule id.
func (s *TurboMSCPChecksServiceOp) GetByRuleID(ctx context.Context, ruleID string) ([]TurboMSCPCheck, *Response, error) {
	if len(ruleID) < 1 {
		return nil, nil, NewArgError("ruleID", "cannot be blank")
	}

	listTMSCPOpt := &listTMSCPOptions{RuleID: ruleID}

	return s.list(ctx, nil, listTMSCPOpt)
}

// Create a new Turbo mSCP check.
func (s *TurboMSCPChecksServiceOp) Create(ctx context.Context, createRequest *TurboMSCPCheckRequest) (*TurboMSCPCheck, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, tmscBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	tmc := new(TurboMSCPCheck)
	resp, err := s.client.Do(ctx, req, tmc)
	if err != nil {
		return nil, resp, err
	}

	return tmc, resp, err
}

// Update a Turbo mSCP check.
func (s *TurboMSCPChecksServiceOp) Update(ctx context.Context, tmcID string, updateRequest *TurboMSCPCheckRequest) (*TurboMSCPCheck, *Response, error) {
	if len(tmcID) < 1 {
		return nil, nil, NewArgError("tmcID", "cannot be blank")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%s/", tmscBasePath, tmcID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	tmc := new(TurboMSCPCheck)
	resp, err := s.client.Do(ctx, req, tmc)
	if err != nil {
		return nil, resp, err
	}

	return tmc, resp, err
}

// Delete a Turbo mSCP check.
func (s *TurboMSCPChecksServiceOp) Delete(ctx context.Context, tmcID string) (*Response, error) {
	if len(tmcID) < 1 {
		return nil, NewArgError("tmcID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", tmscBasePath, tmcID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Turbo mSCP checks
func (s *TurboMSCPChecksServiceOp) list(ctx context.Context, opt *ListOptions, tmcOpt *listTMSCPOptions) ([]TurboMSCPCheck, *Response, error) {
	path := tmscBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, tmcOpt)
	if err != nil {
		return nil, nil, err
	}

	return resolveAllPages[TurboMSCPCheck](ctx, s.client, path)
}
