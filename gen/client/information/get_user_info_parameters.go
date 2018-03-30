// Code generated by go-swagger; DO NOT EDIT.

package information

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
)

// NewGetUserInfoParams creates a new GetUserInfoParams object
// with the default values initialized.
func NewGetUserInfoParams() *GetUserInfoParams {

	return &GetUserInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserInfoParamsWithTimeout creates a new GetUserInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserInfoParamsWithTimeout(timeout time.Duration) *GetUserInfoParams {

	return &GetUserInfoParams{

		timeout: timeout,
	}
}

// NewGetUserInfoParamsWithContext creates a new GetUserInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserInfoParamsWithContext(ctx context.Context) *GetUserInfoParams {

	return &GetUserInfoParams{

		Context: ctx,
	}
}

// NewGetUserInfoParamsWithHTTPClient creates a new GetUserInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserInfoParamsWithHTTPClient(client *http.Client) *GetUserInfoParams {

	return &GetUserInfoParams{
		HTTPClient: client,
	}
}

/*GetUserInfoParams contains all the parameters to send to the API endpoint
for the get user info operation typically these are written to a http.Request
*/
type GetUserInfoParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user info params
func (o *GetUserInfoParams) WithTimeout(timeout time.Duration) *GetUserInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user info params
func (o *GetUserInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user info params
func (o *GetUserInfoParams) WithContext(ctx context.Context) *GetUserInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user info params
func (o *GetUserInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user info params
func (o *GetUserInfoParams) WithHTTPClient(client *http.Client) *GetUserInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user info params
func (o *GetUserInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
