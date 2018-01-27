// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// ThreadDumpReader is a Reader for the ThreadDump structure.
type ThreadDumpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThreadDumpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewThreadDumpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewThreadDumpOK creates a ThreadDumpOK with default headers values
func NewThreadDumpOK() *ThreadDumpOK {
	return &ThreadDumpOK{}
}

/*ThreadDumpOK handles this case with default header values.

No response was specified
*/
type ThreadDumpOK struct {
	Payload *models.SystemThreadDumpResponse
}

func (o *ThreadDumpOK) Error() string {
	return fmt.Sprintf("[GET /cluster/{nodeId}/threaddump][%d] threadDumpOK  %+v", 200, o.Payload)
}

func (o *ThreadDumpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemThreadDumpResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
