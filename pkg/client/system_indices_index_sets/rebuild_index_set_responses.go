// Code generated by go-swagger; DO NOT EDIT.

package system_indices_index_sets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// RebuildIndexSetReader is a Reader for the RebuildIndexSet structure.
type RebuildIndexSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RebuildIndexSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRebuildIndexSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewRebuildIndexSetAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRebuildIndexSetOK creates a RebuildIndexSetOK with default headers values
func NewRebuildIndexSetOK() *RebuildIndexSetOK {
	return &RebuildIndexSetOK{}
}

/*RebuildIndexSetOK handles this case with default header values.

No response was specified
*/
type RebuildIndexSetOK struct {
}

func (o *RebuildIndexSetOK) Error() string {
	return fmt.Sprintf("[POST /system/indices/ranges/index_set/{indexSetId}/rebuild][%d] rebuildIndexSetOK ", 200)
}

func (o *RebuildIndexSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRebuildIndexSetAccepted creates a RebuildIndexSetAccepted with default headers values
func NewRebuildIndexSetAccepted() *RebuildIndexSetAccepted {
	return &RebuildIndexSetAccepted{}
}

/*RebuildIndexSetAccepted handles this case with default header values.

Rebuild/sync systemjob triggered.
*/
type RebuildIndexSetAccepted struct {
}

func (o *RebuildIndexSetAccepted) Error() string {
	return fmt.Sprintf("[POST /system/indices/ranges/index_set/{indexSetId}/rebuild][%d] rebuildIndexSetAccepted ", 202)
}

func (o *RebuildIndexSetAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}