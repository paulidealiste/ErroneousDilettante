package webapp

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDilletanteHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	hndlr := http.HandlerFunc(dilettanteHandler)

	hndlr.ServeHTTP(rr, req)
	fmt.Println(rr.Body)

	if rr.Code != http.StatusOK {
		t.Errorf("The handler returned wrong status code (%v instead of %v)", rr.Code, http.StatusOK)
	}
}

func TestErroneusHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/erroneus", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	hndlr := http.HandlerFunc(erroneusHandler)

	hndlr.ServeHTTP(rr, req)
	fmt.Println(rr.Body)

	if rr.Code != http.StatusOK {
		t.Errorf("The handler returned wrong status code (%v instead of %v)", rr.Code, http.StatusOK)
	}
}
