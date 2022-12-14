// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostOKBody post o k body
//
// swagger:model postOKBody
type PostOKBody struct {

	// mustache
	Mustache *PostOKBodyMustache `json:"mustache,omitempty"`
}

// Validate validates this post o k body
func (m *PostOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMustache(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostOKBody) validateMustache(formats strfmt.Registry) error {
	if swag.IsZero(m.Mustache) { // not required
		return nil
	}

	if m.Mustache != nil {
		if err := m.Mustache.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mustache")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mustache")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post o k body based on the context it is used
func (m *PostOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMustache(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostOKBody) contextValidateMustache(ctx context.Context, formats strfmt.Registry) error {

	if m.Mustache != nil {
		if err := m.Mustache.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mustache")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mustache")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostOKBody) UnmarshalBinary(b []byte) error {
	var res PostOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
