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

// GetReader is a Reader for the Get structure.
type GetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOK creates a GetOK with default headers values
func NewGetOK() *GetOK {
	return &GetOK{}
}

/*GetOK handles this case with default header values.

No response was specified
*/
type GetOK struct {
	Payload *models.OutputSummary
}

func (o *GetOK) Error() string {
	return fmt.Sprintf("[GET /system/outputs/{outputId}][%d] getOK  %+v", 200, o.Payload)
}

func (o *GetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OutputSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotFound creates a GetNotFound with default headers values
func NewGetNotFound() *GetNotFound {
	return &GetNotFound{}
}

/*GetNotFound handles this case with default header values.

No such output on this node.
*/
type GetNotFound struct {
}

func (o *GetNotFound) Error() string {
	return fmt.Sprintf("[GET /system/outputs/{outputId}][%d] getNotFound ", 404)
}

func (o *GetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}