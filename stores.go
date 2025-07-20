package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const storesBasePath = "stores/stores/"

// StoresService is an interface for interfacing with the stores
// endpoints of the Zentral API
type StoresService interface {
	List(context.Context, *ListOptions) ([]Store, *Response, error)
	GetByID(context.Context, string) (*Store, *Response, error)
	GetByName(context.Context, string) (*Store, *Response, error)
	Create(context.Context, *StoreRequest) (*Store, *Response, error)
	Update(context.Context, string, *StoreRequest) (*Store, *Response, error)
	Delete(context.Context, string) (*Response, error)
}

// StoresServiceOp handles communication with the stores related
// methods of the Zentral API.
type StoresServiceOp struct {
	client *Client
}

var _ StoresService = &StoresServiceOp{}

// Store represents a Zentral store

type StoreHTTP struct {
	EndpointURL    string       `json:"endpoint_url"`
	VerifyTLS      bool         `json:"verify_tls"`
	Username       *string      `json:"username"`
	Password       *string      `json:"password"`
	Headers        []HTTPHeader `json:"headers"`
	Concurrency    int          `json:"concurrency"`
	RequestTimeout int          `json:"request_timeout"`
	MaxRetries     int          `json:"max_retries"`
}

type StoreSplunk struct {
	// HEC
	HECURL                    string       `json:"hec_url"`
	HECToken                  string       `json:"hec_token"`
	HECExtraHeaders           []HTTPHeader `json:"hec_extra_headers"`
	HECRequestTimeout         int          `json:"hec_request_timeout"`
	HECIndex                  *string      `json:"hec_index"`
	HECSource                 *string      `json:"hec_source"`
	ComputerNameAsHostSources []string     `json:"computer_name_as_host_sources"`
	CustomHostField           *string      `json:"custom_host_field"`
	SerialNumberField         string       `json:"serial_number_field"`
	BatchSize                 int          `json:"batch_size"`
	// Events URLs
	SearchAppURL *string `json:"search_app_url"`
	// Events search
	SearchURL            *string      `json:"search_url"`
	SearchToken          *string      `json:"search_token"`
	SearchExtraHeaders   []HTTPHeader `json:"search_extra_headers"`
	SearchRequestTimeout int          `json:"search_request_timeout"`
	SearchIndex          *string      `json:"search_index"`
	SearchSource         *string      `json:"search_source"`
	// Common
	VerifyTLS bool `json:"verify_tls"`
}

type Store struct {
	ID                         string          `json:"id"`
	Name                       string          `json:"name"`
	Description                string          `json:"description"`
	AdminConsole               bool            `json:"admin_console"`
	EventsURLAuthorizedRoleIDs []int           `json:"events_url_authorized_roles"`
	EventFilters               *EventFilterSet `json:"event_filters"`
	Backend                    string          `json:"backend"`
	HTTP                       *StoreHTTP      `json:"http_kwargs"`
	Splunk                     *StoreSplunk    `json:"splunk_kwargs"`
	Created                    Timestamp       `json:"created_at"`
	Updated                    Timestamp       `json:"updated_at"`
}

func (s Store) String() string {
	return Stringify(s)
}

// StoreRequest represents a request to create or update a store
type StoreRequest struct {
	Name                       string          `json:"name"`
	Description                string          `json:"description"`
	AdminConsole               bool            `json:"admin_console"`
	EventsURLAuthorizedRoleIDs []int           `json:"events_url_authorized_roles"`
	EventFilters               *EventFilterSet `json:"event_filters"`
	Backend                    string          `json:"backend"`
	HTTP                       *StoreHTTP      `json:"http_kwargs"`
	Splunk                     *StoreSplunk    `json:"splunk_kwargs"`
}

type listSOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the stores
func (s *StoresServiceOp) List(ctx context.Context, opt *ListOptions) ([]Store, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a store by id
func (s *StoresServiceOp) GetByID(ctx context.Context, sID string) (*Store, *Response, error) {
	if len(sID) < 1 {
		return nil, nil, NewArgError("sID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", storesBasePath, sID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	store := new(Store)

	resp, err := s.client.Do(ctx, req, store)
	if err != nil {
		return nil, resp, err
	}

	return store, resp, err
}

// GetByName retrieves a store by name
func (s *StoresServiceOp) GetByName(ctx context.Context, name string) (*Store, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listSOpt := &listSOptions{Name: name}

	stores, resp, err := s.list(ctx, nil, listSOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(stores) < 1 {
		return nil, resp, nil
	}

	return &stores[0], resp, err
}

// Create a new store
func (s *StoresServiceOp) Create(ctx context.Context, createRequest *StoreRequest) (*Store, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, storesBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	store := new(Store)
	resp, err := s.client.Do(ctx, req, store)
	if err != nil {
		return nil, resp, err
	}

	return store, resp, err
}

// Update a store
func (s *StoresServiceOp) Update(ctx context.Context, sID string, updateRequest *StoreRequest) (*Store, *Response, error) {
	if len(sID) < 1 {
		return nil, nil, NewArgError("sID", "cannot be blank")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%s/", storesBasePath, sID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	store := new(Store)
	resp, err := s.client.Do(ctx, req, store)
	if err != nil {
		return nil, resp, err
	}

	return store, resp, err
}

// Delete a store
func (s *StoresServiceOp) Delete(ctx context.Context, sID string) (*Response, error) {
	if len(sID) < 1 {
		return nil, NewArgError("sID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", storesBasePath, sID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing stores
func (s *StoresServiceOp) list(ctx context.Context, opt *ListOptions, sOpt *listSOptions) ([]Store, *Response, error) {
	path := storesBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, sOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var stores []Store
	resp, err := s.client.Do(ctx, req, &stores)
	if err != nil {
		return nil, resp, err
	}

	return stores, resp, err
}
