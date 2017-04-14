package j_stack_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// JStackTemplateSamplesReader is a Reader for the JStackTemplateSamples structure.
type JStackTemplateSamplesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JStackTemplateSamplesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJStackTemplateSamplesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewJStackTemplateSamplesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJStackTemplateSamplesOK creates a JStackTemplateSamplesOK with default headers values
func NewJStackTemplateSamplesOK() *JStackTemplateSamplesOK {
	return &JStackTemplateSamplesOK{}
}

/*JStackTemplateSamplesOK handles this case with default header values.

stacktemplate sample in json and yaml format with default values
*/
type JStackTemplateSamplesOK struct {
	Payload *models.DefaultResponse
}

func (o *JStackTemplateSamplesOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JStackTemplate.samples][%d] jStackTemplateSamplesOK  %+v", 200, o.Payload)
}

func (o *JStackTemplateSamplesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJStackTemplateSamplesUnauthorized creates a JStackTemplateSamplesUnauthorized with default headers values
func NewJStackTemplateSamplesUnauthorized() *JStackTemplateSamplesUnauthorized {
	return &JStackTemplateSamplesUnauthorized{}
}

/*JStackTemplateSamplesUnauthorized handles this case with default header values.

Unauthorized request
*/
type JStackTemplateSamplesUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *JStackTemplateSamplesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JStackTemplate.samples][%d] jStackTemplateSamplesUnauthorized  %+v", 401, o.Payload)
}

func (o *JStackTemplateSamplesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*JStackTemplateSamplesBody j stack template samples body
swagger:model JStackTemplateSamplesBody
*/
type JStackTemplateSamplesBody struct {

	// provider
	// Required: true
	Provider *string `json:"provider"`

	// use defaults
	// Required: true
	UseDefaults bool `json:"useDefaults"`
}

/*JStackTemplateSamplesOKBodyDefaults j stack template samples o k body defaults
swagger:model JStackTemplateSamplesOKBodyDefaults
*/
type JStackTemplateSamplesOKBodyDefaults struct {

	// user inputs
	UserInputs interface{} `json:"userInputs,omitempty"`
}

// Validate validates this j stack template samples o k body defaults
func (o *JStackTemplateSamplesOKBodyDefaults) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

/*JStackTemplateSamplesOKBody j stack template samples o k body
swagger:model JStackTemplateSamplesOKBody
*/
type JStackTemplateSamplesOKBody struct {

	// defaults
	// Required: true
	Defaults *JStackTemplateSamplesOKBodyDefaults `json:"defaults"`

	// json
	// Required: true
	JSON *string `json:"json"`

	// yaml
	// Required: true
	Yaml *string `json:"yaml"`
}

// Validate validates this j stack template samples o k body
func (o *JStackTemplateSamplesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDefaults(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateJSON(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateYaml(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *JStackTemplateSamplesOKBody) validateDefaults(formats strfmt.Registry) error {

	if err := validate.Required("jStackTemplateSamplesOK"+"."+"defaults", "body", o.Defaults); err != nil {
		return err
	}

	if o.Defaults != nil {

		if err := o.Defaults.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("jStackTemplateSamplesOK" + "." + "defaults")
			}
			return err
		}
	}

	return nil
}

func (o *JStackTemplateSamplesOKBody) validateJSON(formats strfmt.Registry) error {

	if err := validate.Required("jStackTemplateSamplesOK"+"."+"json", "body", o.JSON); err != nil {
		return err
	}

	return nil
}

func (o *JStackTemplateSamplesOKBody) validateYaml(formats strfmt.Registry) error {

	if err := validate.Required("jStackTemplateSamplesOK"+"."+"yaml", "body", o.Yaml); err != nil {
		return err
	}

	return nil
}
