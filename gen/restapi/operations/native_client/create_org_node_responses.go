// Code generated by go-swagger; DO NOT EDIT.

package native_client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// CreateOrgNodeOKCode is the HTTP code returned for type CreateOrgNodeOK
const CreateOrgNodeOKCode int = 200

/*CreateOrgNodeOK OK

swagger:response createOrgNodeOK
*/
type CreateOrgNodeOK struct {
}

// NewCreateOrgNodeOK creates CreateOrgNodeOK with default headers values
func NewCreateOrgNodeOK() *CreateOrgNodeOK {

	return &CreateOrgNodeOK{}
}

// WriteResponse to the client
func (o *CreateOrgNodeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
