// Code generated by go-swagger; DO NOT EDIT.

package book

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/kzmake/sample-go-swagger/models"
)

// GetBookByIDOKCode is the HTTP code returned for type GetBookByIDOK
const GetBookByIDOKCode int = 200

/*GetBookByIDOK successful operation

swagger:response getBookByIdOK
*/
type GetBookByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Book `json:"body,omitempty"`
}

// NewGetBookByIDOK creates GetBookByIDOK with default headers values
func NewGetBookByIDOK() *GetBookByIDOK {

	return &GetBookByIDOK{}
}

// WithPayload adds the payload to the get book by Id o k response
func (o *GetBookByIDOK) WithPayload(payload *models.Book) *GetBookByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get book by Id o k response
func (o *GetBookByIDOK) SetPayload(payload *models.Book) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBookByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetBookByIDBadRequestCode is the HTTP code returned for type GetBookByIDBadRequest
const GetBookByIDBadRequestCode int = 400

/*GetBookByIDBadRequest Invalid ID supplied

swagger:response getBookByIdBadRequest
*/
type GetBookByIDBadRequest struct {
}

// NewGetBookByIDBadRequest creates GetBookByIDBadRequest with default headers values
func NewGetBookByIDBadRequest() *GetBookByIDBadRequest {

	return &GetBookByIDBadRequest{}
}

// WriteResponse to the client
func (o *GetBookByIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetBookByIDNotFoundCode is the HTTP code returned for type GetBookByIDNotFound
const GetBookByIDNotFoundCode int = 404

/*GetBookByIDNotFound Book not found

swagger:response getBookByIdNotFound
*/
type GetBookByIDNotFound struct {
}

// NewGetBookByIDNotFound creates GetBookByIDNotFound with default headers values
func NewGetBookByIDNotFound() *GetBookByIDNotFound {

	return &GetBookByIDNotFound{}
}

// WriteResponse to the client
func (o *GetBookByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}