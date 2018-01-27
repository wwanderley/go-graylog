// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ChangeUserReader is a Reader for the ChangeUser structure.
type ChangeUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewChangeUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewChangeUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeUserOK creates a ChangeUserOK with default headers values
func NewChangeUserOK() *ChangeUserOK {
	return &ChangeUserOK{}
}

/*ChangeUserOK handles this case with default header values.

No response was specified
*/
type ChangeUserOK struct {
}

func (o *ChangeUserOK) Error() string {
	return fmt.Sprintf("[PUT /users/{username}][%d] changeUserOK ", 200)
}

func (o *ChangeUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeUserBadRequest creates a ChangeUserBadRequest with default headers values
func NewChangeUserBadRequest() *ChangeUserBadRequest {
	return &ChangeUserBadRequest{}
}

/*ChangeUserBadRequest handles this case with default header values.

Missing or invalid user details.
*/
type ChangeUserBadRequest struct {
}

func (o *ChangeUserBadRequest) Error() string {
	return fmt.Sprintf("[PUT /users/{username}][%d] changeUserBadRequest ", 400)
}

func (o *ChangeUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
