// Code generated by go-swagger; DO NOT EDIT.

package system_lbstatus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ReadGroupsReader is a Reader for the ReadGroups structure.
type ReadGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReadGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReadGroupsOK creates a ReadGroupsOK with default headers values
func NewReadGroupsOK() *ReadGroupsOK {
	return &ReadGroupsOK{}
}

/*ReadGroupsOK handles this case with default header values.

No response was specified
*/
type ReadGroupsOK struct {
	Payload []interface{}
}

func (o *ReadGroupsOK) Error() string {
	return fmt.Sprintf("[GET /system/ldap/groups][%d] readGroupsOK  %+v", 200, o.Payload)
}

func (o *ReadGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
