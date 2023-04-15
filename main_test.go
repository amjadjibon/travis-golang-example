package main

import (
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestIndex(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	index(w, r, nil)

	if w.Code != 200 {
		t.Errorf("Index() = %d; want 200", w.Code)
	}

	if w.Body.String() != "Welcome!\n" {
		t.Errorf("Index() = %s; want Welcome!\n", w.Body.String())
	}
}

func TestHello(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/hello/gopher", nil)
	p := httprouter.Params{httprouter.Param{Key: "name", Value: "gopher"}}

	hello(w, r, p)

	if w.Code != 200 {
		t.Errorf("Hello() = %d; want 200", w.Code)
	}

	if w.Body.String() != "Hello, gopher!\n" {
		t.Errorf("Hello() = %s; want Hello, gopher!\n", w.Body.String())
	}
}
