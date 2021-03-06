// Code generated by go-swagger; DO NOT EDIT.

package search_saved

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// StatsRelativeReader is a Reader for the StatsRelative structure.
type StatsRelativeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatsRelativeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStatsRelativeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewStatsRelativeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStatsRelativeOK creates a StatsRelativeOK with default headers values
func NewStatsRelativeOK() *StatsRelativeOK {
	return &StatsRelativeOK{}
}

/*StatsRelativeOK handles this case with default header values.

No response was specified
*/
type StatsRelativeOK struct {
	Payload *models.FieldStatsResult
}

func (o *StatsRelativeOK) Error() string {
	return fmt.Sprintf("[GET /search/universal/relative/stats][%d] statsRelativeOK  %+v", 200, o.Payload)
}

func (o *StatsRelativeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FieldStatsResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatsRelativeBadRequest creates a StatsRelativeBadRequest with default headers values
func NewStatsRelativeBadRequest() *StatsRelativeBadRequest {
	return &StatsRelativeBadRequest{}
}

/*StatsRelativeBadRequest handles this case with default header values.

Field is not of numeric type.
*/
type StatsRelativeBadRequest struct {
}

func (o *StatsRelativeBadRequest) Error() string {
	return fmt.Sprintf("[GET /search/universal/relative/stats][%d] statsRelativeBadRequest ", 400)
}

func (o *StatsRelativeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
