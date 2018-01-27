// Code generated by go-swagger; DO NOT EDIT.

package system_permissions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// UpdateReader is a Reader for the Update structure.
type UpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateOK creates a UpdateOK with default headers values
func NewUpdateOK() *UpdateOK {
	return &UpdateOK{}
}

/*UpdateOK handles this case with default header values.

No response was specified
*/
type UpdateOK struct {
	Payload *models.Output
}

func (o *UpdateOK) Error() string {
	return fmt.Sprintf("[PUT /system/outputs/{outputId}][%d] updateOK  %+v", 200, o.Payload)
}

func (o *UpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Output)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNotFound creates a UpdateNotFound with default headers values
func NewUpdateNotFound() *UpdateNotFound {
	return &UpdateNotFound{}
}

/*UpdateNotFound handles this case with default header values.

No such output on this node.
*/
type UpdateNotFound struct {
}

func (o *UpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /system/outputs/{outputId}][%d] updateNotFound ", 404)
}

func (o *UpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
