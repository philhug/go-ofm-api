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
)

// NewChangesDbParams creates a new ChangesDbParams object
// with the default values initialized.
func NewChangesDbParams() *ChangesDbParams {
	var ()
	return &ChangesDbParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangesDbParamsWithTimeout creates a new ChangesDbParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangesDbParamsWithTimeout(timeout time.Duration) *ChangesDbParams {
	var ()
	return &ChangesDbParams{

		timeout: timeout,
	}
}

// NewChangesDbParamsWithContext creates a new ChangesDbParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangesDbParamsWithContext(ctx context.Context) *ChangesDbParams {
	var ()
	return &ChangesDbParams{

		Context: ctx,
	}
}

// NewChangesDbParamsWithHTTPClient creates a new ChangesDbParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangesDbParamsWithHTTPClient(client *http.Client) *ChangesDbParams {
	var ()
	return &ChangesDbParams{
		HTTPClient: client,
	}
}

/*ChangesDbParams contains all the parameters to send to the API endpoint
for the changes db operation typically these are written to a http.Request
*/
type ChangesDbParams struct {

	/*Db
	  Database name (e.g. "OAD Pending Changes")

	*/
	Db string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the changes db params
func (o *ChangesDbParams) WithTimeout(timeout time.Duration) *ChangesDbParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the changes db params
func (o *ChangesDbParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the changes db params
func (o *ChangesDbParams) WithContext(ctx context.Context) *ChangesDbParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the changes db params
func (o *ChangesDbParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the changes db params
func (o *ChangesDbParams) WithHTTPClient(client *http.Client) *ChangesDbParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the changes db params
func (o *ChangesDbParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDb adds the db to the changes db params
func (o *ChangesDbParams) WithDb(db string) *ChangesDbParams {
	o.SetDb(db)
	return o
}

// SetDb adds the db to the changes db params
func (o *ChangesDbParams) SetDb(db string) {
	o.Db = db
}

// WriteToRequest writes these params to a swagger request
func (o *ChangesDbParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param db
	if err := r.SetPathParam("db", o.Db); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
