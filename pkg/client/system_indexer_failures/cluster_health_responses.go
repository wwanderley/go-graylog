// Code generated by go-swagger; DO NOT EDIT.

package system_indexer_failures

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// ClusterHealthReader is a Reader for the ClusterHealth structure.
type ClusterHealthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterHealthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewClusterHealthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewClusterHealthOK creates a ClusterHealthOK with default headers values
func NewClusterHealthOK() *ClusterHealthOK {
	return &ClusterHealthOK{}
}

/*ClusterHealthOK handles this case with default header values.

No response was specified
*/
type ClusterHealthOK struct {
	Payload *models.ClusterHealth
}

func (o *ClusterHealthOK) Error() string {
	return fmt.Sprintf("[GET /system/indexer/cluster/health][%d] clusterHealthOK  %+v", 200, o.Payload)
}

func (o *ClusterHealthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
