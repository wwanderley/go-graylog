// Code generated by go-swagger; DO NOT EDIT.

package system_indices_ranges

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// BulkUpdatePatternsReader is a Reader for the BulkUpdatePatterns structure.
type BulkUpdatePatternsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BulkUpdatePatternsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBulkUpdatePatternsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBulkUpdatePatternsOK creates a BulkUpdatePatternsOK with default headers values
func NewBulkUpdatePatternsOK() *BulkUpdatePatternsOK {
	return &BulkUpdatePatternsOK{}
}

/*BulkUpdatePatternsOK handles this case with default header values.

No response was specified
*/
type BulkUpdatePatternsOK struct {
}

func (o *BulkUpdatePatternsOK) Error() string {
	return fmt.Sprintf("[PUT /system/grok][%d] bulkUpdatePatternsOK ", 200)
}

func (o *BulkUpdatePatternsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
