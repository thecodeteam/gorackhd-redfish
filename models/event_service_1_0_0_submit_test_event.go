package models

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*EventService100SubmitTestEvent event service 1 0 0 submit test event

swagger:model EventService.1.0.0_SubmitTestEvent
*/
type EventService100SubmitTestEvent struct {

	/* Link to invoke action
	 */
	Target strfmt.URI `json:"target,omitempty"`

	/* Friendly action name
	 */
	Title string `json:"title,omitempty"`
}

// Validate validates this event service 1 0 0 submit test event
func (m *EventService100SubmitTestEvent) Validate(formats strfmt.Registry) error {
	return nil
}