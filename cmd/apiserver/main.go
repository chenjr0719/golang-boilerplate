package main

import "github.com/chenjr0719/golang-boilerplate/pkg/apiserver"

func main() {
	server := apiserver.NewAPIServer()
	server.Run("0.0.0.0", 8080)
}
