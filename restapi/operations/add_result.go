// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AddResultHandlerFunc turns a function with the right signature into a add result handler
type AddResultHandlerFunc func(AddResultParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddResultHandlerFunc) Handle(params AddResultParams) middleware.Responder {
	return fn(params)
}

// AddResultHandler interface for that can handle valid add result params
type AddResultHandler interface {
	Handle(AddResultParams) middleware.Responder
}

// NewAddResult creates a new http.Handler for the add result operation
func NewAddResult(ctx *middleware.Context, handler AddResultHandler) *AddResult {
	return &AddResult{Context: ctx, Handler: handler}
}

/*AddResult swagger:route POST /result addResult

adds a result item

Adds a result to the system

*/
type AddResult struct {
	Context *middleware.Context
	Handler AddResultHandler
}

func (o *AddResult) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddResultParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}