// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AppRequest30 app request 30
// swagger:model app_request_30
type AppRequest30 struct {

	// automated
	Automated bool `json:"automated"`

	// certificate
	Certificate string `json:"certificate,omitempty"`

	// command
	Command string `json:"command,omitempty"`

	// container count
	ContainerCount int64 `json:"container_count"`

	// container size
	ContainerSize int64 `json:"container_size,omitempty"`

	// destination account
	// Format: uri
	DestinationAccount strfmt.URI `json:"destination_account,omitempty"`

	// destination account id
	DestinationAccountID int64 `json:"destination_account_id,omitempty"`

	// destination region
	DestinationRegion string `json:"destination_region,omitempty"`

	// disk size
	DiskSize int64 `json:"disk_size,omitempty"`

	// docker ref
	DockerRef string `json:"docker_ref,omitempty"`

	// env
	Env interface{} `json:"env,omitempty"`

	// git ref
	GitRef string `json:"git_ref,omitempty"`

	// handle
	Handle string `json:"handle,omitempty"`

	// interactive
	Interactive bool `json:"interactive"`

	// key arn
	KeyArn string `json:"key_arn,omitempty"`

	// private key
	PrivateKey string `json:"private_key,omitempty"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this app request 30
func (m *AppRequest30) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestinationAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppRequest30) validateDestinationAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.DestinationAccount) { // not required
		return nil
	}

	if err := validate.FormatOf("destination_account", "body", "uri", m.DestinationAccount.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AppRequest30) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppRequest30) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppRequest30) UnmarshalBinary(b []byte) error {
	var res AppRequest30
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
