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

// NewSearchUserNodeParams creates a new SearchUserNodeParams object
// with the default values initialized.
func NewSearchUserNodeParams() *SearchUserNodeParams {

	return &SearchUserNodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSearchUserNodeParamsWithTimeout creates a new SearchUserNodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSearchUserNodeParamsWithTimeout(timeout time.Duration) *SearchUserNodeParams {

	return &SearchUserNodeParams{

		timeout: timeout,
	}
}

// NewSearchUserNodeParamsWithContext creates a new SearchUserNodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewSearchUserNodeParamsWithContext(ctx context.Context) *SearchUserNodeParams {

	return &SearchUserNodeParams{

		Context: ctx,
	}
}

// NewSearchUserNodeParamsWithHTTPClient creates a new SearchUserNodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSearchUserNodeParamsWithHTTPClient(client *http.Client) *SearchUserNodeParams {

	return &SearchUserNodeParams{
		HTTPClient: client,
	}
}

/*SearchUserNodeParams contains all the parameters to send to the API endpoint
for the search user node operation typically these are written to a http.Request
*/
type SearchUserNodeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the search user node params
func (o *SearchUserNodeParams) WithTimeout(timeout time.Duration) *SearchUserNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search user node params
func (o *SearchUserNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search user node params
func (o *SearchUserNodeParams) WithContext(ctx context.Context) *SearchUserNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search user node params
func (o *SearchUserNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search user node params
func (o *SearchUserNodeParams) WithHTTPClient(client *http.Client) *SearchUserNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search user node params
func (o *SearchUserNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SearchUserNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}