package route

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorkemsari/golang-rest-api/handler"
	mw "github.com/gorkemsari/golang-rest-api/middleware"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Auth        bool
}

type Routes []Route

var routes = Routes{
	Route{"Token", "POST", "/auth/token", handler.Token, false},
	Route{"CityAll", "GET", "/cities", handler.CityAll, true},
	Route{"City", "GET", "/cities/{id}", handler.City, true},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true).PathPrefix("/api/v1").Subrouter()

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		if route.Auth {
			mw.Auth(handler)
		}

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
