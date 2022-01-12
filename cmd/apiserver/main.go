package main

import (
	"os"

	"github.com/chenjr0719/golang-boilerplate/pkg/apiserver"
	"github.com/chenjr0719/golang-boilerplate/pkg/log"
)

func main() {
	server, err := apiserver.NewAPIServer()
	if err != nil {
		log.Fatal().Err(err).Msg("Create API server failed")
		os.Exit(1)
	}
	err = server.Run("0.0.0.0", 8080)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to run API server")
		os.Exit(1)
	}
}
