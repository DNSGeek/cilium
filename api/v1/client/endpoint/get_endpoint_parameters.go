package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetEndpointParams creates a new GetEndpointParams object
// with the default values initialized.
func NewGetEndpointParams() *GetEndpointParams {

	return &GetEndpointParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEndpointParamsWithTimeout creates a new GetEndpointParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEndpointParamsWithTimeout(timeout time.Duration) *GetEndpointParams {

	return &GetEndpointParams{

		timeout: timeout,
	}
}

// NewGetEndpointParamsWithContext creates a new GetEndpointParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEndpointParamsWithContext(ctx context.Context) *GetEndpointParams {

	return &GetEndpointParams{

		Context: ctx,
	}
}

// NewGetEndpointParamsWithHTTPClient creates a new GetEndpointParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEndpointParamsWithHTTPClient(client *http.Client) *GetEndpointParams {

	return &GetEndpointParams{
		HTTPClient: client,
	}
}

/*GetEndpointParams contains all the parameters to send to the API endpoint
for the get endpoint operation typically these are written to a http.Request
*/
type GetEndpointParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get endpoint params
func (o *GetEndpointParams) WithTimeout(timeout time.Duration) *GetEndpointParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get endpoint params
func (o *GetEndpointParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get endpoint params
func (o *GetEndpointParams) WithContext(ctx context.Context) *GetEndpointParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get endpoint params
func (o *GetEndpointParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get endpoint params
func (o *GetEndpointParams) WithHTTPClient(client *http.Client) *GetEndpointParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get endpoint params
func (o *GetEndpointParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetEndpointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
