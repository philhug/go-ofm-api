// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetMultipleNodesParams creates a new GetMultipleNodesParams object
// no default values defined in spec.
func NewGetMultipleNodesParams() GetMultipleNodesParams {

	return GetMultipleNodesParams{}
}

// GetMultipleNodesParams contains all the bound params for the get multiple nodes operation
// typically these are obtained from a http.Request
//
// swagger:parameters getMultipleNodes
type GetMultipleNodesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Database name (e.g. oad)
	  Required: true
	  In: path
	*/
	Db string
	/*Unique node identifiers
	  Required: true
	  In: query
	*/
	Nodes string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetMultipleNodesParams() beforehand.
func (o *GetMultipleNodesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rDb, rhkDb, _ := route.Params.GetOK("db")
	if err := o.bindDb(rDb, rhkDb, route.Formats); err != nil {
		res = append(res, err)
	}

	qNodes, qhkNodes, _ := qs.GetOK("nodes")
	if err := o.bindNodes(qNodes, qhkNodes, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetMultipleNodesParams) bindDb(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Db = raw

	if err := o.validateDb(formats); err != nil {
		return err
	}

	return nil
}

func (o *GetMultipleNodesParams) validateDb(formats strfmt.Registry) error {

	if err := validate.Enum("db", "path", o.Db, []interface{}{"OAD Pending Changes", "ION originative suite", "OAD Private Workspace", "OAD Static Data", "Documents Libary", "AIS map design", "CFE definition file", "Map Regions", "OAD AIRAC Buffer", "dataOut"}); err != nil {
		return err
	}

	return nil
}

func (o *GetMultipleNodesParams) bindNodes(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("nodes", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("nodes", "query", raw); err != nil {
		return err
	}

	o.Nodes = raw

	return nil
}