// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// CreateNodeOKCode is the HTTP code returned for type CreateNodeOK
const CreateNodeOKCode int = 200

/*CreateNodeOK OK

swagger:response createNodeOK
*/
type CreateNodeOK struct {
}

// NewCreateNodeOK creates CreateNodeOK with default headers values
func NewCreateNodeOK() *CreateNodeOK {

	return &CreateNodeOK{}
}

// WriteResponse to the client
func (o *CreateNodeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// CreateNodeBadRequestCode is the HTTP code returned for type CreateNodeBadRequest
const CreateNodeBadRequestCode int = 400

/*CreateNodeBadRequest Node already exists, searchTags duplicate. searchTags generated server-side. switch to FIR+Type+Designator, returns id of existing object

swagger:response createNodeBadRequest
*/
type CreateNodeBadRequest struct {
}

// NewCreateNodeBadRequest creates CreateNodeBadRequest with default headers values
func NewCreateNodeBadRequest() *CreateNodeBadRequest {

	return &CreateNodeBadRequest{}
}

// WriteResponse to the client
func (o *CreateNodeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
