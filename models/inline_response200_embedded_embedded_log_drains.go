// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineResponse200EmbeddedEmbeddedLogDrains inline response 200 embedded embedded log drains
// swagger:model inline_response_200__embedded__embedded_log_drains
type InlineResponse200EmbeddedEmbeddedLogDrains struct {

	// resource type
	ResourceType string `json:"_type,omitempty"`

	// links
	Links *InlineResponse200EmbeddedEmbeddedLinks `json:"_links,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// drain apps
	DrainApps bool `json:"drain_apps,omitempty"`

	// drain databases
	DrainDatabases bool `json:"drain_databases,omitempty"`

	// drain ephemeral sessions
	DrainEphemeralSessions bool `json:"drain_ephemeral_sessions,omitempty"`

	// drain host
	DrainHost string `json:"drain_host,omitempty"`

	// drain password
	DrainPassword *string `json:"drain_password,omitempty"`

	// drain port
	DrainPort int64 `json:"drain_port,omitempty"`

	// drain proxies
	DrainProxies bool `json:"drain_proxies,omitempty"`

	// drain type
	DrainType string `json:"drain_type,omitempty"`

	// drain username
	DrainUsername *string `json:"drain_username,omitempty"`

	// gentlemanjerry allocation
	GentlemanjerryAllocation []string `json:"gentlemanjerry_allocation"`

	// gentlemanjerry certificate
	GentlemanjerryCertificate *string `json:"gentlemanjerry_certificate,omitempty"`

	// gentlemanjerry docker name
	GentlemanjerryDockerName *string `json:"gentlemanjerry_docker_name,omitempty"`

	// gentlemanjerry host
	GentlemanjerryHost *string `json:"gentlemanjerry_host,omitempty"`

	// gentlemanjerry instance id
	GentlemanjerryInstanceID *string `json:"gentlemanjerry_instance_id,omitempty"`

	// gentlemanjerry port mapping
	GentlemanjerryPortMapping [][]int64 `json:"gentlemanjerry_port_mapping"`

	// handle
	Handle string `json:"handle,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// logging token
	LoggingToken *string `json:"logging_token,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// url
	URL *string `json:"url,omitempty"`
}

// Validate validates this inline response 200 embedded embedded log drains
func (m *InlineResponse200EmbeddedEmbeddedLogDrains) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineResponse200EmbeddedEmbeddedLogDrains) validateLinks(formats strfmt.Registry) error {

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
func (m *InlineResponse200EmbeddedEmbeddedLogDrains) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse200EmbeddedEmbeddedLogDrains) UnmarshalBinary(b []byte) error {
	var res InlineResponse200EmbeddedEmbeddedLogDrains
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
