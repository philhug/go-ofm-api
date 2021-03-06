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

// NewGetMultipleNodesParams creates a new GetMultipleNodesParams object
// with the default values initialized.
func NewGetMultipleNodesParams() *GetMultipleNodesParams {
	var ()
	return &GetMultipleNodesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMultipleNodesParamsWithTimeout creates a new GetMultipleNodesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMultipleNodesParamsWithTimeout(timeout time.Duration) *GetMultipleNodesParams {
	var ()
	return &GetMultipleNodesParams{

		timeout: timeout,
	}
}

// NewGetMultipleNodesParamsWithContext creates a new GetMultipleNodesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMultipleNodesParamsWithContext(ctx context.Context) *GetMultipleNodesParams {
	var ()
	return &GetMultipleNodesParams{

		Context: ctx,
	}
}

// NewGetMultipleNodesParamsWithHTTPClient creates a new GetMultipleNodesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMultipleNodesParamsWithHTTPClient(client *http.Client) *GetMultipleNodesParams {
	var ()
	return &GetMultipleNodesParams{
		HTTPClient: client,
	}
}

/*GetMultipleNodesParams contains all the parameters to send to the API endpoint
for the get multiple nodes operation typically these are written to a http.Request
*/
type GetMultipleNodesParams struct {

	/*Db
	  Database name (e.g. oad)

	*/
	Db string
	/*Nodes
	  Unique node identifiers

	*/
	Nodes string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get multiple nodes params
func (o *GetMultipleNodesParams) WithTimeout(timeout time.Duration) *GetMultipleNodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get multiple nodes params
func (o *GetMultipleNodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get multiple nodes params
func (o *GetMultipleNodesParams) WithContext(ctx context.Context) *GetMultipleNodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get multiple nodes params
func (o *GetMultipleNodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get multiple nodes params
func (o *GetMultipleNodesParams) WithHTTPClient(client *http.Client) *GetMultipleNodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get multiple nodes params
func (o *GetMultipleNodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDb adds the db to the get multiple nodes params
func (o *GetMultipleNodesParams) WithDb(db string) *GetMultipleNodesParams {
	o.SetDb(db)
	return o
}

// SetDb adds the db to the get multiple nodes params
func (o *GetMultipleNodesParams) SetDb(db string) {
	o.Db = db
}

// WithNodes adds the nodes to the get multiple nodes params
func (o *GetMultipleNodesParams) WithNodes(nodes string) *GetMultipleNodesParams {
	o.SetNodes(nodes)
	return o
}

// SetNodes adds the nodes to the get multiple nodes params
func (o *GetMultipleNodesParams) SetNodes(nodes string) {
	o.Nodes = nodes
}

// WriteToRequest writes these params to a swagger request
func (o *GetMultipleNodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param db
	if err := r.SetPathParam("db", o.Db); err != nil {
		return err
	}

	// query param nodes
	qrNodes := o.Nodes
	qNodes := qrNodes
	if qNodes != "" {
		if err := r.SetQueryParam("nodes", qNodes); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
