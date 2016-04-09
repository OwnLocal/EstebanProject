package router

import (
	"github.com/elohr/OwnLocal/controllers/business"
	"github.com/julienschmidt/httprouter"
)

var Router *httprouter.Router

// init creates a Router that will be available on this package
// and assigns the available routes to it
func init() {
	// Router takes care of not allowed methods or 404 errors
	Router = httprouter.New()

	// Businesses
	Router.GET("/businesses/:search", business.List)
	Router.GET("/businesses/:search/:from", business.List)
	Router.GET("/businesses/:search/:from/:size", business.List)
	Router.GET("/business/:id", business.Get)
}

// TODO: if authentication is required, use the following function with a middleware
// TODO: that checks for credentials whenever it is required (eg. Router.GET("/business/:id", use(business.Get, auth.IsLoggedIn)))
//func use(handler http.HandlerFunc, mid ...func(http.Handler) http.HandlerFunc) http.HandlerFunc {
//	for _, m := range mid {
//		handler = m(handler)
//	}
//
//	return handler
//}
