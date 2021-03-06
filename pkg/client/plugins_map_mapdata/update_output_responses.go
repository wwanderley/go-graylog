// Code generated by go-swagger; DO NOT EDIT.

package plugins_map_mapdata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UpdateOutputReader is a Reader for the UpdateOutput structure.
type UpdateOutputReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOutputReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateOutputOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateOutputBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateOutputOK creates a UpdateOutputOK with default headers values
func NewUpdateOutputOK() *UpdateOutputOK {
	return &UpdateOutputOK{}
}

/*UpdateOutputOK handles this case with default header values.

No response was specified
*/
type UpdateOutputOK struct {
}

func (o *UpdateOutputOK) Error() string {
	return fmt.Sprintf("[PUT /plugins/org.graylog.plugins.collector/configurations/{id}/outputs/{output_id}][%d] updateOutputOK ", 200)
}

func (o *UpdateOutputOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateOutputBadRequest creates a UpdateOutputBadRequest with default headers values
func NewUpdateOutputBadRequest() *UpdateOutputBadRequest {
	return &UpdateOutputBadRequest{}
}

/*UpdateOutputBadRequest handles this case with default header values.

The supplied request is not valid.
*/
type UpdateOutputBadRequest struct {
}

func (o *UpdateOutputBadRequest) Error() string {
	return fmt.Sprintf("[PUT /plugins/org.graylog.plugins.collector/configurations/{id}/outputs/{output_id}][%d] updateOutputBadRequest ", 400)
}

func (o *UpdateOutputBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
