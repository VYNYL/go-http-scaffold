package router

import (
	"net/http"

	"github.com/VYNYL/mailer/prehandle"
	"github.com/gorilla/mux"
)

// Route contains route information for multiplexing
type Route struct {
	http.Handler
	Prehandler []prehandle.Prehandler
	Method     string
	Path       string
}

// Test ceates a new mux.Router instance for easy testing. This additionally allows support for Mux params
func (route *Route) Test(w http.ResponseWriter, r *http.Request) {
	m := mux.NewRouter()
	m.Handle(route.Path, prehandle.PreHandle(route.Handler.(http.HandlerFunc), route.Prehandler...)).Methods(route.Method)
	m.ServeHTTP(w, r)
}
