// Code generated by go-swagger; DO NOT EDIT.

package native_client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// SearchUserNodeHandlerFunc turns a function with the right signature into a search user node handler
type SearchUserNodeHandlerFunc func(SearchUserNodeParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn SearchUserNodeHandlerFunc) Handle(params SearchUserNodeParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// SearchUserNodeHandler interface for that can handle valid search user node params
type SearchUserNodeHandler interface {
	Handle(SearchUserNodeParams, interface{}) middleware.Responder
}

// NewSearchUserNode creates a new http.Handler for the search user node operation
func NewSearchUserNode(ctx *middleware.Context, handler SearchUserNodeHandler) *SearchUserNode {
	return &SearchUserNode{Context: ctx, Handler: handler}
}

/*SearchUserNode swagger:route POST /nativeclient/user/_search NativeClient searchUserNode

Search

Return all org objects in database.


*/
type SearchUserNode struct {
	Context *middleware.Context
	Handler SearchUserNodeHandler
}

func (o *SearchUserNode) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSearchUserNodeParams()

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
