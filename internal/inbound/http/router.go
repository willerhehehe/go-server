// Package http DCSwitch API
//
// The purpose of this service is to provide an application to help manage db switch task
//
//      Schemes: http
//      Host: 127.0.0.1
//      Version: 0.0.1
//
//      Security:
//      - api_key:
//
//      SecurityDefinitions:
//      api_key:
//           type: apiKey
//           name: TOKEN
//           in: header
//
// swagger:meta
package http

import (
	handlers "dcswitch/internal/inbound/http/handlers"
	"dcswitch/internal/inbound/http/middlewares"
	"dcswitch/internal/static"
	"dcswitch/pkg/swagger"
	"github.com/gorilla/mux"
	"net/http"
)

// InitHandlers API入口
func InitHandlers() *mux.Router {

	r := mux.NewRouter()
	r.Use(middlewares.RecoverWrap, middlewares.SlowRequestMiddleware, middlewares.LoggingMiddleware, middlewares.CORSMiddleware)

	r.HandleFunc("/healthz", handlers.HealthCheck).Methods("GET")

	r.HandleFunc("/{id:[0-9]+}", handlers.GetMockSlow).Methods("GET")

	// biz
	r.HandleFunc("/switch/versions", handlers.GetAllSwitchVersions).Methods("GET")
	r.HandleFunc("/switch/version/name/{id:[0-9]+}", handlers.EditSwitchVersionName).Methods("PATCH")
	r.HandleFunc("/task/switch/module/detail", handlers.CreateModuleDetailTask).Methods("POST")

	return r
}

// InitDocHandler API文档Handler
func InitDocHandler(r *mux.Router) {
	fs := http.FileServer(http.FS(static.SwaggerFS))
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", fs))
	r.HandleFunc("/docs", swagger.WrappedDocsHandler(swagger.DocsOpts{
		DocsSwaggerJsonURL: "/static/swagger.json",
		DocsJsURL:          "/static/swagger-ui-bundle.js",
		DocsCssURL:         "/static/swagger-ui.css",
		IconURL:            "/static/dcswitch.png",
	})).Methods("GET")
	r.HandleFunc("/redoc", swagger.WrappedReDocHandler(swagger.RedocOpts{
		SpecSwaggerJsonURL: "/static/swagger.json",
		RedocJsURL:         "/static/redoc.standalone.js",
		RedocCssURL:        "/static/redoc.standalone.css",
		IconURL:            "/static/dcswitch.png",
	})).Methods("GET")
}

func HandleFuncWithAuth(path string, f func(http.ResponseWriter, *http.Request), auth func(http.ResponseWriter, *http.Request), methods ...string) {

}
