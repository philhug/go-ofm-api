package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	app "github.com/casualjim/go-app"
	loads "github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"

	ofmdb "github.com/philhug/go-ofm-api"
	"github.com/philhug/go-ofm-api/api/handlers"
	"github.com/philhug/go-ofm-api/gen/restapi"
	"github.com/philhug/go-ofm-api/gen/restapi/operations"
)

// This file was generated by the swagger tool.
// Make sure not to overwrite this file after you generated it because all your edits would be lost!

func main() {

	app, err := app.New("ofmdbapi")
	if err != nil {
		logrus.Fatalln(err)
	}

	log := app.Logger()
	cfg := app.Config()
	cfg.SetDefault("mysql.host", "localhost")
	cfg.SetDefault("dsn", "root:password@tcp(127.0.0.1:3306)/ofm?parseTime=true")

	rt, err := ofmdb.NewRuntime(app)
	if err != nil {
		log.Fatalln(err)
	}

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewOfmdbAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "openflightmaps"
	parser.LongDescription = "Open flightmap data api.\nNote: This is a simple conversion of the current model used by the client for compatbility. It's not indend for public use.\n"

	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	api.BasicAuthAuth = handlers.BasicAuthAuth
	api.APIAuthorizer = handlers.NewAuthorizer(rt)
	api.InformationGetRegionsHandler = handlers.NewInformationGetRegionsHandler(rt)
	api.InformationGetUserInfoHandler = handlers.NewInformationGetUserInfoHandler(rt)
	api.NodesGetNodeHandler = handlers.NewNodesGetNodeHandler(rt)
	api.NodesGetMultipleNodesHandler = handlers.NewNodesGetMultipleNodesHandler(rt)
	api.NodesSearchNodeHandler = handlers.NewNodeSearchNodeHandler(rt)
	api.NativeClientGetBlobHandler = handlers.NewNativeClientGetBlobHandler(rt)

	//server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
