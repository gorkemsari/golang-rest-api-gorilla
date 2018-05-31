package route

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorkemsari/golang-rest-api/handler"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"CityAll", "GET", "/cities", handler.CityAll},
	Route{"City", "GET", "/cities/{id}", handler.City},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true).PathPrefix("/api/v1").Subrouter()

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
