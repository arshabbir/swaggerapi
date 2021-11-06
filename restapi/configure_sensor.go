// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/arshabbir/sensor/models"
	"github.com/arshabbir/sensor/restapi/operations"
	"github.com/arshabbir/sensor/restapi/operations/add_sensordata"
	"github.com/arshabbir/sensor/restapi/operations/sensor_data"
)

//go:generate swagger generate server --target ..\..\sensorapp --name Sensor --spec ..\swagger.yml --principal interface{}

func configureFlags(api *operations.SensorAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SensorAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	//if api.SensorDataGetHandler == nil {
	api.SensorDataGetHandler = sensor_data.GetHandlerFunc(func(params sensor_data.GetParams) middleware.Responder {
		desc := "description"
		return sensor_data.NewGetOK().WithPayload([]*models.Reading{&models.Reading{ID: 7, Description: &desc, Value: rand.Float32()}})
	})
	//}
	//if api.AddSensordataPostAddHandler == nil {
	api.AddSensordataPostAddHandler = add_sensordata.PostAddHandlerFunc(func(params add_sensordata.PostAddParams) middleware.Responder {
		//return middleware.NotImplemented("operation add_sensordata.PostAdd has not yet been implemented")

		fmt.Println("GetHandler invoked ")
		desc := "Sensordata"
		reading := models.Reading{ID: 7, Description: &desc, Value: rand.Float32()}

		return add_sensordata.NewPostAddCreated().WithPayload(&reading)

	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
