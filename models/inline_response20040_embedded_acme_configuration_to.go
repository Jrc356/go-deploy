// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// InlineResponse20040EmbeddedAcmeConfigurationTo inline response 200 40 embedded acme configuration to
// swagger:model inline_response_200_40__embedded_acme_configuration_to
type InlineResponse20040EmbeddedAcmeConfigurationTo struct {

	// legacy
	Legacy bool `json:"legacy,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this inline response 200 40 embedded acme configuration to
func (m *InlineResponse20040EmbeddedAcmeConfigurationTo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse20040EmbeddedAcmeConfigurationTo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse20040EmbeddedAcmeConfigurationTo) UnmarshalBinary(b []byte) error {
	var res InlineResponse20040EmbeddedAcmeConfigurationTo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
