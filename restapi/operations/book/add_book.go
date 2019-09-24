// Code generated by go-swagger; DO NOT EDIT.

package book

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AddBookHandlerFunc turns a function with the right signature into a add book handler
type AddBookHandlerFunc func(AddBookParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn AddBookHandlerFunc) Handle(params AddBookParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// AddBookHandler interface for that can handle valid add book params
type AddBookHandler interface {
	Handle(AddBookParams, interface{}) middleware.Responder
}

// NewAddBook creates a new http.Handler for the add book operation
func NewAddBook(ctx *middleware.Context, handler AddBookHandler) *AddBook {
	return &AddBook{Context: ctx, Handler: handler}
}

/*AddBook swagger:route POST /book book addBook

Add a new book to the bookshlef

*/
type AddBook struct {
	Context *middleware.Context
	Handler AddBookHandler
}

func (o *AddBook) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddBookParams()

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
