// Code generated by go-swagger; DO NOT EDIT.

package native_client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UpdateOrgNodeHandlerFunc turns a function with the right signature into a update org node handler
type UpdateOrgNodeHandlerFunc func(UpdateOrgNodeParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateOrgNodeHandlerFunc) Handle(params UpdateOrgNodeParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateOrgNodeHandler interface for that can handle valid update org node params
type UpdateOrgNodeHandler interface {
	Handle(UpdateOrgNodeParams, interface{}) middleware.Responder
}

// NewUpdateOrgNode creates a new http.Handler for the update org node operation
func NewUpdateOrgNode(ctx *middleware.Context, handler UpdateOrgNodeHandler) *UpdateOrgNode {
	return &UpdateOrgNode{Context: ctx, Handler: handler}
}

/*UpdateOrgNode swagger:route PUT /nativeclient/org/{id} NativeClient updateOrgNode

Replace node

*/
type UpdateOrgNode struct {
	Context *middleware.Context
	Handler UpdateOrgNodeHandler
}

func (o *UpdateOrgNode) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateOrgNodeParams()

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
