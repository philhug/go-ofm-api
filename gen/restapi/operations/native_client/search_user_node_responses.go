// Code generated by go-swagger; DO NOT EDIT.

package native_client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/philhug/go-ofmapi-server/gen/models"
)

// SearchUserNodeOKCode is the HTTP code returned for type SearchUserNodeOK
const SearchUserNodeOKCode int = 200

/*SearchUserNodeOK OK

swagger:response searchUserNodeOK
*/
type SearchUserNodeOK struct {

	/*
	  In: Body
	*/
	Payload *models.NodeNumberList `json:"body,omitempty"`
}

// NewSearchUserNodeOK creates SearchUserNodeOK with default headers values
func NewSearchUserNodeOK() *SearchUserNodeOK {

	return &SearchUserNodeOK{}
}

// WithPayload adds the payload to the search user node o k response
func (o *SearchUserNodeOK) WithPayload(payload *models.NodeNumberList) *SearchUserNodeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the search user node o k response
func (o *SearchUserNodeOK) SetPayload(payload *models.NodeNumberList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SearchUserNodeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
