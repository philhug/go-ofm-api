// Code generated by go-swagger; DO NOT EDIT.

package native_client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetUserNodeHandlerFunc turns a function with the right signature into a get user node handler
type GetUserNodeHandlerFunc func(GetUserNodeParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserNodeHandlerFunc) Handle(params GetUserNodeParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetUserNodeHandler interface for that can handle valid get user node params
type GetUserNodeHandler interface {
	Handle(GetUserNodeParams, interface{}) middleware.Responder
}

// NewGetUserNode creates a new http.Handler for the get user node operation
func NewGetUserNode(ctx *middleware.Context, handler GetUserNodeHandler) *GetUserNode {
	return &GetUserNode{Context: ctx, Handler: handler}
}

/*GetUserNode swagger:route GET /nativeclient/user/{id} NativeClient getUserNode

Get node data

*/
type GetUserNode struct {
	Context *middleware.Context
	Handler GetUserNodeHandler
}

func (o *GetUserNode) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUserNodeParams()

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