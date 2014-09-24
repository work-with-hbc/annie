/*
Memory service through http.
*/

package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/work-with-hbc/annie/pkg/brain"
	ahttp "github.com/work-with-hbc/annie/pkg/utils/http"
)

func rememberSomething(w http.ResponseWriter, r *http.Request) {
	payload := ahttp.GetJsonInput(r)
	if payload == nil {
		ahttp.InvalidInput(w)
		return
	}
	something, ok := payload["thing"]
	if !ok {
		ahttp.InvalidInput(w)
		return
	}

	memory := brain.GetMemoryManager()
	itemId, err := memory.Remember(something.(string))
	if err != nil {
		ahttp.ServerError(w)
		return
	}

	w.WriteHeader(201)
	fmt.Fprintf(w, "{\"id\": \"%s\"}", itemId)
}

func getSomethingById(w http.ResponseWriter, r *http.Request) {
	memory := brain.GetMemoryManager()
	vars := mux.Vars(r)

	itemId := vars["id"]
	item, err := memory.GetFromId(itemId)
	// TODO separate not found error & other error behaviour.
	if err != nil {
		ahttp.NotFound(w)
		return
	}

	fmt.Fprintf(w, "{\"id\": \"%s\", \"value\": \"%s\"}", itemId, item.(string))
}

var (
	RememberSomething = ahttp.HandlerFuncUse(
		rememberSomething,
		ahttp.JsonRequestHandler,
		ahttp.JsonResponseHandler,
	)

	GetSomethingById = ahttp.HandlerFuncUse(
		getSomethingById,
		ahttp.JsonResponseHandler,
	)
)
