package goztl

import (
	"context"
	"fmt"
	"net/http"
)

const mpcBasePath = "mdm/push_certificates/"

// MDMPushCertificatesService is an interface for interfacing with the MDM push certificate
// endpoints of the Zentral API
type MDMPushCertificatesService interface {
	List(context.Context, *ListOptions) ([]MDMPushCertificate, *Response, error)
	GetByID(context.Context, int) (*MDMPushCertificate, *Response, error)
	GetByName(context.Context, string) (*MDMPushCertificate, *Response, error)
}

// MDMPushCertificatesServiceOp handles communication with the MDM push certificates related
// methods of the Zentral API.
type MDMPushCertificatesServiceOp struct {
	client *Client
}

var _ MDMPushCertificatesService = &MDMPushCertificatesServiceOp{}

// MDMPushCertificate represents a Zentral MDM push certificate
type MDMPushCertificate struct {
	ID              int        `json:"id"`
	ProvisioningUID *string    `json:"provisioning_uid"`
	Name            string     `json:"name"`
	Topic           *string    `json:"topic"`
	NotBefore       *Timestamp `json:"not_before"`
	NotAfter        *Timestamp `json:"not_after"`
	Certificate     *string    `json:"certificate"`
	Created         Timestamp  `json:"created_at"`
	Updated         Timestamp  `json:"updated_at"`
}

func (mpc MDMPushCertificate) String() string {
	return Stringify(mpc)
}

type listMPCOptions struct {
	Name string `url:"name,omitempty"`
}

// List lists all the MDM push certificates.
func (s *MDMPushCertificatesServiceOp) List(ctx context.Context, opt *ListOptions) ([]MDMPushCertificate, *Response, error) {
	return s.list(ctx, opt, nil)
}

// GetByID retrieves a MDM push certificate by id.
func (s *MDMPushCertificatesServiceOp) GetByID(ctx context.Context, mpcID int) (*MDMPushCertificate, *Response, error) {
	if mpcID < 1 {
		return nil, nil, NewArgError("mpcID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", mpcBasePath, mpcID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	pc := new(MDMPushCertificate)

	resp, err := s.client.Do(ctx, req, pc)
	if err != nil {
		return nil, resp, err
	}

	return pc, resp, err
}

// GetByName retrieves a MDM push certificate by name.
func (s *MDMPushCertificatesServiceOp) GetByName(ctx context.Context, name string) (*MDMPushCertificate, *Response, error) {
	if len(name) < 1 {
		return nil, nil, NewArgError("name", "cannot be blank")
	}

	listSCOpt := &listMPCOptions{Name: name}

	pcs, resp, err := s.list(ctx, nil, listSCOpt)
	if err != nil {
		return nil, resp, err
	}
	if len(pcs) < 1 {
		return nil, resp, nil
	}

	return &pcs[0], resp, err
}

// Helper method for listing MDM push certificates
func (s *MDMPushCertificatesServiceOp) list(ctx context.Context, opt *ListOptions, mpcOpt *listMPCOptions) ([]MDMPushCertificate, *Response, error) {
	path := mpcBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	path, err = addOptions(path, mpcOpt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var pcs []MDMPushCertificate
	resp, err := s.client.Do(ctx, req, &pcs)
	if err != nil {
		return nil, resp, err
	}

	return pcs, resp, err
}
