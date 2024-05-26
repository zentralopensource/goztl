package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const mmepBasePath = "monolith/manifest_enrollment_packages/"

// MonolithManifestEnrollmentPackagesService is an interface for interfacing with the Monolith manifest enrollment packages
// endpoints of the Zentral API
type MonolithManifestEnrollmentPackagesService interface {
	List(context.Context, *ListOptions) ([]MonolithManifestEnrollmentPackage, *Response, error)
	GetByID(context.Context, int) (*MonolithManifestEnrollmentPackage, *Response, error)
	GetByManifestID(context.Context, int) ([]MonolithManifestEnrollmentPackage, *Response, error)
	Create(context.Context, *MonolithManifestEnrollmentPackageRequest) (*MonolithManifestEnrollmentPackage, *Response, error)
	Update(context.Context, int, *MonolithManifestEnrollmentPackageRequest) (*MonolithManifestEnrollmentPackage, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// MonolithManifestEnrollmentPackagesServiceOp handles commepunication with the Monolith manifest enrollment packages related
// methods of the Zentral API.
type MonolithManifestEnrollmentPackagesServiceOp struct {
	client *Client
}

var _ MonolithManifestEnrollmentPackagesService = &MonolithManifestEnrollmentPackagesServiceOp{}

// MonolithManifestEnrollmentPackage represents a Zentral manifest enrollment package.
type MonolithManifestEnrollmentPackage struct {
	ID           int       `json:"id"`
	ManifestID   int       `json:"manifest"`
	Builder      string    `json:"builder"`
	EnrollmentID int       `json:"enrollment_pk"`
	Version      int       `json:"version"`
	TagIDs       []int     `json:"tags"`
	Created      Timestamp `json:"created_at"`
	Updated      Timestamp `json:"updated_at"`
}

func (se MonolithManifestEnrollmentPackage) String() string {
	return Stringify(se)
}

// MonolithManifestEnrollmentPackageRequest represents a request to create or update a Monolith manifest enrollment package.
type MonolithManifestEnrollmentPackageRequest struct {
	ManifestID   int    `json:"manifest"`
	Builder      string `json:"builder"`
	EnrollmentID int    `json:"enrollment_pk"`
	TagIDs       []int  `json:"tags"`
}

type listMMEPOptions struct {
	ManifestID int `url:"manifest_id,omitempty"`
}

// List lists all the Monolith manifest enrollment packages.
func (s *MonolithManifestEnrollmentPackagesServiceOp) List(ctx context.Context, opt *ListOptions) ([]MonolithManifestEnrollmentPackage, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a Monolith manifest enrollment package by id.
func (s *MonolithManifestEnrollmentPackagesServiceOp) GetByID(ctx context.Context, mmepID int) (*MonolithManifestEnrollmentPackage, *Response, error) {
	if mmepID < 1 {
		return nil, nil, NewArgError("mmepID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mmepBasePath, mmepID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	mmep := new(MonolithManifestEnrollmentPackage)

	resp, err := s.client.Do(ctx, req, mmep)
	if err != nil {
		return nil, resp, err
	}

	return mmep, resp, err
}

// GetByManifestID retrieves Monolith manifest enrollment packages by manifest ID.
func (s *MonolithManifestEnrollmentPackagesServiceOp) GetByManifestID(ctx context.Context, mmID int) ([]MonolithManifestEnrollmentPackage, *Response, error) {
	if mmID < 1 {
		return nil, nil, NewArgError("mmID", "cannot be < 1")
	}

	listMMEPOpt := &listMMEPOptions{ManifestID: mmID}

	return s.list(ctx, nil, listMMEPOpt)
}

// Create a new Monolith manifest enrollment package.
func (s *MonolithManifestEnrollmentPackagesServiceOp) Create(ctx context.Context, createRequest *MonolithManifestEnrollmentPackageRequest) (*MonolithManifestEnrollmentPackage, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, mmepBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	mmep := new(MonolithManifestEnrollmentPackage)
	resp, err := s.client.Do(ctx, req, mmep)
	if err != nil {
		return nil, resp, err
	}

	return mmep, resp, err
}

// Update a Monolith manifest enrollment package.
func (s *MonolithManifestEnrollmentPackagesServiceOp) Update(ctx context.Context, mmepID int, updateRequest *MonolithManifestEnrollmentPackageRequest) (*MonolithManifestEnrollmentPackage, *Response, error) {
	if mmepID < 1 {
		return nil, nil, NewArgError("mmepID", "cannot be less than 1")
	}

	if updateRequest == nil {
		return nil, nil, NewArgError("updateRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", mmepBasePath, mmepID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, nil, err
	}

	mmep := new(MonolithManifestEnrollmentPackage)
	resp, err := s.client.Do(ctx, req, mmep)
	if err != nil {
		return nil, resp, err
	}

	return mmep, resp, err
}

// Delete a Monolith manifest enrollment package.
func (s *MonolithManifestEnrollmentPackagesServiceOp) Delete(ctx context.Context, mmepID int) (*Response, error) {
	if mmepID < 1 {
		return nil, NewArgError("mmepID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mmepBasePath, mmepID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Helper method for listing Monolith manifest enrollment packages.
func (s *MonolithManifestEnrollmentPackagesServiceOp) list(ctx context.Context, opt *ListOptions, mmepOpt *listMMEPOptions) ([]MonolithManifestEnrollmentPackage, *Response, error) {
	path := mmepBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, mmepOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var mmeps []MonolithManifestEnrollmentPackage
	resp, err := s.client.Do(ctx, req, &mmeps)
	if err != nil {
		return nil, resp, err
	}

	return mmeps, resp, err
}
