// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/divanvisagie/go-inventory-tracker/server/models"
	"github.com/divanvisagie/go-inventory-tracker/server/restapi/operations"
	"github.com/divanvisagie/go-inventory-tracker/server/restapi/operations/items"
	"github.com/divanvisagie/go-inventory-tracker/server/services"
	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

var itemService = services.NewItemService()

//go:generate swagger generate server --target .. --name InventoryTracker --spec ../swagger.yml

func configureFlags(api *operations.InventoryTrackerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.InventoryTrackerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.ItemsGetHandler = items.GetHandlerFunc(func(params items.GetParams) middleware.Responder {
		payload := itemService.Get(&params)
		return items.NewGetOK().WithPayload(payload)
	})

	api.ItemsAddOneHandler = items.AddOneHandlerFunc(func(params items.AddOneParams) middleware.Responder {
		if err := itemService.Add(params.Body); err != nil {
			return items.NewAddOneDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return items.NewAddOneCreated()
	})

	api.ItemsDestroyOneHandler = items.DestroyOneHandlerFunc(func(params items.DestroyOneParams) middleware.Responder {
		if err := itemService.Remove(params.ID); err != nil {
			return items.NewDestroyOneDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return items.NewDestroyOneNoContent()
	})
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
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
