package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
)

/*TaskCollectionTaskCollection task collection task collection

swagger:model TaskCollection_TaskCollection
*/
type TaskCollectionTaskCollection struct {

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

	/* Contains the members of this collection.

	Read Only: true
	*/
	Members []TaskTask `json:"Members,omitempty"`

	/* members at odata count
	 */
	MembersAtOdataCount Odata400Count `json:"Members@odata.count,omitempty"`

	/* members at odata navigation link
	 */
	MembersAtOdataNavigationLink *Odata400IDRef `json:"Members@odata.navigationLink,omitempty"`

	/* name
	 */
	Name ResourceName `json:"Name,omitempty"`

	/* This is the manufacturer/provider specific extension moniker used to divide the Oem object into sections.
	 */
	Oem ResourceOem `json:"Oem,omitempty"`
}

// Validate validates this task collection task collection
func (m *TaskCollectionTaskCollection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMembers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskCollectionTaskCollection) validateMembers(formats strfmt.Registry) error {

	if swag.IsZero(m.Members) { // not required
		return nil
	}

	return nil
}