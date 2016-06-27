package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*Manager100CommandConnectTypesSupported manager 1 0 0 command connect types supported

swagger:model Manager.1.0.0_CommandConnectTypesSupported
*/
type Manager100CommandConnectTypesSupported string

// for schema
var manager100CommandConnectTypesSupportedEnum []interface{}

func (m Manager100CommandConnectTypesSupported) validateManager100CommandConnectTypesSupportedEnum(path, location string, value Manager100CommandConnectTypesSupported) error {
	if manager100CommandConnectTypesSupportedEnum == nil {
		var res []Manager100CommandConnectTypesSupported
		if err := json.Unmarshal([]byte(`["SSH","Telnet","IPMI","Oem"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			manager100CommandConnectTypesSupportedEnum = append(manager100CommandConnectTypesSupportedEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, manager100CommandConnectTypesSupportedEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this manager 1 0 0 command connect types supported
func (m Manager100CommandConnectTypesSupported) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateManager100CommandConnectTypesSupportedEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}