// Code generated by go-swagger; DO NOT EDIT.

package dashboards_dashboard_id_widgets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// GetWidgetReader is a Reader for the GetWidget structure.
type GetWidgetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWidgetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetWidgetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetWidgetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetWidgetOK creates a GetWidgetOK with default headers values
func NewGetWidgetOK() *GetWidgetOK {
	return &GetWidgetOK{}
}

/*GetWidgetOK handles this case with default header values.

No response was specified
*/
type GetWidgetOK struct {
	Payload *models.WidgetSummary
}

func (o *GetWidgetOK) Error() string {
	return fmt.Sprintf("[GET /dashboards/{dashboardId}/widgets/{widgetId}][%d] getWidgetOK  %+v", 200, o.Payload)
}

func (o *GetWidgetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WidgetSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWidgetNotFound creates a GetWidgetNotFound with default headers values
func NewGetWidgetNotFound() *GetWidgetNotFound {
	return &GetWidgetNotFound{}
}

/*GetWidgetNotFound handles this case with default header values.

Widget not found.
*/
type GetWidgetNotFound struct {
}

func (o *GetWidgetNotFound) Error() string {
	return fmt.Sprintf("[GET /dashboards/{dashboardId}/widgets/{widgetId}][%d] getWidgetNotFound ", 404)
}

func (o *GetWidgetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
