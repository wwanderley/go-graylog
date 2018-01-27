// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeletePermissionsReader is a Reader for the DeletePermissions structure.
type DeletePermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeletePermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewDeletePermissionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeletePermissionsOK creates a DeletePermissionsOK with default headers values
func NewDeletePermissionsOK() *DeletePermissionsOK {
	return &DeletePermissionsOK{}
}

/*DeletePermissionsOK handles this case with default header values.

No response was specified
*/
type DeletePermissionsOK struct {
}

func (o *DeletePermissionsOK) Error() string {
	return fmt.Sprintf("[DELETE /users/{username}/permissions][%d] deletePermissionsOK ", 200)
}

func (o *DeletePermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePermissionsInternalServerError creates a DeletePermissionsInternalServerError with default headers values
func NewDeletePermissionsInternalServerError() *DeletePermissionsInternalServerError {
	return &DeletePermissionsInternalServerError{}
}

/*DeletePermissionsInternalServerError handles this case with default header values.

When saving the user failed.
*/
type DeletePermissionsInternalServerError struct {
}

func (o *DeletePermissionsInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /users/{username}/permissions][%d] deletePermissionsInternalServerError ", 500)
}

func (o *DeletePermissionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}