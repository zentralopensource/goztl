package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const trjBasePath = "turbo/recurring_jobs/"

// TurboRecurringJobsService is an interface for interfacing with the Turbo recurring jobs
// endpoints of the Zentral API
type TurboRecurringJobsService interface {
	List(context.Context, *ListOptions) ([]TurboRecurringJob, *Response, error)
	GetByID(context.Context, string) (*TurboRecurringJob, *Response, error)
	Create(context.Context, *TurboRecurringJobRequest) (*TurboRecurringJob, *Response, error)
	Update(context.Context, string, *TurboRecurringJobRequest) (*TurboRecurringJob, *Response, error)
	Delete(context.Context, string) (*Response, error)
}

// TurboRecurringJobsServiceOp handles communication with the Turbo recurring jobs related
// methods of the Zentral API.
type TurboRecurringJobsServiceOp struct {
	client *Client
}

var _ TurboRecurringJobsService = &TurboRecurringJobsServiceOp{}

// TurboRecurringJob represents a Zentral Turbo recurring job
type TurboRecurringJob struct {
	ID                    string    `json:"id,omitempty"`
	ConfigurationID       string    `json:"configuration"`
	JobID                 string    `json:"job"`
	Interval              *int      `json:"interval"`
	TagIDs                []int     `json:"tags"`
	ExcludedTagIDs        []int     `json:"excluded_tags"`
	SerialNumbers         []string  `json:"serial_numbers"`
	ExcludedSerialNumbers []string  `json:"excluded_serial_numbers"`
	Created               Timestamp `json:"created_at,omitempty"`
	Updated               Timestamp `json:"updated_at,omitempty"`
}

func (trj TurboRecurringJob) String() string {
	return Stringify(trj)
}

// TurboRecurringJobRequest represents a request to create or update a Turbo recurring job
type TurboRecurringJobRequest struct {
	ConfigurationID       string   `json:"configuration"`
	JobID                 string   `json:"job"`
	Interval              *int     `json:"interval"`
	TagIDs                []int    `json:"tags"`
	ExcludedTagIDs        []int    `json:"excluded_tags"`
	SerialNumbers         []string `json:"serial_numbers"`
	ExcludedSerialNumbers []string `json:"excluded_serial_numbers"`
}

// List lists all the Turbo recurring jobs.
func (s *TurboRecurringJobsServiceOp) List(ctx context.Context, opt *ListOptions) ([]TurboRecurringJob, *Response, error) {
	path := trjBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	return resolveAllPages[TurboRecurringJob](ctx, s.client, path)
}

// GetByID retrieves a Turbo recurring job by id.
func (s *TurboRecurringJobsServiceOp) GetByID(ctx context.Context, trjID string) (*TurboRecurringJob, *Response, error) {
	if len(trjID) < 1 {
		return nil, nil, NewArgError("trjID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", trjBasePath, trjID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	trj := new(TurboRecurringJob)

	resp, err := s.client.Do(ctx, req, trj)
	if err != nil {
		return nil, resp, err
	}

	return trj, resp, err
}

// Create a new Turbo recurring job.
func (s *TurboRecurringJobsServiceOp) Create(ctx context.Context, createRequest *TurboRecurringJobRequest) (*TurboRecurringJob, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, trjBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	trj := new(TurboRecurringJob)
	resp, err := s.client.Do(ctx, req, trj)
	if err != nil {
		return nil, resp, err
	}

	return trj, resp, err
}

// Update a Turbo recurring job.
func (s *TurboRecurringJobsServiceOp) Update(ctx context.Context, trjID string, updateRequest *TurboRecurringJobRequest) (*TurboRecurringJob, *Response, error) {
	if len(trjID) < 1 {
		return nil, nil, NewArgError("trjID", "cannot be blank")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%s/", trjBasePath, trjID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	trj := new(TurboRecurringJob)
	resp, err := s.client.Do(ctx, req, trj)
	if err != nil {
		return nil, resp, err
	}

	return trj, resp, err
}

// Delete a Turbo recurring job.
func (s *TurboRecurringJobsServiceOp) Delete(ctx context.Context, trjID string) (*Response, error) {
	if len(trjID) < 1 {
		return nil, NewArgError("trjID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", trjBasePath, trjID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}
