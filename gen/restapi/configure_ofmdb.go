// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"github.com/philhug/go-ofmapi-server/gen/restapi/operations"
	"github.com/philhug/go-ofmapi-server/gen/restapi/operations/information"
	"github.com/philhug/go-ofmapi-server/gen/restapi/operations/native_client"
	"github.com/philhug/go-ofmapi-server/gen/restapi/operations/nodes"
)

//go:generate swagger generate server --target .. --name Ofmdb --spec ../../swagger.yaml --exclude-main

func configureFlags(api *operations.OfmdbAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.OfmdbAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.BinConsumer = runtime.ByteStreamConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.BinProducer = runtime.ByteStreamProducer()

	// Applies when the Authorization header is set with the Basic scheme
	api.BasicAuthAuth = func(user string, pass string) (interface{}, error) {
		return nil, errors.NotImplemented("basic auth  (BasicAuth) has not yet been implemented")
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	api.AllDbsHandler = operations.AllDbsHandlerFunc(func(params operations.AllDbsParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation .AllDbs has not yet been implemented")
	})
	api.NodesAllDocsHandler = nodes.AllDocsHandlerFunc(func(params nodes.AllDocsParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation nodes.AllDocs has not yet been implemented")
	})
	api.NodesChangesDbHandler = nodes.ChangesDbHandlerFunc(func(params nodes.ChangesDbParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation nodes.ChangesDb has not yet been implemented")
	})
	api.NativeClientCreateBlobHandler = native_client.CreateBlobHandlerFunc(func(params native_client.CreateBlobParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation native_client.CreateBlob has not yet been implemented")
	})
	api.NodesCreateNodeHandler = nodes.CreateNodeHandlerFunc(func(params nodes.CreateNodeParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation nodes.CreateNode has not yet been implemented")
	})
	api.NativeClientCreateOrgNodeHandler = native_client.CreateOrgNodeHandlerFunc(func(params native_client.CreateOrgNodeParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation native_client.CreateOrgNode has not yet been implemented")
	})
	api.NativeClientCreateUserNodeHandler = native_client.CreateUserNodeHandlerFunc(func(params native_client.CreateUserNodeParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation native_client.CreateUserNode has not yet been implemented")
	})
	api.NodesDeleteNodeHandler = nodes.DeleteNodeHandlerFunc(func(params nodes.DeleteNodeParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation nodes.DeleteNode has not yet been implemented")
	})
	api.NativeClientDeleteOrgNodeHandler = native_client.DeleteOrgNodeHandlerFunc(func(params native_client.DeleteOrgNodeParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation native_client.DeleteOrgNode has not yet been implemented")
	})
	api.NativeClientDeleteUserNodeHandler = native_client.DeleteUserNodeHandlerFunc(func(params native_client.DeleteUserNodeParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation native_client.DeleteUserNode has not yet been implemented")
	})
	api.NativeClientGetBlobHandler = native_client.GetBlobHandlerFunc(func(params native_client.GetBlobParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation native_client.GetBlob has not yet been implemented")
	})
	api.NodesGetMultipleNodesHandler = nodes.GetMultipleNodesHandlerFunc(func(params nodes.GetMultipleNodesParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation nodes.GetMultipleNodes has not yet been implemented")
	})
	api.NodesGetNodeHandler = nodes.GetNodeHandlerFunc(func(params nodes.GetNodeParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation nodes.GetNode has not yet been implemented")
	})
	api.NativeClientGetOrgNodeHandler = native_client.GetOrgNodeHandlerFunc(func(params native_client.GetOrgNodeParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation native_client.GetOrgNode has not yet been implemented")
	})
	api.InformationGetRegionsHandler = information.GetRegionsHandlerFunc(func(params information.GetRegionsParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation information.GetRegions has not yet been implemented")
	})
	api.InformationGetUserInfoHandler = information.GetUserInfoHandlerFunc(func(params information.GetUserInfoParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation information.GetUserInfo has not yet been implemented")
	})
	api.NativeClientGetUserNodeHandler = native_client.GetUserNodeHandlerFunc(func(params native_client.GetUserNodeParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation native_client.GetUserNode has not yet been implemented")
	})
	api.NativeClientGetUserPermissionNodeHandler = native_client.GetUserPermissionNodeHandlerFunc(func(params native_client.GetUserPermissionNodeParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation native_client.GetUserPermissionNode has not yet been implemented")
	})
	api.NodesPatchNodeHandler = nodes.PatchNodeHandlerFunc(func(params nodes.PatchNodeParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation nodes.PatchNode has not yet been implemented")
	})
	api.NodesPutDbHandler = nodes.PutDbHandlerFunc(func(params nodes.PutDbParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation nodes.PutDb has not yet been implemented")
	})
	api.NodesSearchNodeHandler = nodes.SearchNodeHandlerFunc(func(params nodes.SearchNodeParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation nodes.SearchNode has not yet been implemented")
	})
	api.NativeClientSearchOrgNodeHandler = native_client.SearchOrgNodeHandlerFunc(func(params native_client.SearchOrgNodeParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation native_client.SearchOrgNode has not yet been implemented")
	})
	api.NativeClientSearchUserNodeHandler = native_client.SearchUserNodeHandlerFunc(func(params native_client.SearchUserNodeParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation native_client.SearchUserNode has not yet been implemented")
	})
	api.NodesUpdateNodeHandler = nodes.UpdateNodeHandlerFunc(func(params nodes.UpdateNodeParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation nodes.UpdateNode has not yet been implemented")
	})
	api.NativeClientUpdateOrgNodeHandler = native_client.UpdateOrgNodeHandlerFunc(func(params native_client.UpdateOrgNodeParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation native_client.UpdateOrgNode has not yet been implemented")
	})
	api.NativeClientUpdateUserNodeHandler = native_client.UpdateUserNodeHandlerFunc(func(params native_client.UpdateUserNodeParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation native_client.UpdateUserNode has not yet been implemented")
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
func configureServer(s *graceful.Server, scheme, addr string) {
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
