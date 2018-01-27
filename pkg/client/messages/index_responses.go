// Code generated by go-swagger; DO NOT EDIT.

package messages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// IndexReader is a Reader for the Index structure.
type IndexReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IndexReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIndexOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIndexOK creates a IndexOK with default headers values
func NewIndexOK() *IndexOK {
	return &IndexOK{}
}

/*IndexOK handles this case with default header values.

No response was specified
*/
type IndexOK struct {
	Payload *models.IndexerOverview
}

func (o *IndexOK) Error() string {
	return fmt.Sprintf("[GET /system/indexer/overview/{indexSetId}][%d] indexOK  %+v", 200, o.Payload)
}

func (o *IndexOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IndexerOverview)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
