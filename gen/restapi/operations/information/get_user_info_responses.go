// Code generated by go-swagger; DO NOT EDIT.

package information

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/philhug/go-ofm-api/gen/models"
)

// GetUserInfoOKCode is the HTTP code returned for type GetUserInfoOK
const GetUserInfoOKCode int = 200

/*GetUserInfoOK OK

swagger:response getUserInfoOK
*/
type GetUserInfoOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewGetUserInfoOK creates GetUserInfoOK with default headers values
func NewGetUserInfoOK() *GetUserInfoOK {

	return &GetUserInfoOK{}
}

// WithPayload adds the payload to the get user info o k response
func (o *GetUserInfoOK) WithPayload(payload *models.User) *GetUserInfoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user info o k response
func (o *GetUserInfoOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserInfoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
