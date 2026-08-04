package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const totjBasePath = "turbo/one_time_jobs/"

// TurboOneTimeJobsService is an interface for interfacing with the Turbo one-time jobs
// endpoints of the Zentral API
type TurboOneTimeJobsService interface {
	List(context.Context, *ListOptions) ([]TurboOneTimeJob, *Response, error)
	GetByID(context.Context, string) (*TurboOneTimeJob, *Response, error)
	Create(context.Context, *TurboOneTimeJobRequest) (*TurboOneTimeJob, *Response, error)
	Update(context.Context, string, *TurboOneTimeJobRequest) (*TurboOneTimeJob, *Response, error)
	Delete(context.Context, string) (*Response, error)
}

// TurboOneTimeJobsServiceOp handles communication with the Turbo one-time jobs related
// methods of the Zentral API.
type TurboOneTimeJobsServiceOp struct {
	client *Client
}

var _ TurboOneTimeJobsService = &TurboOneTimeJobsServiceOp{}

// TurboOneTimeJob represents a Zentral Turbo one-time job
type TurboOneTimeJob struct {
	ID                    string    `json:"id,omitempty"`
	ConfigurationID       string    `json:"configuration"`
	JobID                 string    `json:"job"`
	NotBefore             *string   `json:"not_before"`
	NotAfter              *string   `json:"not_after"`
	TagIDs                []int     `json:"tags"`
	ExcludedTagIDs        []int     `json:"excluded_tags"`
	SerialNumbers         []string  `json:"serial_numbers"`
	ExcludedSerialNumbers []string  `json:"excluded_serial_numbers"`
	Created               Timestamp `json:"created_at,omitempty"`
	Updated               Timestamp `json:"updated_at,omitempty"`
}

func (totj TurboOneTimeJob) String() string {
	return Stringify(totj)
}

// TurboOneTimeJobRequest represents a request to create or update a Turbo one-time job
type TurboOneTimeJobRequest struct {
	ConfigurationID       string   `json:"configuration"`
	JobID                 string   `json:"job"`
	NotBefore             *string  `json:"not_before"`
	NotAfter              *string  `json:"not_after"`
	TagIDs                []int    `json:"tags"`
	ExcludedTagIDs        []int    `json:"excluded_tags"`
	SerialNumbers         []string `json:"serial_numbers"`
	ExcludedSerialNumbers []string `json:"excluded_serial_numbers"`
}

// List lists all the Turbo one-time jobs.
func (s *TurboOneTimeJobsServiceOp) List(ctx context.Context, opt *ListOptions) ([]TurboOneTimeJob, *Response, error) {
	path := totjBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	return resolveAllPages[TurboOneTimeJob](ctx, s.client, path)
}

// GetByID retrieves a Turbo one-time job by id.
func (s *TurboOneTimeJobsServiceOp) GetByID(ctx context.Context, totjID string) (*TurboOneTimeJob, *Response, error) {
	if len(totjID) < 1 {
		return nil, nil, NewArgError("totjID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", totjBasePath, totjID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	totj := new(TurboOneTimeJob)

	resp, err := s.client.Do(ctx, req, totj)
	if err != nil {
		return nil, resp, err
	}

	return totj, resp, err
}

// Create a new Turbo one-time job.
func (s *TurboOneTimeJobsServiceOp) Create(ctx context.Context, createRequest *TurboOneTimeJobRequest) (*TurboOneTimeJob, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, totjBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	totj := new(TurboOneTimeJob)
	resp, err := s.client.Do(ctx, req, totj)
	if err != nil {
		return nil, resp, err
	}

	return totj, resp, err
}

// Update a Turbo one-time job.
func (s *TurboOneTimeJobsServiceOp) Update(ctx context.Context, totjID string, updateRequest *TurboOneTimeJobRequest) (*TurboOneTimeJob, *Response, error) {
	if len(totjID) < 1 {
		return nil, nil, NewArgError("totjID", "cannot be blank")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%s/", totjBasePath, totjID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	totj := new(TurboOneTimeJob)
	resp, err := s.client.Do(ctx, req, totj)
	if err != nil {
		return nil, resp, err
	}

	return totj, resp, err
}

// Delete a Turbo one-time job.
func (s *TurboOneTimeJobsServiceOp) Delete(ctx context.Context, totjID string) (*Response, error) {
	if len(totjID) < 1 {
		return nil, NewArgError("totjID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", totjBasePath, totjID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}
