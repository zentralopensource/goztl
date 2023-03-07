package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const smpiBasePath = "monolith/sub_manifest_pkg_infos/"

// MonolithSubManifestPkgInfosService is an interface for interfacing with the Monolith sub manifest pkg infos
// endpoints of the Zentral API
type MonolithSubManifestPkgInfosService interface {
	List(context.Context, *ListOptions) ([]MonolithSubManifestPkgInfo, *Response, error)
	GetByID(context.Context, int) (*MonolithSubManifestPkgInfo, *Response, error)
	GetBySubManifestID(context.Context, int) ([]MonolithSubManifestPkgInfo, *Response, error)
	Create(context.Context, *MonolithSubManifestPkgInfoRequest) (*MonolithSubManifestPkgInfo, *Response, error)
	Update(context.Context, int, *MonolithSubManifestPkgInfoRequest) (*MonolithSubManifestPkgInfo, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// MonolithSubManifestPkgInfosServiceOp handles cosmpiunication with the Monolith sub manifest pkg infos related
// methods of the Zentral API.
type MonolithSubManifestPkgInfosServiceOp struct {
	client *Client
}

var _ MonolithSubManifestPkgInfosService = &MonolithSubManifestPkgInfosServiceOp{}

// TagShard represents a Zentral tag shard.
type TagShard struct {
	TagID int `json:"tag"`
	Shard int `json:"shard"`
}

// MonolithSubManifestPkgInfo represents a Zentral MonolithSubManifestPkgInfo
type MonolithSubManifestPkgInfo struct {
	ID             int        `json:"id"`
	SubManifestID  int        `json:"sub_manifest"`
	Key            string     `json:"key"`
	PkgInfoName    string     `json:"pkg_info_name"`
	FeaturedItem   bool       `json:"featured_item"`
	ConditionID    *int       `json:"condition"`
	ShardModulo    int        `json:"shard_modulo"`
	DefaultShard   int        `json:"default_shard"`
	ExcludedTagIDs []int      `json:"excluded_tags"`
	TagShards      []TagShard `json:"tag_shards"`
	Created        Timestamp  `json:"created_at,omitempty"`
	Updated        Timestamp  `json:"updated_at,omitempty"`
}

func (se MonolithSubManifestPkgInfo) String() string {
	return Stringify(se)
}

// MonolithSubManifestPkgInfoRequest represents a request to create or update a Monolith sub manifest pkg info
type MonolithSubManifestPkgInfoRequest struct {
	SubManifestID  int        `json:"sub_manifest"`
	Key            string     `json:"key"`
	PkgInfoName    string     `json:"pkg_info_name"`
	FeaturedItem   bool       `json:"featured_item"`
	ConditionID    *int       `json:"condition"`
	ShardModulo    int        `json:"shard_modulo"`
	DefaultShard   int        `json:"default_shard"`
	ExcludedTagIDs []int      `json:"excluded_tags"`
	TagShards      []TagShard `json:"tag_shards"`
}

type listSMPIOptions struct {
	SubManifestID int `url:"sub_manifest_id,omitempty"`
}

// List lists all the Monolith sub manifest pkg infos.
func (s *MonolithSubManifestPkgInfosServiceOp) List(ctx context.Context, opt *ListOptions) ([]MonolithSubManifestPkgInfo, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Monolith sub manifest pkg info by id.
func (s *MonolithSubManifestPkgInfosServiceOp) GetByID(ctx context.Context, smpiID int) (*MonolithSubManifestPkgInfo, *Response, error) {
	if smpiID < 1 {
		return nil, nil, NewArgError("smpiID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", smpiBasePath, smpiID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	smpi := new(MonolithSubManifestPkgInfo)

	resp, err := s.client.Do(ctx, req, smpi)
	if err != nil {
		return nil, resp, err
	}

	return smpi, resp, err
}

// GetBySubManifestID retrieves Monolith sub manifest pkg infos by sub manifest ID.
func (s *MonolithSubManifestPkgInfosServiceOp) GetBySubManifestID(ctx context.Context, smID int) ([]MonolithSubManifestPkgInfo, *Response, error) {
	if smID < 1 {
		return nil, nil, NewArgError("smID", "cannot be < 1")
	}

	listSMPIOpt := &listSMPIOptions{SubManifestID: smID}

	return s.list(ctx, nil, listSMPIOpt)
}

// Create a new Monolith sub manifest pkg info.
func (s *MonolithSubManifestPkgInfosServiceOp) Create(ctx context.Context, createRequest *MonolithSubManifestPkgInfoRequest) (*MonolithSubManifestPkgInfo, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, smpiBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	smpi := new(MonolithSubManifestPkgInfo)
	resp, err := s.client.Do(ctx, req, smpi)
	if err != nil {
		return nil, resp, err
	}

	return smpi, resp, err
}

// Update a Monolith sub manifest pkg info.
func (s *MonolithSubManifestPkgInfosServiceOp) Update(ctx context.Context, smpiID int, updateRequest *MonolithSubManifestPkgInfoRequest) (*MonolithSubManifestPkgInfo, *Response, error) {
	if smpiID < 1 {
		return nil, nil, NewArgError("smpiID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", smpiBasePath, smpiID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	smpi := new(MonolithSubManifestPkgInfo)
	resp, err := s.client.Do(ctx, req, smpi)
	if err != nil {
		return nil, resp, err
	}

	return smpi, resp, err
}

// Delete a Monolith sub manifest pkg info.
func (s *MonolithSubManifestPkgInfosServiceOp) Delete(ctx context.Context, smpiID int) (*Response, error) {
	if smpiID < 1 {
		return nil, NewArgError("smpiID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", smpiBasePath, smpiID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Monolith sub manifest pkg infos
func (s *MonolithSubManifestPkgInfosServiceOp) list(ctx context.Context, opt *ListOptions, smpiOpt *listSMPIOptions) ([]MonolithSubManifestPkgInfo, *Response, error) {
	path := smpiBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, smpiOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var smpis []MonolithSubManifestPkgInfo
	resp, err := s.client.Do(ctx, req, &smpis)
	if err != nil {
		return nil, resp, err
	}

	return smpis, resp, err
}
