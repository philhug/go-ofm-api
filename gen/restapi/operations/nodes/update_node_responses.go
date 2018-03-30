// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// UpdateNodeOKCode is the HTTP code returned for type UpdateNodeOK
const UpdateNodeOKCode int = 200

/*UpdateNodeOK OK

swagger:response updateNodeOK
*/
type UpdateNodeOK struct {
}

// NewUpdateNodeOK creates UpdateNodeOK with default headers values
func NewUpdateNodeOK() *UpdateNodeOK {

	return &UpdateNodeOK{}
}

// WriteResponse to the client
func (o *UpdateNodeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
