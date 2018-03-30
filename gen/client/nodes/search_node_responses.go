// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/philhug/go-ofmapi-server/gen/models"
)

// SearchNodeReader is a Reader for the SearchNode structure.
type SearchNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSearchNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSearchNodeOK creates a SearchNodeOK with default headers values
func NewSearchNodeOK() *SearchNodeOK {
	return &SearchNodeOK{}
}

/*SearchNodeOK handles this case with default header values.

OK
*/
type SearchNodeOK struct {
	Payload *models.NodeNumberList
}

func (o *SearchNodeOK) Error() string {
	return fmt.Sprintf("[POST /node/{db}/_search][%d] searchNodeOK  %+v", 200, o.Payload)
}

func (o *SearchNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NodeNumberList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}