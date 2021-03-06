// Code generated by go-swagger; DO NOT EDIT.

package plugins_map_mapdata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteConfigurationReader is a Reader for the DeleteConfiguration structure.
type DeleteConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteConfigurationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteConfigurationOK creates a DeleteConfigurationOK with default headers values
func NewDeleteConfigurationOK() *DeleteConfigurationOK {
	return &DeleteConfigurationOK{}
}

/*DeleteConfigurationOK handles this case with default header values.

No response was specified
*/
type DeleteConfigurationOK struct {
}

func (o *DeleteConfigurationOK) Error() string {
	return fmt.Sprintf("[DELETE /plugins/org.graylog.plugins.collector/configurations/{id}][%d] deleteConfigurationOK ", 200)
}

func (o *DeleteConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteConfigurationBadRequest creates a DeleteConfigurationBadRequest with default headers values
func NewDeleteConfigurationBadRequest() *DeleteConfigurationBadRequest {
	return &DeleteConfigurationBadRequest{}
}

/*DeleteConfigurationBadRequest handles this case with default header values.

Invalid ObjectId.
*/
type DeleteConfigurationBadRequest struct {
}

func (o *DeleteConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /plugins/org.graylog.plugins.collector/configurations/{id}][%d] deleteConfigurationBadRequest ", 400)
}

func (o *DeleteConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteConfigurationNotFound creates a DeleteConfigurationNotFound with default headers values
func NewDeleteConfigurationNotFound() *DeleteConfigurationNotFound {
	return &DeleteConfigurationNotFound{}
}

/*DeleteConfigurationNotFound handles this case with default header values.

Configuration not found.
*/
type DeleteConfigurationNotFound struct {
}

func (o *DeleteConfigurationNotFound) Error() string {
	return fmt.Sprintf("[DELETE /plugins/org.graylog.plugins.collector/configurations/{id}][%d] deleteConfigurationNotFound ", 404)
}

func (o *DeleteConfigurationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
