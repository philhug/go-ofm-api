// Code generated by go-swagger; DO NOT EDIT.

package native_client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteOrgNodeHandlerFunc turns a function with the right signature into a delete org node handler
type DeleteOrgNodeHandlerFunc func(DeleteOrgNodeParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteOrgNodeHandlerFunc) Handle(params DeleteOrgNodeParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteOrgNodeHandler interface for that can handle valid delete org node params
type DeleteOrgNodeHandler interface {
	Handle(DeleteOrgNodeParams, interface{}) middleware.Responder
}

// NewDeleteOrgNode creates a new http.Handler for the delete org node operation
func NewDeleteOrgNode(ctx *middleware.Context, handler DeleteOrgNodeHandler) *DeleteOrgNode {
	return &DeleteOrgNode{Context: ctx, Handler: handler}
}

/*DeleteOrgNode swagger:route DELETE /nativeclient/org/{id} NativeClient deleteOrgNode

Delete

*/
type DeleteOrgNode struct {
	Context *middleware.Context
	Handler DeleteOrgNodeHandler
}

func (o *DeleteOrgNode) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteOrgNodeParams()

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
