package main

import (
	"flag"
	"log"

	"github.com/bcho/annie/pkg/brain"
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
}

func initBrain(config *jsonconfig.Config) {
	brain.SetupMemoryManagerWithConfig(config.GetConfigObj("memory"))
}
