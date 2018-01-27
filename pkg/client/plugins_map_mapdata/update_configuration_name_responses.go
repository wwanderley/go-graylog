// Code generated by go-swagger; DO NOT EDIT.

package plugins_map_mapdata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UpdateConfigurationNameReader is a Reader for the UpdateConfigurationName structure.
type UpdateConfigurationNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateConfigurationNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateConfigurationNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateConfigurationNameOK creates a UpdateConfigurationNameOK with default headers values
func NewUpdateConfigurationNameOK() *UpdateConfigurationNameOK {
	return &UpdateConfigurationNameOK{}
}

/*UpdateConfigurationNameOK handles this case with default header values.

No response was specified
*/
type UpdateConfigurationNameOK struct {
	Payload UpdateConfigurationNameOKBody
}

func (o *UpdateConfigurationNameOK) Error() string {
	return fmt.Sprintf("[PUT /plugins/org.graylog.plugins.collector/configurations/{id}/name][%d] updateConfigurationNameOK  %+v", 200, o.Payload)
}

func (o *UpdateConfigurationNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateConfigurationNameOKBody update configuration name o k body
swagger:model UpdateConfigurationNameOKBody
*/
type UpdateConfigurationNameOKBody interface{}
