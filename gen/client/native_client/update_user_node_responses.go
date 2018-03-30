// Code generated by go-swagger; DO NOT EDIT.

package native_client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UpdateUserNodeReader is a Reader for the UpdateUserNode structure.
type UpdateUserNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateUserNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateUserNodeOK creates a UpdateUserNodeOK with default headers values
func NewUpdateUserNodeOK() *UpdateUserNodeOK {
	return &UpdateUserNodeOK{}
}

/*UpdateUserNodeOK handles this case with default header values.

OK
*/
type UpdateUserNodeOK struct {
}

func (o *UpdateUserNodeOK) Error() string {
	return fmt.Sprintf("[PUT /nativeclient/user/{id}][%d] updateUserNodeOK ", 200)
}

func (o *UpdateUserNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
