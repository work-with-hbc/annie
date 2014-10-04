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

func makeRememberThingInput(thing string, t *testing.T) string {
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

func makeRememberListInput(things []string, t *testing.T) string {
	var rv struct {
		Things []string `json:"things"`
	}

	rv.Things = things
	input, err := json.Marshal(rv)
	if err != nil {
		t.Fatal(err)
	}

	return string(input)
}

func makePushToListInput(thing string, t *testing.T) string {
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
		t.Errorf("cannot get thing from body: %s", body)
	}

	return rv.Value
}

func getThingsFromBody(body string, t *testing.T) []string {
	var rv struct {
		Value []string `json:"value"`
	}

	err := json.Unmarshal([]byte(body), &rv)
	if err != nil {
		t.Errorf("cannot get things from body: %s", body)
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

func expectThingsEqual(id string, things []string, t *testing.T) {
	route := mux.NewRouter()
	route.Handle("/{id}", GetListById).Methods("GET")

	req := newRequest("GET", "/"+id, "", t)
	rec := httptest.NewRecorder()
	route.ServeHTTP(rec, req)

	if rec.Code != 200 {
		t.Errorf("should be 200 to get a thing, got %d", rec.Code)
	}

	rv := getThingsFromBody(rec.Body.String(), t)
	for i, thing := range things {
		if thing != rv[i] {
			t.Errorf("expected: %s got: %s", thing, rv[i])
		}
	}
}

func TestRememberThingSimpleInput(t *testing.T) {
	setupBrain(t)

	thing := "this_is_a_simple_string"
	req := newRequest("POST", "/foo", makeRememberThingInput(thing, t), t)
	rec := httptest.NewRecorder()

	RememberSomething.ServeHTTP(rec, req)
	if rec.Code != 201 {
		t.Errorf("should be 201 to remember something, got %d", rec.Code)
	}

	id := getIdFromBody(rec.Body.String(), t)
	expectThingEqual(id, thing, t)
}

func TestRememberThingWithComplicateInput(t *testing.T) {
	setupBrain(t)

	thing := "{\"string\":\"test_string\",\"int\":42,\"float\":3.141592653,\"nested\":{\"name\":\"test\"}}"
	req := newRequest("POST", "/foo", makeRememberThingInput(thing, t), t)
	rec := httptest.NewRecorder()

	RememberSomething.ServeHTTP(rec, req)
	if rec.Code != 201 {
		t.Errorf("should be 201 to remember something, got %d", rec.Code)
	}

	id := getIdFromBody(rec.Body.String(), t)
	expectThingEqual(id, thing, t)
}

func TestRememberThingById(t *testing.T) {
	setupBrain(t)

	thingId := "test_key"
	thing := "this_is_a_simple_string"
	req := newRequest("PUT", "/"+thingId, makeRememberThingInput(thing, t), t)
	rec := httptest.NewRecorder()

	route := mux.NewRouter()
	route.Handle("/{id}", RememberSomethingById).Methods("PUT")
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

func TestRememberList(t *testing.T) {
	setupBrain(t)

	things := []string{
		"foo",
		"bar",
		"baz",
	}
	req := newRequest("POST", "/foo", makeRememberListInput(things, t), t)
	rec := httptest.NewRecorder()

	RememberList.ServeHTTP(rec, req)
	if rec.Code != 201 {
		t.Errorf("should be 201 to remember a list, got %d", rec.Code)
	}

	id := getIdFromBody(rec.Body.String(), t)
	expectThingsEqual(id, things, t)
}

func TestPushToList(t *testing.T) {
	setupBrain(t)

	things := []string{
		"foo",
		"bar",
	}
	pushed := "baz"
	req := newRequest("POST", "/foo", makeRememberListInput(things, t), t)
	rec := httptest.NewRecorder()

	RememberList.ServeHTTP(rec, req)
	if rec.Code != 201 {
		t.Errorf("should be 201 to remember a list, got %d", rec.Code)
	}

	things = append(things, pushed)
	id := getIdFromBody(rec.Body.String(), t)

	route := mux.NewRouter()
	route.Handle("/{id}", PushSomethingToListById).Methods("PUT")

	req = newRequest("PUT", "/"+id, makePushToListInput(pushed, t), t)
	rec = httptest.NewRecorder()
	route.ServeHTTP(rec, req)

	expectThingsEqual(id, things, t)
}
