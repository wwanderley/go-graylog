// Code generated by go-swagger; DO NOT EDIT.

package system_indices_retention

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// SaveReader is a Reader for the Save structure.
type SaveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SaveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSaveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewSaveForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSaveOK creates a SaveOK with default headers values
func NewSaveOK() *SaveOK {
	return &SaveOK{}
}

/*SaveOK handles this case with default header values.

No response was specified
*/
type SaveOK struct {
	Payload *models.IndexSetSummary
}

func (o *SaveOK) Error() string {
	return fmt.Sprintf("[POST /system/indices/index_sets][%d] saveOK  %+v", 200, o.Payload)
}

func (o *SaveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IndexSetSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSaveForbidden creates a SaveForbidden with default headers values
func NewSaveForbidden() *SaveForbidden {
	return &SaveForbidden{}
}

/*SaveForbidden handles this case with default header values.

Unauthorized
*/
type SaveForbidden struct {
}

func (o *SaveForbidden) Error() string {
	return fmt.Sprintf("[POST /system/indices/index_sets][%d] saveForbidden ", 403)
}

func (o *SaveForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}