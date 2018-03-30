// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/philhug/go-ofmapi-server/gen/models"
)

// SearchNodeOKCode is the HTTP code returned for type SearchNodeOK
const SearchNodeOKCode int = 200

/*SearchNodeOK OK

swagger:response searchNodeOK
*/
type SearchNodeOK struct {

	/*
	  In: Body
	*/
	Payload *models.NodeNumberList `json:"body,omitempty"`
}

// NewSearchNodeOK creates SearchNodeOK with default headers values
func NewSearchNodeOK() *SearchNodeOK {

	return &SearchNodeOK{}
}

// WithPayload adds the payload to the search node o k response
func (o *SearchNodeOK) WithPayload(payload *models.NodeNumberList) *SearchNodeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the search node o k response
func (o *SearchNodeOK) SetPayload(payload *models.NodeNumberList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SearchNodeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}