// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPutAttachmentParams creates a new PutAttachmentParams object
// no default values defined in spec.
func NewPutAttachmentParams() PutAttachmentParams {

	return PutAttachmentParams{}
}

// PutAttachmentParams contains all the bound params for the put attachment operation
// typically these are obtained from a http.Request
//
// swagger:parameters putAttachment
type PutAttachmentParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Unique blob identifier
	  Required: true
	  In: path
	*/
	Attname string
	/*Database name (e.g. oad)
	  Required: true
	  In: path
	*/
	Db string
	/*Unique node identifier
	  Required: true
	  In: path
	*/
	ID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPutAttachmentParams() beforehand.
func (o *PutAttachmentParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rAttname, rhkAttname, _ := route.Params.GetOK("attname")
	if err := o.bindAttname(rAttname, rhkAttname, route.Formats); err != nil {
		res = append(res, err)
	}

	rDb, rhkDb, _ := route.Params.GetOK("db")
	if err := o.bindDb(rDb, rhkDb, route.Formats); err != nil {
		res = append(res, err)
	}

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutAttachmentParams) bindAttname(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Attname = raw

	return nil
}

func (o *PutAttachmentParams) bindDb(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *PutAttachmentParams) validateDb(formats strfmt.Registry) error {

	if err := validate.Enum("db", "path", o.Db, []interface{}{"OAD Pending Changes", "ION originative suite", "OAD Private Workspace", "OAD Static Data", "Documents Libary", "AIS map design", "CFE definition file", "Map Regions", "OAD AIRAC Buffer", "dataOut"}); err != nil {
		return err
	}

	return nil
}

func (o *PutAttachmentParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ID = raw

	return nil
}