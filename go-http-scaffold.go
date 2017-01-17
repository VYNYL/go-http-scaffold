package main

import (
	"log"
	"net/http"

	"os"

	"github.com/VYNYL/go-http-scaffold/prehandle"
	"github.com/VYNYL/go-http-scaffold/router"
	"github.com/VYNYL/go-http-scaffold/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func doRoute(r *mux.Router, route *router.Route) {
	r.Handle(route.Path, prehandle.PreHandle(route.Handler.(http.HandlerFunc), route.Prehandler...)).Methods(route.Method)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env. Please ensure .env file exists")
		return
	}

	port := os.Getenv("LISTEN_PORT")

	r := mux.NewRouter()

	doRoute(r, routes.GetHello)
	doRoute(r, routes.GetHelloByName)

	http.ListenAndServe(":"+port, r)
}
