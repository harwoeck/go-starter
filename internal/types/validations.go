// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Validate validates this hello world
func (m *HelloWorld) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHello(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HelloWorld) validateHello(formats strfmt.Registry) error {

	if err := validate.Required("hello", "body", m.Hello); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HelloWorld) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HelloWorld) UnmarshalBinary(b []byte) error {
	var res HelloWorld
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// Validate validates this HTTP error
func (m *HTTPError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
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

func (m *HTTPError) validateCode(formats strfmt.Registry) error {

	if swag.IsZero(m.Code) { // not required
		return nil
	}

	if err := validate.MinimumInt("status", "body", int64(m.Code), 100, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("status", "body", int64(m.Code), 599, false); err != nil {
		return err
	}

	return nil
}

func (m *HTTPError) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *HTTPError) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HTTPError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HTTPError) UnmarshalBinary(b []byte) error {
	var res HTTPError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// Validate validates this post login payload
func (m *PostLoginPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostLoginPayload) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *PostLoginPayload) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostLoginPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostLoginPayload) UnmarshalBinary(b []byte) error {
	var res PostLoginPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// Validate validates this post login response
func (m *PostLoginResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiresIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRefreshToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTokenType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostLoginResponse) validateAccessToken(formats strfmt.Registry) error {

	if err := validate.Required("access_token", "body", m.AccessToken); err != nil {
		return err
	}

	return nil
}

func (m *PostLoginResponse) validateExpiresIn(formats strfmt.Registry) error {

	if err := validate.Required("expires_in", "body", m.ExpiresIn); err != nil {
		return err
	}

	return nil
}

func (m *PostLoginResponse) validateRefreshToken(formats strfmt.Registry) error {

	if err := validate.Required("refresh_token", "body", m.RefreshToken); err != nil {
		return err
	}

	return nil
}

func (m *PostLoginResponse) validateTokenType(formats strfmt.Registry) error {

	if err := validate.Required("token_type", "body", m.TokenType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostLoginResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostLoginResponse) UnmarshalBinary(b []byte) error {
	var res PostLoginResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// Validate validates this sample entry
func (m *SampleEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNum(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SampleEntry) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if err := validate.MinLength("data", "body", string(m.Data), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("data", "body", string(m.Data), 10); err != nil {
		return err
	}

	return nil
}

func (m *SampleEntry) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SampleEntry) validateIsActive(formats strfmt.Registry) error {

	if err := validate.Required("isActive", "body", bool(m.IsActive)); err != nil {
		return err
	}

	return nil
}

func (m *SampleEntry) validateMail(formats strfmt.Registry) error {

	if swag.IsZero(m.Mail) { // not required
		return nil
	}

	if err := validate.FormatOf("mail", "body", "email", m.Mail.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SampleEntry) validateNum(formats strfmt.Registry) error {

	if swag.IsZero(m.Num) { // not required
		return nil
	}

	if err := validate.MinimumInt("num", "body", int64(m.Num), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("num", "body", int64(m.Num), 100, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SampleEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SampleEntry) UnmarshalBinary(b []byte) error {
	var res SampleEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}