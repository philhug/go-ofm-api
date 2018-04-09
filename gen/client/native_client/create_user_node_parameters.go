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

	models "github.com/philhug/go-ofm-api/gen/models"
)

// NewCreateUserNodeParams creates a new CreateUserNodeParams object
// with the default values initialized.
func NewCreateUserNodeParams() *CreateUserNodeParams {
	var ()
	return &CreateUserNodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUserNodeParamsWithTimeout creates a new CreateUserNodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateUserNodeParamsWithTimeout(timeout time.Duration) *CreateUserNodeParams {
	var ()
	return &CreateUserNodeParams{

		timeout: timeout,
	}
}

// NewCreateUserNodeParamsWithContext creates a new CreateUserNodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateUserNodeParamsWithContext(ctx context.Context) *CreateUserNodeParams {
	var ()
	return &CreateUserNodeParams{

		Context: ctx,
	}
}

// NewCreateUserNodeParamsWithHTTPClient creates a new CreateUserNodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateUserNodeParamsWithHTTPClient(client *http.Client) *CreateUserNodeParams {
	var ()
	return &CreateUserNodeParams{
		HTTPClient: client,
	}
}

/*CreateUserNodeParams contains all the parameters to send to the API endpoint
for the create user node operation typically these are written to a http.Request
*/
type CreateUserNodeParams struct {

	/*Body
	  Node definition

	*/
	Body *models.Node

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create user node params
func (o *CreateUserNodeParams) WithTimeout(timeout time.Duration) *CreateUserNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create user node params
func (o *CreateUserNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create user node params
func (o *CreateUserNodeParams) WithContext(ctx context.Context) *CreateUserNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create user node params
func (o *CreateUserNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create user node params
func (o *CreateUserNodeParams) WithHTTPClient(client *http.Client) *CreateUserNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create user node params
func (o *CreateUserNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create user node params
func (o *CreateUserNodeParams) WithBody(body *models.Node) *CreateUserNodeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create user node params
func (o *CreateUserNodeParams) SetBody(body *models.Node) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUserNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
