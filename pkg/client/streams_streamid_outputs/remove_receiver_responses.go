// Code generated by go-swagger; DO NOT EDIT.

package streams_streamid_outputs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// RemoveReceiverReader is a Reader for the RemoveReceiver structure.
type RemoveReceiverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveReceiverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRemoveReceiverOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRemoveReceiverBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRemoveReceiverNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveReceiverOK creates a RemoveReceiverOK with default headers values
func NewRemoveReceiverOK() *RemoveReceiverOK {
	return &RemoveReceiverOK{}
}

/*RemoveReceiverOK handles this case with default header values.

No response was specified
*/
type RemoveReceiverOK struct {
}

func (o *RemoveReceiverOK) Error() string {
	return fmt.Sprintf("[DELETE /streams/{streamId}/alerts/receivers][%d] removeReceiverOK ", 200)
}

func (o *RemoveReceiverOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveReceiverBadRequest creates a RemoveReceiverBadRequest with default headers values
func NewRemoveReceiverBadRequest() *RemoveReceiverBadRequest {
	return &RemoveReceiverBadRequest{}
}

/*RemoveReceiverBadRequest handles this case with default header values.

Stream has no email alarm callbacks.
*/
type RemoveReceiverBadRequest struct {
}

func (o *RemoveReceiverBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /streams/{streamId}/alerts/receivers][%d] removeReceiverBadRequest ", 400)
}

func (o *RemoveReceiverBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveReceiverNotFound creates a RemoveReceiverNotFound with default headers values
func NewRemoveReceiverNotFound() *RemoveReceiverNotFound {
	return &RemoveReceiverNotFound{}
}

/*RemoveReceiverNotFound handles this case with default header values.

Stream not found.
*/
type RemoveReceiverNotFound struct {
}

func (o *RemoveReceiverNotFound) Error() string {
	return fmt.Sprintf("[DELETE /streams/{streamId}/alerts/receivers][%d] removeReceiverNotFound ", 404)
}

func (o *RemoveReceiverNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
