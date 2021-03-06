// Code generated by go-swagger; DO NOT EDIT.

package search_transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCreateSearchParams creates a new CreateSearchParams object
// with the default values initialized.
func NewCreateSearchParams() *CreateSearchParams {
	var ()
	return &CreateSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSearchParamsWithTimeout creates a new CreateSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSearchParamsWithTimeout(timeout time.Duration) *CreateSearchParams {
	var ()
	return &CreateSearchParams{

		timeout: timeout,
	}
}

// NewCreateSearchParamsWithContext creates a new CreateSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSearchParamsWithContext(ctx context.Context) *CreateSearchParams {
	var ()
	return &CreateSearchParams{

		Context: ctx,
	}
}

// NewCreateSearchParamsWithHTTPClient creates a new CreateSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSearchParamsWithHTTPClient(client *http.Client) *CreateSearchParams {
	var ()
	return &CreateSearchParams{
		HTTPClient: client,
	}
}

/*CreateSearchParams contains all the parameters to send to the API endpoint
for the create search operation typically these are written to a http.Request
*/
type CreateSearchParams struct {

	/*CreateSearchRequest*/
	CreateSearchRequest CreateSearchBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create search params
func (o *CreateSearchParams) WithTimeout(timeout time.Duration) *CreateSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create search params
func (o *CreateSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create search params
func (o *CreateSearchParams) WithContext(ctx context.Context) *CreateSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create search params
func (o *CreateSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create search params
func (o *CreateSearchParams) WithHTTPClient(client *http.Client) *CreateSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create search params
func (o *CreateSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateSearchRequest adds the createSearchRequest to the create search params
func (o *CreateSearchParams) WithCreateSearchRequest(createSearchRequest CreateSearchBody) *CreateSearchParams {
	o.SetCreateSearchRequest(createSearchRequest)
	return o
}

// SetCreateSearchRequest adds the createSearchRequest to the create search params
func (o *CreateSearchParams) SetCreateSearchRequest(createSearchRequest CreateSearchBody) {
	o.CreateSearchRequest = createSearchRequest
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.CreateSearchRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
