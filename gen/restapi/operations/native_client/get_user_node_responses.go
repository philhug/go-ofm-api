// Code generated by go-swagger; DO NOT EDIT.

package native_client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/philhug/go-ofm-api/gen/models"
)

// GetUserNodeOKCode is the HTTP code returned for type GetUserNodeOK
const GetUserNodeOKCode int = 200

/*GetUserNodeOK OK

swagger:response getUserNodeOK
*/
type GetUserNodeOK struct {

	/*
	  In: Body
	*/
	Payload *models.Node `json:"body,omitempty"`
}

// NewGetUserNodeOK creates GetUserNodeOK with default headers values
func NewGetUserNodeOK() *GetUserNodeOK {

	return &GetUserNodeOK{}
}

// WithPayload adds the payload to the get user node o k response
func (o *GetUserNodeOK) WithPayload(payload *models.Node) *GetUserNodeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user node o k response
func (o *GetUserNodeOK) SetPayload(payload *models.Node) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserNodeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
