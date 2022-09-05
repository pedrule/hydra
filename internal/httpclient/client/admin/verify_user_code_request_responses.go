// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// VerifyUserCodeRequestReader is a Reader for the VerifyUserCodeRequest structure.
type VerifyUserCodeRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VerifyUserCodeRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVerifyUserCodeRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewVerifyUserCodeRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewVerifyUserCodeRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVerifyUserCodeRequestOK creates a VerifyUserCodeRequestOK with default headers values
func NewVerifyUserCodeRequestOK() *VerifyUserCodeRequestOK {
	return &VerifyUserCodeRequestOK{}
}

/* VerifyUserCodeRequestOK describes a response with status code 200, with default header values.

completedRequest
*/
type VerifyUserCodeRequestOK struct {
	Payload *models.CompletedRequest
}

func (o *VerifyUserCodeRequestOK) Error() string {
	return fmt.Sprintf("[PUT /oauth2/auth/requests/device/verify][%d] verifyUserCodeRequestOK  %+v", 200, o.Payload)
}
func (o *VerifyUserCodeRequestOK) GetPayload() *models.CompletedRequest {
	return o.Payload
}

func (o *VerifyUserCodeRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CompletedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVerifyUserCodeRequestNotFound creates a VerifyUserCodeRequestNotFound with default headers values
func NewVerifyUserCodeRequestNotFound() *VerifyUserCodeRequestNotFound {
	return &VerifyUserCodeRequestNotFound{}
}

/* VerifyUserCodeRequestNotFound describes a response with status code 404, with default header values.

jsonError
*/
type VerifyUserCodeRequestNotFound struct {
	Payload *models.JSONError
}

func (o *VerifyUserCodeRequestNotFound) Error() string {
	return fmt.Sprintf("[PUT /oauth2/auth/requests/device/verify][%d] verifyUserCodeRequestNotFound  %+v", 404, o.Payload)
}
func (o *VerifyUserCodeRequestNotFound) GetPayload() *models.JSONError {
	return o.Payload
}

func (o *VerifyUserCodeRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVerifyUserCodeRequestInternalServerError creates a VerifyUserCodeRequestInternalServerError with default headers values
func NewVerifyUserCodeRequestInternalServerError() *VerifyUserCodeRequestInternalServerError {
	return &VerifyUserCodeRequestInternalServerError{}
}

/* VerifyUserCodeRequestInternalServerError describes a response with status code 500, with default header values.

jsonError
*/
type VerifyUserCodeRequestInternalServerError struct {
	Payload *models.JSONError
}

func (o *VerifyUserCodeRequestInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /oauth2/auth/requests/device/verify][%d] verifyUserCodeRequestInternalServerError  %+v", 500, o.Payload)
}
func (o *VerifyUserCodeRequestInternalServerError) GetPayload() *models.JSONError {
	return o.Payload
}

func (o *VerifyUserCodeRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
