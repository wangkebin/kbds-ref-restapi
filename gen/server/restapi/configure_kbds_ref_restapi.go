// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"context"
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/rs/cors"

	viper "github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/wangkebin/kbds-ref-restapi/controller"
	"github.com/wangkebin/kbds-ref-restapi/gen/server/restapi/operations"

	genmodels "github.com/wangkebin/kbds-ref-restapi/gen/server/models"
	models "github.com/wangkebin/kbds-ref-restapi/models"
)

//go:generate swagger generate server --target ../../gen --name KbdsRefRestapi --spec ../../swagger.yml --principal interface{}

func configureFlags(api *operations.KbdsRefRestapiAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.KbdsRefRestapiAPI) http.Handler {
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

	l := zap.Must(zap.NewDevelopment())
	defer l.Sync()
	api.Logger = l.Sugar().Infof

	viper.SetConfigFile("./config/config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		l.Sugar().Errorf("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&models.GlobalConfig)
	if err != nil {
		l.Sugar().Errorf("Environment can't be loaded: ", err)
	}

	api.FilesHandler = operations.FilesHandlerFunc(func(params operations.FilesParams) middleware.Responder {

		return middleware.NotImplemented("operation operations.Files has not yet been implemented")
	})

	api.HealthHandler = operations.HealthHandlerFunc(func(params operations.HealthParams) middleware.Responder {
		return operations.NewHealthOK().WithPayload("200 OK")
	})

	api.SearchHandler = operations.SearchHandlerFunc(func(params operations.SearchParams) middleware.Responder {
		files, err := controller.GetFiles(context.Background(), params.Search.Search, int(params.Search.Page), int(params.Search.Pagesize), l)
		if err != nil {
			msg := err.Error()
			return operations.NewFilesDefault(500).WithPayload(&genmodels.Error{Message: &msg})
		}
		return operations.NewSearchOK().WithPayload(*files.ToResource())
	})

	api.DeleteHandler = operations.DeleteHandlerFunc(func(params operations.DeleteParams) middleware.Responder {
		note, err := controller.DeleteFile(context.Background(), params.Fileid, l)
		if err != nil {
			msg := err.Error()
			return operations.NewDeleteDefault(500).WithPayload(&genmodels.Error{Message: &msg})
		}
		return operations.NewDeleteOK().WithPayload(note)
	})

	api.DeletefilesHandler = operations.DeletefilesHandlerFunc(func(params operations.DeletefilesParams) middleware.Responder {
		finfos := make([]models.File, 0)
		for _, finfo := range params.Files {
			f := models.File{
				Name: *finfo.Name,
				Size: finfo.Size,
			}
			finfos = append(finfos, f)
		}
		note, err := controller.DeleteFiles(context.Background(), &finfos, l)
		if err != nil {
			msg := err.Error()
			return operations.NewDeletefilesDefault(500).WithPayload(&genmodels.Error{Message: &msg})
		}
		return operations.NewDeletefilesOK().WithPayload(note)
	})
	api.DuplicatesHandler = operations.DuplicatesHandlerFunc(func(params operations.DuplicatesParams) middleware.Responder {
		var finfos models.Files
		for _, finfo := range params.Finfos {
			f := models.File{
				Name: *finfo.Name,
				Size: finfo.Size,
			}
			finfos = append(finfos, f)
		}

		files, err := controller.GetDups(context.Background(), finfos, l)
		if err != nil {
			msg := err.Error()
			return operations.NewFilesDefault(500).WithPayload(&genmodels.Error{Message: &msg})
		}
		return operations.NewSearchOK().WithPayload(*files.ToResource())
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
	handleCORS := cors.New(cors.Options{
		AllowedOrigins: models.GlobalConfig.Cors,
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		},
	}).Handler
	return handleCORS(handler)
}
