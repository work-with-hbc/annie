/*
Gateway defination.
*/

package gateway

import (
	"github.com/bcho/annie/pkg/jsonconfig"
)

type GatewayServiceStarter func(config *jsonconfig.Config)

func GetGatewayByName(name string) GatewayServiceStarter {
	switch name {
	case "http":
		return StartHTTPGateway
	}

	return nil
}
