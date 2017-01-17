package routes

import (
	"net/http"

	"encoding/json"

	"github.com/VYNYL/go-http-scaffold/mmodels"
	"github.com/VYNYL/go-http-scaffold/prehandle"
	"github.com/VYNYL/go-http-scaffold/router"
	"github.com/gorilla/mux"
)

// GetHello router.Route
// Path: "/hello",
// Method: "GET"
var GetHello = &router.Route{
	Path:       "/hello",
	Method:     "GET",
	Handler:    http.HandlerFunc(handleGetHello),
	Prehandler: []prehandle.Prehandler{prehandle.SetJSON},
}

func handleGetHello(w http.ResponseWriter, rq *http.Request) {

	json.NewEncoder(w).Encode(map[string]string{
		"Message": "Hello there!",
	})

}

// GetHelloByName router.Route
// Path: "/hello",
// Method: "GET"
var GetHelloByName = &router.Route{
	Path:       "/hello/{name}",
	Method:     "GET",
	Handler:    http.HandlerFunc(handleGetHelloByName),
	Prehandler: []prehandle.Prehandler{prehandle.SetJSON},
}

func handleGetHelloByName(w http.ResponseWriter, rq *http.Request) {

	vars := mux.Vars(rq)

	model := &mmodels.Hello{
		Name: vars["name"],
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"Model":   model,
		"Message": model.Message(),
	})

}
