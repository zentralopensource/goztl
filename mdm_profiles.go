package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const mpBasePath = "mdm/profiles/"

// MDMProfilesService is an interface for interfacing with the MDM profile
// endpoints of the Zentral API
type MDMProfilesService interface {
	List(context.Context, *ListOptions) ([]MDMProfile, *Response, error)
	GetByID(context.Context, string) (*MDMProfile, *Response, error)
	Create(context.Context, *MDMProfileRequest) (*MDMProfile, *Response, error)
	Update(context.Context, string, *MDMProfileRequest) (*MDMProfile, *Response, error)
	Delete(context.Context, string) (*Response, error)
}

// MDMProfilesServiceOp handles communication with the MDM profiles related
// methods of the Zentral API.
type MDMProfilesServiceOp struct {
	client *Client
}

var _ MDMProfilesService = &MDMProfilesServiceOp{}

// MDMProfile represents a Zentral MDM profile
type MDMProfile struct {
	ID               string     `json:"id"`
	ArtifactID       string     `json:"artifact"`
	Source           string     `json:"source"`
	IOS              bool       `json:"ios"`
	IOSMaxVersion    string     `json:"ios_max_version"`
	IOSMinVersion    string     `json:"ios_min_version"`
	IPadOS           bool       `json:"ipados"`
	IPadOSMaxVersion string     `json:"ipados_max_version"`
	IPadOSMinVersion string     `json:"ipados_min_version"`
	MacOS            bool       `json:"macos"`
	MacOSMaxVersion  string     `json:"macos_max_version"`
	MacOSMinVersion  string     `json:"macos_min_version"`
	TVOS             bool       `json:"tvos"`
	TVOSMaxVersion   string     `json:"tvos_max_version"`
	TVOSMinVersion   string     `json:"tvos_min_version"`
	DefaultShard     int        `json:"default_shard"`
	ShardModulo      int        `json:"shard_modulo"`
	ExcludedTagIDs   []int      `json:"excluded_tags"`
	TagShards        []TagShard `json:"tag_shards"`
	Version          int        `json:"version"`
	Created          Timestamp  `json:"created_at"`
	Updated          Timestamp  `json:"updated_at"`
}

func (mp MDMProfile) String() string {
	return Stringify(mp)
}

// MDMProfileRequest represents a request to create or update a MDM profile
type MDMProfileRequest struct {
	ArtifactID       string     `json:"artifact"`
	Source           string     `json:"source"`
	IOS              bool       `json:"ios"`
	IOSMaxVersion    string     `json:"ios_max_version"`
	IOSMinVersion    string     `json:"ios_min_version"`
	IPadOS           bool       `json:"ipados"`
	IPadOSMaxVersion string     `json:"ipados_max_version"`
	IPadOSMinVersion string     `json:"ipados_min_version"`
	MacOS            bool       `json:"macos"`
	MacOSMaxVersion  string     `json:"macos_max_version"`
	MacOSMinVersion  string     `json:"macos_min_version"`
	TVOS             bool       `json:"tvos"`
	TVOSMaxVersion   string     `json:"tvos_max_version"`
	TVOSMinVersion   string     `json:"tvos_min_version"`
	DefaultShard     int        `json:"default_shard"`
	ShardModulo      int        `json:"shard_modulo"`
	ExcludedTagIDs   []int      `json:"excluded_tags"`
	TagShards        []TagShard `json:"tag_shards"`
	Version          int        `json:"version"`
}

type listMPOptions struct{}

// List lists all the MDM profiles.
func (s *MDMProfilesServiceOp) List(ctx context.Context, opt *ListOptions) ([]MDMProfile, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a MDM profile by id.
func (s *MDMProfilesServiceOp) GetByID(ctx context.Context, mpID string) (*MDMProfile, *Response, error) {
	if len(mpID) < 1 {
		return nil, nil, NewArgError("mpID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", mpBasePath, mpID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	mp := new(MDMProfile)

	resp, err := s.client.Do(ctx, req, mp)
	if err != nil {
		return nil, resp, err
	}

	return mp, resp, err
}

// Create a new MDM profile.
func (s *MDMProfilesServiceOp) Create(ctx context.Context, createRequest *MDMProfileRequest) (*MDMProfile, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, mpBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	mp := new(MDMProfile)
	resp, err := s.client.Do(ctx, req, mp)
	if err != nil {
		return nil, resp, err
	}

	return mp, resp, err
}

// Update a MDM profile.
func (s *MDMProfilesServiceOp) Update(ctx context.Context, mpID string, updateRequest *MDMProfileRequest) (*MDMProfile, *Response, error) {
	if len(mpID) < 1 {
		return nil, nil, NewArgError("mpID", "cannot be blank")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%s/", mpBasePath, mpID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	mp := new(MDMProfile)
	resp, err := s.client.Do(ctx, req, mp)
	if err != nil {
		return nil, resp, err
	}

	return mp, resp, err
}

// Delete a MDM profile.
func (s *MDMProfilesServiceOp) Delete(ctx context.Context, mpID string) (*Response, error) {
	if len(mpID) < 1 {
		return nil, NewArgError("mpID", "cannot be blank")
	}

	path := fmt.Sprintf("%s%s/", mpBasePath, mpID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing MDM profiles
func (s *MDMProfilesServiceOp) list(ctx context.Context, opt *ListOptions, mpOpt *listMPOptions) ([]MDMProfile, *Response, error) {
	path := mpBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, mpOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var mps []MDMProfile
	resp, err := s.client.Do(ctx, req, &mps)
	if err != nil {
		return nil, resp, err
	}

	return mps, resp, err
}
