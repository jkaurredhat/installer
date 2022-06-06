// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// NewPcloudV1CloudinstancesCosimagesPostParams creates a new PcloudV1CloudinstancesCosimagesPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudV1CloudinstancesCosimagesPostParams() *PcloudV1CloudinstancesCosimagesPostParams {
	return &PcloudV1CloudinstancesCosimagesPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudV1CloudinstancesCosimagesPostParamsWithTimeout creates a new PcloudV1CloudinstancesCosimagesPostParams object
// with the ability to set a timeout on a request.
func NewPcloudV1CloudinstancesCosimagesPostParamsWithTimeout(timeout time.Duration) *PcloudV1CloudinstancesCosimagesPostParams {
	return &PcloudV1CloudinstancesCosimagesPostParams{
		timeout: timeout,
	}
}

// NewPcloudV1CloudinstancesCosimagesPostParamsWithContext creates a new PcloudV1CloudinstancesCosimagesPostParams object
// with the ability to set a context for a request.
func NewPcloudV1CloudinstancesCosimagesPostParamsWithContext(ctx context.Context) *PcloudV1CloudinstancesCosimagesPostParams {
	return &PcloudV1CloudinstancesCosimagesPostParams{
		Context: ctx,
	}
}

// NewPcloudV1CloudinstancesCosimagesPostParamsWithHTTPClient creates a new PcloudV1CloudinstancesCosimagesPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudV1CloudinstancesCosimagesPostParamsWithHTTPClient(client *http.Client) *PcloudV1CloudinstancesCosimagesPostParams {
	return &PcloudV1CloudinstancesCosimagesPostParams{
		HTTPClient: client,
	}
}

/* PcloudV1CloudinstancesCosimagesPostParams contains all the parameters to send to the API endpoint
   for the pcloud v1 cloudinstances cosimages post operation.

   Typically these are written to a http.Request.
*/
type PcloudV1CloudinstancesCosimagesPostParams struct {

	/* Body.

	   Parameters for the creation of a new cos-image import job
	*/
	Body *models.CreateCosImageImportJob

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud v1 cloudinstances cosimages post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudV1CloudinstancesCosimagesPostParams) WithDefaults() *PcloudV1CloudinstancesCosimagesPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud v1 cloudinstances cosimages post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudV1CloudinstancesCosimagesPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud v1 cloudinstances cosimages post params
func (o *PcloudV1CloudinstancesCosimagesPostParams) WithTimeout(timeout time.Duration) *PcloudV1CloudinstancesCosimagesPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud v1 cloudinstances cosimages post params
func (o *PcloudV1CloudinstancesCosimagesPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud v1 cloudinstances cosimages post params
func (o *PcloudV1CloudinstancesCosimagesPostParams) WithContext(ctx context.Context) *PcloudV1CloudinstancesCosimagesPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud v1 cloudinstances cosimages post params
func (o *PcloudV1CloudinstancesCosimagesPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud v1 cloudinstances cosimages post params
func (o *PcloudV1CloudinstancesCosimagesPostParams) WithHTTPClient(client *http.Client) *PcloudV1CloudinstancesCosimagesPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud v1 cloudinstances cosimages post params
func (o *PcloudV1CloudinstancesCosimagesPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud v1 cloudinstances cosimages post params
func (o *PcloudV1CloudinstancesCosimagesPostParams) WithBody(body *models.CreateCosImageImportJob) *PcloudV1CloudinstancesCosimagesPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud v1 cloudinstances cosimages post params
func (o *PcloudV1CloudinstancesCosimagesPostParams) SetBody(body *models.CreateCosImageImportJob) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud v1 cloudinstances cosimages post params
func (o *PcloudV1CloudinstancesCosimagesPostParams) WithCloudInstanceID(cloudInstanceID string) *PcloudV1CloudinstancesCosimagesPostParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud v1 cloudinstances cosimages post params
func (o *PcloudV1CloudinstancesCosimagesPostParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudV1CloudinstancesCosimagesPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
