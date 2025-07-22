package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const probesActionsBasePath = "probes/actions/"

// ProbesActionsService is an interface for interfacing with the Monolith manifests
// endpoints of the Zentral API
type ProbesActionsService interface {
	List(context.Context, *ListOptions) ([]ProbeAction, *Response, error)
	GetByID(context.Context, string) (*ProbeAction, *Response, error)
	GetByName(context.Context, string) (*ProbeAction, *Response, error)
	Create(context.Context, *ProbeActionRequest) (*ProbeAction, *Response, error)
	Update(context.Context, string, *ProbeActionRequest) (*ProbeAction, *Response, error)
	Delete(context.Context, string) (*Response, error)
}

// ProbesActionsServiceOp handles communication with the probes actions related
// methods of the Zentral API.
type ProbesActionsServiceOp struct {
	client *Client
}

var _ ProbesActionsService = &ProbesActionsServiceOp{}

// ProbeAction represents a Zentral probe action

type ProbeActionHTTPPost struct {
	URL      string       `json:"url"`
	Username *string      `json:"username"`
	Password *string      `json:"password"`
	Headers  []HTTPHeader `json:"headers"`
}

type ProbeActionSlackIncomingWebhook struct {
	URL string `json:"url"`
}

type ProbeAction struct {
	ID                   string                           `json:"id"`
	Name                 string                           `json:"name"`
	Description          string                           `json:"description"`
	Backend              string                           `json:"backend"`
	HTTPPost             *ProbeActionHTTPPost             `json:"http_post_kwargs"`
	SlackIncomingWebhook *ProbeActionSlackIncomingWebhook `json:"slack_incoming_webhook_kwargs"`
	Created              Timestamp                        `json:"created_at"`
	Updated              Timestamp                        `json:"updated_at"`
}

func (pa ProbeAction) String() string {
	return Stringify(pa)
}

// ProbeActionRequest represents a request to create or update a probe action
type ProbeActionRequest struct {
	Name                 string                           `json:"name"`
	Description          string                           `json:"description"`
	Backend              string                           `json:"backend"`
	HTTPPost             *ProbeActionHTTPPost             `json:"http_post_kwargs"`
	SlackIncomingWebhook *ProbeActionSlackIncomingWebhook `json:"slack_incoming_webhook_kwargs"`
}

type listPAOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the probe actions
func (s *ProbesActionsServiceOp) List(ctx context.Context, opt *ListOptions) ([]ProbeAction, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a probe action by id
func (s *ProbesActionsServiceOp) GetByID(ctx context.Context, paID string) (*ProbeAction, *Response, error) {
	if len(paID) < 1 {
		return nil, nil, NewArgError("paID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", probesActionsBasePath, paID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	pa := new(ProbeAction)

	resp, err := s.client.Do(ctx, req, pa)
	if err != nil {
		return nil, resp, err
	}

	return pa, resp, err
}

// GetByName retrieves a probe action by name
func (s *ProbesActionsServiceOp) GetByName(ctx context.Context, name string) (*ProbeAction, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listPAOpt := &listPAOptions{Name: name}

	pas, resp, err := s.list(ctx, nil, listPAOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(pas) < 1 {
		return nil, resp, nil
	}

	return &pas[0], resp, err
}

// Create a new probe action
func (s *ProbesActionsServiceOp) Create(ctx context.Context, createRequest *ProbeActionRequest) (*ProbeAction, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, probesActionsBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	pa := new(ProbeAction)
	resp, err := s.client.Do(ctx, req, pa)
	if err != nil {
		return nil, resp, err
	}

	return pa, resp, err
}

// Update a probe action
func (s *ProbesActionsServiceOp) Update(ctx context.Context, paID string, updateRequest *ProbeActionRequest) (*ProbeAction, *Response, error) {
	if len(paID) < 1 {
		return nil, nil, NewArgError("paID", "cannot be blank")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%s/", probesActionsBasePath, paID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	pa := new(ProbeAction)
	resp, err := s.client.Do(ctx, req, pa)
	if err != nil {
		return nil, resp, err
	}

	return pa, resp, err
}

// Delete a probe action
func (s *ProbesActionsServiceOp) Delete(ctx context.Context, paID string) (*Response, error) {
	if len(paID) < 1 {
		return nil, NewArgError("paID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", probesActionsBasePath, paID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing probe actions
func (s *ProbesActionsServiceOp) list(ctx context.Context, opt *ListOptions, paOpt *listPAOptions) ([]ProbeAction, *Response, error) {
	path := probesActionsBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, paOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var pas []ProbeAction
	resp, err := s.client.Do(ctx, req, &pas)
	if err != nil {
		return nil, resp, err
	}

	return pas, resp, err
}
