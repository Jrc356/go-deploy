// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aptible/go-deploy/models"
)

// GetVhostsVhostIDOperationsReader is a Reader for the GetVhostsVhostIDOperations structure.
type GetVhostsVhostIDOperationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVhostsVhostIDOperationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVhostsVhostIDOperationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetVhostsVhostIDOperationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVhostsVhostIDOperationsOK creates a GetVhostsVhostIDOperationsOK with default headers values
func NewGetVhostsVhostIDOperationsOK() *GetVhostsVhostIDOperationsOK {
	return &GetVhostsVhostIDOperationsOK{}
}

/*GetVhostsVhostIDOperationsOK handles this case with default header values.

successful
*/
type GetVhostsVhostIDOperationsOK struct {
	Payload *models.InlineResponse20031
}

func (o *GetVhostsVhostIDOperationsOK) Error() string {
	return fmt.Sprintf("[GET /vhosts/{vhost_id}/operations][%d] getVhostsVhostIdOperationsOK  %+v", 200, o.Payload)
}

func (o *GetVhostsVhostIDOperationsOK) GetPayload() *models.InlineResponse20031 {
	return o.Payload
}

func (o *GetVhostsVhostIDOperationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20031)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVhostsVhostIDOperationsDefault creates a GetVhostsVhostIDOperationsDefault with default headers values
func NewGetVhostsVhostIDOperationsDefault(code int) *GetVhostsVhostIDOperationsDefault {
	return &GetVhostsVhostIDOperationsDefault{
		_statusCode: code,
	}
}

/*GetVhostsVhostIDOperationsDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetVhostsVhostIDOperationsDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get vhosts vhost ID operations default response
func (o *GetVhostsVhostIDOperationsDefault) Code() int {
	return o._statusCode
}

func (o *GetVhostsVhostIDOperationsDefault) Error() string {
	return fmt.Sprintf("[GET /vhosts/{vhost_id}/operations][%d] GetVhostsVhostIDOperations default  %+v", o._statusCode, o.Payload)
}

func (o *GetVhostsVhostIDOperationsDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetVhostsVhostIDOperationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
