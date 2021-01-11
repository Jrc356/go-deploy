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

// GetMetricDrainsMetricDrainIDContainersReader is a Reader for the GetMetricDrainsMetricDrainIDContainers structure.
type GetMetricDrainsMetricDrainIDContainersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMetricDrainsMetricDrainIDContainersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMetricDrainsMetricDrainIDContainersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMetricDrainsMetricDrainIDContainersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMetricDrainsMetricDrainIDContainersOK creates a GetMetricDrainsMetricDrainIDContainersOK with default headers values
func NewGetMetricDrainsMetricDrainIDContainersOK() *GetMetricDrainsMetricDrainIDContainersOK {
	return &GetMetricDrainsMetricDrainIDContainersOK{}
}

/*GetMetricDrainsMetricDrainIDContainersOK handles this case with default header values.

successful
*/
type GetMetricDrainsMetricDrainIDContainersOK struct {
	Payload *models.InlineResponse20010
}

func (o *GetMetricDrainsMetricDrainIDContainersOK) Error() string {
	return fmt.Sprintf("[GET /metric_drains/{metric_drain_id}/containers][%d] getMetricDrainsMetricDrainIdContainersOK  %+v", 200, o.Payload)
}

func (o *GetMetricDrainsMetricDrainIDContainersOK) GetPayload() *models.InlineResponse20010 {
	return o.Payload
}

func (o *GetMetricDrainsMetricDrainIDContainersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse20010)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMetricDrainsMetricDrainIDContainersDefault creates a GetMetricDrainsMetricDrainIDContainersDefault with default headers values
func NewGetMetricDrainsMetricDrainIDContainersDefault(code int) *GetMetricDrainsMetricDrainIDContainersDefault {
	return &GetMetricDrainsMetricDrainIDContainersDefault{
		_statusCode: code,
	}
}

/*GetMetricDrainsMetricDrainIDContainersDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetMetricDrainsMetricDrainIDContainersDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get metric drains metric drain ID containers default response
func (o *GetMetricDrainsMetricDrainIDContainersDefault) Code() int {
	return o._statusCode
}

func (o *GetMetricDrainsMetricDrainIDContainersDefault) Error() string {
	return fmt.Sprintf("[GET /metric_drains/{metric_drain_id}/containers][%d] GetMetricDrainsMetricDrainIDContainers default  %+v", o._statusCode, o.Payload)
}

func (o *GetMetricDrainsMetricDrainIDContainersDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetMetricDrainsMetricDrainIDContainersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
