package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*LogService100LogService This resource represents the log service for the resource or service to which it is associated.

swagger:model LogService.1.0.0_LogService
*/
type LogService100LogService struct {

	/* at odata context
	 */
	AtOdataContext Odata400Context `json:"@odata.context,omitempty"`

	/* at odata id
	 */
	AtOdataID Odata400ID `json:"@odata.id,omitempty"`

	/* at odata type
	 */
	AtOdataType Odata400Type `json:"@odata.type,omitempty"`

	/* actions
	 */
	Actions *LogService100LogServiceActions `json:"Actions,omitempty"`

	/* The current DateTime (with offset) for the log service, used to set or read time.
	 */
	DateTime strfmt.DateTime `json:"DateTime,omitempty"`

	/* The time offset from UTC that the DateTime property is set to in format: +06:00 .

	Pattern: ([-+][0-1][0-9]:[0-5][0-9])
	*/
	DateTimeLocalOffset string `json:"DateTimeLocalOffset,omitempty"`

	/* description
	 */
	Description ResourceDescription `json:"Description,omitempty"`

	/* References to the log entry collection.

	Read Only: true
	*/
	Entries *LogEntryCollectionLogEntryCollection `json:"Entries,omitempty"`

	/* Id
	 */
	ID ResourceID `json:"Id,omitempty"`

	/* The maximum number of log entries this service can have.

	Read Only: true
	Minimum: 0
	*/
	MaxNumberOfRecords float64 `json:"MaxNumberOfRecords,omitempty"`

	/* name
	 */
	Name ResourceName `json:"Name,omitempty"`

	/* This is the manufacturer/provider specific extension moniker used to divide the Oem object into sections.
	 */
	Oem ResourceOem `json:"Oem,omitempty"`

	/* The overwrite policy for this service that takes place when the log is full.

	Read Only: true
	*/
	OverWritePolicy LogService100OverWritePolicy `json:"OverWritePolicy,omitempty"`

	/* This indicates whether this service is enabled.
	 */
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`

	/* status
	 */
	Status *ResourceStatus `json:"Status,omitempty"`
}

// Validate validates this log service 1 0 0 log service
func (m *LogService100LogService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateTimeLocalOffset(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMaxNumberOfRecords(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOverWritePolicy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogService100LogService) validateDateTimeLocalOffset(formats strfmt.Registry) error {

	if swag.IsZero(m.DateTimeLocalOffset) { // not required
		return nil
	}

	if err := validate.Pattern("DateTimeLocalOffset", "body", string(m.DateTimeLocalOffset), `([-+][0-1][0-9]:[0-5][0-9])`); err != nil {
		return err
	}

	return nil
}

func (m *LogService100LogService) validateMaxNumberOfRecords(formats strfmt.Registry) error {

	if swag.IsZero(m.MaxNumberOfRecords) { // not required
		return nil
	}

	if err := validate.Minimum("MaxNumberOfRecords", "body", float64(m.MaxNumberOfRecords), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *LogService100LogService) validateOverWritePolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.OverWritePolicy) { // not required
		return nil
	}

	if err := m.OverWritePolicy.Validate(formats); err != nil {
		return err
	}

	return nil
}

/*LogService100LogServiceActions The Actions object contains the available custom actions on this resource.

swagger:model LogService100LogServiceActions
*/
type LogService100LogServiceActions struct {

	/* log service clear log
	 */
	NrLogServiceClearLog *LogService100ClearLog `json:"#LogService.ClearLog,omitempty"`

	/* oem
	 */
	Oem interface{} `json:"Oem,omitempty"`
}

// Validate validates this log service100 log service actions
func (m *LogService100LogServiceActions) Validate(formats strfmt.Registry) error {
	return nil
}