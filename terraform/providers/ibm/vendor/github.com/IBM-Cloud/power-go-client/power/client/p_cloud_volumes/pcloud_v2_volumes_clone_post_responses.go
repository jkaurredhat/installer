// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudV2VolumesClonePostReader is a Reader for the PcloudV2VolumesClonePost structure.
type PcloudV2VolumesClonePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudV2VolumesClonePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPcloudV2VolumesClonePostAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudV2VolumesClonePostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudV2VolumesClonePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudV2VolumesClonePostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudV2VolumesClonePostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudV2VolumesClonePostAccepted creates a PcloudV2VolumesClonePostAccepted with default headers values
func NewPcloudV2VolumesClonePostAccepted() *PcloudV2VolumesClonePostAccepted {
	return &PcloudV2VolumesClonePostAccepted{}
}

/* PcloudV2VolumesClonePostAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PcloudV2VolumesClonePostAccepted struct {
	Payload *models.CloneTaskReference
}

func (o *PcloudV2VolumesClonePostAccepted) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes/clone][%d] pcloudV2VolumesClonePostAccepted  %+v", 202, o.Payload)
}
func (o *PcloudV2VolumesClonePostAccepted) GetPayload() *models.CloneTaskReference {
	return o.Payload
}

func (o *PcloudV2VolumesClonePostAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloneTaskReference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumesClonePostBadRequest creates a PcloudV2VolumesClonePostBadRequest with default headers values
func NewPcloudV2VolumesClonePostBadRequest() *PcloudV2VolumesClonePostBadRequest {
	return &PcloudV2VolumesClonePostBadRequest{}
}

/* PcloudV2VolumesClonePostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudV2VolumesClonePostBadRequest struct {
	Payload *models.Error
}

func (o *PcloudV2VolumesClonePostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes/clone][%d] pcloudV2VolumesClonePostBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudV2VolumesClonePostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumesClonePostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumesClonePostUnauthorized creates a PcloudV2VolumesClonePostUnauthorized with default headers values
func NewPcloudV2VolumesClonePostUnauthorized() *PcloudV2VolumesClonePostUnauthorized {
	return &PcloudV2VolumesClonePostUnauthorized{}
}

/* PcloudV2VolumesClonePostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudV2VolumesClonePostUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudV2VolumesClonePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes/clone][%d] pcloudV2VolumesClonePostUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudV2VolumesClonePostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumesClonePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumesClonePostNotFound creates a PcloudV2VolumesClonePostNotFound with default headers values
func NewPcloudV2VolumesClonePostNotFound() *PcloudV2VolumesClonePostNotFound {
	return &PcloudV2VolumesClonePostNotFound{}
}

/* PcloudV2VolumesClonePostNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudV2VolumesClonePostNotFound struct {
	Payload *models.Error
}

func (o *PcloudV2VolumesClonePostNotFound) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes/clone][%d] pcloudV2VolumesClonePostNotFound  %+v", 404, o.Payload)
}
func (o *PcloudV2VolumesClonePostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumesClonePostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumesClonePostInternalServerError creates a PcloudV2VolumesClonePostInternalServerError with default headers values
func NewPcloudV2VolumesClonePostInternalServerError() *PcloudV2VolumesClonePostInternalServerError {
	return &PcloudV2VolumesClonePostInternalServerError{}
}

/* PcloudV2VolumesClonePostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudV2VolumesClonePostInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudV2VolumesClonePostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes/clone][%d] pcloudV2VolumesClonePostInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudV2VolumesClonePostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumesClonePostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
