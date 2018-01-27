// Code generated by go-swagger; DO NOT EDIT.

package system_indexer_overview

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// OpenReader is a Reader for the Open structure.
type OpenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OpenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewOpenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOpenOK creates a OpenOK with default headers values
func NewOpenOK() *OpenOK {
	return &OpenOK{}
}

/*OpenOK handles this case with default header values.

No response was specified
*/
type OpenOK struct {
	Payload *models.OpenIndicesInfo
}

func (o *OpenOK) Error() string {
	return fmt.Sprintf("[GET /system/indexer/indices/open][%d] openOK  %+v", 200, o.Payload)
}

func (o *OpenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenIndicesInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
