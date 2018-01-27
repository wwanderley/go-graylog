// Code generated by go-swagger; DO NOT EDIT.

package plugins_system_pipelines_pipeline

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ConnectStreamsReader is a Reader for the ConnectStreams structure.
type ConnectStreamsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConnectStreamsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewConnectStreamsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewConnectStreamsOK creates a ConnectStreamsOK with default headers values
func NewConnectStreamsOK() *ConnectStreamsOK {
	return &ConnectStreamsOK{}
}

/*ConnectStreamsOK handles this case with default header values.

No response was specified
*/
type ConnectStreamsOK struct {
	Payload []interface{}
}

func (o *ConnectStreamsOK) Error() string {
	return fmt.Sprintf("[POST /plugins/org.graylog.plugins.pipelineprocessor/system/pipelines/connections/to_pipeline][%d] connectStreamsOK  %+v", 200, o.Payload)
}

func (o *ConnectStreamsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
