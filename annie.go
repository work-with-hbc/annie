package main

import (
	"flag"
	"log"
	"time"

	"github.com/bcho/annie/pkg/brain"
	"github.com/bcho/annie/pkg/gateway"
	"github.com/bcho/annie/pkg/jsonconfig"
)

const DEFAULT_CONFIG = "config/basic.json"

func main() {
	configPath := flag.String("config", DEFAULT_CONFIG, "config json")
	flag.Parse()

	config, err := jsonconfig.LoadFromFile(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	initBrain(config)
	initGateway(config)

	// TODO listen for channel
	for {
		time.Sleep(1 * time.Hour)
	}
}

func initBrain(config *jsonconfig.Config) {
	brain.SetupMemoryManagerWithConfig(config.GetConfigObj("memory"))
}

func initGateway(config *jsonconfig.Config) {
	enabledGatewayConfigs := config.GetArrayConfigObj("gateway.enabled")

	for _, config := range enabledGatewayConfigs {
		gatewayName := config.GetString("name")
		gatewayStarter := gateway.GetGatewayByName(gatewayName)
		if gatewayStarter == nil {
			log.Fatalf("gateway not found: %s", gatewayName)
		}

		go gatewayStarter(config.GetConfigObj("options"))
	}
}
