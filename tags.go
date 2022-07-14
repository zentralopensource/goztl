package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const tagBasePath = "inventory/tags/"

// TagsService is an interface for interfacing with the tags
// endpoints of the Zentral API
type TagsService interface {
	List(context.Context, *ListOptions) ([]Tag, *Response, error)
	GetByID(context.Context, int) (*Tag, *Response, error)
	GetByName(context.Context, string) (*Tag, *Response, error)
	Create(context.Context, *TagCreateRequest) (*Tag, *Response, error)
	Update(context.Context, int, *TagUpdateRequest) (*Tag, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// TagsServiceOp handles communication with the tags related
// methods of the Zentral API.
type TagsServiceOp struct {
	client *Client
}

var _ TagsService = &TagsServiceOp{}

// Tag represents a Zentral Tag
type Tag struct {
	ID                 int    `json:"id,omitempty"`
	TaxonomyID         int    `json:"taxonomy,omitempty"`
	MetaBusinessUnitID int    `json:"meta_business_unit,omitempty"`
	Name               string `json:"name,omitempty"`
	Slug               string `json:"slug,omitempty"`
	Color              string `json:"color,omitempty"`
}

// TagCreateRequest represents a request to create a tag.
type TagCreateRequest struct {
	Name               string `json:"name"`
	TaxonomyID         int    `json:"taxonomy,omitempty"`
	MetaBusinessUnitID int    `json:"meta_business_unit,omitempty"`
	Color              string `json:"color,omitempty"`
}

// TagUpdateRequest represents a request to create a tag.
type TagUpdateRequest struct {
	Name               string `json:"name"`
	TaxonomyID         int    `json:"taxonomy,omitempty"`
	MetaBusinessUnitID int    `json:"meta_business_unit,omitempty"`
	Color              string `json:"color,omitempty"`
}

func (tag Tag) String() string {
	return Stringify(tag)
}

type listTagOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the tags.
func (s *TagsServiceOp) List(ctx context.Context, opt *ListOptions) ([]Tag, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a tag by id.
func (s *TagsServiceOp) GetByID(ctx context.Context, tagID int) (*Tag, *Response, error) {
	if tagID < 1 {
		return nil, nil, NewArgError("tagID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", tagBasePath, tagID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	tag := new(Tag)

	resp, err := s.client.Do(ctx, req, tag)
	if err != nil {
		return nil, resp, err
	}

	return tag, resp, err
}

// GetByName retrieves a tag by name.
func (s *TagsServiceOp) GetByName(ctx context.Context, name string) (*Tag, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listTagOpt := &listTagOptions{Name: name}

	tags, resp, err := s.list(ctx, nil, listTagOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(tags) < 1 {
		return nil, resp, nil
	}

	return &tags[0], resp, err
}

// Create a new tag.
func (s *TagsServiceOp) Create(ctx context.Context, createRequest *TagCreateRequest) (*Tag, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, tagBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	tag := new(Tag)
	resp, err := s.client.Do(ctx, req, tag)
	if err != nil {
		return nil, resp, err
	}

	return tag, resp, err
}

// Update a tag.
func (s *TagsServiceOp) Update(ctx context.Context, tagID int, updateRequest *TagUpdateRequest) (*Tag, *Response, error) {
	if tagID < 1 {
		return nil, nil, NewArgError("tagID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", tagBasePath, tagID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	tag := new(Tag)
	resp, err := s.client.Do(ctx, req, tag)
	if err != nil {
		return nil, resp, err
	}

	return tag, resp, err
}

// Delete a tag.
func (s *TagsServiceOp) Delete(ctx context.Context, tagID int) (*Response, error) {
	if tagID < 1 {
		return nil, NewArgError("tagID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", tagBasePath, tagID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing tags
func (s *TagsServiceOp) list(ctx context.Context, opt *ListOptions, tagOpt *listTagOptions) ([]Tag, *Response, error) {
	path := tagBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, tagOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var tags []Tag
	resp, err := s.client.Do(ctx, req, &tags)
	if err != nil {
		return nil, resp, err
	}

	return tags, resp, err
}
