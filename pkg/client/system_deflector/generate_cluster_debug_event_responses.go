// Code generated by go-swagger; DO NOT EDIT.

package system_deflector

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GenerateClusterDebugEventReader is a Reader for the GenerateClusterDebugEvent structure.
type GenerateClusterDebugEventReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenerateClusterDebugEventReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGenerateClusterDebugEventOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGenerateClusterDebugEventOK creates a GenerateClusterDebugEventOK with default headers values
func NewGenerateClusterDebugEventOK() *GenerateClusterDebugEventOK {
	return &GenerateClusterDebugEventOK{}
}

/*GenerateClusterDebugEventOK handles this case with default header values.

No response was specified
*/
type GenerateClusterDebugEventOK struct {
}

func (o *GenerateClusterDebugEventOK) Error() string {
	return fmt.Sprintf("[POST /system/debug/events/cluster][%d] generateClusterDebugEventOK ", 200)
}

func (o *GenerateClusterDebugEventOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
