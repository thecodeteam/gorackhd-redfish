package models

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*VirtualMediaVirtualMedia virtual media virtual media

swagger:model VirtualMedia_VirtualMedia
*/
type VirtualMediaVirtualMedia struct {
	VirtualMedia100VirtualMedia
}

// Validate validates this virtual media virtual media
func (m *VirtualMediaVirtualMedia) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.VirtualMedia100VirtualMedia.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}