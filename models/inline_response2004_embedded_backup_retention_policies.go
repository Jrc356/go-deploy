// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse2004EmbeddedBackupRetentionPolicies inline response 200 4 embedded backup retention policies
// swagger:model inline_response_200_4__embedded_backup_retention_policies
type InlineResponse2004EmbeddedBackupRetentionPolicies struct {

	// resource type
	ResourceType string `json:"_type,omitempty"`

	// links
	Links *InlineResponse2004EmbeddedLinks `json:"_links,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// daily
	Daily int64 `json:"daily,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// keep final
	KeepFinal bool `json:"keep_final,omitempty"`

	// make copy
	MakeCopy bool `json:"make_copy,omitempty"`

	// monthly
	Monthly int64 `json:"monthly,omitempty"`
}

// Validate validates this inline response 200 4 embedded backup retention policies
func (m *InlineResponse2004EmbeddedBackupRetentionPolicies) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse2004EmbeddedBackupRetentionPolicies) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse2004EmbeddedBackupRetentionPolicies) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse2004EmbeddedBackupRetentionPolicies) UnmarshalBinary(b []byte) error {
	var res InlineResponse2004EmbeddedBackupRetentionPolicies
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
