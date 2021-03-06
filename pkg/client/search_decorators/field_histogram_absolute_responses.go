// Code generated by go-swagger; DO NOT EDIT.

package search_decorators

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// FieldHistogramAbsoluteReader is a Reader for the FieldHistogramAbsolute structure.
type FieldHistogramAbsoluteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FieldHistogramAbsoluteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFieldHistogramAbsoluteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewFieldHistogramAbsoluteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewFieldHistogramAbsoluteOK creates a FieldHistogramAbsoluteOK with default headers values
func NewFieldHistogramAbsoluteOK() *FieldHistogramAbsoluteOK {
	return &FieldHistogramAbsoluteOK{}
}

/*FieldHistogramAbsoluteOK handles this case with default header values.

No response was specified
*/
type FieldHistogramAbsoluteOK struct {
	Payload *models.HistogramResult
}

func (o *FieldHistogramAbsoluteOK) Error() string {
	return fmt.Sprintf("[GET /search/universal/absolute/fieldhistogram][%d] fieldHistogramAbsoluteOK  %+v", 200, o.Payload)
}

func (o *FieldHistogramAbsoluteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HistogramResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFieldHistogramAbsoluteBadRequest creates a FieldHistogramAbsoluteBadRequest with default headers values
func NewFieldHistogramAbsoluteBadRequest() *FieldHistogramAbsoluteBadRequest {
	return &FieldHistogramAbsoluteBadRequest{}
}

/*FieldHistogramAbsoluteBadRequest handles this case with default header values.

Field is not of numeric type.
*/
type FieldHistogramAbsoluteBadRequest struct {
}

func (o *FieldHistogramAbsoluteBadRequest) Error() string {
	return fmt.Sprintf("[GET /search/universal/absolute/fieldhistogram][%d] fieldHistogramAbsoluteBadRequest ", 400)
}

func (o *FieldHistogramAbsoluteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
