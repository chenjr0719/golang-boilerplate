package main

import (
	"github.com/chenjr0719/golang-boilerplate/pkg/log"
	"github.com/chenjr0719/golang-boilerplate/pkg/worker"
)

func main() {
	err := worker.NewWorker()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start worker")
	}
}
