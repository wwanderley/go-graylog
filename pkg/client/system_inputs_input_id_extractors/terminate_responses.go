// Code generated by go-swagger; DO NOT EDIT.

package system_inputs_input_id_extractors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// TerminateReader is a Reader for the Terminate structure.
type TerminateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TerminateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTerminateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTerminateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewTerminateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTerminateOK creates a TerminateOK with default headers values
func NewTerminateOK() *TerminateOK {
	return &TerminateOK{}
}

/*TerminateOK handles this case with default header values.

No response was specified
*/
type TerminateOK struct {
}

func (o *TerminateOK) Error() string {
	return fmt.Sprintf("[DELETE /system/inputs/{inputId}/extractors/{extractorId}][%d] terminateOK ", 200)
}

func (o *TerminateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTerminateBadRequest creates a TerminateBadRequest with default headers values
func NewTerminateBadRequest() *TerminateBadRequest {
	return &TerminateBadRequest{}
}

/*TerminateBadRequest handles this case with default header values.

Invalid request.
*/
type TerminateBadRequest struct {
}

func (o *TerminateBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /system/inputs/{inputId}/extractors/{extractorId}][%d] terminateBadRequest ", 400)
}

func (o *TerminateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTerminateNotFound creates a TerminateNotFound with default headers values
func NewTerminateNotFound() *TerminateNotFound {
	return &TerminateNotFound{}
}

/*TerminateNotFound handles this case with default header values.

Extractor not found.
*/
type TerminateNotFound struct {
}

func (o *TerminateNotFound) Error() string {
	return fmt.Sprintf("[DELETE /system/inputs/{inputId}/extractors/{extractorId}][%d] terminateNotFound ", 404)
}

func (o *TerminateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}