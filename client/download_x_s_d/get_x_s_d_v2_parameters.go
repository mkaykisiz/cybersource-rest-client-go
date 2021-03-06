// Code generated by go-swagger; DO NOT EDIT.

package download_x_s_d

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

// NewGetXSDV2Params creates a new GetXSDV2Params object
// with the default values initialized.
func NewGetXSDV2Params() *GetXSDV2Params {
	var ()
	return &GetXSDV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetXSDV2ParamsWithTimeout creates a new GetXSDV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetXSDV2ParamsWithTimeout(timeout time.Duration) *GetXSDV2Params {
	var ()
	return &GetXSDV2Params{

		timeout: timeout,
	}
}

// NewGetXSDV2ParamsWithContext creates a new GetXSDV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetXSDV2ParamsWithContext(ctx context.Context) *GetXSDV2Params {
	var ()
	return &GetXSDV2Params{

		Context: ctx,
	}
}

// NewGetXSDV2ParamsWithHTTPClient creates a new GetXSDV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetXSDV2ParamsWithHTTPClient(client *http.Client) *GetXSDV2Params {
	var ()
	return &GetXSDV2Params{
		HTTPClient: client,
	}
}

/*GetXSDV2Params contains all the parameters to send to the API endpoint
for the get x s d v2 operation typically these are written to a http.Request
*/
type GetXSDV2Params struct {

	/*ReportDefinitionNameVersion
	  Name and version of XSD file to download. Some XSDs only have one version. In that case version name is not needed. Some example values are DecisionManagerDetailReport, DecisionManagerTypes

	*/
	ReportDefinitionNameVersion string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get x s d v2 params
func (o *GetXSDV2Params) WithTimeout(timeout time.Duration) *GetXSDV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get x s d v2 params
func (o *GetXSDV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get x s d v2 params
func (o *GetXSDV2Params) WithContext(ctx context.Context) *GetXSDV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get x s d v2 params
func (o *GetXSDV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get x s d v2 params
func (o *GetXSDV2Params) WithHTTPClient(client *http.Client) *GetXSDV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get x s d v2 params
func (o *GetXSDV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReportDefinitionNameVersion adds the reportDefinitionNameVersion to the get x s d v2 params
func (o *GetXSDV2Params) WithReportDefinitionNameVersion(reportDefinitionNameVersion string) *GetXSDV2Params {
	o.SetReportDefinitionNameVersion(reportDefinitionNameVersion)
	return o
}

// SetReportDefinitionNameVersion adds the reportDefinitionNameVersion to the get x s d v2 params
func (o *GetXSDV2Params) SetReportDefinitionNameVersion(reportDefinitionNameVersion string) {
	o.ReportDefinitionNameVersion = reportDefinitionNameVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetXSDV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param reportDefinitionNameVersion
	if err := r.SetPathParam("reportDefinitionNameVersion", o.ReportDefinitionNameVersion); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
