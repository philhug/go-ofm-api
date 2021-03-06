// Code generated by go-swagger; DO NOT EDIT.

package native_client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// SearchOrgNodeHandlerFunc turns a function with the right signature into a search org node handler
type SearchOrgNodeHandlerFunc func(SearchOrgNodeParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn SearchOrgNodeHandlerFunc) Handle(params SearchOrgNodeParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// SearchOrgNodeHandler interface for that can handle valid search org node params
type SearchOrgNodeHandler interface {
	Handle(SearchOrgNodeParams, interface{}) middleware.Responder
}

// NewSearchOrgNode creates a new http.Handler for the search org node operation
func NewSearchOrgNode(ctx *middleware.Context, handler SearchOrgNodeHandler) *SearchOrgNode {
	return &SearchOrgNode{Context: ctx, Handler: handler}
}

/*SearchOrgNode swagger:route POST /nativeclient/org/_search NativeClient searchOrgNode

Search

Return all org objects in database.


*/
type SearchOrgNode struct {
	Context *middleware.Context
	Handler SearchOrgNodeHandler
}

func (o *SearchOrgNode) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSearchOrgNodeParams()

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
