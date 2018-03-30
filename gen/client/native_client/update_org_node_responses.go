// Code generated by go-swagger; DO NOT EDIT.

package native_client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UpdateOrgNodeReader is a Reader for the UpdateOrgNode structure.
type UpdateOrgNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOrgNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateOrgNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateOrgNodeOK creates a UpdateOrgNodeOK with default headers values
func NewUpdateOrgNodeOK() *UpdateOrgNodeOK {
	return &UpdateOrgNodeOK{}
}

/*UpdateOrgNodeOK handles this case with default header values.

OK
*/
type UpdateOrgNodeOK struct {
}

func (o *UpdateOrgNodeOK) Error() string {
	return fmt.Sprintf("[PUT /nativeclient/org/{id}][%d] updateOrgNodeOK ", 200)
}

func (o *UpdateOrgNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}