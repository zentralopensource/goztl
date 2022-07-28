package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const jmespathCheckBasePath = "inventory/jmespath_checks/"

// JMESPathChecksService is an interface for interfacing with the JMESPath checks
// endpoints of the Zentral API.
type JMESPathChecksService interface {
	List(context.Context, *ListOptions) ([]JMESPathCheck, *Response, error)
	GetByID(context.Context, int) (*JMESPathCheck, *Response, error)
	GetByName(context.Context, string) (*JMESPathCheck, *Response, error)
	Create(context.Context, *JMESPathCheckCreateRequest) (*JMESPathCheck, *Response, error)
	Update(context.Context, int, *JMESPathCheckUpdateRequest) (*JMESPathCheck, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// JMESPathChecksServiceOp handles communication with the jmespath_checks related
// methods of the Zentral API.
type JMESPathChecksServiceOp struct {
	client *Client
}

var _ JMESPathChecksService = &JMESPathChecksServiceOp{}

// JMESPathCheck represents a Zentral JMESPath check.
type JMESPathCheck struct {
	ID                 int       `json:"id"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	SourceName         string    `json:"source_name"`
	Platforms          []string  `json:"platforms"`
	TagIDs             []int     `json:"tags"`
	JMESPathExpression string    `json:"jmespath_expression"`
	Version            int       `json:"version"`
	Created            Timestamp `json:"created_at"`
	Updated            Timestamp `json:"updated_at"`
}

// JMESPathCheckCreateRequest represents a request to create a JMESPath check.
type JMESPathCheckCreateRequest struct {
	Name               string   `json:"name"`
	Description        string   `json:"description"`
	SourceName         string   `json:"source_name"`
	Platforms          []string `json:"platforms"`
	TagIDs             []int    `json:"tags"`
	JMESPathExpression string   `json:"jmespath_expression"`
}

// JMESPathCheckUpdateRequest represents a request to update a JMESPath check.
type JMESPathCheckUpdateRequest struct {
	Name               string   `json:"name"`
	Description        string   `json:"description"`
	SourceName         string   `json:"source_name"`
	Platforms          []string `json:"platforms"`
	TagIDs             []int    `json:"tags"`
	JMESPathExpression string   `json:"jmespath_expression"`
}

func (jmespath_check JMESPathCheck) String() string {
	return Stringify(jmespath_check)
}

type listJMESPathCheckOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the jmespath_checks.
func (s *JMESPathChecksServiceOp) List(ctx context.Context, opt *ListOptions) ([]JMESPathCheck, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a jmespath_check by id.
func (s *JMESPathChecksServiceOp) GetByID(ctx context.Context, jmespathCheckID int) (*JMESPathCheck, *Response, error) {
	if jmespathCheckID < 1 {
		return nil, nil, NewArgError("jmespathCheckID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", jmespathCheckBasePath, jmespathCheckID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	jmespath_check := new(JMESPathCheck)

	resp, err := s.client.Do(ctx, req, jmespath_check)
	if err != nil {
		return nil, resp, err
	}

	return jmespath_check, resp, err
}

// GetByName retrieves a jmespath_check by name.
func (s *JMESPathChecksServiceOp) GetByName(ctx context.Context, name string) (*JMESPathCheck, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listJMESPathCheckOpt := &listJMESPathCheckOptions{Name: name}

	jmespath_checks, resp, err := s.list(ctx, nil, listJMESPathCheckOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(jmespath_checks) < 1 {
		return nil, resp, nil
	}

	return &jmespath_checks[0], resp, err
}

// Create a new jmespath_check.
func (s *JMESPathChecksServiceOp) Create(ctx context.Context, createRequest *JMESPathCheckCreateRequest) (*JMESPathCheck, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, jmespathCheckBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	jmespath_check := new(JMESPathCheck)
	resp, err := s.client.Do(ctx, req, jmespath_check)
	if err != nil {
		return nil, resp, err
	}

	return jmespath_check, resp, err
}

// Update a jmespath_check.
func (s *JMESPathChecksServiceOp) Update(ctx context.Context, jmespathCheckID int, updateRequest *JMESPathCheckUpdateRequest) (*JMESPathCheck, *Response, error) {
	if jmespathCheckID < 1 {
		return nil, nil, NewArgError("jmespathCheckID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", jmespathCheckBasePath, jmespathCheckID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	jmespath_check := new(JMESPathCheck)
	resp, err := s.client.Do(ctx, req, jmespath_check)
	if err != nil {
		return nil, resp, err
	}

	return jmespath_check, resp, err
}

// Delete a jmespath_check.
func (s *JMESPathChecksServiceOp) Delete(ctx context.Context, jmespathCheckID int) (*Response, error) {
	if jmespathCheckID < 1 {
		return nil, NewArgError("jmespathCheckID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", jmespathCheckBasePath, jmespathCheckID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing jmespath_checks
func (s *JMESPathChecksServiceOp) list(ctx context.Context, opt *ListOptions, jmespathCheckOpt *listJMESPathCheckOptions) ([]JMESPathCheck, *Response, error) {
	path := jmespathCheckBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, jmespathCheckOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var jmespath_checks []JMESPathCheck
	resp, err := s.client.Do(ctx, req, &jmespath_checks)
	if err != nil {
		return nil, resp, err
	}

	return jmespath_checks, resp, err
}
