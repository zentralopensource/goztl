package goztl

import (
	"context"
	"net/http"
)

const mlaBasePath = "mdm/location_assets/"

// MDMLocationAssetsService is an interface for interfacing with the MDM location assets
// endpoint of the Zentral API
type MDMLocationAssetsService interface {
	List(context.Context, *ListOptions) ([]MDMLocationAsset, *Response, error)
	Get(context.Context, int, string, string) (*MDMLocationAsset, *Response, error)
}

// MDMLocationAssetsServiceOp handles communication with the MDM location assets related
// methods of the Zentral API.
type MDMLocationAssetsServiceOp struct {
	client *Client
}

var _ MDMLocationAssetsService = &MDMLocationAssetsServiceOp{}

// MDMLocation represents a Zentral MDM location asset
type MDMLocationAsset struct {
	ID             int       `json:"id"`
	LocationID     int       `json:"location"`
	AssetID        int       `json:"asset"`
	AdamID         string    `json:"adam_id"`
	PricingParam   string    `json:"pricing_param"`
	AssignedCount  int       `json:"assigned_count"`
	AvailableCount int       `json:"available_count"`
	RetiredCount   int       `json:"retired_count"`
	TotalCount     int       `json:"total_count"`
	Created        Timestamp `json:"created_at"`
	Updated        Timestamp `json:"updated_at"`
}

func (mla MDMLocationAsset) String() string {
	return Stringify(mla)
}

type listMLAOptions struct {
	LocationID   int    `url:"location_id,omitempty"`
	AdamID       string `url:"adam_id,omitempty"`
	PricingParam string `url:"pricing_param"`
}

// List lists all the MDM location assets.
func (s *MDMLocationAssetsServiceOp) List(ctx context.Context, opt *ListOptions) ([]MDMLocationAsset, *Response, error) {
	return s.list(ctx, opt, nil)
}

// Get retrieves a MDM location asset by location ID, Adam ID, and pricing param
func (s *MDMLocationAssetsServiceOp) Get(ctx context.Context, lid int, aid string, pp string) (*MDMLocationAsset, *Response, error) {
	if lid < 1 {
		return nil, nil, NewArgError("lid", "cannot be less than 1")
	}

	if len(aid) < 1 {
		return nil, nil, NewArgError("aid", "cannot be empty")
	}

	if len(pp) < 1 {
		return nil, nil, NewArgError("pp", "cannot be empty")
	}

	listMLAOpt := &listMLAOptions{LocationID: lid, AdamID: aid, PricingParam: pp}

	las, resp, err := s.list(ctx, nil, listMLAOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(las) < 1 {
		return nil, resp, nil
	}

	return &las[0], resp, err
}

// Helper method for listing MDM locations
func (s *MDMLocationAssetsServiceOp) list(ctx context.Context, opt *ListOptions, mlaOpt *listMLAOptions) ([]MDMLocationAsset, *Response, error) {
	path := mlaBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, mlaOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var las []MDMLocationAsset
	resp, err := s.client.Do(ctx, req, &las)
	if err != nil {
		return nil, resp, err
	}

	return las, resp, err
}
