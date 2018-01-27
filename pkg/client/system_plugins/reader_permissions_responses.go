// Code generated by go-swagger; DO NOT EDIT.

package system_plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// ReaderPermissionsReader is a Reader for the ReaderPermissions structure.
type ReaderPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReaderPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReaderPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReaderPermissionsOK creates a ReaderPermissionsOK with default headers values
func NewReaderPermissionsOK() *ReaderPermissionsOK {
	return &ReaderPermissionsOK{}
}

/*ReaderPermissionsOK handles this case with default header values.

No response was specified
*/
type ReaderPermissionsOK struct {
	Payload *models.ReaderPermissionResponse
}

func (o *ReaderPermissionsOK) Error() string {
	return fmt.Sprintf("[GET /system/permissions/reader/{username}][%d] readerPermissionsOK  %+v", 200, o.Payload)
}

func (o *ReaderPermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReaderPermissionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
