// Code generated by go-swagger; DO NOT EDIT.

package native_client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteUserNodeHandlerFunc turns a function with the right signature into a delete user node handler
type DeleteUserNodeHandlerFunc func(DeleteUserNodeParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteUserNodeHandlerFunc) Handle(params DeleteUserNodeParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteUserNodeHandler interface for that can handle valid delete user node params
type DeleteUserNodeHandler interface {
	Handle(DeleteUserNodeParams, interface{}) middleware.Responder
}

// NewDeleteUserNode creates a new http.Handler for the delete user node operation
func NewDeleteUserNode(ctx *middleware.Context, handler DeleteUserNodeHandler) *DeleteUserNode {
	return &DeleteUserNode{Context: ctx, Handler: handler}
}

/*DeleteUserNode swagger:route DELETE /nativeclient/user/{id} NativeClient deleteUserNode

Delete

*/
type DeleteUserNode struct {
	Context *middleware.Context
	Handler DeleteUserNodeHandler
}

func (o *DeleteUserNode) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteUserNodeParams()

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
