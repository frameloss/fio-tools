// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetFeeVotesMultiplierProducerHandlerFunc turns a function with the right signature into a get fee votes multiplier producer handler
type GetFeeVotesMultiplierProducerHandlerFunc func(GetFeeVotesMultiplierProducerParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetFeeVotesMultiplierProducerHandlerFunc) Handle(params GetFeeVotesMultiplierProducerParams) middleware.Responder {
	return fn(params)
}

// GetFeeVotesMultiplierProducerHandler interface for that can handle valid get fee votes multiplier producer params
type GetFeeVotesMultiplierProducerHandler interface {
	Handle(GetFeeVotesMultiplierProducerParams) middleware.Responder
}

// NewGetFeeVotesMultiplierProducer creates a new http.Handler for the get fee votes multiplier producer operation
func NewGetFeeVotesMultiplierProducer(ctx *middleware.Context, handler GetFeeVotesMultiplierProducerHandler) *GetFeeVotesMultiplierProducer {
	return &GetFeeVotesMultiplierProducer{Context: ctx, Handler: handler}
}

/* GetFeeVotesMultiplierProducer swagger:route GET /fee/votes/multiplier/{producer} getFeeVotesMultiplierProducer

list a producer's multiplier

*/
type GetFeeVotesMultiplierProducer struct {
	Context *middleware.Context
	Handler GetFeeVotesMultiplierProducerHandler
}

func (o *GetFeeVotesMultiplierProducer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetFeeVotesMultiplierProducerParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetFeeVotesMultiplierProducerOKBody get fee votes multiplier producer o k body
//
// swagger:model GetFeeVotesMultiplierProducerOKBody
type GetFeeVotesMultiplierProducerOKBody struct {

	// multiplier
	Multiplier float64 `json:"multiplier,omitempty"`
}

// Validate validates this get fee votes multiplier producer o k body
func (o *GetFeeVotesMultiplierProducerOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get fee votes multiplier producer o k body based on context it is used
func (o *GetFeeVotesMultiplierProducerOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetFeeVotesMultiplierProducerOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetFeeVotesMultiplierProducerOKBody) UnmarshalBinary(b []byte) error {
	var res GetFeeVotesMultiplierProducerOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}