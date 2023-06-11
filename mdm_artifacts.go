package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const maBasePath = "mdm/artifacts/"

// MDMArtifactsService is an interface for interfacing with the MDM artifact
// endpoints of the Zentral API
type MDMArtifactsService interface {
	List(context.Context, *ListOptions) ([]MDMArtifact, *Response, error)
	GetByID(context.Context, string) (*MDMArtifact, *Response, error)
	GetByName(context.Context, string) (*MDMArtifact, *Response, error)
	Create(context.Context, *MDMArtifactRequest) (*MDMArtifact, *Response, error)
	Update(context.Context, string, *MDMArtifactRequest) (*MDMArtifact, *Response, error)
	Delete(context.Context, string) (*Response, error)
}

// MDMArtifactsServiceOp handles communication with the MDM artifacts related
// methods of the Zentral API.
type MDMArtifactsServiceOp struct {
	client *Client
}

var _ MDMArtifactsService = &MDMArtifactsServiceOp{}

// MDMArtifact represents a Zentral MDM artifact
type MDMArtifact struct {
	ID                          string    `json:"id"`
	Name                        string    `json:"name"`
	Type                        string    `json:"type"`
	Channel                     string    `json:"channel"`
	Platforms                   []string  `json:"platforms"`
	InstallDuringSetupAssistant bool      `json:"install_during_setup_assistant"`
	AutoUpdate                  bool      `json:"auto_update"`
	ReinstallInterval           int       `json:"reinstall_interval"`
	ReinstallOnOSUpdate         string    `json:"reinstall_on_os_update"`
	Requires                    []string  `json:"requires"`
	Created                     Timestamp `json:"created_at,omitempty"`
	Updated                     Timestamp `json:"updated_at,omitempty"`
}

func (ma MDMArtifact) String() string {
	return Stringify(ma)
}

// MDMArtifactRequest represents a request to create or update a MDM artifact
type MDMArtifactRequest struct {
	Name                        string   `json:"name"`
	Type                        string   `json:"type"`
	Channel                     string   `json:"channel"`
	Platforms                   []string `json:"platforms"`
	InstallDuringSetupAssistant bool     `json:"install_during_setup_assistant"`
	AutoUpdate                  bool     `json:"auto_update"`
	ReinstallInterval           int      `json:"reinstall_interval"`
	ReinstallOnOSUpdate         string   `json:"reinstall_on_os_update"`
	Requires                    []string `json:"requires"`
}

type listMAOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the MDM artifacts.
func (s *MDMArtifactsServiceOp) List(ctx context.Context, opt *ListOptions) ([]MDMArtifact, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a MDM artifact by id.
func (s *MDMArtifactsServiceOp) GetByID(ctx context.Context, maID string) (*MDMArtifact, *Response, error) {
	if len(maID) < 1 {
		return nil, nil, NewArgError("maID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", maBasePath, maID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	ma := new(MDMArtifact)

	resp, err := s.client.Do(ctx, req, ma)
	if err != nil {
		return nil, resp, err
	}

	return ma, resp, err
}

// GetByName retrieves a MDM artifact by name.
func (s *MDMArtifactsServiceOp) GetByName(ctx context.Context, name string) (*MDMArtifact, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listMAOpt := &listMAOptions{Name: name}

	mas, resp, err := s.list(ctx, nil, listMAOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(mas) < 1 {
		return nil, resp, nil
	}

	return &mas[0], resp, err
}

// Create a new MDM artifact.
func (s *MDMArtifactsServiceOp) Create(ctx context.Context, createRequest *MDMArtifactRequest) (*MDMArtifact, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, maBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	ma := new(MDMArtifact)
	resp, err := s.client.Do(ctx, req, ma)
	if err != nil {
		return nil, resp, err
	}

	return ma, resp, err
}

// Update a MDM artifact.
func (s *MDMArtifactsServiceOp) Update(ctx context.Context, maID string, updateRequest *MDMArtifactRequest) (*MDMArtifact, *Response, error) {
	if len(maID) < 1 {
		return nil, nil, NewArgError("maID", "cannot be blank")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%s/", maBasePath, maID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	ma := new(MDMArtifact)
	resp, err := s.client.Do(ctx, req, ma)
	if err != nil {
		return nil, resp, err
	}

	return ma, resp, err
}

// Delete a MDM artifact.
func (s *MDMArtifactsServiceOp) Delete(ctx context.Context, maID string) (*Response, error) {
	if len(maID) < 1 {
		return nil, NewArgError("maID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", maBasePath, maID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing MDM artifacts
func (s *MDMArtifactsServiceOp) list(ctx context.Context, opt *ListOptions, maOpt *listMAOptions) ([]MDMArtifact, *Response, error) {
	path := maBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, maOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var mas []MDMArtifact
	resp, err := s.client.Do(ctx, req, &mas)
	if err != nil {
		return nil, resp, err
	}

	return mas, resp, err
}
