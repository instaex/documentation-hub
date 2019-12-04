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

// DecodedCalldata decoded calldata
// swagger:model DecodedCalldata
type DecodedCalldata struct {

	// arguments
	// Required: true
	Arguments []interface{} `json:"arguments"`

	// function
	// Required: true
	Function *string `json:"function"`
}

// Validate validates this decoded calldata
func (m *DecodedCalldata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArguments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFunction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DecodedCalldata) validateArguments(formats strfmt.Registry) error {

	if err := validate.Required("arguments", "body", m.Arguments); err != nil {
		return err
	}

	return nil
}

func (m *DecodedCalldata) validateFunction(formats strfmt.Registry) error {

	if err := validate.Required("function", "body", m.Function); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DecodedCalldata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DecodedCalldata) UnmarshalBinary(b []byte) error {
	var res DecodedCalldata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}