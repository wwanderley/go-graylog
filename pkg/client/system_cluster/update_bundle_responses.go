// Code generated by go-swagger; DO NOT EDIT.

package system_cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UpdateBundleReader is a Reader for the UpdateBundle structure.
type UpdateBundleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateBundleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateBundleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateBundleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUpdateBundleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateBundleOK creates a UpdateBundleOK with default headers values
func NewUpdateBundleOK() *UpdateBundleOK {
	return &UpdateBundleOK{}
}

/*UpdateBundleOK handles this case with default header values.

No response was specified
*/
type UpdateBundleOK struct {
}

func (o *UpdateBundleOK) Error() string {
	return fmt.Sprintf("[PUT /system/bundles/{bundleId}][%d] updateBundleOK ", 200)
}

func (o *UpdateBundleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateBundleBadRequest creates a UpdateBundleBadRequest with default headers values
func NewUpdateBundleBadRequest() *UpdateBundleBadRequest {
	return &UpdateBundleBadRequest{}
}

/*UpdateBundleBadRequest handles this case with default header values.

Missing or invalid content pack
*/
type UpdateBundleBadRequest struct {
}

func (o *UpdateBundleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /system/bundles/{bundleId}][%d] updateBundleBadRequest ", 400)
}

func (o *UpdateBundleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateBundleInternalServerError creates a UpdateBundleInternalServerError with default headers values
func NewUpdateBundleInternalServerError() *UpdateBundleInternalServerError {
	return &UpdateBundleInternalServerError{}
}

/*UpdateBundleInternalServerError handles this case with default header values.

Error while updating content pack
*/
type UpdateBundleInternalServerError struct {
}

func (o *UpdateBundleInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /system/bundles/{bundleId}][%d] updateBundleInternalServerError ", 500)
}

func (o *UpdateBundleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
