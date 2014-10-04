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

type getListByIdResponse struct {
	Id    string   `json:"id"`
	Value []string `json:"value"`
}

type pushSomethingToListResponse struct {
	Id string `json:"id"`
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
	itemId, err := memory.StoreString(something.(string))
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

func rememberSomethingById(w http.ResponseWriter, r *http.Request) {
	memory := brain.GetMemoryManager()

	vars := mux.Vars(r)
	itemId := vars["id"]

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

	err := memory.StoreStringByKey(itemId, something.(string))
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

	w.WriteHeader(200)
	w.Write(rv)
}

func getSomethingById(w http.ResponseWriter, r *http.Request) {
	memory := brain.GetMemoryManager()
	vars := mux.Vars(r)

	itemId := vars["id"]
	item, err := memory.GetString(itemId)
	// TODO separate not found error & other error behaviour.
	if err != nil {
		ahttp.NotFound(w)
		return
	}

	resp := new(getSomethingByIdResponse)
	resp.Id = itemId
	resp.Value = item
	rv, err := json.Marshal(resp)
	if err != nil {
		ahttp.ServerError(w)
		return
	}

	w.WriteHeader(200)
	w.Write(rv)
}

func rememberList(w http.ResponseWriter, r *http.Request) {
	payload := ahttp.GetJsonInput(r)
	if payload == nil {
		ahttp.InvalidInput(w)
		return
	}
	listRaw, ok := payload["things"]
	if !ok {
		ahttp.InvalidInput(w)
		return
	}
	var list []string
	for _, thing := range listRaw.([]interface{}) {
		list = append(list, thing.(string))
	}

	memory := brain.GetMemoryManager()
	id, err := memory.StoreList(list)
	if err != nil {
		ahttp.ServerError(w)
		return
	}

	resp := new(rememberSomethingResponse)
	resp.Id = id
	rv, err := json.Marshal(resp)
	if err != nil {
		ahttp.ServerError(w)
		return
	}

	w.WriteHeader(201)
	w.Write(rv)
}

func rememberListById(w http.ResponseWriter, r *http.Request) {
	memory := brain.GetMemoryManager()

	vars := mux.Vars(r)
	id := vars["id"]

	payload := ahttp.GetJsonInput(r)
	if payload == nil {
		ahttp.InvalidInput(w)
		return
	}

	listRaw, ok := payload["things"]
	if !ok {
		ahttp.InvalidInput(w)
		return
	}
	var list []string
	for _, thing := range listRaw.([]interface{}) {
		list = append(list, thing.(string))
	}

	err := memory.StoreListByKey(id, list)
	if err != nil {
		ahttp.ServerError(w)
		return
	}

	resp := new(rememberSomethingResponse)
	resp.Id = id
	rv, err := json.Marshal(resp)
	if err != nil {
		ahttp.ServerError(w)
		return
	}

	w.WriteHeader(200)
	w.Write(rv)
}

func pushSomethingToListById(w http.ResponseWriter, r *http.Request) {
	payload := ahttp.GetJsonInput(r)
	if payload == nil {
		ahttp.InvalidInput(w)
		return
	}

	memory := brain.GetMemoryManager()
	vars := mux.Vars(r)

	id := vars["id"]
	if !memory.Has(id) {
		ahttp.NotFound(w)
		return
	}

	something, ok := payload["thing"]
	if !ok {
		ahttp.InvalidInput(w)
		return
	}
	err := memory.PushList(id, something.(string))
	if err != nil {
		ahttp.ServerError(w)
		return
	}

	resp := new(pushSomethingToListResponse)
	resp.Id = id
	rv, err := json.Marshal(resp)
	if err != nil {
		ahttp.ServerError(w)
		return
	}

	w.WriteHeader(200)
	w.Write(rv)
}

func getListById(w http.ResponseWriter, r *http.Request) {
	memory := brain.GetMemoryManager()
	vars := mux.Vars(r)

	var list []string
	id := vars["id"]
	err := memory.GetList(id, &list)
	// TODO separate not found error & other error behaviour.
	if err != nil {
		ahttp.NotFound(w)
		return
	}

	resp := new(getListByIdResponse)
	resp.Id = id
	resp.Value = list
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

	RememberSomethingById = ahttp.HandlerFuncUse(
		rememberSomethingById,
		ahttp.JsonRequestHandler,
		ahttp.JsonResponseHandler,
	)

	GetSomethingById = ahttp.HandlerFuncUse(
		getSomethingById,
		ahttp.JsonResponseHandler,
	)

	RememberList = ahttp.HandlerFuncUse(
		rememberList,
		ahttp.JsonRequestHandler,
		ahttp.JsonResponseHandler,
	)

	RememberListById = ahttp.HandlerFuncUse(
		rememberListById,
		ahttp.JsonRequestHandler,
		ahttp.JsonResponseHandler,
	)

	PushSomethingToListById = ahttp.HandlerFuncUse(
		pushSomethingToListById,
		ahttp.JsonRequestHandler,
		ahttp.JsonResponseHandler,
	)

	GetListById = ahttp.HandlerFuncUse(
		getListById,
		ahttp.JsonResponseHandler,
	)
)
