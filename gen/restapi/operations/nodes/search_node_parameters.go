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
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/philhug/go-ofm-api/gen/models"
)

// NewSearchNodeParams creates a new SearchNodeParams object
// with the default values initialized.
func NewSearchNodeParams() SearchNodeParams {

	var (
		// initialize parameters with default values

		deletedDefault = bool(false)
	)

	return SearchNodeParams{
		Deleted: &deletedDefault,
	}
}

// SearchNodeParams contains all the bound params for the search node operation
// typically these are obtained from a http.Request
//
// swagger:parameters searchNode
type SearchNodeParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*bounding box (left,bottom,right,top)
	  In: query
	*/
	Bbox *string
	/*Database name (e.g. "OAD Pending Changes")
	  Required: true
	  In: path
	*/
	Db string
	/*return deleted nodes
	  In: query
	  Default: false
	*/
	Deleted *bool
	/*full text search in Search Tags
	  In: query
	*/
	Fulltext *string
	/*Query with name value pairs. exact matches. e.g. {region: 'LSAS', type: 'Ase/Abd'
	  Required: true
	  In: body
	*/
	Query *models.NodeQuery
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewSearchNodeParams() beforehand.
func (o *SearchNodeParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qBbox, qhkBbox, _ := qs.GetOK("bbox")
	if err := o.bindBbox(qBbox, qhkBbox, route.Formats); err != nil {
		res = append(res, err)
	}

	rDb, rhkDb, _ := route.Params.GetOK("db")
	if err := o.bindDb(rDb, rhkDb, route.Formats); err != nil {
		res = append(res, err)
	}

	qDeleted, qhkDeleted, _ := qs.GetOK("deleted")
	if err := o.bindDeleted(qDeleted, qhkDeleted, route.Formats); err != nil {
		res = append(res, err)
	}

	qFulltext, qhkFulltext, _ := qs.GetOK("fulltext")
	if err := o.bindFulltext(qFulltext, qhkFulltext, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.NodeQuery
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("query", "body"))
			} else {
				res = append(res, errors.NewParseError("query", "body", "", err))
			}
		} else {

			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Query = &body
			}
		}
	} else {
		res = append(res, errors.Required("query", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SearchNodeParams) bindBbox(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Bbox = &raw

	return nil
}

func (o *SearchNodeParams) bindDb(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *SearchNodeParams) validateDb(formats strfmt.Registry) error {

	if err := validate.Enum("db", "path", o.Db, []interface{}{"OAD Pending Changes", "ION originative suite", "OAD Private Workspace", "OAD Static Data", "Documents Libary", "AIS map design", "CFE definition file", "Map Regions", "OAD AIRAC Buffer", "dataOut"}); err != nil {
		return err
	}

	return nil
}

func (o *SearchNodeParams) bindDeleted(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewSearchNodeParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("deleted", "query", "bool", raw)
	}
	o.Deleted = &value

	return nil
}

func (o *SearchNodeParams) bindFulltext(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Fulltext = &raw

	return nil
}
