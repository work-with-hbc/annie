/*
HTTP gateway.
*/

package gateway

import (
	"fmt"
	"log"
	"net/http"

	httpApi "github.com/work-with-hbc/annie/pkg/behaviour/http"
	"github.com/work-with-hbc/annie/pkg/jsonconfig"
	ahttp "github.com/work-with-hbc/annie/pkg/utils/http"

	"github.com/gorilla/mux"
)

func StartHTTPGateway(config *jsonconfig.Config) {
	http.Handle("/", ahttp.HandlerUse(
		setupRoute(),
		ahttp.CORSHeaderHandler,
	))

	dest := fmt.Sprintf(
		"%s:%d",
		config.GetDefaultString("host", "127.0.0.1"),
		config.GetDefaultInt("port", 8080),
	)
	log.Printf("HTTP gateway started on: %s", dest)
	http.ListenAndServe(dest, nil)
}

func setupRoute() *mux.Router {
	route := mux.NewRouter()
	apiRoute := route.PathPrefix("/api/v1").Subrouter()

	thingRoute := apiRoute.PathPrefix("/thing").Subrouter()
	thingRoute.Handle("/{id}", httpApi.GetSomethingById).Methods("GET")
	thingRoute.Handle("/{id}", httpApi.RememberSomethingById).Methods("PUT")
	thingRoute.Handle("/", httpApi.RememberSomething).Methods("POST")

	listOfThingRoute := apiRoute.PathPrefix("/list/thing").Subrouter()
	listOfThingRoute.Handle("/{id}", httpApi.GetListById).Methods("GET")
	listOfThingRoute.Handle("/{id}", httpApi.RememberListById).Methods("PUT")
	listOfThingRoute.Handle("/{id}/item", httpApi.PushSomethingToListById).Methods("PUT")
	listOfThingRoute.Handle("/", httpApi.RememberList).Methods("POST")

	return route
}
