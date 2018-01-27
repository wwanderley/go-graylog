// Code generated by go-swagger; DO NOT EDIT.

package search_universal_absolute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// AddMemberReader is a Reader for the AddMember structure.
type AddMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddMemberOK creates a AddMemberOK with default headers values
func NewAddMemberOK() *AddMemberOK {
	return &AddMemberOK{}
}

/*AddMemberOK handles this case with default header values.

No response was specified
*/
type AddMemberOK struct {
}

func (o *AddMemberOK) Error() string {
	return fmt.Sprintf("[PUT /roles/{rolename}/members/{username}][%d] addMemberOK ", 200)
}

func (o *AddMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}