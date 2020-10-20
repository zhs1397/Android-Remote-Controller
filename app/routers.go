/*
 * demo
 *
 * simple demo for testing android
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package app

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

//Route display all routes
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes contain lists of Route
type Routes []Route

//NewRouter create new Routes
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

//Index show homepage
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Android controller demo")
}

var routes = Routes{
	//    root/v1/cloud-phone/package
	Route{
		"Index",
		"GET",
		"/v1/cloud-phone/",
		Index,
	},

	Route{
		"InstallApp",
		strings.ToUpper("Post"),
		"/v1/cloud-phone/package/install-package/install",
		InstallPackage,
	},

	Route{
		"InstallApp",
		"GET",
		"/v1/cloud-phone/package/install-package/",
		InstallPackage,
	},

	Route{
		"Open",
		"GET",
		"/v1/cloud-phone/package/open-package/",
		OpenPackage,
	},

	Route{
		"Open",
		strings.ToUpper("Post"),
		"/v1/cloud-phone/package/open-package/open",
		OpenPackage,
	},

	Route{
		"ShowPackage",
		strings.ToUpper("Get"),
		"/v1/cloud-phone/package/list-package/",
		ShowPackage,
	},

	//    root/v1/cloud-phone/device
	Route{
		"ShowBattery",
		strings.ToUpper("Get"),
		"/v1/cloud-phone/device/list-battery/",
		Battery,
	},

	Route{
		"ShowScreen",
		strings.ToUpper("Get"),
		"/v1/cloud-phone/device/list-screen/",
		Screen,
	},

	Route{
		"ShowSystem",
		"GET",
		"/v1/cloud-phone/device/list-system",
		System,
	},
}
