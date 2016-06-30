package redfish_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/emccode/gorackhd-redfish/models"
)

// GetManagersIdentifierVirtualMediaIndexReader is a Reader for the GetManagersIdentifierVirtualMediaIndex structure.
type GetManagersIdentifierVirtualMediaIndexReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetManagersIdentifierVirtualMediaIndexReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetManagersIdentifierVirtualMediaIndexOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetManagersIdentifierVirtualMediaIndexBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetManagersIdentifierVirtualMediaIndexUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetManagersIdentifierVirtualMediaIndexForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetManagersIdentifierVirtualMediaIndexNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetManagersIdentifierVirtualMediaIndexInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetManagersIdentifierVirtualMediaIndexOK creates a GetManagersIdentifierVirtualMediaIndexOK with default headers values
func NewGetManagersIdentifierVirtualMediaIndexOK() *GetManagersIdentifierVirtualMediaIndexOK {
	return &GetManagersIdentifierVirtualMediaIndexOK{}
}

/*GetManagersIdentifierVirtualMediaIndexOK handles this case with default header values.

Success
*/
type GetManagersIdentifierVirtualMediaIndexOK struct {
	Payload *models.VirtualMedia100VirtualMedia
}

func (o *GetManagersIdentifierVirtualMediaIndexOK) Error() string {
	return fmt.Sprintf("[GET /Managers/{identifier}/VirtualMedia/{index}][%d] getManagersIdentifierVirtualMediaIndexOK  %+v", 200, o.Payload)
}

func (o *GetManagersIdentifierVirtualMediaIndexOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMedia100VirtualMedia)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetManagersIdentifierVirtualMediaIndexBadRequest creates a GetManagersIdentifierVirtualMediaIndexBadRequest with default headers values
func NewGetManagersIdentifierVirtualMediaIndexBadRequest() *GetManagersIdentifierVirtualMediaIndexBadRequest {
	return &GetManagersIdentifierVirtualMediaIndexBadRequest{}
}

/*GetManagersIdentifierVirtualMediaIndexBadRequest handles this case with default header values.

The request could not be processed because it contains missing or invalid information  (such as validation error on an input field, a missing required value, and so on).  An extended error shall be returned in the response body, as defined in section Extended  Error Handling.

*/
type GetManagersIdentifierVirtualMediaIndexBadRequest struct {
}

func (o *GetManagersIdentifierVirtualMediaIndexBadRequest) Error() string {
	return fmt.Sprintf("[GET /Managers/{identifier}/VirtualMedia/{index}][%d] getManagersIdentifierVirtualMediaIndexBadRequest ", 400)
}

func (o *GetManagersIdentifierVirtualMediaIndexBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetManagersIdentifierVirtualMediaIndexUnauthorized creates a GetManagersIdentifierVirtualMediaIndexUnauthorized with default headers values
func NewGetManagersIdentifierVirtualMediaIndexUnauthorized() *GetManagersIdentifierVirtualMediaIndexUnauthorized {
	return &GetManagersIdentifierVirtualMediaIndexUnauthorized{}
}

/*GetManagersIdentifierVirtualMediaIndexUnauthorized handles this case with default header values.

The authentication credentials included with this request are missing or invalid.

*/
type GetManagersIdentifierVirtualMediaIndexUnauthorized struct {
}

func (o *GetManagersIdentifierVirtualMediaIndexUnauthorized) Error() string {
	return fmt.Sprintf("[GET /Managers/{identifier}/VirtualMedia/{index}][%d] getManagersIdentifierVirtualMediaIndexUnauthorized ", 401)
}

func (o *GetManagersIdentifierVirtualMediaIndexUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetManagersIdentifierVirtualMediaIndexForbidden creates a GetManagersIdentifierVirtualMediaIndexForbidden with default headers values
func NewGetManagersIdentifierVirtualMediaIndexForbidden() *GetManagersIdentifierVirtualMediaIndexForbidden {
	return &GetManagersIdentifierVirtualMediaIndexForbidden{}
}

/*GetManagersIdentifierVirtualMediaIndexForbidden handles this case with default header values.

The server recognized the credentials in the request, but those credentials do not  possess authorization to perform this request.

*/
type GetManagersIdentifierVirtualMediaIndexForbidden struct {
}

func (o *GetManagersIdentifierVirtualMediaIndexForbidden) Error() string {
	return fmt.Sprintf("[GET /Managers/{identifier}/VirtualMedia/{index}][%d] getManagersIdentifierVirtualMediaIndexForbidden ", 403)
}

func (o *GetManagersIdentifierVirtualMediaIndexForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetManagersIdentifierVirtualMediaIndexNotFound creates a GetManagersIdentifierVirtualMediaIndexNotFound with default headers values
func NewGetManagersIdentifierVirtualMediaIndexNotFound() *GetManagersIdentifierVirtualMediaIndexNotFound {
	return &GetManagersIdentifierVirtualMediaIndexNotFound{}
}

/*GetManagersIdentifierVirtualMediaIndexNotFound handles this case with default header values.

The request specified a URI of a resource that does not exist.

*/
type GetManagersIdentifierVirtualMediaIndexNotFound struct {
}

func (o *GetManagersIdentifierVirtualMediaIndexNotFound) Error() string {
	return fmt.Sprintf("[GET /Managers/{identifier}/VirtualMedia/{index}][%d] getManagersIdentifierVirtualMediaIndexNotFound ", 404)
}

func (o *GetManagersIdentifierVirtualMediaIndexNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetManagersIdentifierVirtualMediaIndexInternalServerError creates a GetManagersIdentifierVirtualMediaIndexInternalServerError with default headers values
func NewGetManagersIdentifierVirtualMediaIndexInternalServerError() *GetManagersIdentifierVirtualMediaIndexInternalServerError {
	return &GetManagersIdentifierVirtualMediaIndexInternalServerError{}
}

/*GetManagersIdentifierVirtualMediaIndexInternalServerError handles this case with default header values.

Error
*/
type GetManagersIdentifierVirtualMediaIndexInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetManagersIdentifierVirtualMediaIndexInternalServerError) Error() string {
	return fmt.Sprintf("[GET /Managers/{identifier}/VirtualMedia/{index}][%d] getManagersIdentifierVirtualMediaIndexInternalServerError  %+v", 500, o.Payload)
}

func (o *GetManagersIdentifierVirtualMediaIndexInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}