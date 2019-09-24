// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/kzmake/sample-go-swagger/restapi/operations"
	"github.com/kzmake/sample-go-swagger/restapi/operations/book"
	"github.com/kzmake/sample-go-swagger/restapi/operations/user"
)

//go:generate swagger generate server --target ../../sample-go-swagger --name SampleGoSwagger --spec ../swagger/swagger.yaml

func configureFlags(api *operations.SampleGoSwaggerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SampleGoSwaggerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.UrlformConsumer = runtime.DiscardConsumer

	api.XMLConsumer = runtime.XMLConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.XMLProducer = runtime.XMLProducer()

	// Applies when the "api_key" header is set
	api.APIKeyAuth = func(token string) (interface{}, error) {
		return nil, errors.NotImplemented("api key auth (api_key) api_key from header param [api_key] has not yet been implemented")
	}
	api.BookshelfAuthAuth = func(token string, scopes []string) (interface{}, error) {
		return nil, errors.NotImplemented("oauth2 bearer auth (bookshelf_auth) has not yet been implemented")
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	if api.BookAddBookHandler == nil {
		api.BookAddBookHandler = book.AddBookHandlerFunc(func(params book.AddBookParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation book.AddBook has not yet been implemented")
		})
	}
	if api.UserCreateUserHandler == nil {
		api.UserCreateUserHandler = user.CreateUserHandlerFunc(func(params user.CreateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.CreateUser has not yet been implemented")
		})
	}
	if api.BookDeleteBookHandler == nil {
		api.BookDeleteBookHandler = book.DeleteBookHandlerFunc(func(params book.DeleteBookParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation book.DeleteBook has not yet been implemented")
		})
	}
	if api.UserDeleteUserHandler == nil {
		api.UserDeleteUserHandler = user.DeleteUserHandlerFunc(func(params user.DeleteUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.DeleteUser has not yet been implemented")
		})
	}
	if api.BookGetBookByIDHandler == nil {
		api.BookGetBookByIDHandler = book.GetBookByIDHandlerFunc(func(params book.GetBookByIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation book.GetBookByID has not yet been implemented")
		})
	}
	if api.UserGetUserByNameHandler == nil {
		api.UserGetUserByNameHandler = user.GetUserByNameHandlerFunc(func(params user.GetUserByNameParams) middleware.Responder {
			return middleware.NotImplemented("operation user.GetUserByName has not yet been implemented")
		})
	}
	if api.UserLoginUserHandler == nil {
		api.UserLoginUserHandler = user.LoginUserHandlerFunc(func(params user.LoginUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.LoginUser has not yet been implemented")
		})
	}
	if api.UserLogoutUserHandler == nil {
		api.UserLogoutUserHandler = user.LogoutUserHandlerFunc(func(params user.LogoutUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.LogoutUser has not yet been implemented")
		})
	}
	if api.BookUpdateBookHandler == nil {
		api.BookUpdateBookHandler = book.UpdateBookHandlerFunc(func(params book.UpdateBookParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation book.UpdateBook has not yet been implemented")
		})
	}
	if api.BookUpdateBookWithFormHandler == nil {
		api.BookUpdateBookWithFormHandler = book.UpdateBookWithFormHandlerFunc(func(params book.UpdateBookWithFormParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation book.UpdateBookWithForm has not yet been implemented")
		})
	}
	if api.UserUpdateUserHandler == nil {
		api.UserUpdateUserHandler = user.UpdateUserHandlerFunc(func(params user.UpdateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UpdateUser has not yet been implemented")
		})
	}

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
