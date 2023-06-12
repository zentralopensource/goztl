package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const mbaBasePath = "mdm/blueprint_artifacts/"

// MDMBlueprintArtifactsService is an interface for interfacing with the MDM blueprint artifact
// endpoints of the Zentral API
type MDMBlueprintArtifactsService interface {
	List(context.Context, *ListOptions) ([]MDMBlueprintArtifact, *Response, error)
	GetByID(context.Context, int) (*MDMBlueprintArtifact, *Response, error)
	Create(context.Context, *MDMBlueprintArtifactRequest) (*MDMBlueprintArtifact, *Response, error)
	Update(context.Context, int, *MDMBlueprintArtifactRequest) (*MDMBlueprintArtifact, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// MDMBlueprintArtifactsServiceOp handles communication with the MDM blueprint artifacts related
// methods of the Zentral API.
type MDMBlueprintArtifactsServiceOp struct {
	client *Client
}

var _ MDMBlueprintArtifactsService = &MDMBlueprintArtifactsServiceOp{}

// MDMBlueprintArtifact represents a Zentral MDM blueprint artifact
type MDMBlueprintArtifact struct {
	ID               int        `json:"id"`
	BlueprintID      int        `json:"blueprint"`
	ArtifactID       string     `json:"artifact"`
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
	Created          Timestamp  `json:"created_at"`
	Updated          Timestamp  `json:"updated_at"`
}

func (mba MDMBlueprintArtifact) String() string {
	return Stringify(mba)
}

// MDMBlueprintArtifactRequest represents a request to create or update a MDM blueprint artifact
type MDMBlueprintArtifactRequest struct {
	BlueprintID      int        `json:"blueprint"`
	ArtifactID       string     `json:"artifact"`
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
}

type listMBAOptions struct{}

// List lists all the MDM blueprint artifacts.
func (s *MDMBlueprintArtifactsServiceOp) List(ctx context.Context, opt *ListOptions) ([]MDMBlueprintArtifact, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a MDM blueprint artifact by id.
func (s *MDMBlueprintArtifactsServiceOp) GetByID(ctx context.Context, mbaID int) (*MDMBlueprintArtifact, *Response, error) {
	if mbaID < 1 {
		return nil, nil, NewArgError("mbaID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mbaBasePath, mbaID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	mp := new(MDMBlueprintArtifact)

	resp, err := s.client.Do(ctx, req, mp)
	if err != nil {
		return nil, resp, err
	}

	return mp, resp, err
}

// Create a new MDM blueprint artifact.
func (s *MDMBlueprintArtifactsServiceOp) Create(ctx context.Context, createRequest *MDMBlueprintArtifactRequest) (*MDMBlueprintArtifact, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, mbaBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	mp := new(MDMBlueprintArtifact)
	resp, err := s.client.Do(ctx, req, mp)
	if err != nil {
		return nil, resp, err
	}

	return mp, resp, err
}

// Update a MDM blueprint artifact.
func (s *MDMBlueprintArtifactsServiceOp) Update(ctx context.Context, mbaID int, updateRequest *MDMBlueprintArtifactRequest) (*MDMBlueprintArtifact, *Response, error) {
	if mbaID < 1 {
		return nil, nil, NewArgError("mbaID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", mbaBasePath, mbaID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	mba := new(MDMBlueprintArtifact)
	resp, err := s.client.Do(ctx, req, mba)
	if err != nil {
		return nil, resp, err
	}

	return mba, resp, err
}

// Delete a MDM blueprint artifact.
func (s *MDMBlueprintArtifactsServiceOp) Delete(ctx context.Context, mbaID int) (*Response, error) {
	if mbaID < 1 {
		return nil, NewArgError("mbaID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mbaBasePath, mbaID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing MDM blueprint artifacts
func (s *MDMBlueprintArtifactsServiceOp) list(ctx context.Context, opt *ListOptions, mpOpt *listMBAOptions) ([]MDMBlueprintArtifact, *Response, error) {
	path := mbaBasePath
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

	var mbas []MDMBlueprintArtifact
	resp, err := s.client.Do(ctx, req, &mbas)
	if err != nil {
		return nil, resp, err
	}

	return mbas, resp, err
}
