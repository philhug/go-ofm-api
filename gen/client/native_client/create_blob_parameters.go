// Code generated by go-swagger; DO NOT EDIT.

package native_client

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

// NewCreateBlobParams creates a new CreateBlobParams object
// with the default values initialized.
func NewCreateBlobParams() *CreateBlobParams {
	var ()
	return &CreateBlobParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBlobParamsWithTimeout creates a new CreateBlobParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateBlobParamsWithTimeout(timeout time.Duration) *CreateBlobParams {
	var ()
	return &CreateBlobParams{

		timeout: timeout,
	}
}

// NewCreateBlobParamsWithContext creates a new CreateBlobParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateBlobParamsWithContext(ctx context.Context) *CreateBlobParams {
	var ()
	return &CreateBlobParams{

		Context: ctx,
	}
}

// NewCreateBlobParamsWithHTTPClient creates a new CreateBlobParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateBlobParamsWithHTTPClient(client *http.Client) *CreateBlobParams {
	var ()
	return &CreateBlobParams{
		HTTPClient: client,
	}
}

/*CreateBlobParams contains all the parameters to send to the API endpoint
for the create blob operation typically these are written to a http.Request
*/
type CreateBlobParams struct {

	/*Body
	  Binary data

	*/
	Body *strfmt.Base64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create blob params
func (o *CreateBlobParams) WithTimeout(timeout time.Duration) *CreateBlobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create blob params
func (o *CreateBlobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create blob params
func (o *CreateBlobParams) WithContext(ctx context.Context) *CreateBlobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create blob params
func (o *CreateBlobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create blob params
func (o *CreateBlobParams) WithHTTPClient(client *http.Client) *CreateBlobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create blob params
func (o *CreateBlobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create blob params
func (o *CreateBlobParams) WithBody(body *strfmt.Base64) *CreateBlobParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create blob params
func (o *CreateBlobParams) SetBody(body *strfmt.Base64) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBlobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
