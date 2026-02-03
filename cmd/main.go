package main

import (
	"log"
	"morse-server/internal/server"
	"os"
)

var logger log.Logger

func main() {
	logger.SetOutput(os.Stdout)

	serv := server.Create(logger)

	err := serv.Server.ListenAndServe()
	if err != nil {
		serv.Logger.Fatal(err)
	}
}
