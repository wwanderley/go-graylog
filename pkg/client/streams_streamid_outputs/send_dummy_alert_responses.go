// Code generated by go-swagger; DO NOT EDIT.

package streams_streamid_outputs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// SendDummyAlertReader is a Reader for the SendDummyAlert structure.
type SendDummyAlertReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendDummyAlertReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSendDummyAlertOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSendDummyAlertBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSendDummyAlertNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSendDummyAlertOK creates a SendDummyAlertOK with default headers values
func NewSendDummyAlertOK() *SendDummyAlertOK {
	return &SendDummyAlertOK{}
}

/*SendDummyAlertOK handles this case with default header values.

No response was specified
*/
type SendDummyAlertOK struct {
}

func (o *SendDummyAlertOK) Error() string {
	return fmt.Sprintf("[POST /streams/{streamId}/alerts/sendDummyAlert][%d] sendDummyAlertOK ", 200)
}

func (o *SendDummyAlertOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSendDummyAlertBadRequest creates a SendDummyAlertBadRequest with default headers values
func NewSendDummyAlertBadRequest() *SendDummyAlertBadRequest {
	return &SendDummyAlertBadRequest{}
}

/*SendDummyAlertBadRequest handles this case with default header values.

Stream has no alarm callbacks
*/
type SendDummyAlertBadRequest struct {
}

func (o *SendDummyAlertBadRequest) Error() string {
	return fmt.Sprintf("[POST /streams/{streamId}/alerts/sendDummyAlert][%d] sendDummyAlertBadRequest ", 400)
}

func (o *SendDummyAlertBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSendDummyAlertNotFound creates a SendDummyAlertNotFound with default headers values
func NewSendDummyAlertNotFound() *SendDummyAlertNotFound {
	return &SendDummyAlertNotFound{}
}

/*SendDummyAlertNotFound handles this case with default header values.

Stream not found.
*/
type SendDummyAlertNotFound struct {
}

func (o *SendDummyAlertNotFound) Error() string {
	return fmt.Sprintf("[POST /streams/{streamId}/alerts/sendDummyAlert][%d] sendDummyAlertNotFound ", 404)
}

func (o *SendDummyAlertNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}