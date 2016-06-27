package models

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*Manager100Reset manager 1 0 0 reset

swagger:model Manager.1.0.0_Reset
*/
type Manager100Reset struct {

	/* Link to invoke action
	 */
	Target strfmt.URI `json:"target,omitempty"`

	/* Friendly action name
	 */
	Title string `json:"title,omitempty"`
}

// Validate validates this manager 1 0 0 reset
func (m *Manager100Reset) Validate(formats strfmt.Registry) error {
	return nil
}