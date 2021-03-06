// Code generated by go-swagger; DO NOT EDIT.

package dashboards_dashboard_id_widgets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UpdateDescriptionReader is a Reader for the UpdateDescription structure.
type UpdateDescriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDescriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateDescriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewUpdateDescriptionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateDescriptionOK creates a UpdateDescriptionOK with default headers values
func NewUpdateDescriptionOK() *UpdateDescriptionOK {
	return &UpdateDescriptionOK{}
}

/*UpdateDescriptionOK handles this case with default header values.

No response was specified
*/
type UpdateDescriptionOK struct {
}

func (o *UpdateDescriptionOK) Error() string {
	return fmt.Sprintf("[PUT /dashboards/{dashboardId}/widgets/{widgetId}/description][%d] updateDescriptionOK ", 200)
}

func (o *UpdateDescriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateDescriptionNotFound creates a UpdateDescriptionNotFound with default headers values
func NewUpdateDescriptionNotFound() *UpdateDescriptionNotFound {
	return &UpdateDescriptionNotFound{}
}

/*UpdateDescriptionNotFound handles this case with default header values.

Widget not found.
*/
type UpdateDescriptionNotFound struct {
}

func (o *UpdateDescriptionNotFound) Error() string {
	return fmt.Sprintf("[PUT /dashboards/{dashboardId}/widgets/{widgetId}/description][%d] updateDescriptionNotFound ", 404)
}

func (o *UpdateDescriptionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
