/*
Memory service through http.
*/

package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/work-with-hbc/annie/pkg/brain"
	ahttp "github.com/work-with-hbc/annie/pkg/utils/http"
)

type rememberSomethingResponse struct {
	Id string `json:"id"`
}

type getSomethingByIdResponse struct {
	Id    string `json:"id"`
	Value string `json:"value"`
}

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

	resp := new(rememberSomethingResponse)
	resp.Id = itemId
	rv, err := json.Marshal(resp)
	if err != nil {
		ahttp.ServerError(w)
		return
	}

	w.WriteHeader(201)
	w.Write(rv)
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

	resp := new(getSomethingByIdResponse)
	resp.Id = itemId
	resp.Value = item.(string)
	rv, err := json.Marshal(resp)
	if err != nil {
		ahttp.ServerError(w)
		return
	}

	w.WriteHeader(200)
	w.Write(rv)
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
