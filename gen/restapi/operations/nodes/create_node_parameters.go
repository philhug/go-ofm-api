// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/philhug/go-ofm-api/gen/models"
)

// NewCreateNodeParams creates a new CreateNodeParams object
// no default values defined in spec.
func NewCreateNodeParams() CreateNodeParams {

	return CreateNodeParams{}
}

// CreateNodeParams contains all the bound params for the create node operation
// typically these are obtained from a http.Request
//
// swagger:parameters createNode
type CreateNodeParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Node definition
	  Required: true
	  In: body
	*/
	Body *models.Node
	/*Database name (e.g. "OAD Pending Changes")
	  Required: true
	  In: path
	*/
	Db string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateNodeParams() beforehand.
func (o *CreateNodeParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Node
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body"))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
		} else {

			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	} else {
		res = append(res, errors.Required("body", "body"))
	}
	rDb, rhkDb, _ := route.Params.GetOK("db")
	if err := o.bindDb(rDb, rhkDb, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNodeParams) bindDb(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *CreateNodeParams) validateDb(formats strfmt.Registry) error {

	if err := validate.Enum("db", "path", o.Db, []interface{}{"OAD Pending Changes", "ION originative suite", "OAD Private Workspace", "OAD Static Data", "Documents Libary", "AIS map design", "CFE definition file", "Map Regions", "OAD AIRAC Buffer", "dataOut"}); err != nil {
		return err
	}

	return nil
}
