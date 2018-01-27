// Code generated by go-swagger; DO NOT EDIT.

package system_authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// LocalesReader is a Reader for the Locales structure.
type LocalesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocalesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewLocalesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLocalesOK creates a LocalesOK with default headers values
func NewLocalesOK() *LocalesOK {
	return &LocalesOK{}
}

/*LocalesOK handles this case with default header values.

No response was specified
*/
type LocalesOK struct {
	Payload *models.LocalesResponse
}

func (o *LocalesOK) Error() string {
	return fmt.Sprintf("[GET /system/locales][%d] localesOK  %+v", 200, o.Payload)
}

func (o *LocalesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LocalesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
