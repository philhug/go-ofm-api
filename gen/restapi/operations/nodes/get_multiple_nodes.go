// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetMultipleNodesHandlerFunc turns a function with the right signature into a get multiple nodes handler
type GetMultipleNodesHandlerFunc func(GetMultipleNodesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetMultipleNodesHandlerFunc) Handle(params GetMultipleNodesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetMultipleNodesHandler interface for that can handle valid get multiple nodes params
type GetMultipleNodesHandler interface {
	Handle(GetMultipleNodesParams, interface{}) middleware.Responder
}

// NewGetMultipleNodes creates a new http.Handler for the get multiple nodes operation
func NewGetMultipleNodes(ctx *middleware.Context, handler GetMultipleNodesHandler) *GetMultipleNodes {
	return &GetMultipleNodes{Context: ctx, Handler: handler}
}

/*GetMultipleNodes swagger:route GET /node/{db}/_multi Nodes getMultipleNodes

Get multiple node data

*/
type GetMultipleNodes struct {
	Context *middleware.Context
	Handler GetMultipleNodesHandler
}

func (o *GetMultipleNodes) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetMultipleNodesParams()

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
