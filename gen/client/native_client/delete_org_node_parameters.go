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

// NewDeleteOrgNodeParams creates a new DeleteOrgNodeParams object
// with the default values initialized.
func NewDeleteOrgNodeParams() *DeleteOrgNodeParams {
	var ()
	return &DeleteOrgNodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOrgNodeParamsWithTimeout creates a new DeleteOrgNodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteOrgNodeParamsWithTimeout(timeout time.Duration) *DeleteOrgNodeParams {
	var ()
	return &DeleteOrgNodeParams{

		timeout: timeout,
	}
}

// NewDeleteOrgNodeParamsWithContext creates a new DeleteOrgNodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteOrgNodeParamsWithContext(ctx context.Context) *DeleteOrgNodeParams {
	var ()
	return &DeleteOrgNodeParams{

		Context: ctx,
	}
}

// NewDeleteOrgNodeParamsWithHTTPClient creates a new DeleteOrgNodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteOrgNodeParamsWithHTTPClient(client *http.Client) *DeleteOrgNodeParams {
	var ()
	return &DeleteOrgNodeParams{
		HTTPClient: client,
	}
}

/*DeleteOrgNodeParams contains all the parameters to send to the API endpoint
for the delete org node operation typically these are written to a http.Request
*/
type DeleteOrgNodeParams struct {

	/*ID
	  Unique identifier

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete org node params
func (o *DeleteOrgNodeParams) WithTimeout(timeout time.Duration) *DeleteOrgNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete org node params
func (o *DeleteOrgNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete org node params
func (o *DeleteOrgNodeParams) WithContext(ctx context.Context) *DeleteOrgNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete org node params
func (o *DeleteOrgNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete org node params
func (o *DeleteOrgNodeParams) WithHTTPClient(client *http.Client) *DeleteOrgNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete org node params
func (o *DeleteOrgNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete org node params
func (o *DeleteOrgNodeParams) WithID(id string) *DeleteOrgNodeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete org node params
func (o *DeleteOrgNodeParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOrgNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}