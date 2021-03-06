// Code generated by go-swagger; DO NOT EDIT.

package external

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

	models "github.com/aeternity/aepp-sdk-go/v5/swagguard/node/models"
)

// NewPostTransactionParams creates a new PostTransactionParams object
// with the default values initialized.
func NewPostTransactionParams() *PostTransactionParams {
	var ()
	return &PostTransactionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostTransactionParamsWithTimeout creates a new PostTransactionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostTransactionParamsWithTimeout(timeout time.Duration) *PostTransactionParams {
	var ()
	return &PostTransactionParams{

		timeout: timeout,
	}
}

// NewPostTransactionParamsWithContext creates a new PostTransactionParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostTransactionParamsWithContext(ctx context.Context) *PostTransactionParams {
	var ()
	return &PostTransactionParams{

		Context: ctx,
	}
}

// NewPostTransactionParamsWithHTTPClient creates a new PostTransactionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostTransactionParamsWithHTTPClient(client *http.Client) *PostTransactionParams {
	var ()
	return &PostTransactionParams{
		HTTPClient: client,
	}
}

/*PostTransactionParams contains all the parameters to send to the API endpoint
for the post transaction operation typically these are written to a http.Request
*/
type PostTransactionParams struct {

	/*Body
	  The new transaction

	*/
	Body *models.Tx

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post transaction params
func (o *PostTransactionParams) WithTimeout(timeout time.Duration) *PostTransactionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post transaction params
func (o *PostTransactionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post transaction params
func (o *PostTransactionParams) WithContext(ctx context.Context) *PostTransactionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post transaction params
func (o *PostTransactionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post transaction params
func (o *PostTransactionParams) WithHTTPClient(client *http.Client) *PostTransactionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post transaction params
func (o *PostTransactionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post transaction params
func (o *PostTransactionParams) WithBody(body *models.Tx) *PostTransactionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post transaction params
func (o *PostTransactionParams) SetBody(body *models.Tx) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostTransactionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
