package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
)

/*Task100Task This is the schema definition for a Task resource.

swagger:model Task.1.0.0_Task
*/
type Task100Task struct {

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

	/* The date-time stamp that the task was last completed.

	Read Only: true
	*/
	EndTime strfmt.DateTime `json:"EndTime,omitempty"`

	/* Id
	 */
	ID ResourceID `json:"Id,omitempty"`

	/* This is an array of messages associated with the task.

	Read Only: true
	*/
	Messages []*Message100Message `json:"Messages,omitempty"`

	/* name
	 */
	Name ResourceName `json:"Name,omitempty"`

	/* This is the manufacturer/provider specific extension moniker used to divide the Oem object into sections.
	 */
	Oem ResourceOem `json:"Oem,omitempty"`

	/* The date-time stamp that the task was last started.

	Read Only: true
	*/
	StartTime strfmt.DateTime `json:"StartTime,omitempty"`

	/* The state of the task.

	Read Only: true
	*/
	TaskState Task100TaskState `json:"TaskState,omitempty"`

	/* This is the completion status of the task.

	Read Only: true
	*/
	TaskStatus ResourceHealth `json:"TaskStatus,omitempty"`
}

// Validate validates this task 1 0 0 task
func (m *Task100Task) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessages(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTaskState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTaskStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Task100Task) validateMessages(formats strfmt.Registry) error {

	if swag.IsZero(m.Messages) { // not required
		return nil
	}

	return nil
}

func (m *Task100Task) validateTaskState(formats strfmt.Registry) error {

	if swag.IsZero(m.TaskState) { // not required
		return nil
	}

	if err := m.TaskState.Validate(formats); err != nil {
		return err
	}

	return nil
}

func (m *Task100Task) validateTaskStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.TaskStatus) { // not required
		return nil
	}

	if err := m.TaskStatus.Validate(formats); err != nil {
		return err
	}

	return nil
}