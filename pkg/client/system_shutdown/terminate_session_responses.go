// Code generated by go-swagger; DO NOT EDIT.

package system_shutdown

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// TerminateSessionReader is a Reader for the TerminateSession structure.
type TerminateSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TerminateSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTerminateSessionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTerminateSessionOK creates a TerminateSessionOK with default headers values
func NewTerminateSessionOK() *TerminateSessionOK {
	return &TerminateSessionOK{}
}

/*TerminateSessionOK handles this case with default header values.

No response was specified
*/
type TerminateSessionOK struct {
}

func (o *TerminateSessionOK) Error() string {
	return fmt.Sprintf("[DELETE /system/sessions/{sessionId}][%d] terminateSessionOK ", 200)
}

func (o *TerminateSessionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}