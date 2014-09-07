/*
JSON request body handler testcases
*/

package http

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var okHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
})

var echoNameHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	payload := GetInput(r)
	if payload == nil {
		http.Error(w, "not ok", 400)
		return
	}

	w.Write([]byte(payload["name"].(string)))
})

var notHandler = func(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("not "))
		h.ServeHTTP(w, r)
	})
}

func newRequest(method, url, content string, t *testing.T) *http.Request {
	req, err := http.NewRequest(method, url, strings.NewReader(content))
	if err != nil {
		t.Fatal(err)
	}
	return req
}

func TestJsonRequestHandlerWithNormalJson(t *testing.T) {
	req := newRequest("GET", "/foo", "{}", t)
	rec := httptest.NewRecorder()
	JsonRequestHandler(okHandler).ServeHTTP(rec, req)
	if rec.Code != 200 {
		t.Errorf("should be ok for normal json")
	}
	if rec.Body.String() != "ok" {
		t.Errorf("cannot handle normal json")
	}
}

func TestJsonRequestHandlerWithInvalidJson(t *testing.T) {
	req := newRequest("GET", "/foo", "{", t)
	rec := httptest.NewRecorder()
	JsonRequestHandler(okHandler).ServeHTTP(rec, req)
	if rec.Code != 400 {
		t.Errorf("should raise 400 for invalid json")
	}
}

func TestJsonRequestHandlerTakeInput(t *testing.T) {
	req := newRequest("GET", "/foo", "{\"name\": \"foo\"}", t)
	rec := httptest.NewRecorder()
	JsonRequestHandler(echoNameHandler).ServeHTTP(rec, req)

	if rec.Code != 200 {
		t.Errorf("should be ok for normal json")
	}
	if rec.Body.String() != "foo" {
		t.Errorf("cannot take json input")
	}
}

func TestHandlerUse(t *testing.T) {
	req := newRequest("GET", "/foo", "{\"name\": \"foo\"}", t)
	rec := httptest.NewRecorder()

	HandlerUse(okHandler, notHandler).ServeHTTP(rec, req)

	if rec.Code != 200 {
		t.Errorf("cannot handle normal request")
	}

	respBody := rec.Body.String()
	if respBody != "not ok" {
		t.Errorf("cannot use middleware: %s", respBody)
	}
}

func TestHandlerFuncUse(t *testing.T) {
	req := newRequest("GET", "/foo", "{\"name\": \"foo\"}", t)
	rec := httptest.NewRecorder()

	HandlerFuncUse(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
		},
		notHandler,
	).ServeHTTP(rec, req)

	if rec.Code != 200 {
		t.Errorf("cannot handle normal request")
	}

	respBody := rec.Body.String()
	if respBody != "not ok" {
		t.Errorf("cannot use middleware: %s", respBody)
	}
}
