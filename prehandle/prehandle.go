package prehandle

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/VYNYL/mailer/merror"
)

// Prehandler type is exactly the same as http.HandlerFunc except that a return bool is expected to indicate success/failure
type Prehandler func(http.ResponseWriter, *http.Request) bool

// PreHandle accepts an http.HandlerFunc and preprocesses it with n-prehandlers.
// If any prehandler returns false, the process will be aborted and the handler will never be reached
func PreHandle(handle http.HandlerFunc, prehandlers ...Prehandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		for _, pre := range prehandlers {
			if !pre(w, r) {
				// The prehandler signals a halt
				return
			}
		}
		handle(w, r)
	}

}

// SetJSON sets the Content-Type to application/json
func SetJSON(w http.ResponseWriter, r *http.Request) bool {
	w.Header().Set("Content-Type", "application/json")
	return true
}

// RequireBody forces a body to exist with a maximum length limit
// If the body does not exist, an http.StatusBadRequest is returned. This is required for POST requests
// This prehandler protects against overflows and null-pointer exceptions
func RequireBody(limit int64) Prehandler {
	return func(w http.ResponseWriter, r *http.Request) bool {
		if r.Body == nil {
			merror.Respond(w, &merror.MailerSimpleError{
				Code:    http.StatusBadRequest,
				Message: "EMPTY_BODY",
			})
			return false
		}

		body, err := ioutil.ReadAll(io.LimitReader(r.Body, limit))
		if err != nil {
			merror.Respond(w, &merror.MailerSimpleError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
			return false
		}

		r.Header.Set("X-Body", string(body))

		return true
	}
}
