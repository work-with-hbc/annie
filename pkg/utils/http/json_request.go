/*
JSON request body handler
*/

package http

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/context"
)

func parseJsonBody(r *http.Request) (interface{}, error) {
	var reader io.Reader = r.Body
	rawBody, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	var body map[string]interface{}
	if err = json.Unmarshal(rawBody, &body); err != nil {
		return nil, err
	}

	return body, err
}

const JSON_REQUEST_BODY_KEY = "JSON_BODY"

func JsonRequestHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := parseJsonBody(r)
		if err == nil {
			context.Set(r, JSON_REQUEST_BODY_KEY, body)
			defer context.Clear(r)
			h.ServeHTTP(w, r)
		} else {
			http.Error(w, "Invalid json payload body", 400)
		}
	})
}

func GetJsonInput(r *http.Request) map[string]interface{} {
	payload, ok := context.GetOk(r, JSON_REQUEST_BODY_KEY)
	if !ok {
		return nil
	}

	return payload.(map[string]interface{})
}
