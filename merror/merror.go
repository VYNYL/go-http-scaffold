package merror

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type IMailerError interface {
	Parse() (int, string)
}

type MailerSimpleError struct {
	Code    int
	Message interface{}
	Log     string        `json:"-"`
	Req     *http.Request `json:"-"`
}

// Log writes text to the error log at ./log/err.log. If the error log does not exist, it creates a new one.
func Log(lines ...string) error {
	path, err := filepath.Abs("./log/err.log")
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		return err
	}

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		return err
	}
	defer f.Close()

	log.SetOutput(f)

	for _, line := range lines {
		log.Println(line)
	}

	return nil
}

// Parse will json.Marshal an error and, if it has a Log value, prints it to the error log
func (simple *MailerSimpleError) Parse() (int, string) {
	if simple.Log != "" {
		Log(fmt.Sprintf("%s\n", simple.Log),
			fmt.Sprintf("Dump %+v\n\n", simple.Req))
	}

	encoded, err := json.Marshal(simple)
	if err != nil {
		fmt.Println("Problem encoding error", err)
	}

	return simple.Code, string(encoded)
}

// Respond with a marshalled error
func Respond(w http.ResponseWriter, err IMailerError) {
	code, msg := err.Parse()
	w.WriteHeader(code)
	fmt.Fprintln(w, msg)
}


// Handle a server 500 error by reporting a problem to the logger
func ServerError(w http.ResponseWriter, rq *http.Request, err error) {
	Respond(w, &MailerSimpleError{
		Code: http.StatusInternalServerError,
		Log:  err.Error(),
		Req:  rq,
	})
}