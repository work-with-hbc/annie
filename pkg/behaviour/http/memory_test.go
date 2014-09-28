/*
Memory http service testcases
*/

package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/work-with-hbc/annie/pkg/brain"
	"github.com/work-with-hbc/annie/pkg/jsonconfig"
)

func setupBrain(t *testing.T) {
	config, _ := jsonconfig.LoadFromString("{}")
	brain.SetupMemoryManagerWithConfig(config)
}

func newRequest(method, url, content string, t *testing.T) *http.Request {
	req, err := http.NewRequest(method, url, strings.NewReader(content))
	if err != nil {
		t.Fatal(err)
	}
	return req
}

func makeInput(thing string, t *testing.T) string {
	var rv struct {
		Thing string `json:"thing"`
	}

	rv.Thing = thing
	input, err := json.Marshal(rv)
	if err != nil {
		t.Fatal(err)
	}

	return string(input)
}

func getIdFromBody(body string, t *testing.T) string {
	var rv struct {
		Id string `json:"id"`
	}

	err := json.Unmarshal([]byte(body), &rv)
	if err != nil {
		t.Errorf("cannot get id from body: %s", body)
	}

	return rv.Id
}

func getThingFromBody(body string, t *testing.T) string {
	var rv struct {
		Value string `json:"value"`
	}

	err := json.Unmarshal([]byte(body), &rv)
	if err != nil {
		t.Errorf("cannot get id from body: %s", body)
	}

	return rv.Value
}

func expectThingEqual(id string, thing string, t *testing.T) {
	route := mux.NewRouter()
	route.Handle("/{id}", GetSomethingById).Methods("GET")

	req := newRequest("GET", "/"+id, "", t)
	rec := httptest.NewRecorder()
	route.ServeHTTP(rec, req)

	if rec.Code != 200 {
		t.Errorf("should be 200 to get a thing, got %d", rec.Code)
	}
	rv := getThingFromBody(rec.Body.String(), t)
	if rv != thing {
		t.Errorf("got unexpected thing: %s", rv)
	}
}

func TestMemorySimpleInput(t *testing.T) {
	setupBrain(t)

	thing := "this_is_a_simple_string"
	req := newRequest("POST", "/foo", makeInput(thing, t), t)
	rec := httptest.NewRecorder()

	RememberSomething.ServeHTTP(rec, req)
	if rec.Code != 201 {
		t.Errorf("should be 201 to remember something, got %d", rec.Code)
	}

	id := getIdFromBody(rec.Body.String(), t)
	expectThingEqual(id, thing, t)
}

func TestMemoryWithComplicateInput(t *testing.T) {
	setupBrain(t)

	thing := "{\"string\":\"test_string\",\"int\":42,\"float\":3.141592653,\"nested\":{\"name\":\"test\"}}"
	req := newRequest("POST", "/foo", makeInput(thing, t), t)
	rec := httptest.NewRecorder()

	RememberSomething.ServeHTTP(rec, req)
	if rec.Code != 201 {
		t.Errorf("should be 201 to remember something, got %d", rec.Code)
	}

	id := getIdFromBody(rec.Body.String(), t)
	expectThingEqual(id, thing, t)
}

func TestRememberWithName(t *testing.T) {
	setupBrain(t)

	thingId := "test_key"
	thing := "this_is_a_simple_string"
	req := newRequest("PUT", "/"+thingId, makeInput(thing, t), t)
	rec := httptest.NewRecorder()

	route := mux.NewRouter()
	route.Handle("/{id}", RememberSomethingWithName).Methods("PUT")
	route.ServeHTTP(rec, req)
	if rec.Code != 200 {
		t.Errorf("should be 200 to remember something with name, got %d", rec.Code)
	}

	id := getIdFromBody(rec.Body.String(), t)
	if id != thingId {
		t.Errorf("stored item id should be %s, got %s", thingId, id)
	}
	expectThingEqual(id, thing, t)
}
