// Code generated by go-swagger; DO NOT EDIT.

package dashboards_dashboard_id_widgets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UpdateWidgetReader is a Reader for the UpdateWidget structure.
type UpdateWidgetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateWidgetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateWidgetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewUpdateWidgetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateWidgetOK creates a UpdateWidgetOK with default headers values
func NewUpdateWidgetOK() *UpdateWidgetOK {
	return &UpdateWidgetOK{}
}

/*UpdateWidgetOK handles this case with default header values.

No response was specified
*/
type UpdateWidgetOK struct {
}

func (o *UpdateWidgetOK) Error() string {
	return fmt.Sprintf("[PUT /dashboards/{dashboardId}/widgets/{widgetId}][%d] updateWidgetOK ", 200)
}

func (o *UpdateWidgetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateWidgetNotFound creates a UpdateWidgetNotFound with default headers values
func NewUpdateWidgetNotFound() *UpdateWidgetNotFound {
	return &UpdateWidgetNotFound{}
}

/*UpdateWidgetNotFound handles this case with default header values.

Widget not found.
*/
type UpdateWidgetNotFound struct {
}

func (o *UpdateWidgetNotFound) Error() string {
	return fmt.Sprintf("[PUT /dashboards/{dashboardId}/widgets/{widgetId}][%d] updateWidgetNotFound ", 404)
}

func (o *UpdateWidgetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}