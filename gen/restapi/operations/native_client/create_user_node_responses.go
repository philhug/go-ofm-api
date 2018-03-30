// Code generated by go-swagger; DO NOT EDIT.

package native_client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// CreateUserNodeOKCode is the HTTP code returned for type CreateUserNodeOK
const CreateUserNodeOKCode int = 200

/*CreateUserNodeOK OK

swagger:response createUserNodeOK
*/
type CreateUserNodeOK struct {
}

// NewCreateUserNodeOK creates CreateUserNodeOK with default headers values
func NewCreateUserNodeOK() *CreateUserNodeOK {

	return &CreateUserNodeOK{}
}

// WriteResponse to the client
func (o *CreateUserNodeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
