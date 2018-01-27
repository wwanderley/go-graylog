// Code generated by go-swagger; DO NOT EDIT.

package system_fields

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CycleReader is a Reader for the Cycle structure.
type CycleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CycleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCycleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCycleOK creates a CycleOK with default headers values
func NewCycleOK() *CycleOK {
	return &CycleOK{}
}

/*CycleOK handles this case with default header values.

No response was specified
*/
type CycleOK struct {
}

func (o *CycleOK) Error() string {
	return fmt.Sprintf("[POST /system/deflector/{indexSetId}/cycle][%d] cycleOK ", 200)
}

func (o *CycleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
