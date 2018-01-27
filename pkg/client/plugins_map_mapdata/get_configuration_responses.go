// Code generated by go-swagger; DO NOT EDIT.

package plugins_map_mapdata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetConfigurationReader is a Reader for the GetConfiguration structure.
type GetConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetConfigurationNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetConfigurationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetConfigurationOK creates a GetConfigurationOK with default headers values
func NewGetConfigurationOK() *GetConfigurationOK {
	return &GetConfigurationOK{}
}

/*GetConfigurationOK handles this case with default header values.

No response was specified
*/
type GetConfigurationOK struct {
}

func (o *GetConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /plugins/org.graylog.plugins.collector/{collectorId}][%d] getConfigurationOK ", 200)
}

func (o *GetConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetConfigurationNotModified creates a GetConfigurationNotModified with default headers values
func NewGetConfigurationNotModified() *GetConfigurationNotModified {
	return &GetConfigurationNotModified{}
}

/*GetConfigurationNotModified handles this case with default header values.

Configuration didn't update.
*/
type GetConfigurationNotModified struct {
}

func (o *GetConfigurationNotModified) Error() string {
	return fmt.Sprintf("[GET /plugins/org.graylog.plugins.collector/{collectorId}][%d] getConfigurationNotModified ", 304)
}

func (o *GetConfigurationNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetConfigurationBadRequest creates a GetConfigurationBadRequest with default headers values
func NewGetConfigurationBadRequest() *GetConfigurationBadRequest {
	return &GetConfigurationBadRequest{}
}

/*GetConfigurationBadRequest handles this case with default header values.

Invalid ObjectId.
*/
type GetConfigurationBadRequest struct {
}

func (o *GetConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[GET /plugins/org.graylog.plugins.collector/{collectorId}][%d] getConfigurationBadRequest ", 400)
}

func (o *GetConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetConfigurationNotFound creates a GetConfigurationNotFound with default headers values
func NewGetConfigurationNotFound() *GetConfigurationNotFound {
	return &GetConfigurationNotFound{}
}

/*GetConfigurationNotFound handles this case with default header values.

Collector not found.
*/
type GetConfigurationNotFound struct {
}

func (o *GetConfigurationNotFound) Error() string {
	return fmt.Sprintf("[GET /plugins/org.graylog.plugins.collector/{collectorId}][%d] getConfigurationNotFound ", 404)
}

func (o *GetConfigurationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
