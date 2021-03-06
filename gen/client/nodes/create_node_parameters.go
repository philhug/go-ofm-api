// Code generated by go-swagger; DO NOT EDIT.

package nodes

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

// NewCreateNodeParams creates a new CreateNodeParams object
// with the default values initialized.
func NewCreateNodeParams() *CreateNodeParams {
	var ()
	return &CreateNodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNodeParamsWithTimeout creates a new CreateNodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateNodeParamsWithTimeout(timeout time.Duration) *CreateNodeParams {
	var ()
	return &CreateNodeParams{

		timeout: timeout,
	}
}

// NewCreateNodeParamsWithContext creates a new CreateNodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateNodeParamsWithContext(ctx context.Context) *CreateNodeParams {
	var ()
	return &CreateNodeParams{

		Context: ctx,
	}
}

// NewCreateNodeParamsWithHTTPClient creates a new CreateNodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateNodeParamsWithHTTPClient(client *http.Client) *CreateNodeParams {
	var ()
	return &CreateNodeParams{
		HTTPClient: client,
	}
}

/*CreateNodeParams contains all the parameters to send to the API endpoint
for the create node operation typically these are written to a http.Request
*/
type CreateNodeParams struct {

	/*Body
	  Node definition

	*/
	Body *models.Node
	/*Db
	  Database name (e.g. "OAD Pending Changes")

	*/
	Db string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create node params
func (o *CreateNodeParams) WithTimeout(timeout time.Duration) *CreateNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create node params
func (o *CreateNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create node params
func (o *CreateNodeParams) WithContext(ctx context.Context) *CreateNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create node params
func (o *CreateNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create node params
func (o *CreateNodeParams) WithHTTPClient(client *http.Client) *CreateNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create node params
func (o *CreateNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create node params
func (o *CreateNodeParams) WithBody(body *models.Node) *CreateNodeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create node params
func (o *CreateNodeParams) SetBody(body *models.Node) {
	o.Body = body
}

// WithDb adds the db to the create node params
func (o *CreateNodeParams) WithDb(db string) *CreateNodeParams {
	o.SetDb(db)
	return o
}

// SetDb adds the db to the create node params
func (o *CreateNodeParams) SetDb(db string) {
	o.Db = db
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param db
	if err := r.SetPathParam("db", o.Db); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
