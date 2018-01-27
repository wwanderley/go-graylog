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

// TermsAbsoluteReader is a Reader for the TermsAbsolute structure.
type TermsAbsoluteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TermsAbsoluteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTermsAbsoluteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewTermsAbsoluteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTermsAbsoluteOK creates a TermsAbsoluteOK with default headers values
func NewTermsAbsoluteOK() *TermsAbsoluteOK {
	return &TermsAbsoluteOK{}
}

/*TermsAbsoluteOK handles this case with default header values.

No response was specified
*/
type TermsAbsoluteOK struct {
	Payload *models.TermsResult
}

func (o *TermsAbsoluteOK) Error() string {
	return fmt.Sprintf("[GET /search/universal/absolute/terms][%d] termsAbsoluteOK  %+v", 200, o.Payload)
}

func (o *TermsAbsoluteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TermsResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTermsAbsoluteBadRequest creates a TermsAbsoluteBadRequest with default headers values
func NewTermsAbsoluteBadRequest() *TermsAbsoluteBadRequest {
	return &TermsAbsoluteBadRequest{}
}

/*TermsAbsoluteBadRequest handles this case with default header values.

Invalid timerange parameters provided.
*/
type TermsAbsoluteBadRequest struct {
}

func (o *TermsAbsoluteBadRequest) Error() string {
	return fmt.Sprintf("[GET /search/universal/absolute/terms][%d] termsAbsoluteBadRequest ", 400)
}

func (o *TermsAbsoluteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
