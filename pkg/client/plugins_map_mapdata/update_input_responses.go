// Code generated by go-swagger; DO NOT EDIT.

package plugins_map_mapdata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UpdateInputReader is a Reader for the UpdateInput structure.
type UpdateInputReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateInputReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateInputOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateInputBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateInputOK creates a UpdateInputOK with default headers values
func NewUpdateInputOK() *UpdateInputOK {
	return &UpdateInputOK{}
}

/*UpdateInputOK handles this case with default header values.

No response was specified
*/
type UpdateInputOK struct {
}

func (o *UpdateInputOK) Error() string {
	return fmt.Sprintf("[PUT /plugins/org.graylog.plugins.collector/configurations/{id}/inputs/{input_id}][%d] updateInputOK ", 200)
}

func (o *UpdateInputOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateInputBadRequest creates a UpdateInputBadRequest with default headers values
func NewUpdateInputBadRequest() *UpdateInputBadRequest {
	return &UpdateInputBadRequest{}
}

/*UpdateInputBadRequest handles this case with default header values.

The supplied request is not valid.
*/
type UpdateInputBadRequest struct {
}

func (o *UpdateInputBadRequest) Error() string {
	return fmt.Sprintf("[PUT /plugins/org.graylog.plugins.collector/configurations/{id}/inputs/{input_id}][%d] updateInputBadRequest ", 400)
}

func (o *UpdateInputBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}