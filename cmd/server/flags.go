package main

import (
	"flag"
	"os"
)

var flagAddress string

func parseFlags() {
	flag.StringVar(&flagAddress, "a", "localhost:8080", "address with port for the server to run on")
	flag.Parse()

	if envServerAddress := os.Getenv("ADDRESS"); envServerAddress != "" {
		flagAddress = envServerAddress
	}
}
