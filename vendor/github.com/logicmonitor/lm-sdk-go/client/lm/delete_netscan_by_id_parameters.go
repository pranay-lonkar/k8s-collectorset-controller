// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteNetscanByIDParams creates a new DeleteNetscanByIDParams object
// with the default values initialized.
func NewDeleteNetscanByIDParams() *DeleteNetscanByIDParams {
	var ()
	return &DeleteNetscanByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNetscanByIDParamsWithTimeout creates a new DeleteNetscanByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteNetscanByIDParamsWithTimeout(timeout time.Duration) *DeleteNetscanByIDParams {
	var ()
	return &DeleteNetscanByIDParams{

		timeout: timeout,
	}
}

// NewDeleteNetscanByIDParamsWithContext creates a new DeleteNetscanByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteNetscanByIDParamsWithContext(ctx context.Context) *DeleteNetscanByIDParams {
	var ()
	return &DeleteNetscanByIDParams{

		Context: ctx,
	}
}

// NewDeleteNetscanByIDParamsWithHTTPClient creates a new DeleteNetscanByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteNetscanByIDParamsWithHTTPClient(client *http.Client) *DeleteNetscanByIDParams {
	var ()
	return &DeleteNetscanByIDParams{
		HTTPClient: client,
	}
}

/*DeleteNetscanByIDParams contains all the parameters to send to the API endpoint
for the delete netscan by Id operation typically these are written to a http.Request
*/
type DeleteNetscanByIDParams struct {

	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete netscan by Id params
func (o *DeleteNetscanByIDParams) WithTimeout(timeout time.Duration) *DeleteNetscanByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete netscan by Id params
func (o *DeleteNetscanByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete netscan by Id params
func (o *DeleteNetscanByIDParams) WithContext(ctx context.Context) *DeleteNetscanByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete netscan by Id params
func (o *DeleteNetscanByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete netscan by Id params
func (o *DeleteNetscanByIDParams) WithHTTPClient(client *http.Client) *DeleteNetscanByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete netscan by Id params
func (o *DeleteNetscanByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete netscan by Id params
func (o *DeleteNetscanByIDParams) WithID(id int32) *DeleteNetscanByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete netscan by Id params
func (o *DeleteNetscanByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNetscanByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}