// Code generated by go-swagger; DO NOT EDIT.

package system_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetAllReader is a Reader for the GetAll structure.
type GetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAllOK creates a GetAllOK with default headers values
func NewGetAllOK() *GetAllOK {
	return &GetAllOK{}
}

/*GetAllOK handles this case with default header values.

No response was specified
*/
type GetAllOK struct {
	Payload GetAllOKBody
}

func (o *GetAllOK) Error() string {
	return fmt.Sprintf("[GET /system/codecs/types/all][%d] getAllOK  %+v", 200, o.Payload)
}

func (o *GetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetAllOKBody get all o k body
swagger:model GetAllOKBody
*/
type GetAllOKBody interface{}
