// Code generated by go-swagger; DO NOT EDIT.

package information

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetRegionsHandlerFunc turns a function with the right signature into a get regions handler
type GetRegionsHandlerFunc func(GetRegionsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRegionsHandlerFunc) Handle(params GetRegionsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetRegionsHandler interface for that can handle valid get regions params
type GetRegionsHandler interface {
	Handle(GetRegionsParams, interface{}) middleware.Responder
}

// NewGetRegions creates a new http.Handler for the get regions operation
func NewGetRegions(ctx *middleware.Context, handler GetRegionsHandler) *GetRegions {
	return &GetRegions{Context: ctx, Handler: handler}
}

/*GetRegions swagger:route GET /info/regions Information getRegions

List all Regions

Returns a list of available airspace regions / FIRs

*/
type GetRegions struct {
	Context *middleware.Context
	Handler GetRegionsHandler
}

func (o *GetRegions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRegionsParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
