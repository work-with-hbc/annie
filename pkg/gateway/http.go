/*
HTTP gateway.
*/

package gateway

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bcho/annie/pkg/jsonconfig"
	"github.com/gorilla/mux"
)

func StartHTTPGateway(config *jsonconfig.Config) {
	route := mux.NewRouter()

	http.Handle("/", route)

	dest := fmt.Sprintf(
		"%s:%d",
		config.GetDefaultString("host", "127.0.0.1"),
		config.GetDefaultInt("port", 8080),
	)
	log.Printf("HTTP gateway started on: %s", dest)
	http.ListenAndServe(dest, nil)
}
