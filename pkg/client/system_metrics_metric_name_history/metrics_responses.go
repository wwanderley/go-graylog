// Code generated by go-swagger; DO NOT EDIT.

package system_metrics_metric_name_history

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// MetricsReader is a Reader for the Metrics structure.
type MetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewMetricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMetricsOK creates a MetricsOK with default headers values
func NewMetricsOK() *MetricsOK {
	return &MetricsOK{}
}

/*MetricsOK handles this case with default header values.

No response was specified
*/
type MetricsOK struct {
	Payload MetricsOKBody
}

func (o *MetricsOK) Error() string {
	return fmt.Sprintf("[GET /system/metrics][%d] metricsOK  %+v", 200, o.Payload)
}

func (o *MetricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*MetricsOKBody metrics o k body
swagger:model MetricsOKBody
*/
type MetricsOKBody interface{}
