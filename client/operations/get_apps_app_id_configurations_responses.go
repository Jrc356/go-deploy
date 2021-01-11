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

// GetAppsAppIDConfigurationsReader is a Reader for the GetAppsAppIDConfigurations structure.
type GetAppsAppIDConfigurationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppsAppIDConfigurationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAppsAppIDConfigurationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAppsAppIDConfigurationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAppsAppIDConfigurationsOK creates a GetAppsAppIDConfigurationsOK with default headers values
func NewGetAppsAppIDConfigurationsOK() *GetAppsAppIDConfigurationsOK {
	return &GetAppsAppIDConfigurationsOK{}
}

/*GetAppsAppIDConfigurationsOK handles this case with default header values.

successful
*/
type GetAppsAppIDConfigurationsOK struct {
	Payload *models.InlineResponse2009
}

func (o *GetAppsAppIDConfigurationsOK) Error() string {
	return fmt.Sprintf("[GET /apps/{app_id}/configurations][%d] getAppsAppIdConfigurationsOK  %+v", 200, o.Payload)
}

func (o *GetAppsAppIDConfigurationsOK) GetPayload() *models.InlineResponse2009 {
	return o.Payload
}

func (o *GetAppsAppIDConfigurationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse2009)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppsAppIDConfigurationsDefault creates a GetAppsAppIDConfigurationsDefault with default headers values
func NewGetAppsAppIDConfigurationsDefault(code int) *GetAppsAppIDConfigurationsDefault {
	return &GetAppsAppIDConfigurationsDefault{
		_statusCode: code,
	}
}

/*GetAppsAppIDConfigurationsDefault handles this case with default header values.

Error response. Often a 4xx or 5xx status code
*/
type GetAppsAppIDConfigurationsDefault struct {
	_statusCode int

	Payload *models.InlineResponseDefault
}

// Code gets the status code for the get apps app ID configurations default response
func (o *GetAppsAppIDConfigurationsDefault) Code() int {
	return o._statusCode
}

func (o *GetAppsAppIDConfigurationsDefault) Error() string {
	return fmt.Sprintf("[GET /apps/{app_id}/configurations][%d] GetAppsAppIDConfigurations default  %+v", o._statusCode, o.Payload)
}

func (o *GetAppsAppIDConfigurationsDefault) GetPayload() *models.InlineResponseDefault {
	return o.Payload
}

func (o *GetAppsAppIDConfigurationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponseDefault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
