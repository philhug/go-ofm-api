// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/philhug/go-ofmapi-server/gen/models"
)

// GetNodeOKCode is the HTTP code returned for type GetNodeOK
const GetNodeOKCode int = 200

/*GetNodeOK OK

swagger:response getNodeOK
*/
type GetNodeOK struct {

	/*
	  In: Body
	*/
	Payload *models.Node `json:"body,omitempty"`
}

// NewGetNodeOK creates GetNodeOK with default headers values
func NewGetNodeOK() *GetNodeOK {

	return &GetNodeOK{}
}

// WithPayload adds the payload to the get node o k response
func (o *GetNodeOK) WithPayload(payload *models.Node) *GetNodeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get node o k response
func (o *GetNodeOK) SetPayload(payload *models.Node) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNodeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}