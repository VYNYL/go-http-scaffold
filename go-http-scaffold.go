package main

import (
	"log"
	"net/http"

	"os"

	"github.com/VYNYL/mailer/prehandle"
	"github.com/VYNYL/mailer/router"
	"github.com/VYNYL/mailer/routes"
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

	doRoute(r, routes.PostSend)

	http.ListenAndServe(":"+port, r)
}
