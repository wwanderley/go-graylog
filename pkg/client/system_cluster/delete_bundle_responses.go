// Code generated by go-swagger; DO NOT EDIT.

package system_cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteBundleReader is a Reader for the DeleteBundle structure.
type DeleteBundleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBundleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteBundleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteBundleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteBundleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteBundleOK creates a DeleteBundleOK with default headers values
func NewDeleteBundleOK() *DeleteBundleOK {
	return &DeleteBundleOK{}
}

/*DeleteBundleOK handles this case with default header values.

No response was specified
*/
type DeleteBundleOK struct {
}

func (o *DeleteBundleOK) Error() string {
	return fmt.Sprintf("[DELETE /system/bundles/{bundleId}][%d] deleteBundleOK ", 200)
}

func (o *DeleteBundleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteBundleNotFound creates a DeleteBundleNotFound with default headers values
func NewDeleteBundleNotFound() *DeleteBundleNotFound {
	return &DeleteBundleNotFound{}
}

/*DeleteBundleNotFound handles this case with default header values.

Missing or invalid content pack
*/
type DeleteBundleNotFound struct {
}

func (o *DeleteBundleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /system/bundles/{bundleId}][%d] deleteBundleNotFound ", 404)
}

func (o *DeleteBundleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteBundleInternalServerError creates a DeleteBundleInternalServerError with default headers values
func NewDeleteBundleInternalServerError() *DeleteBundleInternalServerError {
	return &DeleteBundleInternalServerError{}
}

/*DeleteBundleInternalServerError handles this case with default header values.

Error while applying content pack
*/
type DeleteBundleInternalServerError struct {
}

func (o *DeleteBundleInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /system/bundles/{bundleId}][%d] deleteBundleInternalServerError ", 500)
}

func (o *DeleteBundleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
