// Code generated by go-swagger; DO NOT EDIT.

package plugins_map_mapdata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CopyOutputReader is a Reader for the CopyOutput structure.
type CopyOutputReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CopyOutputReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCopyOutputOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCopyOutputBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCopyOutputNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCopyOutputOK creates a CopyOutputOK with default headers values
func NewCopyOutputOK() *CopyOutputOK {
	return &CopyOutputOK{}
}

/*CopyOutputOK handles this case with default header values.

No response was specified
*/
type CopyOutputOK struct {
}

func (o *CopyOutputOK) Error() string {
	return fmt.Sprintf("[POST /plugins/org.graylog.plugins.collector/configurations/{id}/outputs/{outputId}/{name}][%d] copyOutputOK ", 200)
}

func (o *CopyOutputOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCopyOutputBadRequest creates a CopyOutputBadRequest with default headers values
func NewCopyOutputBadRequest() *CopyOutputBadRequest {
	return &CopyOutputBadRequest{}
}

/*CopyOutputBadRequest handles this case with default header values.

Invalid ObjectId.
*/
type CopyOutputBadRequest struct {
}

func (o *CopyOutputBadRequest) Error() string {
	return fmt.Sprintf("[POST /plugins/org.graylog.plugins.collector/configurations/{id}/outputs/{outputId}/{name}][%d] copyOutputBadRequest ", 400)
}

func (o *CopyOutputBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCopyOutputNotFound creates a CopyOutputNotFound with default headers values
func NewCopyOutputNotFound() *CopyOutputNotFound {
	return &CopyOutputNotFound{}
}

/*CopyOutputNotFound handles this case with default header values.

Configuration or Output not found.
*/
type CopyOutputNotFound struct {
}

func (o *CopyOutputNotFound) Error() string {
	return fmt.Sprintf("[POST /plugins/org.graylog.plugins.collector/configurations/{id}/outputs/{outputId}/{name}][%d] copyOutputNotFound ", 404)
}

func (o *CopyOutputNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
