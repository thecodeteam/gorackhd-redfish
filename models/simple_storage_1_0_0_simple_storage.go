package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
)

/*SimpleStorage100SimpleStorage This is the schema definition for the Simple Storage resource.  It represents the properties of a storage controller and its directly-attached devices.

swagger:model SimpleStorage.1.0.0_SimpleStorage
*/
type SimpleStorage100SimpleStorage struct {

	/* at odata context
	 */
	AtOdataContext Odata400Context `json:"@odata.context,omitempty"`

	/* at odata id
	 */
	AtOdataID Odata400ID `json:"@odata.id,omitempty"`

	/* at odata type
	 */
	AtOdataType Odata400Type `json:"@odata.type,omitempty"`

	/* description
	 */
	Description ResourceDescription `json:"Description,omitempty"`

	/* The storage devices associated with this resource

	Read Only: true
	*/
	Devices []*SimpleStorage100Device `json:"Devices,omitempty"`

	/* Id
	 */
	ID ResourceID `json:"Id,omitempty"`

	/* name
	 */
	Name ResourceName `json:"Name,omitempty"`

	/* This is the manufacturer/provider specific extension moniker used to divide the Oem object into sections.
	 */
	Oem ResourceOem `json:"Oem,omitempty"`

	/* status
	 */
	Status *ResourceStatus `json:"Status,omitempty"`

	/* The UEFI device path used to access this storage controller.

	Read Only: true
	*/
	UefiDevicePath string `json:"UefiDevicePath,omitempty"`
}

// Validate validates this simple storage 1 0 0 simple storage
func (m *SimpleStorage100SimpleStorage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDevices(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SimpleStorage100SimpleStorage) validateDevices(formats strfmt.Registry) error {

	if swag.IsZero(m.Devices) { // not required
		return nil
	}

	return nil
}