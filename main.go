package main

import (
	"context"
	"flag"
	"log"

	"github.com/coboshm/go_api_skeleton/pkg/server"
)

// Handle CLI arguments.
var env = flag.String("env", "testing", "Set application environment: -env=testing")
var debug = flag.Bool("debug", false, "Enable debug -debug")
var configDir = flag.String("config-dir", "config", "Set the application config dir: -config=/opt/config")

func init() {
}

func main() {
	config := server.LoadConfig(*debug, *env, *configDir)

	server, err := server.NewServer(config)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	err = server.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
