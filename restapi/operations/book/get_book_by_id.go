// Code generated by go-swagger; DO NOT EDIT.

package book

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetBookByIDHandlerFunc turns a function with the right signature into a get book by Id handler
type GetBookByIDHandlerFunc func(GetBookByIDParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetBookByIDHandlerFunc) Handle(params GetBookByIDParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetBookByIDHandler interface for that can handle valid get book by Id params
type GetBookByIDHandler interface {
	Handle(GetBookByIDParams, interface{}) middleware.Responder
}

// NewGetBookByID creates a new http.Handler for the get book by Id operation
func NewGetBookByID(ctx *middleware.Context, handler GetBookByIDHandler) *GetBookByID {
	return &GetBookByID{Context: ctx, Handler: handler}
}

/*GetBookByID swagger:route GET /book/{bookId} book getBookById

Find book by ID

Returns a single book

*/
type GetBookByID struct {
	Context *middleware.Context
	Handler GetBookByIDHandler
}

func (o *GetBookByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetBookByIDParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
