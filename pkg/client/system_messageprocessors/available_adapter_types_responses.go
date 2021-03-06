// Code generated by go-swagger; DO NOT EDIT.

package system_messageprocessors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// AvailableAdapterTypesReader is a Reader for the AvailableAdapterTypes structure.
type AvailableAdapterTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AvailableAdapterTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAvailableAdapterTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAvailableAdapterTypesOK creates a AvailableAdapterTypesOK with default headers values
func NewAvailableAdapterTypesOK() *AvailableAdapterTypesOK {
	return &AvailableAdapterTypesOK{}
}

/*AvailableAdapterTypesOK handles this case with default header values.

No response was specified
*/
type AvailableAdapterTypesOK struct {
	Payload AvailableAdapterTypesOKBody
}

func (o *AvailableAdapterTypesOK) Error() string {
	return fmt.Sprintf("[GET /system/lookup/types/adapters][%d] availableAdapterTypesOK  %+v", 200, o.Payload)
}

func (o *AvailableAdapterTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AvailableAdapterTypesOKBody available adapter types o k body
swagger:model AvailableAdapterTypesOKBody
*/
type AvailableAdapterTypesOKBody interface{}
