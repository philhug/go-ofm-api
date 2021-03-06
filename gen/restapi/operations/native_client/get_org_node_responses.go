// Code generated by go-swagger; DO NOT EDIT.

package native_client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/philhug/go-ofm-api/gen/models"
)

// GetOrgNodeOKCode is the HTTP code returned for type GetOrgNodeOK
const GetOrgNodeOKCode int = 200

/*GetOrgNodeOK OK

swagger:response getOrgNodeOK
*/
type GetOrgNodeOK struct {

	/*
	  In: Body
	*/
	Payload *models.Organization `json:"body,omitempty"`
}

// NewGetOrgNodeOK creates GetOrgNodeOK with default headers values
func NewGetOrgNodeOK() *GetOrgNodeOK {

	return &GetOrgNodeOK{}
}

// WithPayload adds the payload to the get org node o k response
func (o *GetOrgNodeOK) WithPayload(payload *models.Organization) *GetOrgNodeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get org node o k response
func (o *GetOrgNodeOK) SetPayload(payload *models.Organization) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOrgNodeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
