// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UpdateNodeReader is a Reader for the UpdateNode structure.
type UpdateNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateNodeOK creates a UpdateNodeOK with default headers values
func NewUpdateNodeOK() *UpdateNodeOK {
	return &UpdateNodeOK{}
}

/*UpdateNodeOK handles this case with default header values.

OK
*/
type UpdateNodeOK struct {
}

func (o *UpdateNodeOK) Error() string {
	return fmt.Sprintf("[PUT /node/{db}/{id}][%d] updateNodeOK ", 200)
}

func (o *UpdateNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
