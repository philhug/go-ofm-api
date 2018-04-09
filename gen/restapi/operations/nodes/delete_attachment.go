// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteAttachmentHandlerFunc turns a function with the right signature into a delete attachment handler
type DeleteAttachmentHandlerFunc func(DeleteAttachmentParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteAttachmentHandlerFunc) Handle(params DeleteAttachmentParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteAttachmentHandler interface for that can handle valid delete attachment params
type DeleteAttachmentHandler interface {
	Handle(DeleteAttachmentParams, interface{}) middleware.Responder
}

// NewDeleteAttachment creates a new http.Handler for the delete attachment operation
func NewDeleteAttachment(ctx *middleware.Context, handler DeleteAttachmentHandler) *DeleteAttachment {
	return &DeleteAttachment{Context: ctx, Handler: handler}
}

/*DeleteAttachment swagger:route DELETE /node/{db}/{id}/{attname} Nodes deleteAttachment

Get Blob

*/
type DeleteAttachment struct {
	Context *middleware.Context
	Handler DeleteAttachmentHandler
}

func (o *DeleteAttachment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteAttachmentParams()

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
