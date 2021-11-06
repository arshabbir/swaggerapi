// Code generated by go-swagger; DO NOT EDIT.

package sensor_data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/arshabbir/sensor/models"
)

// GetOKCode is the HTTP code returned for type GetOK
const GetOKCode int = 200

/*GetOK list the sensor data

swagger:response getOK
*/
type GetOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Reading `json:"body,omitempty"`
}

// NewGetOK creates GetOK with default headers values
func NewGetOK() *GetOK {

	return &GetOK{}
}

// WithPayload adds the payload to the get o k response
func (o *GetOK) WithPayload(payload []*models.Reading) *GetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get o k response
func (o *GetOK) SetPayload(payload []*models.Reading) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Reading, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetDefault API error response

swagger:response getDefault
*/
type GetDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Apierror `json:"body,omitempty"`
}

// NewGetDefault creates GetDefault with default headers values
func NewGetDefault(code int) *GetDefault {
	if code <= 0 {
		code = 500
	}

	return &GetDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get default response
func (o *GetDefault) WithStatusCode(code int) *GetDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get default response
func (o *GetDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get default response
func (o *GetDefault) WithPayload(payload *models.Apierror) *GetDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get default response
func (o *GetDefault) SetPayload(payload *models.Apierror) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
