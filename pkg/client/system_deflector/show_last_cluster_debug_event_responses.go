// Code generated by go-swagger; DO NOT EDIT.

package system_deflector

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// ShowLastClusterDebugEventReader is a Reader for the ShowLastClusterDebugEvent structure.
type ShowLastClusterDebugEventReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowLastClusterDebugEventReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewShowLastClusterDebugEventOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewShowLastClusterDebugEventOK creates a ShowLastClusterDebugEventOK with default headers values
func NewShowLastClusterDebugEventOK() *ShowLastClusterDebugEventOK {
	return &ShowLastClusterDebugEventOK{}
}

/*ShowLastClusterDebugEventOK handles this case with default header values.

No response was specified
*/
type ShowLastClusterDebugEventOK struct {
	Payload *models.DebugEvent
}

func (o *ShowLastClusterDebugEventOK) Error() string {
	return fmt.Sprintf("[GET /system/debug/events/cluster][%d] showLastClusterDebugEventOK  %+v", 200, o.Payload)
}

func (o *ShowLastClusterDebugEventOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DebugEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
