// Code generated by go-swagger; DO NOT EDIT.

package system_lbstatus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// GetLdapSettingsReader is a Reader for the GetLdapSettings structure.
type GetLdapSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLdapSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLdapSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLdapSettingsOK creates a GetLdapSettingsOK with default headers values
func NewGetLdapSettingsOK() *GetLdapSettingsOK {
	return &GetLdapSettingsOK{}
}

/*GetLdapSettingsOK handles this case with default header values.

No response was specified
*/
type GetLdapSettingsOK struct {
	Payload *models.LdapSettingsResponse
}

func (o *GetLdapSettingsOK) Error() string {
	return fmt.Sprintf("[GET /system/ldap/settings][%d] getLdapSettingsOK  %+v", 200, o.Payload)
}

func (o *GetLdapSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LdapSettingsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
