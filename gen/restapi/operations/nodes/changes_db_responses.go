// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ChangesDbOKCode is the HTTP code returned for type ChangesDbOK
const ChangesDbOKCode int = 200

/*ChangesDbOK OK

swagger:response changesDbOK
*/
type ChangesDbOK struct {
}

// NewChangesDbOK creates ChangesDbOK with default headers values
func NewChangesDbOK() *ChangesDbOK {

	return &ChangesDbOK{}
}

// WriteResponse to the client
func (o *ChangesDbOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
