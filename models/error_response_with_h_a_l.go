// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ErrorResponseWithHAL error response with h a l
// swagger:model ErrorResponseWithHAL
type ErrorResponseWithHAL []*ErrorResponseWithHALItems0

// Validate validates this error response with h a l
func (m ErrorResponseWithHAL) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ErrorResponseWithHAL) validateLinks(formats strfmt.Registry) error {

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

func (m *ErrorResponseWithHAL) validateLinks(formats strfmt.Registry) error {

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

// ErrorResponseWithHALItems0 error response with h a l items0
// swagger:model ErrorResponseWithHALItems0
type ErrorResponseWithHALItems0 struct {

	// details
	Details *ErrorResponseWithHALItems0Details `json:"details,omitempty"`

	// The detailed message related to the type stated above.
	Message string `json:"message,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this error response with h a l items0
func (m *ErrorResponseWithHALItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ErrorResponseWithHALItems0) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.Details) { // not required
		return nil
	}

	if m.Details != nil {
		if err := m.Details.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("details")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ErrorResponseWithHALItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorResponseWithHALItems0) UnmarshalBinary(b []byte) error {
	var res ErrorResponseWithHALItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ErrorResponseWithHALItems0Details error response with h a l items0 details
// swagger:model ErrorResponseWithHALItems0Details
type ErrorResponseWithHALItems0Details struct {

	// The location of the field that threw the error.
	Location string `json:"location,omitempty"`

	// The name of the field that threw the error.
	Name string `json:"name,omitempty"`
}

// Validate validates this error response with h a l items0 details
func (m *ErrorResponseWithHALItems0Details) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ErrorResponseWithHALItems0Details) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorResponseWithHALItems0Details) UnmarshalBinary(b []byte) error {
	var res ErrorResponseWithHALItems0Details
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ErrorResponseWithHALLinks error response with h a l links
// swagger:model ErrorResponseWithHALLinks
type ErrorResponseWithHALLinks struct {

	// payment instruments
	PaymentInstruments *ErrorResponseWithHALLinksPaymentInstruments `json:"paymentInstruments,omitempty"`
}

// Validate validates this error response with h a l links
func (m *ErrorResponseWithHALLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaymentInstruments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ErrorResponseWithHALLinks) validatePaymentInstruments(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentInstruments) { // not required
		return nil
	}

	if m.PaymentInstruments != nil {
		if err := m.PaymentInstruments.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "paymentInstruments")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ErrorResponseWithHALLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorResponseWithHALLinks) UnmarshalBinary(b []byte) error {
	var res ErrorResponseWithHALLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ErrorResponseWithHALLinksPaymentInstruments error response with h a l links payment instruments
// swagger:model ErrorResponseWithHALLinksPaymentInstruments
type ErrorResponseWithHALLinksPaymentInstruments struct {

	// href
	Href string `json:"href,omitempty"`
}

// Validate validates this error response with h a l links payment instruments
func (m *ErrorResponseWithHALLinksPaymentInstruments) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ErrorResponseWithHALLinksPaymentInstruments) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorResponseWithHALLinksPaymentInstruments) UnmarshalBinary(b []byte) error {
	var res ErrorResponseWithHALLinksPaymentInstruments
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}