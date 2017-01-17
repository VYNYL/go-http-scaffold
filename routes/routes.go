package routes

import (
	"net/http"

	"github.com/VYNYL/mailer/prehandle"
	"github.com/VYNYL/mailer/router"
)

// PostSend router.Route
// Path: "/send",
// Method: "POST"
var PostSend = &router.Route{
	Path:       "/send",
	Method:     "POST",
	Handler:    http.HandlerFunc(handlePostSend),
	Prehandler: []prehandle.Prehandler{prehandle.SetJSON},
}

func handlePostSend(w http.ResponseWriter, rq *http.Request) {
}
