// Code generated by go-swagger; DO NOT EDIT.

package native_client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CreateUserNodeReader is a Reader for the CreateUserNode structure.
type CreateUserNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateUserNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateUserNodeOK creates a CreateUserNodeOK with default headers values
func NewCreateUserNodeOK() *CreateUserNodeOK {
	return &CreateUserNodeOK{}
}

/*CreateUserNodeOK handles this case with default header values.

OK
*/
type CreateUserNodeOK struct {
}

func (o *CreateUserNodeOK) Error() string {
	return fmt.Sprintf("[PUT /nativeclient/user/_create][%d] createUserNodeOK ", 200)
}

func (o *CreateUserNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
