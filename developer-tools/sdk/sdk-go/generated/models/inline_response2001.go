// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// InlineResponse2001 inline response 200 1
// swagger:model inline_response_200_1
type InlineResponse2001 struct {

	// Height
	Height uint64 `json:"height,omitempty"`
}

// Validate validates this inline response 200 1
func (m *InlineResponse2001) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InlineResponse2001) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineResponse2001) UnmarshalBinary(b []byte) error {
	var res InlineResponse2001
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}