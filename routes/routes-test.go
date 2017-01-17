package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGetHello tests the route GetHello
func TestGetHello(t *testing.T) {
	w := httptest.NewRecorder()

	// If you wish to test a POST request, replace the third argument with a bytes.NewBuffer([]byte())
	r, err := http.NewRequest(GetHello.Method, GetHello.Path, nil)

	if err != nil {
		t.Fatal(err)
	}

	GetHello.Test(w, r)

	if w.Code != http.StatusAccepted {
		t.Fatalf("%d status code when %d expected", w.Code, http.StatusAccepted)
	}
}

// TestGetHelloByName tests the route GetHelloByName
func TestGetHelloByName(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest(GetHello.Method, GetHello.Path, nil)

	if err != nil {
		t.Fatal(err)
	}

	GetHelloByName.Test(w, r)

	if w.Code != http.StatusAccepted {
		t.Fatalf("%d status code when %d expected", w.Code, http.StatusAccepted)
	}
}
