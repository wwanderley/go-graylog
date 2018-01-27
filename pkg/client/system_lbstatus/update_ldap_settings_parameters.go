// Code generated by go-swagger; DO NOT EDIT.

package system_lbstatus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// NewUpdateLdapSettingsParams creates a new UpdateLdapSettingsParams object
// with the default values initialized.
func NewUpdateLdapSettingsParams() *UpdateLdapSettingsParams {
	var ()
	return &UpdateLdapSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateLdapSettingsParamsWithTimeout creates a new UpdateLdapSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateLdapSettingsParamsWithTimeout(timeout time.Duration) *UpdateLdapSettingsParams {
	var ()
	return &UpdateLdapSettingsParams{

		timeout: timeout,
	}
}

// NewUpdateLdapSettingsParamsWithContext creates a new UpdateLdapSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateLdapSettingsParamsWithContext(ctx context.Context) *UpdateLdapSettingsParams {
	var ()
	return &UpdateLdapSettingsParams{

		Context: ctx,
	}
}

// NewUpdateLdapSettingsParamsWithHTTPClient creates a new UpdateLdapSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateLdapSettingsParamsWithHTTPClient(client *http.Client) *UpdateLdapSettingsParams {
	var ()
	return &UpdateLdapSettingsParams{
		HTTPClient: client,
	}
}

/*UpdateLdapSettingsParams contains all the parameters to send to the API endpoint
for the update ldap settings operation typically these are written to a http.Request
*/
type UpdateLdapSettingsParams struct {

	/*JSONBody*/
	JSONBody *models.LdapSettingsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update ldap settings params
func (o *UpdateLdapSettingsParams) WithTimeout(timeout time.Duration) *UpdateLdapSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update ldap settings params
func (o *UpdateLdapSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update ldap settings params
func (o *UpdateLdapSettingsParams) WithContext(ctx context.Context) *UpdateLdapSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update ldap settings params
func (o *UpdateLdapSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update ldap settings params
func (o *UpdateLdapSettingsParams) WithHTTPClient(client *http.Client) *UpdateLdapSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update ldap settings params
func (o *UpdateLdapSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJSONBody adds the jSONBody to the update ldap settings params
func (o *UpdateLdapSettingsParams) WithJSONBody(jSONBody *models.LdapSettingsRequest) *UpdateLdapSettingsParams {
	o.SetJSONBody(jSONBody)
	return o
}

// SetJSONBody adds the jsonBody to the update ldap settings params
func (o *UpdateLdapSettingsParams) SetJSONBody(jSONBody *models.LdapSettingsRequest) {
	o.JSONBody = jSONBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateLdapSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.JSONBody != nil {
		if err := r.SetBodyParam(o.JSONBody); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
